package infrastructures

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// IHTTPClient ...
type IHTTPClient interface {
	Get(url string, header map[string]string) (*HTTPResponse, error)
	Post(url string, header map[string]string, body []byte) (*HTTPResponse, error)
	Delete(url string, header map[string]string) (*HTTPResponse, error)
}

// HTTPResponse ...
type HTTPResponse struct {
	Body       []byte
	StatusCode int
}

// HTTPClient ...
type HTTPClient struct {
	send func(req *http.Request) (*http.Response, error)
}

// NewHTTPClient ...
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		send: new(http.Client).Do,
	}
}

func (c *HTTPClient) request(m, url string, header map[string]string, body []byte) (*HTTPResponse, error) {
	req, err := http.NewRequest(m, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	res, err := c.send(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &HTTPResponse{
		Body:       bs,
		StatusCode: res.StatusCode,
	}, nil
}

// Get ...
func (c *HTTPClient) Get(url string, header map[string]string) (*HTTPResponse, error) {
	return c.request("GET", url, header, nil)
}

// Post ...
func (c *HTTPClient) Post(url string, header map[string]string, body []byte) (*HTTPResponse, error) {
	return c.request("POST", url, header, body)
}

// Delete ...
func (c *HTTPClient) Delete(url string, header map[string]string) (*HTTPResponse, error) {
	return c.request("DELETE", url, header, nil)
}
