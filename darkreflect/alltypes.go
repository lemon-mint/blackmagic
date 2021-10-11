package darkreflect

import (
	"reflect"
)

func AllTypes() []reflect.Type {
	sections, offset := Typelinks()
	var types []reflect.Type
	for i, offs := range offset {
		for _, off := range offs {
			t := RtypeOff(sections[i], off)
			types = append(types, ToType(t))
		}
	}
	return types
}
