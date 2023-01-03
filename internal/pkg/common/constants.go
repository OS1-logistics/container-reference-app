package constants

type key int

const (
	AppName             = "container-reference-app-api"
	TenantIDHeaderName  = "X-Coreos-Tid"
	RequestIDHeaderName = "X-Coreos-Request-Id"
	UserInfoHeaderName  = "X-Coreos-Userinfo"
	AcessHeaderName     = "X-Coreos-Access"
	BasePath            = "/"
)

const UserStruct = "user"

var ShipmentsEntityCode = "0003"

const StateMachineFilePath = "domain/init/state-machnine.json"
const ShipmentCreationEventCode = "E-001"
const ShipmentCreationReasonCode = "R-0044"
const ClientContainerIdKey = "clientContainerId"
