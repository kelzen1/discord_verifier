package utils

import (
	"reflect"
	"strings"
)

func getFieldName(field reflect.StructField) (name string, valid bool) {
	tag, found := field.Tag.Lookup("json")

	if !found {
		return
	}

	separated := strings.Split(tag, ",")

	for _, value := range separated {
		// если есть еще какие то другие поля то гг
		if value == "omitempty" || value == "-" {
			return
		}
	}

	return tag, true
}

func getFields(Struct any) (retVal map[string]interface{}) {
	retVal = make(map[string]interface{})

	value := reflect.ValueOf(Struct)
	valueT := value.Type()

	for i := 0; i < value.Type().NumField(); i++ {
		field := valueT.Field(i)
		fieldName, found := getFieldName(field)

		if found {
			retVal[fieldName] = byte(0)
		}

	}

	return
}

func CompareJSONToStruct(mapped map[string]interface{}, Struct any) (retVal bool) {

	structFields := getFields(Struct)

	if len(structFields) == 0 {
		return true
	}

	for key := range structFields {
		_, found := mapped[key]

		if !found {
			return
		}

	}

	return true
}
