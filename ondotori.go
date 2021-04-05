package ondotori

import (
	"fmt"
	"net/http"
	"context"
	"encoding/json"
	"io"
	"os"
	"errors"
	"bytes"
)

type OndotoriError struct {
	Code int
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
	token string
	loginId string
	loginPass string
	httpclient httpClient
}

type CurrentParams struct {
	remote []string
	base []string
}

func (p *CurrentParams) validate() error {
	if len(p.remote) > 0 && len(p.base) > 0 {
		return &OndotoriError{9999, "子機のシリアルと親機のシリアルは同時に指定できません"}
	}
	return nil
}

func New(token string, login_id string, login_pass string, opts ...Option) (*Client, error) {
	s := &Client{
		token: token,
		loginId: login_id,
		loginPass: login_pass,
		httpclient: &http.Client{},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s, nil
}

type ResponseBody struct {
	Text string `json:"devices"`
}

func (client *Client) Get(ctx context.Context) (*Devices, error) {
	jsonReq := `{"api-key":"` + client.token + `","login-id":"` + client.loginId + `","login-pass":"` + client.loginPass + `"}`
	fmt.Println("Request :", jsonReq)
	req, err := http.NewRequest(http.MethodPost, "https://api.webstorage.jp/v1/devices/current", bytes.NewBuffer([]byte(jsonReq)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-HTTP-Method-Override", "GET")
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Authorization", fmt.Sprintf("Bearere %s", client.token))

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

