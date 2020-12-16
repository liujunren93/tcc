package tcc

type Status interface {
	GetCode() int32
	GetMsg() string
}
type status struct {
	code int32
	msg  string
}

func (s status) GetCode() int32 {
	return s.code
}

func (s status) GetMsg() string {
	return s.msg
}

var StatusOk = status{200, "ok"}

func NewStatus(code int32, msg string) status {
	return status{code, msg}
}
