package types

const (
	// ModuleName defines the module name
	ModuleName = "vm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_vm"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ContractsKey      = "Contracts/value/"
	ContractsCountKey = "Contracts/count/"
)
