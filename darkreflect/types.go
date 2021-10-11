package darkreflect

import "reflect"

func GetOffset(t reflect.Type, f string) (uintptr, bool) {
	field, ok := t.FieldByName(f)
	if !ok {
		return 0, false
	}
	return field.Offset, true
}

func GetType(name string) reflect.Type {
	types := TypesByString(name)
	if len(types) == 0 {
		return nil
	}
	return ToType(types[0])
}

func MustGetType(name string) reflect.Type {
	t := GetType(name)
	if t == nil {
		panic("type not found")
	}
	return t
}

func MustGetOffset(t reflect.Type, f string) uintptr {
	offset, ok := GetOffset(t, f)
	if !ok {
		panic("field not found")
	}
	return offset
}
