package http_model

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/valyala/fasthttp"
)

type QueryInterface interface {
	BuildParams() string
}

type HttpTaskInterface interface {
	GetData() QueryInterface
	GetUrl() string
}

func buildQueryUrl(url string, query QueryInterface) string {
	ret := fmt.Sprintf("%s?%s", url, query.BuildParams())
	return ret
}

func requestGet(s HttpTaskInterface) ([]byte, error) {
	query := s.GetData()
	url := s.GetUrl()
	furl := buildQueryUrl(url, query)
	fmt.Println(furl)
	statusCode, body, err := fasthttp.Get(nil, furl)
	if err != nil || statusCode != 200 {
		hlog.Infof("request get: %s error, code: %d, error: %s",
			furl, statusCode, err.Error())
		return []byte{}, err
	} else {
		return body, nil
	}
}

func requestPost(s HttpTaskInterface) {

}
