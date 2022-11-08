package constants

type Event string

const (
	Create Event = "Create"
	Read   Event = "Read"
	Update Event = "Update"
	Delete Event = "Delete"
	List   Event = "List"
)
