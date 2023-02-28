package main

import (
    "io/fs"
    "log"
    "os"
    "strings"
    "time"

    "github.com/anaskhan96/soup"
    "github.com/go-rod/rod"
    "github.com/go-rod/rod/lib/launcher"
    "github.com/go-rod/stealth"
)

func rbcParse() (RBCNews, error) {
    launcher.NewBrowser().MustGet()
    browser := rod.New().Timeout(time.Minute).MustConnect()
    defer browser.MustClose()

    page := stealth.MustPage(browser)

    page.MustNavigate("https://www.rbc.ru/").MustWaitLoad()

    page.MustElement("div > div > div.td1083bb7 > div.o65fb5a61").MustClick().MustWaitLoad()
    page.MustElement("body > div.live-tv-popup.js-live-tv-popup.active > div.live-tv-popup__head > div").MustClick().MustWaitLoad()
    page.MustElement("body > div.push-allow.js-push-allow > div.push-allow__block.js-push-allow-block.active > div.push-allow__controls > div:nth-child(2) > a").MustClick().MustWaitLoad()

    if DEBUG {
        page.MustScreenshotFullPage("")
        os.WriteFile("html1.html", []byte(page.MustHTML()), fs.ModePerm)
    }

    err := page.Mouse.Scroll(0, 1000000, 10)
    if err != nil {
        return RBCNews{}, err
    }

    if DEBUG {
        page.MustScreenshotFullPage("")
        os.WriteFile("html2.html", []byte(page.MustHTML()), fs.ModePerm)
    }

    html := page.MustHTML()

    news := parseHTML(html)

    return news, nil
}

func parseHTML(html string) RBCNews {
    news := RBCNews{}

    doc := soup.HTMLParse(html)

    mainNewsBlock := doc.FindStrict("div", "class", "main js-main-reload")
    if mainNewsBlock.Pointer != nil {
        mainNews := mainNewsBlock.FindStrict("div", "class", "main__big js-main-reload-item")
        if mainNews.Pointer != nil {
            tmp := mainNews.FindStrict("a", "class", "main__big__link js-yandex-counter")
            if tmp.Pointer != nil {
                url := tmp.Attrs()["href"]
                text := tmp.FullText()
                log.Println(url, strings.TrimSpace(text))
                news.MainNews = News{URL: url, Title: strings.TrimSpace(text)}
            }
        }
    }

    topNewsBlock := doc.FindStrict("div", "class", "main__list")
    if topNewsBlock.Pointer != nil {
        topNewsList := topNewsBlock.FindAllStrict("div", "class", "main__inner l-col-center")
        for _, topNews := range topNewsList {
            newsList := topNews.FindAllStrict("div", "class", "main__feed js-main-reload-item")
            for _, item := range newsList {
                tmp := item.FindStrict("a", "class", "main__feed__link js-yandex-counter js-visited")
                if tmp.Pointer != nil {
                    url := tmp.Attrs()["href"]
                    text := tmp.FullText()
                    log.Println(url, strings.TrimSpace(text))
                    news.TopNews = append(news.TopNews, News{URL: url, Title: strings.TrimSpace(text)})
                }
            }

        }
    }

    centralNewsBlock := doc.FindStrict("div", "class", "js-index-central-column")
    if centralNewsBlock.Pointer != nil {
        centralNewsList := centralNewsBlock.FindAll("div", "class", "js-index-doscroll")
        for _, item := range centralNewsList {
            tmp := item.Find("a", "class", "js-index-central-column-io")
            if tmp.Pointer != nil {
                url := tmp.Attrs()["href"]
                text := tmp.FullText()
                log.Println(url, strings.TrimSpace(text))
                news.CentralNews = append(news.CentralNews, News{URL: url, Title: strings.TrimSpace(text)})
            }
        }
    }

    return news
}
