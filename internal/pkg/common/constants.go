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

type Operation struct {
	Name       string
	EventCode  string
	ReasonCode string
}

var (
	OPERATION_OPEN = Operation{
		Name:       "open",
		EventCode:  "E-025",
		ReasonCode: "R-0001",
	}
	OPERATION_CLOSE = Operation{
		Name:       "close",
		EventCode:  "E-026",
		ReasonCode: "R-0001",
	}
	OPERATION_COMPLETE = Operation{
		Name:       "complete",
		EventCode:  "E-027",
		ReasonCode: "R-0001",
	}
	OPERATION_DEAD = Operation{
		Name:       "dead",
		EventCode:  "E-028",
		ReasonCode: "R-0001",
	}
)

const UserStruct = "user"
const StateMachineFilePath = "domain/init/state-machnine.json"
const PackageContainerTypeName = "PackageCsa"
const BagContainerTypeName = "BagT"
