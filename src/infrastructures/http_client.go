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

type HttpClient struct {
	http_Client_Do func(req *http.Request) (*http.Response, error)
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		http_Client_Do: new(http.Client).Do,
	}
}

func (c *HttpClient) request(m, url string, header map[string]string, body []byte) (*HttpResponse, error) {
	req, err := http.NewRequest(m, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	res, err := c.http_Client_Do(req)
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

func (c *HttpClient) Get(url string, header map[string]string) (*HttpResponse, error) {
	return c.request("GET", url, header, nil)
}

func (c *HttpClient) Post(url string, header map[string]string, body []byte) (*HttpResponse, error) {
	return c.request("POST", url, header, body)
}

func (c *HttpClient) Delete(url string, header map[string]string) (*HttpResponse, error) {
	return c.request("DELETE", url, header, nil)
}
