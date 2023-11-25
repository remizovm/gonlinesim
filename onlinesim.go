package gonlinesim

import (
	"errors"
)

const (
	defaultHost = "https://onlinesim.io"

	LangFR = "fr"
	LangDE = "de"
	LangRU = "ru"
	LangEN = "en"
	LangZH = "zh"
)

var (
	ErrAccountBlocked    = errors.New("account blocked")
	ErrWrongKey          = errors.New("wrong key")
	ErrNoKey             = errors.New("no key")
	ErrNoService         = errors.New("no service or invalid service name")
	ErrRequestNotFound   = errors.New("request not found")
	ErrAPIAccessDisabled = errors.New("api access disabled")
	ErrAPIAccessIP       = errors.New("api access blocked for this IP")
	ErrLowBalance        = errors.New("low balance")
)

func checkErr(val string) error {
	switch val {
	case "1":
		return nil
	case "ACCOUNT_BLOCKED":
		return ErrAccountBlocked
	case "ERROR_WRONG_KEY":
		return ErrWrongKey
	case "ERROR_NO_KEY":
		return ErrNoKey
	case "ERROR_NO_SERVICE":
		return ErrNoService
	case "REQUEST_NOT_FOUND":
		return ErrRequestNotFound
	case "API_ACCESS_DISABLED":
		return ErrAPIAccessDisabled
	case "API_ACCESS_IP":
		return ErrAPIAccessIP
	case "WARNING_LOW_BALANCE":
		return ErrLowBalance
	default:
		return nil
	}
}
