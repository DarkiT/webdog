package message

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
