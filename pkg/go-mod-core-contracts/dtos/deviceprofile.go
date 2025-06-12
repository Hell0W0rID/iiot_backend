//
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"errors"
	"fmt"
	"strings"

	"iiot-backend/pkg/go-mod-core-contracts/common"
	iiotErrors "iiot-backend/pkg/go-mod-core-contracts/errors"
	"iiot-backend/pkg/go-mod-core-contracts/models"
)

type DeviceTemplate struct {
	DeviceTemplateBasicInfo `json:",inline" yaml:",inline"`
	DeviceResources        []DeviceResource `json:"deviceResources" yaml:"deviceResources" validate:"dive"`
	DeviceActions         []DeviceAction  `json:"deviceCommands" yaml:"deviceCommands" validate:"dive"`
	ApiVersion             string           `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
}

// Validate satisfies the Validator interface
func (dp *DeviceTemplate) Validate() error {
	err := common.Validate(dp)
	if err != nil {
		// The DeviceTemplateBasicInfo is the internal struct in Golang programming, not in the Profile model,
		// so it should be hidden from the error messages.
		err = errors.New(strings.ReplaceAll(err.Error(), ".DeviceTemplateBasicInfo", ""))
		return iiotErrors.NewCommonIIOT(iiotErrors.KindContractInvalid, "Invalid DeviceTemplate.", err)
	}
	return ValidateDeviceTemplateDTO(*dp)
}

// UnmarshalYAML implements the Unmarshaler interface for the DeviceTemplate type
func (dp *DeviceTemplate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var alias struct {
		DeviceTemplateBasicInfo `yaml:",inline"`
		DeviceResources        []DeviceResource `yaml:"deviceResources"`
		DeviceActions         []DeviceAction  `yaml:"deviceCommands"`
		ApiVersion             string           `yaml:"apiVersion"`
	}
	if err := unmarshal(&alias); err != nil {
		return iiotErrors.NewCommonIIOT(iiotErrors.KindContractInvalid, "failed to unmarshal request body as YAML.", err)
	}
	*dp = DeviceTemplate(alias)

	if err := dp.Validate(); err != nil {
		return iiotErrors.NewCommonIIOTWrapper(err)
	}

	// Normalize resource's value type
	for i, resource := range dp.DeviceResources {
		valueType, err := common.NormalizeValueType(resource.Properties.ValueType)
		if err != nil {
			return iiotErrors.NewCommonIIOTWrapper(err)
		}
		dp.DeviceResources[i].Properties.ValueType = valueType
	}
	return nil
}

// ToDeviceTemplateModel transforms the DeviceTemplate DTO to the DeviceTemplate model
func ToDeviceTemplateModel(deviceProfileDTO DeviceTemplate) models.DeviceTemplate {
	return models.DeviceTemplate{
		DBTimestamp:     models.DBTimestamp(deviceProfileDTO.DBTimestamp),
		Id:              deviceProfileDTO.Id,
		Name:            deviceProfileDTO.Name,
		Description:     deviceProfileDTO.Description,
		Manufacturer:    deviceProfileDTO.Manufacturer,
		Model:           deviceProfileDTO.Model,
		Labels:          deviceProfileDTO.Labels,
		DeviceResources: ToDeviceResourceModels(deviceProfileDTO.DeviceResources),
		DeviceActions:  ToDeviceActionModels(deviceProfileDTO.DeviceActions),
		ApiVersion:      deviceProfileDTO.ApiVersion,
	}
}

// FromDeviceTemplateModelToDTO transforms the DeviceTemplate Model to the DeviceTemplate DTO
func FromDeviceTemplateModelToDTO(deviceProfile models.DeviceTemplate) DeviceTemplate {
	if deviceProfile.ApiVersion == "" {
		deviceProfile.ApiVersion = common.ApiVersion
	}
	return DeviceTemplate{
		DeviceTemplateBasicInfo: DeviceTemplateBasicInfo{
			DBTimestamp:  DBTimestamp(deviceProfile.DBTimestamp),
			Id:           deviceProfile.Id,
			Name:         deviceProfile.Name,
			Description:  deviceProfile.Description,
			Manufacturer: deviceProfile.Manufacturer,
			Model:        deviceProfile.Model,
			Labels:       deviceProfile.Labels,
		},
		DeviceResources: FromDeviceResourceModelsToDTOs(deviceProfile.DeviceResources),
		DeviceActions:  FromDeviceActionModelsToDTOs(deviceProfile.DeviceActions),
		ApiVersion:      deviceProfile.ApiVersion,
	}
}

// FromDeviceTemplateModelToBasicInfoDTO transforms the DeviceTemplate Model to the DeviceTemplateBasicInfo DTO
func FromDeviceTemplateModelToBasicInfoDTO(deviceProfile models.DeviceTemplate) DeviceTemplateBasicInfo {
	return DeviceTemplateBasicInfo{
		DBTimestamp:  DBTimestamp(deviceProfile.DBTimestamp),
		Id:           deviceProfile.Id,
		Name:         deviceProfile.Name,
		Description:  deviceProfile.Description,
		Manufacturer: deviceProfile.Manufacturer,
		Model:        deviceProfile.Model,
		Labels:       deviceProfile.Labels,
	}
}

func ValidateDeviceTemplateDTO(profile DeviceTemplate) error {
	// deviceResources validation
	dupCheck := make(map[string]bool)
	for _, resource := range profile.DeviceResources {
		if resource.Properties.ValueType == common.ValueTypeBinary &&
			strings.Contains(resource.Properties.ReadWrite, common.ReadWrite_W) {
			return iiotErrors.NewCommonIIOT(iiotErrors.KindContractInvalid, fmt.Sprintf("write permission not support %s value type for resource '%s'", common.ValueTypeBinary, resource.Name), nil)
		}
		// deviceResource name should not duplicated
		if dupCheck[resource.Name] {
			return iiotErrors.NewCommonIIOT(iiotErrors.KindContractInvalid, fmt.Sprintf("device resource %s is duplicated", resource.Name), nil)
		}
		dupCheck[resource.Name] = true
	}
	// deviceCommands validation
	dupCheck = make(map[string]bool)
	for _, command := range profile.DeviceActions {
		// deviceCommand name should not duplicated
		if dupCheck[command.Name] {
			return iiotErrors.NewCommonIIOT(iiotErrors.KindContractInvalid, fmt.Sprintf("device command %s is duplicated", command.Name), nil)
		}
		dupCheck[command.Name] = true

		resourceOperations := command.ResourceOperations
		for _, ro := range resourceOperations {
			// ResourceOperations referenced in deviceCommands must exist
			if !deviceResourcesContains(profile.DeviceResources, ro.DeviceResource) {
				return iiotErrors.NewCommonIIOT(iiotErrors.KindContractInvalid, fmt.Sprintf("device command's resource %s doesn't match any device resource", ro.DeviceResource), nil)
			}
			// Check the ReadWrite whether is align to the deviceResource
			if !validReadWritePermission(profile.DeviceResources, ro.DeviceResource, command.ReadWrite) {
				return iiotErrors.NewCommonIIOT(iiotErrors.KindContractInvalid, fmt.Sprintf("device command's ReadWrite permission '%s' doesn't align the device resource", command.ReadWrite), nil)
			}
		}
	}
	return nil
}

func deviceResourcesContains(resources []DeviceResource, name string) bool {
	contains := false
	for _, resource := range resources {
		if resource.Name == name {
			contains = true
			break
		}
	}
	return contains
}

func validReadWritePermission(resources []DeviceResource, name string, readWrite string) bool {
	valid := true
	for _, resource := range resources {
		if resource.Name == name {
			if resource.Properties.ReadWrite != common.ReadWrite_RW && resource.Properties.ReadWrite != common.ReadWrite_WR &&
				resource.Properties.ReadWrite != readWrite {
				valid = false
				break
			}
		}
	}
	return valid
}
