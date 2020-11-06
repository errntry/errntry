package i18n

import "strings"

type Response struct {
	Code    int    `json:"code"`
	Err     string `json:"error,omitempty"`
	Msg     string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
}

func (r Response) String() string {
	if r.Err != "" {
		return r.Err
	}
	return r.Msg
}

func (r Response) LowerString() string {
	return strings.ToLower(r.String())
}

func (r Response) Error() string {
	return r.Err
}

func (r Response) Success() bool {
	return r.Err == "" && r.Code < 400
}

func NewResponse(code int, msg string) Response {
	if code < 400 {
		return Response{Code: code, Msg: msg}
	}
	return Response{Code: code, Err: msg}
}
