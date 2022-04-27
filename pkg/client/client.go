package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/torwald-sergesson/app-a/pkg/dto/v2"
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
		return dto.User{}, fmt.Errorf("fail to build url: %w", err)
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return dto.User{}, fmt.Errorf("fail to build request: %w", err)
	}
	log.Printf("request; method=%s; url=%s\n", req.Method, req.URL.String())
	resp, err := cli.httpClient.Do(req)
	if err != nil {
		return dto.User{}, fmt.Errorf("fail to do request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dto.User{}, fmt.Errorf("fail to read body: %w", err)
	}
	log.Printf("body: %s\n", body)
	if resp.StatusCode != http.StatusOK {
		return dto.User{}, fmt.Errorf("wrong status code; expected=%d; received=%d", http.StatusOK, resp.StatusCode)
	}
	var user dto.User
	if err = json.Unmarshal(body, &user); err != nil {
		return dto.User{}, fmt.Errorf("fail to unmarshal body: %w", err)
	}
	log.Printf("user %d tags: [%s]\n", user.ID, strings.Join(user.Tags, " "))
	return user, nil
}

func (cli *Client) MyGroup() (dto.Group, error) {
	u, err := cli.url("/api/me")
	if err != nil {
		return dto.Group{}, fmt.Errorf("fail to build url: %w", err)
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return dto.Group{}, fmt.Errorf("fail to build request: %w", err)
	}
	log.Printf("request; method=%s; url=%s\n", req.Method, req.URL.String())
	resp, err := cli.httpClient.Do(req)
	if err != nil {
		return dto.Group{}, fmt.Errorf("fail to do request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dto.Group{}, fmt.Errorf("fail to read body: %w", err)
	}
	log.Printf("body: %s\n", body)
	if resp.StatusCode != http.StatusOK {
		return dto.Group{}, fmt.Errorf("wrong status code; expected=%d; received=%d", http.StatusOK, resp.StatusCode)
	}
	var group dto.Group
	if err = json.Unmarshal(body, &group); err != nil {
		return dto.Group{}, fmt.Errorf("fail to unmarshal body: %w", err)
	}
	log.Printf("group %d\n", group.ID)
	return group, nil
}
