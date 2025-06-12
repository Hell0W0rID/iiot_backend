//
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// DeviceTemplateClient defines the interface for interactions with the DeviceTemplate endpoint on the IIOT Foundry iiot-metadata service.
type DeviceTemplateClient interface {
	// Add adds new profiles
	Add(ctx context.Context, reqs []requests.DeviceTemplateRequest) ([]common.BaseWithIdResponse, errors.IIOT)
	// Update updates profiles
	Update(ctx context.Context, reqs []requests.DeviceTemplateRequest) ([]common.BaseResponse, errors.IIOT)
	// AddByYaml adds new profile by uploading a file in YAML format
	AddByYaml(ctx context.Context, yamlFilePath string) (common.BaseWithIdResponse, errors.IIOT)
	// UpdateByYaml updates profile by uploading a file in YAML format
	UpdateByYaml(ctx context.Context, yamlFilePath string) (common.BaseResponse, errors.IIOT)
	// DeleteByName deletes profile by name
	DeleteByName(ctx context.Context, name string) (common.BaseResponse, errors.IIOT)
	// DeviceTemplateByName queries profile by name
	DeviceTemplateByName(ctx context.Context, name string) (responses.DeviceTemplateResponse, errors.IIOT)
	// AllDeviceTemplates queries all profiles
	AllDeviceTemplates(ctx context.Context, labels []string, offset int, limit int) (responses.MultiDeviceTemplatesResponse, errors.IIOT)
	// AllDeviceTemplateBasicInfos queries all profile basic infos
	AllDeviceTemplateBasicInfos(ctx context.Context, labels []string, offset int, limit int) (responses.MultiDeviceTemplateBasicInfoResponse, errors.IIOT)
	// DeviceTemplatesByModel queries profiles by model
	DeviceTemplatesByModel(ctx context.Context, model string, offset int, limit int) (responses.MultiDeviceTemplatesResponse, errors.IIOT)
	// DeviceTemplatesByManufacturer queries profiles by manufacturer
	DeviceTemplatesByManufacturer(ctx context.Context, manufacturer string, offset int, limit int) (responses.MultiDeviceTemplatesResponse, errors.IIOT)
	// DeviceTemplatesByManufacturerAndModel queries profiles by manufacturer and model
	DeviceTemplatesByManufacturerAndModel(ctx context.Context, manufacturer string, model string, offset int, limit int) (responses.MultiDeviceTemplatesResponse, errors.IIOT)
	// DeviceResourceByProfileNameAndResourceName queries the device resource by profileName and resourceName
	DeviceResourceByProfileNameAndResourceName(ctx context.Context, profileName string, resourceName string) (responses.DeviceResourceResponse, errors.IIOT)
	// UpdateDeviceTemplateBasicInfo updates existing profile's basic info
	UpdateDeviceTemplateBasicInfo(ctx context.Context, reqs []requests.DeviceTemplateBasicInfoRequest) ([]common.BaseResponse, errors.IIOT)
	// AddDeviceTemplateResource adds new device resource to an existing profile
	AddDeviceTemplateResource(ctx context.Context, reqs []requests.AddDeviceResourceRequest) ([]common.BaseResponse, errors.IIOT)
	// UpdateDeviceTemplateResource updates existing device resource
	UpdateDeviceTemplateResource(ctx context.Context, reqs []requests.UpdateDeviceResourceRequest) ([]common.BaseResponse, errors.IIOT)
	// DeleteDeviceResourceByName deletes device resource by name
	DeleteDeviceResourceByName(ctx context.Context, profileName string, resourceName string) (common.BaseResponse, errors.IIOT)
	// AddDeviceTemplateDeviceAction adds new device command to an existing profile
	AddDeviceTemplateDeviceAction(ctx context.Context, reqs []requests.AddDeviceActionRequest) ([]common.BaseResponse, errors.IIOT)
	// UpdateDeviceTemplateDeviceAction updates existing device command
	UpdateDeviceTemplateDeviceAction(ctx context.Context, reqs []requests.UpdateDeviceActionRequest) ([]common.BaseResponse, errors.IIOT)
	// DeleteDeviceActionByName deletes device command by name
	DeleteDeviceActionByName(ctx context.Context, profileName string, commandName string) (common.BaseResponse, errors.IIOT)
}
