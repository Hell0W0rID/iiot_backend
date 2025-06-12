# EdgeX Go to IIoT Backend Code Mapping

This document tracks the systematic transformation from EdgeX Go code to our IIoT Backend implementation using pkg/ shared libraries.

## Import Path Mappings

### Bootstrap and Core Libraries
| EdgeX Go Import | IIoT Backend Import | Status |
|-----------------|-------------------|--------|
| `github.com/edgexfoundry/go-mod-bootstrap/v4/bootstrap` | `iiot-backend/pkg/go-mod-bootstrap/bootstrap` | ✓ Mapped |
| `github.com/edgexfoundry/go-mod-bootstrap/v4/bootstrap/container` | `iiot-backend/pkg/go-mod-bootstrap/bootstrap/container` | ✓ Mapped |
| `github.com/edgexfoundry/go-mod-bootstrap/v4/di` | `iiot-backend/pkg/go-mod-bootstrap/di` | ✓ Mapped |
| `github.com/edgexfoundry/go-mod-core-contracts/v4/common` | `iiot-backend/pkg/go-mod-core-contracts/common` | ✓ Mapped |
| `github.com/edgexfoundry/go-mod-core-contracts/v4/dtos` | `iiot-backend/pkg/go-mod-core-contracts/dtos` | ✓ Mapped |
| `github.com/edgexfoundry/go-mod-core-contracts/v4/errors` | `iiot-backend/pkg/go-mod-core-contracts/errors` | ✓ Mapped |

### Service-Specific Imports
| EdgeX Go Import | IIoT Backend Import | Status |
|-----------------|-------------------|--------|
| `github.com/edgexfoundry/edgex-go/internal/core/command/container` | `iiot-backend/services/core/command/container` | ✓ Mapped |
| `github.com/edgexfoundry/edgex-go/internal/core/command/config` | `iiot-backend/services/core/command/config` | ✓ Mapped |

## Client Name Mappings

### Bootstrap Container Client Functions
| EdgeX Go Function | IIoT Backend Function | Status |
|------------------|-------------------|--------|
| `bootstrapContainer.DeviceClientFrom()` | `bootstrapContainer.DeviceClientFrom()` | ✓ Same |
| `bootstrapContainer.DeviceProfileClientFrom()` | `bootstrapContainer.DeviceTemplateClientFrom()` | ✓ Mapped |
| `bootstrapContainer.DeviceServiceClientFrom()` | `bootstrapContainer.DeviceHandlerClientFrom()` | ✓ Mapped |
| `bootstrapContainer.DeviceServiceCommandClientFrom()` | `bootstrapContainer.DeviceHandlerCommandClientFrom()` | ✓ Mapped |

### Client Constants
| EdgeX Go Constant | IIoT Backend Constant | Status |
|------------------|-------------------|--------|
| `DeviceServiceClientName` | `DeviceHandlerClientName` | ✓ Mapped |
| `DeviceProfileClientName` | `DeviceTemplateClientName` | ✓ Mapped |
| `DeviceServiceCommandClientName` | `DeviceHandlerCommandClientName` | ✓ Mapped |
| `CoreCommandServiceKey` | `CoreCommandServiceName` | ✓ Mapped |

## Method Name Mappings

### Client Method Names
| EdgeX Go Method | IIoT Backend Method | Status |
|----------------|-------------------|--------|
| `dpc.DeviceProfileByName()` | `dpc.DeviceTemplateByName()` | ✓ Mapped |
| `dsc.DeviceServiceByName()` | `dsc.DeviceHandlerByName()` | ✓ Mapped |
| `dc.AllDevices()` | `dc.AllDevices()` | ✓ Same |
| `dc.DeviceByName()` | `dc.DeviceByName()` | ✓ Same |

## DTO and Response Types

### Data Transfer Objects
| EdgeX Go DTO | IIoT Backend DTO | Status |
|-------------|----------------|--------|
| `dtos.DeviceCoreCommand` | `dtos.DeviceCoreCommand` | ✓ Same |
| `dtos.CoreCommand` | `dtos.CoreCommand` | ✓ Same |
| `dtos.DeviceProfile` | `dtos.DeviceTemplate` | 🔄 To Verify |
| `responses.EventResponse` | `responses.EventResponse` | ✓ Same |
| `commonDTO.BaseResponse` | `commonDTO.BaseResponse` | ✓ Same |

## Service and Configuration Mappings

### Service Names
| EdgeX Go Service | IIoT Backend Service | Status |
|-----------------|-------------------|--------|
| `common.CoreCommandServiceKey` | `common.CoreCommandServiceName` | ✓ Mapped |

### Configuration Access
| EdgeX Go Pattern | IIoT Backend Pattern | Status |
|-----------------|-------------------|--------|
| `commandContainer.ConfigurationFrom(dic.Get)` | `commandContainer.ConfigurationFrom(dic.Get)` | ✓ Same |
| `configuration.Service.Url()` | `configuration.Service.Url()` | ✓ Same |

## Notes
- ✓ Mapped: Successfully transformed and verified
- 🔄 To Verify: Transformation applied, needs runtime verification
- ❌ Issues: Requires further investigation

## Error Interface Transformations

### EdgeX to IIOT Error Types
| EdgeX Go Error Type | IIoT Backend Error Type | Status |
|---------------------|------------------------|--------|
| `errors.EdgeX` | `errors.IIOT` | ✓ Transformed |
| `errors.NewCommonEdgeX()` | `errors.NewCommonIIOT()` | ✓ Transformed |
| `errors.NewCommonEdgeXWrapper()` | `errors.NewCommonIIOTWrapper()` | ✓ Transformed |

### Function Signature Updates
All command service application layer functions now use `errors.IIOT` interface:
- `AllCommands() (deviceCoreCommands []dtos.DeviceCoreCommand, totalCount uint32, err errors.IIOT)`
- `CommandsByDeviceName() (deviceCoreCommand dtos.DeviceCoreCommand, err errors.IIOT)`
- `IssueGetCommandByName() (res *responses.EventResponse, err errors.IIOT)`
- `IssueSetCommandByName() (response commonDTO.BaseResponse, err errors.IIOT)`
- `coreCommandParameters() ([]dtos.CoreCommandParameter, errors.IIOT)`
- `buildCoreCommands() ([]dtos.CoreCommand, errors.IIOT)`

### Error Constructor Replacements
All error constructors systematically replaced:
- 25+ instances of `errors.NewCommonEdgeX` → `errors.NewCommonIIOT`
- 15+ instances of `errors.NewCommonEdgeXWrapper` → `errors.NewCommonIIOTWrapper`

## Transformation Completion Status
- ✅ **Complete EdgeX Go Architecture**: Exact directory structure and patterns implemented
- ✅ **Client Interface Transformation**: All EdgeX client names mapped to IIOT equivalents
- ✅ **HTTP Controller Integration**: Controllers call actual application layer business logic
- ✅ **Business Logic Implementation**: Full command operations with proper error handling
- ✅ **Error Type Standardization**: All EdgeX error types replaced with IIOT equivalents
- ✅ **Keyword Transformation**: Systematic replacement of EdgeX terminology throughout codebase

## Last Updated
Updated during core command service restructuring - January 2025