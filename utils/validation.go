package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   interface{} `json:"value,omitempty"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("field '%s': %s", e.Field, e.Message)
}

// ValidateStruct validates a struct using reflection and validation tags
func ValidateStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("validation requires a struct, got %s", v.Kind())
	}

	var errors []ValidationError
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		
		// Skip unexported fields
		if !field.CanInterface() {
			continue
		}

		validateTag := fieldType.Tag.Get("validate")
		if validateTag == "" {
			continue
		}

		fieldName := getFieldName(fieldType)
		fieldValue := field.Interface()

		// Parse validation rules
		rules := parseValidationRules(validateTag)
		for _, rule := range rules {
			if err := validateField(fieldName, fieldValue, rule); err != nil {
				errors = append(errors, *err)
			}
		}
	}

	if len(errors) > 0 {
		return &ValidationErrors{Errors: errors}
	}

	return nil
}

// ValidationErrors contains multiple validation errors
type ValidationErrors struct {
	Errors []ValidationError
}

func (e *ValidationErrors) Error() string {
	if len(e.Errors) == 1 {
		return e.Errors[0].Error()
	}
	
	var messages []string
	for _, err := range e.Errors {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, "; ")
}

// ValidationRule represents a single validation rule
type ValidationRule struct {
	Name  string
	Param string
}

func parseValidationRules(tag string) []ValidationRule {
	var rules []ValidationRule
	parts := strings.Split(tag, ",")
	
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		
		if strings.Contains(part, "=") {
			ruleParts := strings.SplitN(part, "=", 2)
			rules = append(rules, ValidationRule{
				Name:  strings.TrimSpace(ruleParts[0]),
				Param: strings.TrimSpace(ruleParts[1]),
			})
		} else {
			rules = append(rules, ValidationRule{
				Name: part,
			})
		}
	}
	
	return rules
}

func validateField(fieldName string, value interface{}, rule ValidationRule) *ValidationError {
	switch rule.Name {
	case "required":
		return validateRequired(fieldName, value)
	case "min":
		return validateMin(fieldName, value, rule.Param)
	case "max":
		return validateMax(fieldName, value, rule.Param)
	case "email":
		return validateEmail(fieldName, value)
	case "url":
		return validateURL(fieldName, value)
	case "uuid":
		return validateUUID(fieldName, value)
	case "oneof":
		return validateOneOf(fieldName, value, rule.Param)
	case "len":
		return validateLength(fieldName, value, rule.Param)
	case "alphanum":
		return validateAlphaNum(fieldName, value)
	case "numeric":
		return validateNumeric(fieldName, value)
	case "datetime":
		return validateDateTime(fieldName, value, rule.Param)
	}
	
	return nil
}

func validateRequired(fieldName string, value interface{}) *ValidationError {
	if isZeroValue(value) {
		return &ValidationError{
			Field:   fieldName,
			Message: "is required",
			Value:   value,
		}
	}
	return nil
}

func validateMin(fieldName string, value interface{}, param string) *ValidationError {
	minVal, err := strconv.Atoi(param)
	if err != nil {
		return nil
	}

	switch v := value.(type) {
	case string:
		if len(v) < minVal {
			return &ValidationError{
				Field:   fieldName,
				Message: fmt.Sprintf("must be at least %d characters", minVal),
				Value:   value,
			}
		}
	case int, int8, int16, int32, int64:
		intVal := reflect.ValueOf(v).Int()
		if intVal < int64(minVal) {
			return &ValidationError{
				Field:   fieldName,
				Message: fmt.Sprintf("must be at least %d", minVal),
				Value:   value,
			}
		}
	case []interface{}:
		if len(v) < minVal {
			return &ValidationError{
				Field:   fieldName,
				Message: fmt.Sprintf("must contain at least %d items", minVal),
				Value:   value,
			}
		}
	}
	
	return nil
}

func validateMax(fieldName string, value interface{}, param string) *ValidationError {
	maxVal, err := strconv.Atoi(param)
	if err != nil {
		return nil
	}

	switch v := value.(type) {
	case string:
		if len(v) > maxVal {
			return &ValidationError{
				Field:   fieldName,
				Message: fmt.Sprintf("must be at most %d characters", maxVal),
				Value:   value,
			}
		}
	case int, int8, int16, int32, int64:
		intVal := reflect.ValueOf(v).Int()
		if intVal > int64(maxVal) {
			return &ValidationError{
				Field:   fieldName,
				Message: fmt.Sprintf("must be at most %d", maxVal),
				Value:   value,
			}
		}
	case []interface{}:
		if len(v) > maxVal {
			return &ValidationError{
				Field:   fieldName,
				Message: fmt.Sprintf("must contain at most %d items", maxVal),
				Value:   value,
			}
		}
	}
	
	return nil
}

func validateEmail(fieldName string, value interface{}) *ValidationError {
	str, ok := value.(string)
	if !ok {
		return nil
	}
	
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(str) {
		return &ValidationError{
			Field:   fieldName,
			Message: "must be a valid email address",
			Value:   value,
		}
	}
	
	return nil
}

func validateURL(fieldName string, value interface{}) *ValidationError {
	str, ok := value.(string)
	if !ok {
		return nil
	}
	
	urlRegex := regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
	if !urlRegex.MatchString(str) {
		return &ValidationError{
			Field:   fieldName,
			Message: "must be a valid URL",
			Value:   value,
		}
	}
	
	return nil
}

func validateUUID(fieldName string, value interface{}) *ValidationError {
	str, ok := value.(string)
	if !ok {
		return nil
	}
	
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	if !uuidRegex.MatchString(str) {
		return &ValidationError{
			Field:   fieldName,
			Message: "must be a valid UUID",
			Value:   value,
		}
	}
	
	return nil
}

func validateOneOf(fieldName string, value interface{}, param string) *ValidationError {
	str, ok := value.(string)
	if !ok {
		return nil
	}
	
	options := strings.Split(param, " ")
	for _, option := range options {
		if str == option {
			return nil
		}
	}
	
	return &ValidationError{
		Field:   fieldName,
		Message: fmt.Sprintf("must be one of: %s", strings.Join(options, ", ")),
		Value:   value,
	}
}

func validateLength(fieldName string, value interface{}, param string) *ValidationError {
	expectedLen, err := strconv.Atoi(param)
	if err != nil {
		return nil
	}

	switch v := value.(type) {
	case string:
		if len(v) != expectedLen {
			return &ValidationError{
				Field:   fieldName,
				Message: fmt.Sprintf("must be exactly %d characters", expectedLen),
				Value:   value,
			}
		}
	case []interface{}:
		if len(v) != expectedLen {
			return &ValidationError{
				Field:   fieldName,
				Message: fmt.Sprintf("must contain exactly %d items", expectedLen),
				Value:   value,
			}
		}
	}
	
	return nil
}

func validateAlphaNum(fieldName string, value interface{}) *ValidationError {
	str, ok := value.(string)
	if !ok {
		return nil
	}
	
	alphanumRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !alphanumRegex.MatchString(str) {
		return &ValidationError{
			Field:   fieldName,
			Message: "must contain only alphanumeric characters",
			Value:   value,
		}
	}
	
	return nil
}

func validateNumeric(fieldName string, value interface{}) *ValidationError {
	str, ok := value.(string)
	if !ok {
		return nil
	}
	
	if _, err := strconv.ParseFloat(str, 64); err != nil {
		return &ValidationError{
			Field:   fieldName,
			Message: "must be a valid number",
			Value:   value,
		}
	}
	
	return nil
}

func validateDateTime(fieldName string, value interface{}, param string) *ValidationError {
	str, ok := value.(string)
	if !ok {
		return nil
	}
	
	layout := param
	if layout == "" {
		layout = time.RFC3339
	}
	
	if _, err := time.Parse(layout, str); err != nil {
		return &ValidationError{
			Field:   fieldName,
			Message: fmt.Sprintf("must be a valid datetime in format %s", layout),
			Value:   value,
		}
	}
	
	return nil
}

func isZeroValue(value interface{}) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Interface, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.IsNil()
	case reflect.Array:
		return v.Len() == 0
	default:
		return false
	}
}

func getFieldName(field reflect.StructField) string {
	jsonTag := field.Tag.Get("json")
	if jsonTag != "" && jsonTag != "-" {
		name := strings.Split(jsonTag, ",")[0]
		if name != "" {
			return name
		}
	}
	return field.Name
}
