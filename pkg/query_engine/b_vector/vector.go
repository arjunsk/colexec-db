package vector

import (
	types "colexecdb/pkg/query_engine/a_types"
	"fmt"
)
import "github.com/RoaringBitmap/roaring"

type Vector struct {
	typ    *types.Type
	col    any
	length int
	nsp    *roaring.Bitmap
}

func NewVec(typ types.Type) *Vector {
	v := &Vector{
		typ: &typ,
		nsp: roaring.New(),
	}

	switch typ.Oid {
	case types.T_int32:
		v.col = make([]int32, 0)
	case types.T_int64:
		v.col = make([]int64, 0)
	}
	return v
}

func Append[T any](vec *Vector, val T, isNull bool) error {
	length := vec.length
	vec.length++
	if isNull {
		vec.nsp.Add(uint32(length))
	} else {
		//TODO: very simple append instead of mpool based pre-extend etc.
		col := vec.col.([]T)
		col = append(col, val)
		vec.col = col
	}
	return nil
}

func (v *Vector) Free() {
	v.nsp.Clear()
	v.col = nil
	v.length = 0
	v.typ = nil
}

func (v *Vector) Length() int {
	return v.length
}

func (v *Vector) GetType() *types.Type {
	return v.typ
}

func (v *Vector) String() string {
	switch v.typ.Oid {
	case types.T_int32:
		return vecToString[int32](v)
	case types.T_int64:
		return vecToString[int64](v)
	default:
		panic("vec to string unknown types.")
	}
}

func vecToString[T types.FixedSizeT](v *Vector) string {
	col := v.col.([]T)
	if v.nsp.GetCardinality() > 0 {
		return fmt.Sprintf("%v-%s", col, v.nsp.String())
	} else {
		return fmt.Sprintf("%v", col)
	}
}

func Get[T any](vec *Vector, i uint32) (res T, isNull bool) {
	if vec.nsp.Contains(i) {
		return res, true
	}
	return vec.col.([]T)[i], false
}
