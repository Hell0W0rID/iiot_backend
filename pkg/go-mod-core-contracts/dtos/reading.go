//
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/fxamacker/cbor/v2"
	"github.com/google/uuid"

	"iiot-backend/pkg/go-mod-core-contracts/common"
	iiotErrors "iiot-backend/pkg/go-mod-core-contracts/errors"
	"iiot-backend/pkg/go-mod-core-contracts/models"
)

type BaseMeasurement struct {
	Id            string `json:"id,omitempty"`
	Origin        int64  `json:"origin" validate:"required"`
	DeviceName    string `json:"deviceName" validate:"required,iiot-dto-none-empty-string"`
	ResourceName  string `json:"resourceName" validate:"required,iiot-dto-none-empty-string"`
	ProfileName   string `json:"profileName" validate:"required,iiot-dto-none-empty-string"`
	ValueType     string `json:"valueType" validate:"required,iiot-dto-value-type"`
	Units         string `json:"units,omitempty"`
	Tags          Tags   `json:"tags,omitempty"`
	BinaryMeasurement `json:",inline" validate:"-"`
	SimpleMeasurement `json:",inline" validate:"-"`
	ObjectMeasurement `json:",inline" validate:"-"`
	NullMeasurement   `json:",inline" validate:"-"`
}

type SimpleMeasurement struct {
	Value string `json:"value"`
}

type BinaryMeasurement struct {
	BinaryValue []byte `json:"binaryValue,omitempty" validate:"omitempty"`
	MediaType   string `json:"mediaType,omitempty" validate:"required_with=BinaryValue"`
}

type ObjectMeasurement struct {
	ObjectValue any `json:"objectValue,omitempty"`
}

type NullMeasurement struct {
	isNull bool // indicate the reading value should be null in the JSON payload
}

func (b BaseMeasurement) IsNull() bool {
	return b.isNull
}

func newBaseMeasurement(profileName string, deviceName string, resourceName string, valueType string) BaseMeasurement {
	return BaseMeasurement{
		Id:           uuid.NewString(),
		Origin:       time.Now().UnixNano(),
		DeviceName:   deviceName,
		ResourceName: resourceName,
		ProfileName:  profileName,
		ValueType:    valueType,
	}
}

// NewSimpleMeasurement creates and returns a new initialized BaseMeasurement with its SimpleMeasurement initialized
func NewSimpleMeasurement(profileName string, deviceName string, resourceName string, valueType string, value any) (BaseMeasurement, error) {
	stringValue, err := convertInterfaceValue(valueType, value)
	if err != nil {
		return BaseMeasurement{}, err
	}

	reading := newBaseMeasurement(profileName, deviceName, resourceName, valueType)
	reading.SimpleMeasurement = SimpleMeasurement{
		Value: stringValue,
	}
	return reading, nil
}

// NewBinaryMeasurement creates and returns a new initialized BaseMeasurement with its BinaryMeasurement initialized
func NewBinaryMeasurement(profileName string, deviceName string, resourceName string, binaryValue []byte, mediaType string) BaseMeasurement {
	reading := newBaseMeasurement(profileName, deviceName, resourceName, common.ValueTypeBinary)
	reading.BinaryMeasurement = BinaryMeasurement{
		BinaryValue: binaryValue,
		MediaType:   mediaType,
	}
	return reading
}

// NewObjectMeasurement creates and returns a new initialized BaseMeasurement with its ObjectMeasurement initialized
func NewObjectMeasurement(profileName string, deviceName string, resourceName string, objectValue any) BaseMeasurement {
	reading := newBaseMeasurement(profileName, deviceName, resourceName, common.ValueTypeObject)
	reading.ObjectMeasurement = ObjectMeasurement{
		ObjectValue: objectValue,
	}
	return reading
}

