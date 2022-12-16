package types

type GasValues struct {
	Init   uint64
	Import uint64
	Write  uint64
	Read   uint64
}

func DefaultGasValues() GasValues {
	return GasValues{
		Init:   5000,
		Import: 2000,
		Write:  2000,
		Read:   1000,
	}
}
