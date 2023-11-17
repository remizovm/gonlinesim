package gonlinesim

import (
	"strconv"

	"github.com/go-resty/resty/v2"
)

type Option func(*resty.Request)

func WithCountryCode(val uint) Option {
	return func(r *resty.Request) {
		r.SetQueryParam("country", strconv.Itoa(int(val)))
	}
}

func WithLang(val string) Option {
	return func(r *resty.Request) {
		r.SetQueryParam("lang", val)
	}
}

func WithPhone(val string) Option {
	return func(r *resty.Request) {
		r.SetQueryParam("number", val)
	}
}
