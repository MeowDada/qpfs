package drive

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Client is a http client to interacts with drive daemon.
type Client struct {
	addr string
	clnt *http.Client
}

// NewClient creates an instance of drive Client.
func NewClient(addr string) (API, error) {
	clnt := &http.Client{}
	return &Client{
		addr: addr,
		clnt: clnt,
	}, nil
}

func (c *Client) Add(ctx context.Context, key, fpath string) error {
	url, err := url.Parse(c.addr + "/add")
	if err != nil {
		return err
	}

	q := url.Query()

	q.Add("key", key)
	q.Add("path", fpath)
	url.RawQuery = q.Encode()

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return err
	}

	resp, err := c.clnt.Do(r)
	if err != nil {
		return err
	}

	m := map[string]interface{}{}
	json.NewDecoder(resp.Body).Decode(&m)

	return fmt.Errorf("%v", m)
}

func (c *Client) Get(ctx context.Context, key string) error {
	return nil
}

func (c *Client) List(ctx context.Context, prefix string) (ListResult, error) {
	url, err := url.Parse(c.addr + "/list")
	if err != nil {
		return ListResult{}, err
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return ListResult{}, err
	}

	resp, err := c.clnt.Do(r)
	if err != nil {
		return ListResult{}, err
	}

	var lr ListResult

	if err := json.NewDecoder(resp.Body).Decode(&lr); err != nil {
		return ListResult{}, err
	}

	return lr, nil
}

func (c *Client) Stop(ctx context.Context) error {
	url, err := url.Parse(c.addr + "/stop")
	if err != nil {
		return err
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return err
	}

	resp, err := c.clnt.Do(r)
	if err != nil {
		return err
	}

	m := map[string]interface{}{}
	json.NewDecoder(resp.Body).Decode(&m)

	fmt.Println(m)

	return nil
}
