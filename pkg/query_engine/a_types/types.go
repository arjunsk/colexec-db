package types

type T uint8

const (
	T_int32 T = iota
	T_int64
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
	case T_int64:
		typ.Size = 8
	default:
		panic("Unknown type")
	}
	return typ

}

type FixedSizeT interface {
	~int32
}
