package quivr_go_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

const (
	DefaultReadTimeout = 5 * time.Minute

	POST   = "POST"
	GET    = "GET"
	DELETE = "DELETE"
	PUT    = "PUT"
)

type Client struct {
	addr       string
	apiKey     string
	timeout    time.Duration
	httpClient *http.Client
}

func NewClient(addr, apiKey string, opts ...Option) *Client {
	c := &Client{
		addr:    addr,
		apiKey:  apiKey,
		timeout: DefaultReadTimeout,
	}
	for _, opt := range opts {
		opt(c)
	}
	c.httpClient = &http.Client{
		Timeout: c.timeout,
	}
	return c
}

func (c *Client) do(ctx context.Context, method string, api string, input interface{}) ([]byte, error) {
	var b io.Reader
	var body []byte
	var err error
	if input != nil {
		body, err = json.Marshal(input)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal input: %v", err)
		}
		// turn the body into an io.Reader
		b = bytes.NewReader(body)
	}

	// create a request with all headers required for authentication.
	req, err := http.NewRequest(method, c.addr+api, b)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.WithContext(ctx)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	// run the request and return the response body.
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request api: %s, err: %v", api, err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body, api: %s, err: %v", api, err)
	}
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("api response err, api: %s, http code: %d, err: %s", api, resp.StatusCode, gjson.GetBytes(respBody, "detail").String())
	}
	fmt.Println("response:", string(respBody))
	return respBody, nil
}

func (c *Client) upload(ctx context.Context, api string, filename string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", c.addr+api, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 创建一个multipart表单文件域
	part, err := writer.CreateFormFile("uploadFile", filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %v", err)
	}
	_, _ = part.Write(data)

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close writer: %v", err)
	}

	req.WithContext(ctx)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Body = io.NopCloser(body)

	// run the request and return the response body.
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request api: %s, err: %v", api, err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body, api: %s, err: %v", api, err)
	}
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("api response err, api: %s, http code: %d, err: %s", api, resp.StatusCode, gjson.GetBytes(respBody, "detail").String())
	}
	return respBody, nil
}

func Do[T any](ctx context.Context, c *Client, method string, api string, input interface{}) (*T, error) {
	bs, err := c.do(ctx, method, api, input)
	if err != nil {
		return nil, err
	}
	resp := new(T)
	err = json.Unmarshal(bs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Upload[T any](ctx context.Context, c *Client, api, filename string, data []byte) (*T, error) {
	bs, err := c.upload(ctx, api, filename, data)
	if err != nil {
		return nil, err
	}
	resp := new(T)
	err = json.Unmarshal(bs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
