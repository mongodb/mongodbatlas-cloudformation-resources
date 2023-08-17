package enums

type AtlasPrivateEndpointStatus string

const (
	ReservationRequested AtlasPrivateEndpointStatus = "RESERVATION_REQUESTED"
	Reserved             AtlasPrivateEndpointStatus = "RESERVED"
	Initiating           AtlasPrivateEndpointStatus = "INITIATING"
	Available            AtlasPrivateEndpointStatus = "AVAILABLE"
	Failed               AtlasPrivateEndpointStatus = "FAILED"
	Deleting             AtlasPrivateEndpointStatus = "DELETING"
)
