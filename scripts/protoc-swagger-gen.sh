#!/usr/bin/env sh

# Generates the Swagger (OpenAPI v2) spec for the REST (gRPC gateway) API and
# writes it to client/docs/swagger-ui/swagger.yaml.
#
# The spec combines:
#   - the Cosmos SDK spec shipped with the Cosmos SDK module (pre-generated),
#   - IBC and IBC rate-limiting query services (from their Go module protos),
#   - the Cosmos EVM and Qubetics query services in ./proto.
#
# Run through the Makefile (make proto-swagger-gen), which copies the third-party
# files into ./tmp-swagger-deps and runs this script in the proto-builder image.
set -eo pipefail

OUT=/workspace/tmp-swagger-gen
TEMPLATE='{"version":"v1","plugins":[{"name":"swagger","out":"'"$OUT"'","opt":"logtostderr=true,fqn_for_swagger_name=true,simple_operation_ids=true"}]}'

rm -rf "$OUT"
mkdir -p "$OUT"

# generate swagger files for every proto file that defines a query service
generate() {
	local root=$1
	shift
	(
		cd "$root"
		for file in $(find "$@" -name '*.proto' | sort); do
			if grep -qE '^service (Query|QueryService) ' "$file"; then
				buf generate --template "$TEMPLATE" "$file"
			fi
		done
	)
}

generate /workspace/proto cosmos qubetics
generate /workspace/tmp-swagger-deps/ibc-go/proto ibc
generate /workspace/tmp-swagger-deps/ratelimit/proto ratelimit

# Prefix operation IDs with their package (e.g. QubeticsNodeV3Params) so that
# services sharing method names (Params, Nodes, ...) do not collide, then write
# the swagger-combine configuration.
node - "$OUT" /workspace/tmp-swagger-deps/cosmos-sdk/swagger.yaml /workspace/tmp-swagger-gen/config.json <<'JS'
const fs = require("fs");
const path = require("path");
const [outDir, sdkSpec, configFile] = process.argv.slice(2);

const walk = (dir) =>
  fs.readdirSync(dir, { withFileTypes: true }).flatMap((e) =>
    e.isDirectory() ? walk(path.join(dir, e.name)) : [path.join(dir, e.name)]
  );

const specs = walk(outDir).filter((f) => f.endsWith(".swagger.json")).sort();
const apis = [{ url: sdkSpec }];

for (const file of specs) {
  const rel = path.relative(outDir, path.dirname(file));
  const prefix = rel
    .split(/[\/_-]/)
    .filter(Boolean)
    .map((p) => p[0].toUpperCase() + p.slice(1))
    .join("");
  const spec = JSON.parse(fs.readFileSync(file, "utf8"));
  if (!spec.paths || Object.keys(spec.paths).length === 0) continue;
  for (const ops of Object.values(spec.paths)) {
    for (const op of Object.values(ops)) {
      if (op && op.operationId) op.operationId = prefix + op.operationId;
    }
  }
  fs.writeFileSync(file, JSON.stringify(spec, null, 2));
  apis.push({ url: file });
}

const config = {
  swagger: "2.0",
  info: {
    title: "Qubetics - gRPC Gateway docs",
    description: "A REST interface for state queries",
    version: "1.0.0",
  },
  apis,
};
fs.writeFileSync(configFile, JSON.stringify(config, null, 2));
JS

# combine the swagger files into a single spec
swagger-combine "$OUT/config.json" -o "$OUT/combined.yaml" -f yaml --continueOnConflictingPaths true --includeDefinitions true

# swagger-combine inlines $refs, which turns recursive types (e.g. the Cosmos SDK
# tx ModeInfo, whose multisig variant contains a list of ModeInfo) into
# self-referencing YAML aliases that Swagger UI cannot load ("does not specify a
# valid version field"). Replace each recursive back-reference with a generic
# object and write plain YAML without anchors/aliases.
NODE_PATH="$(npm root -g)/swagger-combine/node_modules" node - "$OUT/combined.yaml" /workspace/client/docs/swagger-ui/swagger.yaml <<'JS'
const fs = require("fs");
const yaml = require("js-yaml");
const [input, output] = process.argv.slice(2);

const spec = yaml.load(fs.readFileSync(input, "utf8"));

let cut = 0;
const ancestors = new Set();
const breakCycles = (node) => {
  if (node === null || typeof node !== "object") return node;
  if (ancestors.has(node)) {
    cut++;
    return { type: "object", description: "Recursive type; see the enclosing schema." };
  }
  ancestors.add(node);
  const copy = Array.isArray(node) ? [] : {};
  for (const [k, v] of Object.entries(node)) copy[k] = breakCycles(v);
  ancestors.delete(node);
  return copy;
};

const acyclic = breakCycles(spec);
fs.writeFileSync(output, yaml.dump(acyclic, { noRefs: true, lineWidth: -1 }));
console.log(`swagger.yaml written (${cut} recursive references replaced)`);
JS

# clean up the intermediate files
rm -rf "$OUT"
