package telerik

import (
	"encoding/json"
	"github.com/dsoprea/go-logging"
	"reflect"
)

func (el *ToJavaScriptConverter) ToSJsonSchema(elementName, jsType string, element reflect.Value) []byte {
	var fieldName, tagDescription, tagEnum, tagRequired, tagPattern, tagMinimum, tagMaximum, tagType string

	var required []string
	var properties map[string]map[string]interface{}

	if len(properties) == 0 {
		properties = make(map[string]map[string]interface{})
	}

	if len(required) == 0 {
		required = make([]string, 0)
	}

	for i := 0; i < element.NumField(); i += 1 {
		field := element.Field(i)
		typeField := element.Type().Field(i)
		tag := typeField.Tag

		_, fieldName = el.getTagData(tag)
		if fieldName == "-" {
			continue
		}

		_, tagDescription = el.getTagDataByName("description", tag)

		_, tagEnum = el.getTagDataByName("enum", tag)

		_, tagRequired = el.getTagDataByName("required", tag)

		_, tagPattern = el.getTagDataByName("pattern", tag)

		_, tagMinimum = el.getTagDataByName("minimum", tag)

		_, tagMaximum = el.getTagDataByName("maximum", tag)

		_, tagType = el.getTagDataByName("type", tag)

		properties[fieldName] = make(map[string]interface{})

		if tagDescription != "" {
			properties[fieldName]["description"] = tagDescription
		}

		if tagEnum != "" {
			properties[fieldName]["enum"] = tagEnum
		}

		if tagPattern != "" {
			properties[fieldName]["pattern"] = tagPattern
		}

		if tagMinimum != "" {
			properties[fieldName]["minimum"] = tagMinimum
		}

		if tagMaximum != "" {
			properties[fieldName]["maximum"] = tagMaximum
		}

		if tagRequired == "true" || tagRequired == "True" || tagRequired == "TRUE" || tagRequired == "1" {
			required = append(required, fieldName)
		}

		switch field.Interface().(type) {

		case string:
			properties[fieldName]["type"] = "string"
		case Boolean:
			properties[fieldName]["type"] = "boolean"
		case int:
			properties[fieldName]["type"] = "number"
		}

		if tagType != "" {
			properties[fieldName]["type"] = tagType
		}
	}

	ret, err := json.Marshal(map[string]interface{}{
		elementName: map[string]interface{}{
			"type":       jsType,
			"properties": properties,
			"required":   required,
		},
	})

	if err != nil {
		log.Panic(err)
	}

	return ret
}
