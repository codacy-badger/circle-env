package infrastructures

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type IHttpClient interface {
	Get(url string, header map[string]string) (*HttpResponse, error)
	Post(url string, header map[string]string, body []byte) (*HttpResponse, error)
	Delete(url string, header map[string]string) (*HttpResponse, error)
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
	if err != nil {
		return nil, err
	}

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
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

func (c *HttpClient) Post(u string, header map[string]string, body []byte) (*HttpResponse, error) {
	cl := new(http.Client)

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
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

func (c *HttpClient) Delete(url string, header map[string]string) (*HttpResponse, error) {
	cl := new(http.Client)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
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
