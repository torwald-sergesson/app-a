package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/torwald-sergesson/app-a/pkg/dto"
)

type Client struct {
	addr       string
	httpClient http.Client
}

func NewClient(addr string, timeout time.Duration) *Client {
	return &Client{
		addr: addr,
		httpClient: http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       timeout,
		},
	}
}

func (cli *Client) url(path string) (url.URL, error) {
	u := url.URL{
		Scheme: "http",
		Host:   cli.addr,
		Path:   path,
	}
	return u, nil
}

func (cli *Client) Me() (dto.User, error) {
	u, err := cli.url("/api/me")
	if err != nil {
		return dto.User{}, nil
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return dto.User{}, err
	}
	resp, err := cli.httpClient.Do(req)
	if err != nil {
		return dto.User{}, err
	}
	var body []byte
	defer resp.Body.Close()

	if _, err = resp.Body.Read(body); err != nil {
		return dto.User{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return dto.User{}, fmt.Errorf("")
	}
	var user dto.User
	if err = json.Unmarshal(body, &user); err != nil {
		return dto.User{}, err
	}
	return user, nil
}