// NewObjectMeasurementWithArray creates and returns a new initialized BaseMeasurement with its ObjectMeasurement initialized with ObjectArray valueType
func NewObjectMeasurementWithArray(profileName string, deviceName string, resourceName string, objectValue any) BaseMeasurement {
	reading := newBaseMeasurement(profileName, deviceName, resourceName, common.ValueTypeObjectArray)
	reading.ObjectMeasurement = ObjectMeasurement{
		ObjectValue: objectValue,
	}
	return reading
}

// NewNullMeasurement creates and returns a new initialized BaseMeasurement with null
func NewNullMeasurement(profileName string, deviceName string, resourceName string, valueType string) BaseMeasurement {
	reading := newBaseMeasurement(profileName, deviceName, resourceName, valueType)
	reading.isNull = true
	return reading
}

func convertInterfaceValue(valueType string, value any) (string, error) {
	switch valueType {
	case common.ValueTypeBool:
		return convertSimpleValue(valueType, reflect.Bool, value)
	case common.ValueTypeString:
		return convertSimpleValue(valueType, reflect.String, value)

	case common.ValueTypeUint8:
		return convertSimpleValue(valueType, reflect.Uint8, value)
	case common.ValueTypeUint16:
		return convertSimpleValue(valueType, reflect.Uint16, value)
	case common.ValueTypeUint32:
		return convertSimpleValue(valueType, reflect.Uint32, value)
	case common.ValueTypeUint64:
		return convertSimpleValue(valueType, reflect.Uint64, value)

	case common.ValueTypeInt8:
		return convertSimpleValue(valueType, reflect.Int8, value)
	case common.ValueTypeInt16:
		return convertSimpleValue(valueType, reflect.Int16, value)
	case common.ValueTypeInt32:
		return convertSimpleValue(valueType, reflect.Int32, value)
	case common.ValueTypeInt64:
		return convertSimpleValue(valueType, reflect.Int64, value)

	case common.ValueTypeFloat32:
		return convertFloatValue(valueType, reflect.Float32, value)
	case common.ValueTypeFloat64:
		return convertFloatValue(valueType, reflect.Float64, value)

	case common.ValueTypeBoolArray:
		return convertSimpleArrayValue(valueType, reflect.Bool, value)
	case common.ValueTypeStringArray:
		return convertSimpleArrayValue(valueType, reflect.String, value)

	case common.ValueTypeUint8Array:
		return convertSimpleArrayValue(valueType, reflect.Uint8, value)
	case common.ValueTypeUint16Array:
		return convertSimpleArrayValue(valueType, reflect.Uint16, value)
	case common.ValueTypeUint32Array:
		return convertSimpleArrayValue(valueType, reflect.Uint32, value)
	case common.ValueTypeUint64Array:
		return convertSimpleArrayValue(valueType, reflect.Uint64, value)

	case common.ValueTypeInt8Array:
		return convertSimpleArrayValue(valueType, reflect.Int8, value)
	case common.ValueTypeInt16Array:
		return convertSimpleArrayValue(valueType, reflect.Int16, value)
	case common.ValueTypeInt32Array:
		return convertSimpleArrayValue(valueType, reflect.Int32, value)
	case common.ValueTypeInt64Array:
		return convertSimpleArrayValue(valueType, reflect.Int64, value)

	case common.ValueTypeFloat32Array:
		arrayValue, ok := value.([]float32)
		if !ok {
			return "", fmt.Errorf("unable to cast value to []float32 for %s", valueType)
		}

		return convertFloat32ArrayValue(arrayValue)
	case common.ValueTypeFloat64Array:
		arrayValue, ok := value.([]float64)
		if !ok {
			return "", fmt.Errorf("unable to cast value to []float64 for %s", valueType)
		}

		return convertFloat64ArrayValue(arrayValue)

	default:
		return "", fmt.Errorf("invalid simple reading type of %s", valueType)
	}
}

func convertSimpleValue(valueType string, kind reflect.Kind, value any) (string, error) {
	if err := validateType(valueType, kind, value); err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", value), nil
}

