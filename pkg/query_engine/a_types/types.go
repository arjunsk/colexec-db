package types

type T uint8

const (
	T_int32 T = iota
)

type Type struct {
	Oid T
}

type FixedSizeT interface {
	~int32
}
