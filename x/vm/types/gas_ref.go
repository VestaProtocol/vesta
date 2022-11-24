package types

type GasValues struct {
	Import uint64
	Write  uint64
	Read   uint64
}

func DefaultGasValues() GasValues {
	return GasValues{
		Import: 20,
		Write:  2,
		Read:   1,
	}
}
