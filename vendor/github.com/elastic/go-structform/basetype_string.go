// Code generated by "stringer -type=BaseType"; DO NOT EDIT

package structform

import "fmt"

const _BaseType_name = "AnyTypeByteTypeStringTypeBoolTypeZeroTypeIntTypeInt8TypeInt16TypeInt32TypeInt64TypeUintTypeUint8TypeUint16TypeUint32TypeUint64TypeFloat32TypeFloat64Type"

var _BaseType_index = [...]uint8{0, 7, 15, 25, 33, 41, 48, 56, 65, 74, 83, 91, 100, 110, 120, 130, 141, 152}

func (i BaseType) String() string {
	if i >= BaseType(len(_BaseType_index)-1) {
		return fmt.Sprintf("BaseType(%d)", i)
	}
	return _BaseType_name[_BaseType_index[i]:_BaseType_index[i+1]]
}
