package constants

import (
	"fmt"
)

type EventStatus string

const (
	CreationInit                   EventStatus = "CREATION_INIT"
	CreatingPrivateEndpointService EventStatus = "CREATING_PRIVATE_ENDPOINT_SERVICE"
	CreatingPrivateEndpoint        EventStatus = "CREATING_PRIVATE_ENDPOINT"
)

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
