package constants

const (
	BindPassword          = "BindPassword"
	AuthenticationEnabled = "AuthenticationEnabled"
	BindUsername          = "BindUsername"
	Hostname              = "Hostname"
	Port                  = "Port"
	UserToDNMapping       = "UserToDNMapping"
)

/*
func ParseEventStatus(eventStatus string) (EventStatus, error) {
	for _, v := range getValues() {
		if string(v) == eventStatus {
			return v, nil
		}
	}

	return "", fmt.Errorf("unable to parse EventStatus %s", eventStatus)
}

func getValues() []EventStatus {
	return []EventStatus{
		CreationInit,
		CreatingPrivateEndpoint,
		CreatingPrivateEndpointService,
	}
}
*/
