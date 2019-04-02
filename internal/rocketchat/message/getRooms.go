package message

import "github.com/google/uuid"

type GetRooms struct {
	baseMethodMessage
	Parameters []GetRoomsParam
}

type GetRoomsParam struct {
	Date string `json:"$date"`
}

type GetRoomsResponse struct {
	Results []GetRoomsResponseResult `json:"result"`
}

type GetRoomsResponseResult struct {
	ID   string `json:"_id"`
	Type string `json:"t"`
	Name string `json:"name"`
}

func NewGetRooms() GetRooms {
	return GetRooms{
		baseMethodMessage: baseMethodMessage{
			Message: "method",
			Method:  "rooms/get",
			ID:      uuid.New().String(),
		},
		Parameters: []GetRoomsParam{
			{
				Date: "0",
			},
		},
	}
}
