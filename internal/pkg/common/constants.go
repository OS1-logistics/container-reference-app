package common

type key int

const (
	TenantIDHeaderName                 = "X-COREOS-TENANT-ID"
	RequestIDHeaderName                = "X-COREOS-REQUEST-ID"
	AcessHeaderName                    = "X-COREOS-ACCESS"
	UserInfoHeaderName                 = "X-COREOS-USERINFO"
	ApiKeyHeaderName                   = "X-API-KEY"
	BasePath                           = "/api/v1"
	OS1AuthBasePath                    = "/core/api/v1/ext-auth/token/validate"
	PackageContainerTypeName           = "PackageCsa"
	BagContainerTypeName               = "BagT"
	CONTAINER_OPERATION_CONTAINERIZE   = "CONTAINERIZE"
	CONTAINER_OPERATION_DECONTAINERIZE = "DECONTAINERIZE"
)

type ContainerStateOperation struct {
	Name       string
	EventCode  string
	ReasonCode string
}

var (
	STATE_OPERATION_OPEN = ContainerStateOperation{
		Name:       "open",
		EventCode:  "E-025",
		ReasonCode: "R-0001",
	}
	STATE_OPERATION_CLOSE = ContainerStateOperation{
		Name:       "close",
		EventCode:  "E-026",
		ReasonCode: "R-0001",
	}
	STATE_OPERATION_COMPLETE = ContainerStateOperation{
		Name:       "complete",
		EventCode:  "E-027",
		ReasonCode: "R-0001",
	}
	STATE_OPERATION_DEAD = ContainerStateOperation{
		Name:       "dead",
		EventCode:  "E-028",
		ReasonCode: "R-0001",
	}
)