func convertFloatValue(valueType string, kind reflect.Kind, value any) (string, error) {
	if err := validateType(valueType, kind, value); err != nil {
		return "", err
	}

	switch kind {
	case reflect.Float32:
		// as above has validated the value type/kind/value, it is safe to cast the value to float32 here
		float32Val, ok := value.(float32)
		if !ok {
			return "", fmt.Errorf("unable to cast value to float32 for %s", valueType)
		}
		return strconv.FormatFloat(float64(float32Val), 'e', -1, 32), nil
	case reflect.Float64:
		// as above has validated the value type/kind/value, it is safe to cast the value to float64 here
		float64Val, ok := value.(float64)
		if !ok {
			return "", fmt.Errorf("unable to cast value to float64 for %s", valueType)
		}
		return strconv.FormatFloat(float64Val, 'e', -1, 64), nil
	default:
		return "", fmt.Errorf("invalid kind %s to convert float value to string", kind.String())
	}
}

func convertSimpleArrayValue(valueType string, kind reflect.Kind, value any) (string, error) {
	if err := validateType(valueType, kind, value); err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v", value)
	result = strings.ReplaceAll(result, " ", ", ")
	return result, nil
}

func convertFloat32ArrayValue(values []float32) (string, error) {
	var result strings.Builder
	result.WriteString("[")
	first := true
	for _, value := range values {
		if first {
			floatValue, err := convertFloatValue(common.ValueTypeFloat32, reflect.Float32, value)
			if err != nil {
				return "", err
			}
			result.WriteString(floatValue)
			first = false
			continue
		}

		floatValue, err := convertFloatValue(common.ValueTypeFloat32, reflect.Float32, value)
		if err != nil {
			return "", err
		}
		result.WriteString(", " + floatValue)
	}

	result.WriteString("]")
	return result.String(), nil
}

func convertFloat64ArrayValue(values []float64) (string, error) {
	var result strings.Builder
	result.WriteString("[")
	first := true
	for _, value := range values {
		if first {
			floatValue, err := convertFloatValue(common.ValueTypeFloat64, reflect.Float64, value)
			if err != nil {
				return "", err
			}
			result.WriteString(floatValue)
			first = false
			continue
		}

		floatValue, err := convertFloatValue(common.ValueTypeFloat64, reflect.Float64, value)
		if err != nil {
			return "", err
		}
		result.WriteString(", " + floatValue)
	}

	result.WriteString("]")
	return result.String(), nil
}

func validateType(valueType string, kind reflect.Kind, value any) error {
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		if kind != reflect.TypeOf(value).Elem().Kind() {
			return fmt.Errorf("slice of type of value `%s` not a match for specified ValueType '%s", kind.String(), valueType)
		}
		return nil
	}

	if kind != reflect.TypeOf(value).Kind() {
		return fmt.Errorf("type of value `%s` not a match for specified ValueType '%s", kind.String(), valueType)
	}

	return nil
}

// Validate satisfies the Validator interface
func (b BaseMeasurement) Validate() error {
	if b.isNull {
		return nil
	}
	if b.ValueType == common.ValueTypeBinary {
		// validate the inner BinaryMeasurement struct
		binaryMeasurement := b.BinaryMeasurement
		if err := common.Validate(binaryMeasurement); err != nil {
			return err
		}
	} else if b.ValueType == common.ValueTypeObject || b.ValueType == common.ValueTypeObjectArray {
		// validate the inner ObjectMeasurement struct
		objectMeasurement := b.ObjectMeasurement
		if err := common.Validate(objectMeasurement); err != nil {
			return err
		}
	} else {
		// validate the inner SimpleMeasurement struct
		simpleMeasurement := b.SimpleMeasurement
		if err := common.Validate(simpleMeasurement); err != nil {
			return err
		}
		if err := ValidateValue(b.ValueType, simpleMeasurement.Value); err != nil {
			return iiotErrors.NewCommonIIOT(iiotErrors.KindContractInvalid, fmt.Sprintf("The value does not match the %v valueType", b.ValueType), nil)
		}
	}

	return nil
}

