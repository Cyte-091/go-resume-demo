package errno

const (
	Success      = 0
	ParamErr     = 10001
	UserNotLogin = 10002
	BizErr       = 10003
	PageNotFound = 404
)

var msg = map[int]string{
	Success:      "success",
	ParamErr:     "参数错误",
	UserNotLogin: "未登录",
	BizErr:       "业务异常",
	PageNotFound: "页面不存在",
}

type ErrNo struct {
	Code int
	Msg  string
}

func (e *ErrNo) Error() string {
	return e.Msg
}

func GetMsg(code int) string { return msg[code] }

func New(code int, desc string) error {
	return &ErrNo{Code: code, Msg: desc}
}
