package esb

import "net/http"

type Client interface {
	Do(method HttpMethod, path string, params interface{}) (response *http.Response, err error)
	DoWithOptions(method HttpMethod, path string, body interface{}, opts ...Option) (response *http.Response, err error)
}
type HttpMethod string

const (
	PostMethod = "post"
	GetMethod  = "get"
)
