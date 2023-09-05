package types

type T uint8

const (
	T_int32 T = iota
	T_float32
)

type Type struct {
	Oid T
}

type FixedSizeT interface {
	~float32 | ~int32
}
