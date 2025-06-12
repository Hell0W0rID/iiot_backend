//
//
// SPDX-License-Identifier: Apache-2.0

package models

type BaseMeasurement struct {
	Id           string
	Origin       int64
	DeviceName   string
	ResourceName string
	ProfileName  string
	ValueType    string
	Units        string
	Tags         map[string]any
}

type BinaryMeasurement struct {
	BaseMeasurement `json:",inline"`
	BinaryValue []byte
	MediaType   string
}

type SimpleMeasurement struct {
	BaseMeasurement `json:",inline"`
	Value       string
}

type NullMeasurement struct {
	BaseMeasurement `json:",inline"`
	Value       any
}

func NewNullMeasurement(b BaseMeasurement) NullMeasurement {
	return NullMeasurement{
		BaseMeasurement: b,
		Value:       nil,
	}
}

type ObjectMeasurement struct {
	BaseMeasurement `json:",inline"`
	ObjectValue any
}

// Measurement is an abstract interface to be implemented by BinaryMeasurement/SimpleMeasurement
type Measurement interface {
	GetBaseMeasurement() BaseMeasurement
}

// Implement GetBaseMeasurement() method in order for BinaryMeasurement and SimpleMeasurement, ObjectMeasurement structs to implement the
// abstract Measurement interface and then be used as a Measurement.
// Also, the Measurement interface can access the BaseMeasurement fields.
// This is Golang's way to implement inheritance.
func (b BinaryMeasurement) GetBaseMeasurement() BaseMeasurement { return b.BaseMeasurement }
func (s SimpleMeasurement) GetBaseMeasurement() BaseMeasurement { return s.BaseMeasurement }
func (o ObjectMeasurement) GetBaseMeasurement() BaseMeasurement { return o.BaseMeasurement }
func (n NullMeasurement) GetBaseMeasurement() BaseMeasurement   { return n.BaseMeasurement }
