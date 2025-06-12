package interfaces

import (
	"context"
	"iiot-backend/pkg/go-mod-core-contracts/dtos"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// CommandClient defines the interface for device command operations
type CommandClient interface {
	AllDeviceCoreCommands(ctx context.Context, offset int, limit int) (dtos.MultiDeviceCoreCommandsResponse, errors.IIOTError)
	DeviceCoreCommandsByDeviceName(ctx context.Context, deviceName string) (dtos.DeviceResponse, errors.IIOTError)
	IssueGetCommandByName(ctx context.Context, deviceName string, commandName string, queryParams string) (*dtos.DataEventResponse, errors.IIOTError)
	IssueSetCommandByName(ctx context.Context, deviceName string, commandName string, queryParams string, settings map[string]interface{}) (dtos.BaseResponse, errors.IIOTError)
	IssueGetCommandByNameWithObject(ctx context.Context, deviceName string, commandName string, queryParams string) (*dtos.DataEventResponse, errors.IIOTError)
	IssueSetCommandByNameWithObject(ctx context.Context, deviceName string, commandName string, queryParams string, settings interface{}) (dtos.BaseResponse, errors.IIOTError)
}