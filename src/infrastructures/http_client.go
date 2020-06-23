package infrastructures

import (
	"io/ioutil"
	"net/http"
)

type IHttpClient interface {
	Get(url string, header map[string]string) (*HttpResponse, error)
}

type HttpResponse struct {
	Body       []byte
	StatusCode int
}

type HttpClient struct{}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (c *HttpClient) Get(url string, header map[string]string) (*HttpResponse, error) {
	cl := new(http.Client)

	req, err := http.NewRequest("GET", url, nil)
	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	if err != nil {
		return nil, err
	}

	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &HttpResponse{
		Body:       bs,
		StatusCode: res.StatusCode,
	}, nil
}
