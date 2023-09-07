package vector

import types "colexecdb/pkg/query_engine/a_types"
import "github.com/RoaringBitmap/roaring"

type Vector struct {
	typ    types.Type
	col    any
	length int
	nsp    roaring.Bitmap
}

func NewVec(typ types.Type) *Vector {
	return &Vector{
		typ: typ,
	}
}

func Append[T any](vec *Vector, val T, isNull bool) error {
	length := vec.length
	vec.length++
	if isNull {
		vec.nsp.Add(uint32(length))
	} else {
		col := vec.col.([]T)
		col[length] = val
	}
	return nil
}

func (v *Vector) Free() {

}

func Get[T any](vec *Vector, i uint32) (res T, isNull bool) {
	if vec.nsp.Contains(i) {
		return res, true
	}
	return vec.col.([]T)[i], false
}