// ToMeasurementModel converts Measurement DTO to Measurement Model
func ToMeasurementModel(r BaseMeasurement) models.Measurement {
	var readingModel models.Measurement
	br := models.BaseMeasurement{
		Id:           r.Id,
		Origin:       r.Origin,
		DeviceName:   r.DeviceName,
		ResourceName: r.ResourceName,
		ProfileName:  r.ProfileName,
		ValueType:    r.ValueType,
		Units:        r.Units,
		Tags:         r.Tags,
	}
	if r.NullMeasurement.isNull {
		return models.NewNullMeasurement(br)
	}
	if r.ValueType == common.ValueTypeBinary {
		readingModel = models.BinaryMeasurement{
			BaseMeasurement: br,
			BinaryValue: r.BinaryValue,
			MediaType:   r.MediaType,
		}
	} else if r.ValueType == common.ValueTypeObject || r.ValueType == common.ValueTypeObjectArray {
		readingModel = models.ObjectMeasurement{
			BaseMeasurement: br,
			ObjectValue: r.ObjectValue,
		}
	} else {
		readingModel = models.SimpleMeasurement{
			BaseMeasurement: br,
			Value:       r.Value,
		}
	}
	return readingModel
}

func FromMeasurementModelToDTO(reading models.Measurement) BaseMeasurement {
	var baseMeasurement BaseMeasurement
	switch r := reading.(type) {
	case models.BinaryMeasurement:
		baseMeasurement = BaseMeasurement{
			Id:            r.Id,
			Origin:        r.Origin,
			DeviceName:    r.DeviceName,
			ResourceName:  r.ResourceName,
			ProfileName:   r.ProfileName,
			ValueType:     r.ValueType,
			Units:         r.Units,
			Tags:          r.Tags,
			BinaryMeasurement: BinaryMeasurement{BinaryValue: r.BinaryValue, MediaType: r.MediaType},
		}
	case models.ObjectMeasurement:
		baseMeasurement = BaseMeasurement{
			Id:            r.Id,
			Origin:        r.Origin,
			DeviceName:    r.DeviceName,
			ResourceName:  r.ResourceName,
			ProfileName:   r.ProfileName,
			ValueType:     r.ValueType,
			Units:         r.Units,
			Tags:          r.Tags,
			ObjectMeasurement: ObjectMeasurement{ObjectValue: r.ObjectValue},
		}
	case models.SimpleMeasurement:
		baseMeasurement = BaseMeasurement{
			Id:            r.Id,
			Origin:        r.Origin,
			DeviceName:    r.DeviceName,
			ResourceName:  r.ResourceName,
			ProfileName:   r.ProfileName,
			ValueType:     r.ValueType,
			Units:         r.Units,
			Tags:          r.Tags,
			SimpleMeasurement: SimpleMeasurement{Value: r.Value},
		}
	case models.NullMeasurement:
		baseMeasurement = BaseMeasurement{
			Id:           r.Id,
			Origin:       r.Origin,
			DeviceName:   r.DeviceName,
			ResourceName: r.ResourceName,
			ProfileName:  r.ProfileName,
			ValueType:    r.ValueType,
			Units:        r.Units,
			Tags:         r.Tags,
			NullMeasurement:  NullMeasurement{isNull: true},
		}
	}

	return baseMeasurement
}

// ValidateValue used to check whether the value and valueType are matched
func ValidateValue(valueType string, value string) error {
	if strings.Contains(valueType, "Array") {
		return parseArrayValue(valueType, value)
	} else {
		return parseSimpleValue(valueType, value)
	}
}

