package telerik

import (
	"encoding/json"
	"fmt"
	"github.com/dsoprea/go-logging"
	"reflect"
)

func (el *ToJavaScriptConverter) ToSJsonSchema(elementName, jsType string, element reflect.Value) []byte {
	var fieldName, tagKeyNewName, tagDescription, tagEnum, tagRequired, tagPattern, tagMinimum, tagMaximum, tagType, tagComplex string

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

		_, tagDescription = el.getTagDataByName("jsonSchema_description", tag)
		if tagDescription == "-" {
			continue
		}

		_, tagKeyNewName = el.getTagDataByName("jsonSchema_keyNewName", tag)
		if tagKeyNewName != "" {
			fieldName = tagKeyNewName
		}

		_, tagEnum = el.getTagDataByName("jsonSchema_enum", tag)

		_, tagRequired = el.getTagDataByName("jsonSchema_required", tag)

		_, tagPattern = el.getTagDataByName("jsonSchema_pattern", tag)

		_, tagMinimum = el.getTagDataByName("jsonSchema_minimum", tag)

		_, tagMaximum = el.getTagDataByName("jsonSchema_minimum", tag)

		_, tagType = el.getTagDataByName("jsonSchema_type", tag)

		_, tagComplex = el.getTagDataByName("jsonSchema_complex", tag)
		if tagComplex != "" {
			var tagComplexMap map[string]interface{}
			err := json.Unmarshal([]byte(tagComplex), &tagComplexMap)
			if err != nil {
				log.Panicf("error: %v", err.Error())
			}

			properties[fieldName] = tagComplexMap
			continue
		}

		properties[fieldName] = make(map[string]interface{})

		if tagDescription != "" {
			properties[fieldName]["description"] = tagDescription
		}

		if tagEnum != "" {

			var tagEnumArr []string
			err := json.Unmarshal([]byte(tagEnum), &tagEnumArr)
			if err != nil {
				log.Panicf("error: %v", err.Error())
			}

			properties[fieldName]["enum"] = tagEnumArr
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

		switch converted := field.Interface().(type) {

		case TypeHtmlDropZone:
			properties[fieldName] = map[string]interface{}{
				"type": "string",
				"enum": TypeHtmlDropZones[1:],
			}

		case HtmlGlobalAttributes:

			newTagMapStringInterface := make(map[string]interface{})
			err := json.Unmarshal([]byte(el.ToSJsonSchema("", "object", reflect.ValueOf(converted))), &newTagMapStringInterface)
			if err != nil {
				log.Panicf("error: %v", err.Error())
			}
			properties[fieldName] = newTagMapStringInterface

		case map[string]interface{}:

			properties[fieldName] = map[string]interface{}{
				"type": "object",
			}

		case map[string]string:

			properties[fieldName] = map[string]interface{}{
				"type": "object",
			}

		case string:
			properties[fieldName]["type"] = "string"
		case Boolean:
			properties[fieldName]["type"] = "boolean"
		case int:
			properties[fieldName]["type"] = "number"

		default:
			fmt.Printf("\n\n%golangToSJsonSchema.go: T\n\n", converted)

		}

		if tagType != "" {
			properties[fieldName]["type"] = tagType
		}
	}

	var ret []byte
	var err error
	if elementName == "" {
		ret, err = json.Marshal(map[string]interface{}{
			"type":       jsType,
			"properties": properties,
			"required":   required,
		})
	} else {
		ret, err = json.Marshal(map[string]interface{}{
			elementName: map[string]interface{}{
				"type":       jsType,
				"properties": properties,
				"required":   required,
			},
		})
	}

	if err != nil {
		log.Panic(err)
	}

	return ret
}
