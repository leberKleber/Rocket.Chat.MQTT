package message

type Connect struct {
	Message string   `json:"msg"`
	Version string   `json:"version"`
	Support []string `json:"support"`
}

func NewConnect() Connect {
	return Connect{
		Message: "connect",
		Version: "1",
		Support: []string{"1"},
	}
}
