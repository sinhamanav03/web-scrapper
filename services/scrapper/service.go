package scrapper

import (
	"context"
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/sinhamanav03/web-scrapper/internal/config"
	"github.com/sinhamanav03/web-scrapper/services/scrapper/model"
)

type ScrapperService interface {
	ScrapeUrl(ctx context.Context, url string) (*model.ScrapeUrlResponse, error)
}

type scrapperService struct {
	conf *config.Config
	coll *colly.Collector
}

func NewScrapperService(conf *config.Config, coll *colly.Collector) ScrapperService {
	scSvc := scrapperService{
		conf: conf,
		coll: coll,
	}
	return &scSvc
}

func (svc *scrapperService) ScrapeUrl(ctx context.Context, url string) (*model.ScrapeUrlResponse, error) {
	resp := model.ScrapeUrlResponse{}

	svc.coll.OnRequest(func(r *colly.Request) {
		fmt.Println("Link of the page:", r.URL)
	})

	//get Landing image
	svc.coll.OnHTML(svc.conf.Attributes.Image, func(h *colly.HTMLElement) {
		resp.ImageURL = h.Attr("src")
	})

	//get product name
	svc.coll.OnHTML(svc.conf.Attributes.Name, func(h *colly.HTMLElement) {
		resp.Name = strings.TrimSpace(h.Text)
	})

	//get porduct price
	svc.coll.OnHTML(svc.conf.Attributes.Price, func(h *colly.HTMLElement) {
		resp.Price = h.Text
	})

	//get product review count
	svc.coll.OnHTML(svc.conf.Attributes.Reviews, func(h *colly.HTMLElement) {
		resp.Reviews = strings.Split(h.Text, " ")[0]
	})

	//get product description
	svc.coll.OnHTML(svc.conf.Attributes.Description, func(h *colly.HTMLElement) {
		resp.Description = resp.Description + h.Text

	})

	fmt.Println(resp)

	svc.coll.Visit(url)

	return &resp, nil
}
