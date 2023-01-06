package common

type key int

const (
	AppName                            = "container-reference-app"
	TenantIDHeaderName                 = "X-Coreos-Tid"
	RequestIDHeaderName                = "X-Coreos-Request-Id"
	UserInfoHeaderName                 = "X-Coreos-Userinfo"
	AcessHeaderName                    = "X-Coreos-Access"
	BasePath                           = "/api/v1"
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
