package slg

import (
	"encoding/xml"
	"errors"
	"github.com/h2non/gentleman"
	"github.com/h2non/gentleman/plugins/query"
	"github.com/h2non/gentleman/plugins/url"
	"github.com/h2non/gentleman/context"
	"log"
	"net/http"
	"strconv"
)

type Search struct {
	Housings []*Housing `xml:"annonces>annonce"`
}

type SearchParams struct {
	PageSize   *int
	PageNumber *int
	MinPrice   *float64
	MaxPrice   *float64
	ZipCode    string
}

func NewSearchParams() SearchParams {
	return SearchParams{}
}

func (p SearchParams) SetPageSize(size int) SearchParams {
	p.PageSize = &size
	return p
}

func (p SearchParams) SetPageNumber(number int) SearchParams {
	p.PageNumber = &number
	return p
}

func (p SearchParams) SetMinPrice(price float64) SearchParams {
	p.MinPrice = &price
	return p
}

func (p SearchParams) SetMaxPrice(price float64) SearchParams {
	p.MaxPrice = &price
	return p
}

func (p SearchParams) SetZipCode(zip string) SearchParams {
	p.ZipCode = zip
	return p
}

type SearchService struct {
	client *Client
}

func NewSearchService(client *Client) *SearchService {
	return &SearchService{client: client}
}

func (s *SearchService) Search(params SearchParams) (*Search, error) {
	searcher, err := s.client.httpClient()
	if err != nil {
		return nil, err
	}
	searcher.Use(url.Path("/5_1,search.xml"))

	addQuery(searcher, "SEARCHpi", params.PageSize)
	addQuery(searcher, "SEARCHpg", params.PageNumber)
	addQuery(searcher, "pxmin", params.MinPrice)
	addQuery(searcher, "pxmax", params.MaxPrice)
	addQuery(searcher, "cp", params.ZipCode)
	addQuery(searcher, "idtt", "1")
	addQuery(searcher, "idtypebien", "1,2")
	addQuery(searcher, "nbslots", "3")

	res, err := searcher.Request().UseRequest(func(ctx *context.Context, h context.Handler) {
		log.Printf("query=%v", ctx.Request.URL.Query().Encode())
		h.Next(ctx)
	}).Method("GET").Send()

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("authentication required")
	}

	if !res.Ok {
		return nil, errors.New("http request failed")
	}

	log.Println(res.String())
	search := &Search{}
	err = xml.Unmarshal(res.Bytes(), search)
	if err != nil {
		return nil, err
	}

	return search, nil
}

func addQuery(c *gentleman.Client, key string, value interface{}) {
	queryValue := ""
	switch v := value.(type) {
	case *int:
		if value != nil {
			queryValue = strconv.Itoa(*v)
		}
	case *float64:
		if value != nil {
			queryValue = strconv.FormatFloat(*v, 'f', 6, 64)
		}
	case string:
		if value != nil {
			queryValue = v
		}
	}
	c.Use(query.Set(key, queryValue))
}