func parseSimpleValue(valueType string, value string) (err error) {
	switch valueType {
	case common.ValueTypeBool:
		_, err = strconv.ParseBool(value)

	case common.ValueTypeUint8:
		_, err = strconv.ParseUint(value, 10, 8)
	case common.ValueTypeUint16:
		_, err = strconv.ParseUint(value, 10, 16)
	case common.ValueTypeUint32:
		_, err = strconv.ParseUint(value, 10, 32)
	case common.ValueTypeUint64:
		_, err = strconv.ParseUint(value, 10, 64)

	case common.ValueTypeInt8:
		_, err = strconv.ParseInt(value, 10, 8)
	case common.ValueTypeInt16:
		_, err = strconv.ParseInt(value, 10, 16)
	case common.ValueTypeInt32:
		_, err = strconv.ParseInt(value, 10, 32)
	case common.ValueTypeInt64:
		_, err = strconv.ParseInt(value, 10, 64)

	case common.ValueTypeFloat32:
		_, err = strconv.ParseFloat(value, 32)
	case common.ValueTypeFloat64:
		_, err = strconv.ParseFloat(value, 64)
	}

	if err != nil {
		return err
	}
	return nil
}

func parseArrayValue(valueType string, value string) (err error) {
	arrayValue := strings.Split(value[1:len(value)-1], ",") // trim "[" and "]"

	for _, v := range arrayValue {
		v = strings.TrimSpace(v)
		switch valueType {
		case common.ValueTypeBoolArray:
			err = parseSimpleValue(common.ValueTypeBool, v)

		case common.ValueTypeUint8Array:
			err = parseSimpleValue(common.ValueTypeUint8, v)
		case common.ValueTypeUint16Array:
			err = parseSimpleValue(common.ValueTypeUint16, v)
		case common.ValueTypeUint32Array:
			err = parseSimpleValue(common.ValueTypeUint32, v)
		case common.ValueTypeUint64Array:
			err = parseSimpleValue(common.ValueTypeUint64, v)

		case common.ValueTypeInt8Array:
			err = parseSimpleValue(common.ValueTypeInt8, v)
		case common.ValueTypeInt16Array:
			err = parseSimpleValue(common.ValueTypeInt16, v)
		case common.ValueTypeInt32Array:
			err = parseSimpleValue(common.ValueTypeInt32, v)
		case common.ValueTypeInt64Array:
			err = parseSimpleValue(common.ValueTypeInt64, v)

		case common.ValueTypeFloat32Array:
			err = parseSimpleValue(common.ValueTypeFloat32, v)
		case common.ValueTypeFloat64Array:
			err = parseSimpleValue(common.ValueTypeFloat64, v)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// UnmarshalObjectValue is a helper function used to unmarshal the ObjectValue of a reading to the passed in target type.
// Note that this function will only work on readings with 'Object' or 'ObjectArray' valueType.  An error will be returned when invoking
// this function on a reading with valueType other than 'Object' or 'ObjectArray'.
func (b BaseMeasurement) UnmarshalObjectValue(target any) error {
	if b.ValueType == common.ValueTypeObject || b.ValueType == common.ValueTypeObjectArray {
		// marshal the current reading ObjectValue to JSON
		jsonEncodedData, err := json.Marshal(b.ObjectValue)
		if err != nil {
			return iiotErrors.NewCommonIIOT(iiotErrors.KindContractInvalid, "failed to encode the object value of reading to JSON", err)
		}
		// unmarshal the JSON into the passed in target
		err = json.Unmarshal(jsonEncodedData, target)
		if err != nil {
			return iiotErrors.NewCommonIIOT(iiotErrors.KindContractInvalid, fmt.Sprintf("failed to unmarshall the object value of reading into type %v", reflect.TypeOf(target).String()), err)
		}
	} else {
		return iiotErrors.NewCommonIIOT(iiotErrors.KindContractInvalid, fmt.Sprintf("invalid usage of UnmarshalObjectValue function invocation on reading with %v valueType", b.ValueType), nil)
	}

	return nil
}

func (b BaseMeasurement) MarshalJSON() ([]byte, error) {
	return b.marshal(json.Marshal)
}

func (b BaseMeasurement) MarshalCBOR() ([]byte, error) {
	return b.marshal(cbor.Marshal)
}

func (b BaseMeasurement) marshal(marshal func(any) ([]byte, error)) ([]byte, error) {
	type reading struct {
		Id           string `json:"id,omitempty"`
		Origin       int64  `json:"origin"`
		DeviceName   string `json:"deviceName"`
		ResourceName string `json:"resourceName"`
		ProfileName  string `json:"profileName"`
		ValueType    string `json:"valueType"`
		Units        string `json:"units,omitempty"`
		Tags         Tags   `json:"tags,omitempty"`
	}
	if b.isNull {
		return marshal(&struct {
			reading     `json:",inline"`
			Value       any `json:"value"`
			BinaryValue any `json:"binaryValue"`
			ObjectValue any `json:"objectValue"`
		}{
			reading: reading{
				Id:           b.Id,
				Origin:       b.Origin,
				DeviceName:   b.DeviceName,
				ResourceName: b.ResourceName,
				ProfileName:  b.ProfileName,
				ValueType:    b.ValueType,
				Units:        b.Units,
				Tags:         b.Tags,
			},
			Value:       nil,
			BinaryValue: nil,
			ObjectValue: nil,
		})
	}
	r := reading{
		Id:           b.Id,
		Origin:       b.Origin,
		DeviceName:   b.DeviceName,
		ResourceName: b.ResourceName,
		ProfileName:  b.ProfileName,
		ValueType:    b.ValueType,
		Units:        b.Units,
		Tags:         b.Tags,
	}
	switch b.ValueType {
	case common.ValueTypeObject, common.ValueTypeObjectArray:
		return marshal(&struct {
			reading       `json:",inline"`
			ObjectMeasurement `json:",inline" validate:"-"`
		}{
			reading:       r,
			ObjectMeasurement: b.ObjectMeasurement,
		})
	case common.ValueTypeBinary:
		return marshal(&struct {
			reading       `json:",inline"`
			BinaryMeasurement `json:",inline" validate:"-"`
		}{
			reading:       r,
			BinaryMeasurement: b.BinaryMeasurement,
		})
	default:
		return marshal(&struct {
			reading       `json:",inline"`
			SimpleMeasurement `json:",inline" validate:"-"`
		}{
			reading:       r,
			SimpleMeasurement: b.SimpleMeasurement,
		})
	}
}

func (b *BaseMeasurement) UnmarshalJSON(data []byte) error {
	return b.Unmarshal(data, json.Unmarshal)
}

func (b *BaseMeasurement) UnmarshalCBOR(data []byte) error {
	return b.Unmarshal(data, cbor.Unmarshal)
}

func (b *BaseMeasurement) Unmarshal(data []byte, unmarshal func([]byte, any) error) error {
	var aux struct {
		Id           string
		Origin       int64
		DeviceName   string
		ResourceName string
		ProfileName  string
		ValueType    string
		Units        string
		Tags         Tags
		Value        any
		BinaryMeasurement
		ObjectMeasurement
	}
	if err := unmarshal(data, &aux); err != nil {
		return err
	}

	b.Id = aux.Id
	b.Origin = aux.Origin
	b.DeviceName = aux.DeviceName
	b.ResourceName = aux.ResourceName
	b.ProfileName = aux.ProfileName
	b.ValueType = aux.ValueType
	b.Units = aux.Units
	b.Tags = aux.Tags
	b.BinaryMeasurement = aux.BinaryMeasurement
	if aux.Value != nil {
		b.SimpleMeasurement = SimpleMeasurement{Value: fmt.Sprintf("%s", aux.Value)}
	}
	b.ObjectMeasurement = aux.ObjectMeasurement

	switch aux.ValueType {
	case common.ValueTypeObject, common.ValueTypeObjectArray:
		if aux.ObjectValue == nil {
			b.isNull = true
		}
	case common.ValueTypeBinary:
		if aux.BinaryValue == nil {
			b.isNull = true
		}
	default:
		if aux.Value == nil {
			b.isNull = true
		}
	}
	return nil
}
