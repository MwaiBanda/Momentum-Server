package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v3"
)

// BindAndValidate decodes a JSON request body and verifies fields marked as
// required in the DTO's binding tags are present.
func BindAndValidate[T any](ctx fiber.Ctx) (T, error) {
	var input T
	if err := ctx.Bind().Body(&input); err != nil {
		return input, fiber.ErrBadRequest
	}

	if err := validateRequiredFields(ctx.Body(), reflect.TypeOf(input), ""); err != nil {
		return input, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return input, nil
}

func validateRequiredFields(body []byte, typ reflect.Type, path string) error {
	body = bytes.TrimSpace(body)
	if len(body) == 0 || bytes.Equal(body, []byte("null")) {
		return fmt.Errorf("invalid request body")
	}

	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct || typ.PkgPath() == "time" {
		return nil
	}

	var object map[string]json.RawMessage
	if err := json.Unmarshal(body, &object); err != nil || object == nil {
		return fmt.Errorf("invalid request body")
	}

	for fieldIndex := 0; fieldIndex < typ.NumField(); fieldIndex++ {
		field := typ.Field(fieldIndex)
		if field.PkgPath != "" {
			continue
		}

		name, omitted := jsonFieldName(field)
		if omitted {
			continue
		}

		fieldPath := name
		if path != "" {
			fieldPath = path + "." + name
		}

		rawValue, present := object[name]
		if isRequired(field) && (!present || bytes.Equal(bytes.TrimSpace(rawValue), []byte("null"))) {
			return fmt.Errorf("missing required field %q", fieldPath)
		}
		if !present || bytes.Equal(bytes.TrimSpace(rawValue), []byte("null")) {
			continue
		}

		fieldType := field.Type
		for fieldType.Kind() == reflect.Pointer {
			fieldType = fieldType.Elem()
		}

		switch fieldType.Kind() {
		case reflect.Struct:
			if err := validateRequiredFields(rawValue, fieldType, fieldPath); err != nil {
				return err
			}
		case reflect.Array, reflect.Slice:
			var values []json.RawMessage
			if err := json.Unmarshal(rawValue, &values); err != nil {
				return fmt.Errorf("invalid request body")
			}
			elementType := fieldType.Elem()
			for index, value := range values {
				if err := validateRequiredFields(value, elementType, fmt.Sprintf("%s[%d]", fieldPath, index)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func isRequired(field reflect.StructField) bool {
	tag := field.Tag.Get("binding")
	for _, option := range strings.Split(tag, ",") {
		if option == "required" {
			return true
		}
	}
	return false
}

func jsonFieldName(field reflect.StructField) (string, bool) {
	tag := field.Tag.Get("json")
	name := strings.Split(tag, ",")[0]
	if name == "-" {
		return "", true
	}
	if name == "" {
		return field.Name, false
	}
	return name, false
}
