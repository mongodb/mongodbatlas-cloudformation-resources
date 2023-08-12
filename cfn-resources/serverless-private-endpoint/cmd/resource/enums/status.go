package enums

type Status string

const (
	ReservationRequested Status = "RESERVATION_REQUESTED"
	Reserved             Status = "RESERVED"
	Initiating           Status = "INITIATING"
	Available            Status = "AVAILABLE"
	Failed               Status = "FAILED"
	Deleting             Status = "DELETING"
)
