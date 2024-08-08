package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"opnsense-cli/internal/config"
	"strings"
)

type ApiClient struct {
	Host      string
	ApiKey    string
	ApiSecret string
	client    *http.Client
}

func New(cfg config.Config) ApiClient {
	c := ApiClient{
		Host:      cfg.Host,
		ApiKey:    cfg.ApiKey,
		ApiSecret: cfg.ApiSecret,
		client:    http.DefaultClient,
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return c
}

func (a ApiClient) ReconfigureUnbound() error {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://%s/api/unbound/service/reconfigure", a.Host), nil)
	if err != nil {
		return fmt.Errorf("error creating http request: %w", err)
	}
	req.SetBasicAuth(a.ApiKey, a.ApiSecret)
	res, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("error performing request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status code 200, got: %d", res.StatusCode)
	}
	return nil
}

func (a ApiClient) PerformRequest(method, module, controller, command string, body io.Reader, args ...string) (io.ReadCloser, error) {
	var url *url.URL
	var err error
	if args == nil {
		url, err = url.Parse(fmt.Sprintf("https://%s/api/%s/%s/%s", a.Host, module, controller, command))
	} else {
		url, err = url.Parse(fmt.Sprintf("https://%s/api/%s/%s/%s%s", a.Host, module, controller, command, mapArgs(args)))
	}

	if err != nil {
		return nil, fmt.Errorf("unable to parse url with provided values: %w", err)
	}
	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	req.SetBasicAuth(a.ApiKey, a.ApiSecret)
	res, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status code 200, got: %d", res.StatusCode)
	}
	return res.Body, nil
}

func mapArgs(args []string) string {
	b := strings.Builder{}
	for _, v := range args {
		b.WriteString(fmt.Sprintf("/%s", v))
	}
	return b.String()
}
