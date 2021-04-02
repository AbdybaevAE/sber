package esb

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/abdybaevae/guid"
)

type impl struct {
	client  *http.Client
	options *Options
}

func (i *impl) setParams(params interface{}) {
	i.options.Body = &body{
		Header: bodyHeader{
			Service:      i.options.Service,
			Platform:     i.options.Platform,
			SourceSystem: i.options.SourceSystem,
			RqUid:        guid.Gen(),
			OperUid:      guid.Gen(),
			RqTm:         time.Now().Format(time.RFC3339),
		},
		Params: params,
	}
}
func (i *impl) Do(method HttpMethod, path string, params interface{}) (*http.Response, error) {
	var body io.Reader
	if method == PostMethod {
		i.setParams(params)
		bodyBytes, err := json.Marshal(i.options.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequest(string(method), i.options.Host+path, body)
	if err != nil {
		return nil, err
	}
	i.client.Timeout = i.options.Timeout
	return i.client.Do(req)
}

func (i *impl) DoWithOptions(method HttpMethod, path string, params interface{}, opts ...Option) (*http.Response, error) {
	newImpl := NewClient(opts...)
	return newImpl.Do(method, path, params)
}
func NewClient(opts ...Option) Client {
	options := NewOptions(opts...)
	client := &http.Client{
		Timeout: options.Timeout,
	}
	return &impl{
		options: options,
		client:  client,
	}
}
