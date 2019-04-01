package message

type Pong struct {
	Message string `json:"msg"`
}

func NewPong() Pong {
	return Pong{
		Message: "ping",
	}
}
