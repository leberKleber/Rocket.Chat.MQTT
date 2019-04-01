package message

type baseMethodMessage struct {
	ID      string `json:"id"`
	Message string `json:"msg"`
	Method  string `json:"method"`
}

type GeneralMessage struct {
	Message string `json:"msg"`
	ID      string `json:"id"`
	Method  string `json:"method"`
}
