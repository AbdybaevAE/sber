package esb

import (
	"time"
)

const DefaultTimeout = time.Second * 60

type bodyHeader struct {
	Service      string `json:"service"`
	Platform     string `json:"platform"`
	SourceSystem string `json:"sourceSystem"`
	RqUid        string `json:"rqUid"`
	OperUid      string `json:"operUid"`
	RqTm         string `json:"rqTm"`
}

type body struct {
	Header bodyHeader  `json:"header"`
	Params interface{} `json:"rqParms"`
}
type Options struct {
	Host         string
	Timeout      time.Duration
	Body         *body
	Service      string
	Platform     string
	SourceSystem string
}
type Option func(op *Options)

func SetHeader(service, platform, sourceSystem string) Option {
	return func(op *Options) {
		op.Service = service
		op.Platform = platform
		op.SourceSystem = sourceSystem
	}
}
func SetTimeout(timeout time.Duration) Option {
	return func(op *Options) {
		op.Timeout = timeout
	}
}
func SetHost(host string) Option {
	return func(op *Options) {
		op.Host = host
	}
}
func NewOptions(opts ...Option) *Options {
	options := &Options{
		Timeout: DefaultTimeout,
	}
	for _, o := range opts {
		o(options)
	}
	return options
}

// func (o *Options) Copy() *Options {
// 	return &Options{
// 		Host:    o.Host,
// 		Timeout: o.Timeout,
// 	}
// }
