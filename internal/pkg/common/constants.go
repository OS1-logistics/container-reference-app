package common

type key int

const (
	AppName             = "container-reference-app"
	TenantIDHeaderName  = "X-Coreos-Tid"
	RequestIDHeaderName = "X-Coreos-Request-Id"
	UserInfoHeaderName  = "X-Coreos-Userinfo"
	AcessHeaderName     = "X-Coreos-Access"
	BasePath            = "/"
)

const UserStruct = "user"
const StateMachineFilePath = "domain/init/state-machnine.json"
const PackageContainerTypeName = "PackageCsa"
const BagContainerTypeName = "BagT"
