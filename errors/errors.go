package errors

import (
	"encoding/json"
	"fmt"
)

const (
	errorStringFormat = "ErrCode [%d]\nErrMsg: %s\nErrMsgEn: %s"
)

var (
	errLogger = newLogger()
)

// Error holds the error message, this message never really changes
type Error struct {
	ErrCode ErrCode `json:"err_code"`
	// The message of the error.
	ErrMsg string `json:"err_msg"`
	// The message of the error.
	ErrMsgEn string `json:"err_msg_en"`
}

type CustomErrMsg BaseErrMsg

// New creates and returns an Error with a pre-defined user output message
func New(errCode ErrCode, errMsg ...CustomErrMsg) Error {
	if _, ok := baseErrors[errCode]; ok {
		if len(errMsg) > 1 {
			errLogger.Warning("redundant errMsg parameters")
		}

		if len(errMsg) > 0 {
			var tmpErrMsg string
			var tmpErrMsgEn string

			if errMsg[0].ErrMsg != "" {
				tmpErrMsg = errMsg[0].ErrMsg
			} else {
				tmpErrMsg = baseErrors[errCode].ErrMsg
			}
			if errMsg[0].ErrMsgEn != "" {
				tmpErrMsgEn = errMsg[0].ErrMsgEn
			} else {
				tmpErrMsgEn = baseErrors[errCode].ErrMsgEn
			}
			return Error{
				ErrCode:  errCode,
				ErrMsg:   tmpErrMsg,
				ErrMsgEn: tmpErrMsgEn,
			}
		} else {
			return Error{
				ErrCode:  errCode,
				ErrMsg:   baseErrors[errCode].ErrMsg,
				ErrMsgEn: baseErrors[errCode].ErrMsgEn,
			}
		}
	} else {
		if len(errMsg) == 0 {
			errLogger.Error("lack of err msg info when defining custom error")
			return Error{
				ErrCode:  errCode,
				ErrMsg:   "未知错误",
				ErrMsgEn: "unknown error",
			}
		}
		if errMsg[0].ErrMsg == "" {
			errLogger.Error("lack of err msg info when defining custom error")
			return Error{
				ErrCode:  errCode,
				ErrMsg:   "未知错误",
				ErrMsgEn: "unknown error",
			}
		}
		return Error{
			ErrCode:  errCode,
			ErrMsg:   errMsg[0].ErrMsg,
			ErrMsgEn: errMsg[0].ErrMsgEn,
		}
	}
}

func NewFormat(errCode ErrCode, errMsg ...interface{}) Error {
	if oriError, ok := baseErrors[errCode]; ok {
		if len(errMsg) > 0 {
			var tmpErrMsg string
			tmpErrMsg = fmt.Sprintf(oriError.ErrMsg, errMsg...)

			return Error{
				ErrCode:  errCode,
				ErrMsg:   tmpErrMsg,
				ErrMsgEn: oriError.ErrMsgEn,
			}
		} else {
			return Error{
				ErrCode:  errCode,
				ErrMsg:   oriError.ErrMsg,
				ErrMsgEn: oriError.ErrMsgEn,
			}
		}
	} else {
		return Error{
			ErrCode:  errCode,
			ErrMsg:   "未知错误",
			ErrMsgEn: "unknown error",
		}
	}
}

func (e Error) Error() string {
	return fmt.Sprintf(errorStringFormat, e.ErrCode, e.ErrMsg, e.ErrMsgEn)
}

func (e Error) Marshal() []byte {
	j, _ := json.Marshal(e)
	return j
}

func (e Error) MarshalString() string {
	return string(e.Marshal())
}
