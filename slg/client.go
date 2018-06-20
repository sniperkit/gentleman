package slg

import (
	"errors"

	"github.com/h2non/gentleman"
	"github.com/h2non/gentleman/plugins/headers"
	"github.com/h2non/gentleman/plugins/url"
)

const (
	baseUrl = "https://ws-seloger.svc.groupe-seloger.com"
)

type ClientOpts struct {
	AppGuid  string
	AppToken string
}

type Client struct {
	opts           ClientOpts
	token          string
	httpBaseClient *gentleman.Client
	search         *SearchService
}

func NewClient(opts ClientOpts) *Client {
	httpBaseClient := gentleman.New()
	httpBaseClient.Use(url.BaseURL(baseUrl))
	client := &Client{
		opts:           opts,
		httpBaseClient: httpBaseClient,
	}
	client.search = NewSearchService(client)
	return client
}

func (c *Client) Search(params SearchParams) (*Search, error) {
	return c.search.Search(params)
}

func (c *Client) httpClient() (*gentleman.Client, error) {
	if c.token == "" {
		if err := c.Authenticate(); err != nil {
			return nil, err
		}
	}
	client := gentleman.New().UseParent(c.httpBaseClient)
	client.Use(headers.Set("AppToken", c.token))
	client.Use(headers.Del("User-Agent"))
	client.Use(headers.Del("Content-Type"))
	return client, nil
}

func (c *Client) Authenticate() error {
	authenticator := gentleman.New()
	authenticator.UseParent(c.httpBaseClient)

	authenticator.Use(url.Path("/5_1,identification.xml"))

	authenticator.Use(headers.Set("AppGuid", c.opts.AppGuid))
	authenticator.Use(headers.Set("AppToken", c.opts.AppToken))

	res, err := authenticator.Request().Method("GET").Send()

	if err != nil {
		return err
	}

	if !res.Ok {
		return errors.New("HTTP request failed")
	}

	c.token = res.Header["Apptoken"][0]
	return nil
}
