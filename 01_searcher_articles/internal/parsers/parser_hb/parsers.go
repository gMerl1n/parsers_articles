package parserhb

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/entities"
	"go.uber.org/zap"
	"golang.org/x/net/html/charset"
)

var Headers = map[string][]string{
	"Accept":                    {"application/json"},
	"Accept-Language":           {"en-US,en;q=0.9,ru;q=0.8"},
	"Cache-Control":             {"max-age=0"},
	"Connection":                {"keep-alive"},
	"Content-Length":            {"99"},
	"Cookie":                    {"_gcl_au=1.1.1517141233.1729442364; _ga_S28W1WC23F=GS1.1.1735150451.55.1.1735150949.43.0.91050444; _ga=GA1.1.1574279796.1729442364; hl=ru; fl=ru; habr_uuid=jyaQHosraeKJeyTvAtE8dDR4V4oCIdN%2BiN%2FuEq0%2Fwedz0TA1i%2FXul2Bf%2FCeSY%2BOgRl4ok3iV%2BF70p06dCMIOVA; visited_articles=493088:552278:688546:137089:859990:801981:72757:596071:773088:415829; _ym_uid=1729442932197542418; _ym_d=1729442932; _ga_8ZVM81B7DF=GS1.1.1733245981.2.1.1733246001.40.0.1008508144; mp_e2d341d0f1fa432ebeafb8f954b334b2_mixpanel=%7B%22distinct_id%22%3A%20%22%24device%3A192c312f2c4183c-0ee99eeb9a092a-44300520-384000-192c312f2c4183c%22%2C%22%24device_id%22%3A%20%22192c312f2c4183c-0ee99eeb9a092a-44300520-384000-192c312f2c4183c%22%2C%22%24initial_referrer%22%3A%20%22https%3A%2F%2Fcareer.habr.com%2F%22%2C%22%24initial_referring_domain%22%3A%20%22career.habr.com%22%2C%22__mps%22%3A%20%7B%7D%2C%22__mpso%22%3A%20%7B%22%24initial_referrer%22%3A%20%22https%3A%2F%2Fcareer.habr.com%2F%22%2C%22%24initial_referring_domain%22%3A%20%22career.habr.com%22%7D%2C%22__mpus%22%3A%20%7B%7D%2C%22__mpa%22%3A%20%7B%7D%2C%22__mpu%22%3A%20%7B%7D%2C%22__mpr%22%3A%20%5B%5D%2C%22__mpap%22%3A%20%5B%5D%7D; _ga_P5KMDQF7GZ=GS1.1.1732453959.1.0.1732453959.60.0.0; _ga_HPBXGJ309J=GS1.1.1732547714.1.0.1732547714.0.0.0; _ga_NLL0MBRHY7=GS1.1.1732661102.1.0.1732661102.0.0.0; habr_web_home_feed=/articles/"},
	"Host":                      {"habr.com"},
	"Sec-Fetch-Dest":            {"document"},
	"Sec-Fetch-Mode":            {"navigate"},
	"Sec-Fetch-Site":            {"cross-site"},
	"Sec-Fetch-User":            {"?1"},
	"Upgrade-Insecure-Requests": {"1"},
	"User-Agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"},
	"sec-ch-ua":                 {"'Google Chrome';v='117', 'Not;A=Brand';v='8', 'Chromium';v='117'"},
	"sec-ch-ua-mobile":          {"?0"},
	"sec-ch-ua-platform":        {"Windows"},
}

type ParserHabr struct {
	logger *zap.Logger
}

func NewParserHabr(log *zap.Logger) *ParserHabr {
	return &ParserHabr{
		logger: log,
	}
}

func (p *ParserHabr) MakeRequest(url string) (*goquery.Document, error) {

	// Создаем новый http запрос
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		p.logger.Warn("Failed to make request", zap.Error(err))
		return nil, err
	}

	request.Header = http.Header(Headers)

	response, err := client.Do(request)

	if err != nil {
		p.logger.Warn("Failed to make request", zap.Error(err))
		return nil, err
	}

	if response.StatusCode != 200 {
		p.logger.Warn("Failed to make request", zap.Int("Status code", response.StatusCode))
		return nil, err
	}

	p.logger.Info("Successful", zap.Int("Status code", response.StatusCode))

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

	p.logger.Info("Successfulley received document")
	return document, nil

}

func (p *ParserHabr) GetNumPages(pageHTML *goquery.Document) string {

	var num string

	pageHTML.Find(".tm-pagination__pages").Each(func(i int, s *goquery.Selection) {
		num = s.Find("a").Last().Text()
	})

	return num

}

func (p *ParserHabr) IsDeepExceeded(article *entities.Article, deep int64) bool {

	// if article.PublishedAt > deep {
	// 	p.logger.Info("Deep has been exceeded", zap.String("Title article", article.Title))
	// 	return true
	// } else {
	// 	return false
	// }
	return true
}

func (p *ParserHabr) ParseArticle(urlPage, providerSign string) (entities.Article, error) {

	pageHTML, err := p.MakeRequest(urlPage)
	if err != nil {
		return entities.Article{}, err
	}

	author := pageHTML.Find(".tm-user-info__username").Text()
	title := pageHTML.Find(".tm-title tm-title_h1").Text()
	body := pageHTML.Find(".tm-article-body").Text()
	publisedAt := pageHTML.Find(".tm-article-datetime-published").Text()

	fmt.Println("publisedAt")
	fmt.Println(publisedAt)

	// _, err := time.Parse("2006-01-02, 03:04AM", publisedAt)
	// if err != nil {
	// 	return entities.Article{}, err
	// }

	return entities.Article{
		Author:       author,
		Title:        title,
		ProviderSign: providerSign,
		URL:          urlPage,
		Body:         body,
		PublishedAt:  time.Now(),
	}, nil

}

func (p *ParserHabr) GetArticleUrls(numPages int, urlCategory string) []string {

	articlesUrls := make([]string, 0)

	for page := 0; page < numPages; page++ {

		if page == 0 {
			articlesUrls = append(articlesUrls, urlCategory)
		} else {
			urlCat := urlCategory + "page" + strconv.Itoa(page) + "/"
			articlesUrls = append(articlesUrls, urlCat)
		}

	}

	return articlesUrls

}

func (p *ParserHabr) ParseLoop(data *entities.DataForParsing) (*entities.DataForParsing, error) {

	urlCategory := data.UrlCategory

	pageHTML, err := p.MakeRequest(urlCategory)
	if err != nil {
		return nil, err
	}

	numPages := p.GetNumPages(pageHTML)
	if numPages == "" {
		p.logger.Warn("Failed to get num pages")
	}

	numPagesInt, err := strconv.Atoi(numPages)
	if err != nil {
		p.logger.Warn("Failed to cast string numPages to int", zap.Error(err))
		return nil, err
	}

	urlPages := p.GetArticleUrls(numPagesInt, urlCategory)

	for _, url := range urlPages[0:2] {
		parsedArticle, err := p.ParseArticle(url, data.Provider)
		if err != nil {
			fmt.Print(err)
		}

		data.Articles = append(data.Articles, parsedArticle)

	}

	fmt.Println(urlPages[0:2])

	return data, nil

}
