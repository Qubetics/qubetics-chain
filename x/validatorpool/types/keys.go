package types

const (
	// ModuleName defines the module name
	ModuleName = "validatorpool"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_validatorpool"
)

// Store keys
var (
	ValidatorPoolKey = []byte{0x01} // Key for validator pool state
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
