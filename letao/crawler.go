package letao

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Filter struct {
	Category  string // 目標分類 ID
	ViewCount int    // 每頁顯示數量（最大為 40）
	IsRecent  bool   // 是否只顯示最新商品
}

type AuctionItem struct {
	URL      string // 商品連結
	ImageURL string // 商品圖片連結
	Title    string // 商品標題
	PriceMP  string // 價格資訊(mp)
	PriceM   string // 價格資訊(m)
	BidsInfo string // 出價資訊
	TimeInfo string // 時間資訊
}

var (
	searhFinalLink = ""
)

func LetaoCrawler(baseURL string, filter Filter) ([]AuctionItem, error) {
	// Limit viewcount to maximum 40
	viewCount := min(filter.ViewCount, 40)

	// Ensure BaseURL ends with /
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	// Build target URL
	url := fmt.Sprintf("%s?p=&category=%s&seller=&viewcount=%d", baseURL, filter.Category, viewCount)
	if filter.IsRecent {
		url += "&new=1"
	}

	respBody, err := sendGetRequest(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(respBody)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var items []AuctionItem
	shouldStop := false
	setURL := true

	doc.Find("div.item").Each(func(i int, s *goquery.Selection) {
		if shouldStop {
			return
		}

		href, exists := s.Find("div.imgInfo a").First().Attr("href")
		// filter non-product tags
		if !exists || strings.TrimSpace(href) == "" {
			return
		}
		href = "https:" + href

		if searhFinalLink != "" && href == searhFinalLink {
			shouldStop = true
			return
		}

		if setURL {
			searhFinalLink = href
			setURL = false
		}

		imgHref, exists := s.Find("div.imgInfo a img").First().Attr("src")
		if !exists {
			return // Skip items without link
		}
		// Get title text
		title := strings.TrimSpace(s.Find("div.titleInfo div.title a").First().Text())

		// Get price info
		priceMP := strings.TrimSpace(s.Find("div.priceInfo div.cbid div.mp").First().Text())
		priceM := strings.TrimSpace(s.Find("div.priceInfo div.cbid div.m").First().Text())

		// Get bids info
		bidsInfo := strings.TrimSpace(s.Find("div.bidsInfo").First().Text())

		// Get time info
		timeInfo := strings.TrimSpace(s.Find("div.timeInfo").First().Text())

		// Create auction item
		item := AuctionItem{
			URL:      href,
			ImageURL: imgHref,
			Title:    title,
			PriceMP:  priceMP,
			PriceM:   priceM,
			BidsInfo: bidsInfo,
			TimeInfo: timeInfo,
		}

		items = append(items, item)
	})

	return items, nil
}
