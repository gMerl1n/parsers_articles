package parsers

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
)

type RequestOpts struct {
	Method  string
	Head    map[string]string
	Cookies map[string]string
	Proxy   map[string]string
}

type ParserHB struct {
	UrlHB string
}

func (hb *ParserHB) GetHtmlPage(reqOptions *RequestOpts) (*goquery.Document, error) {

	if reqOptions.Method == "" {
		return nil, fmt.Errorf("no method name, impossible to make")
	}

	if hb.UrlHB == "" {
		return nil, fmt.Errorf("no url, impossible to make")
	}

	// Создаем новый http запрос
	client := &http.Client{}

	request, err := http.NewRequest(reqOptions.Method, hb.UrlHB, nil)
	if err != nil {
		fmt.Println("Ошибка во время исполнения запрос", err)
	}

	if reqOptions.Head != nil {
		key, ok := reqOptions.Head["Host"]
		if !ok {
			key, ok := reqOptions.Head["host"]
			if ok {
				request.Header.Set(key, reqOptions.Head["host"])
			}

		} else {
			request.Header.Set(key, reqOptions.Head["Host"])
		}

		for key, value := range reqOptions.Head {
			request.Header.Set(key, value)
		}
	}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Ошибка во время исполнения запроса", err)
	}

	fmt.Println(response.StatusCode)

	utf8, err := charset.NewReader(response.Body, response.Header.Get("Content-Type"))
	if err != nil {
		println("Ошибка charset", err)
		return nil, err
	}

	document, err := goquery.NewDocumentFromReader(utf8)
	if err != nil {
		println("Ошибка парсинга документа из UTF-8", err)
	}

	defer response.Body.Close()

	return document, nil

}

func (hb *ParserHB) GetUrlsArticles(pageHTML *goquery.Document) []string {

	arrayUrls := make([]string, 0)

	// Через цикл получаем ссылки на все статьи с главной страницы
	pageHTML.Find(".mg-card__title").Each(func(index int, item *goquery.Selection) {
		url, ok := item.Find(".mg-card__link").Attr("href")
		if ok {
			arrayUrls = append(arrayUrls, url)
		}

	})

	return arrayUrls

}
