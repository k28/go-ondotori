package ondotori

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type OndotoriError struct {
	Code    int
	Message string
}

func (e *OndotoriError) Error() string {
	return fmt.Sprintf("%v [%d]", e.Message, e.Code)
}

// httpClient defines the minimal interface need for http.Client to be implemented.
type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Option func(*Client)

func OptionHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.httpclient = hc
	}
}

type Client struct {
	baseParam  BaseParam
	httpclient httpClient
}

func New(token string, login_id string, login_pass string, opts ...Option) (*Client, error) {
	b := BaseParam{
		Token:     token,
		LoginId:   login_id,
		LoginPass: login_pass,
	}
	s := &Client{
		baseParam:  b,
		httpclient: &http.Client{},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s, nil
}

func (client *Client) Get(param makeParam, ctx context.Context) (*Devices, error) {
	jsonReq, err := json.Marshal(param.MakeJsonMap(client.baseParam))
	if err != nil {
		return nil, err
	}
	// fmt.Println("Request :", string(jsonReq))
	u := param.MakeUri(client.baseParam)
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer([]byte(string(jsonReq))))
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-HTTP-Method-Override", "GET")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.request(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("bad response status code %d", resp.StatusCode)
	}

	//b, err := io.ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var r io.Reader = resp.Body
	r = io.TeeReader(r, os.Stderr)

	var body Devices
	if err := json.NewDecoder(r).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (client *Client) request(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)

	respCh := make(chan *http.Response)
	errCh := make(chan error)

	go func() {
		resp, err := client.httpclient.Do(req)
		if err != nil {
			errCh <- err
			return
		}

		respCh <- resp
	}()

	select {
	case resp := <-respCh:
		return resp, nil
	case err := <-errCh:
		return nil, err
	case <-ctx.Done():
		return nil, errors.New("HTTP request cancelled")
	}
}
