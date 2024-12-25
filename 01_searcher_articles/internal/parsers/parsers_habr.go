package parsers

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/entities"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/services"
	"go.uber.org/zap"
	"golang.org/x/net/html/charset"
)

type ParserHabr struct {
	dataForParsing entities.DataForParsing
	service        services.ServiceArticle
	logger         *zap.Logger
}

func NewParserHabr(data entities.DataForParsing, serv services.ServiceArticle, log *zap.Logger) *ParserHabr {
	return &ParserHabr{
		dataForParsing: data,
		service:        serv,
		logger:         log,
	}
}

func (p *ParserHabr) MakeRequest() (*goquery.Document, error) {

	// Создаем новый http запрос
	client := &http.Client{}

	request, err := http.NewRequest("GET", p.dataForParsing.UrlCategory, nil)
	if err != nil {
		p.logger.Warn("Failed to make request", zap.Error(err))
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil {
		p.logger.Warn("Failed to make request", zap.Error(err))
		return nil, err
	}

	defer response.Body.Close()

	// Получаем содержимое html пакета
	utf8, err := charset.NewReader(response.Body, response.Header.Get("Content-Type"))
	if err != nil {
		p.logger.Warn("Failed to decode response body", zap.Error(err))
		return nil, err
	}

	document, err := goquery.NewDocumentFromReader(utf8)
	if err != nil {
		p.logger.Warn("Failed to make document from utf 8", zap.Error(err))
		return nil, err
	}

	return document, nil

}

func (p *ParserHabr) GetNumPages(pageHTML *goquery.Document) {

	numPages := pageHTML.Find("tm-pagination__page")
	fmt.Println(numPages)
}
