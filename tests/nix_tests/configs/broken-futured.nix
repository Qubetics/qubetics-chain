{ pkgs ? import ../../../nix { } }:
let qubeticsd = (pkgs.callPackage ../../../. { });
in
qubeticsd.overrideAttrs (oldAttrs: {
  patches = oldAttrs.patches or [ ] ++ [
    ./broken-qubeticsd.patch
  ];
})
