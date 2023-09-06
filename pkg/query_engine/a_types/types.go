package types

type T uint8

const (
	T_int32 T = iota
)

type Type struct {
	Oid  T
	Size int32
}

func (t T) ToType() Type {
	var typ Type

	typ.Oid = t
	switch t {
	case T_int32:
		typ.Size = 4
	default:
		panic("Unknown type")
	}
	return typ

}

type FixedSizeT interface {
	~int32
}
