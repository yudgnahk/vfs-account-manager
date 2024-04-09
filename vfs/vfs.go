package vfs

import (
	"fmt"
	httputils "github.com/yudgnahk/go-common-utils/http"
	"github.com/yudgnahk/vfs-account-manager/utils"
	"net/http"
	"time"
)

type Client struct {
	Session string
}

func NewClient(session string) *Client {
	return &Client{
		Session: session,
	}
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9,vi;q=0.8,ru;q=0.7")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Length", "0")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Origin", "https://trading.vfs.com.vn")
	req.Header.Add("Referer", "https://trading.vfs.com.vn/member/default")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("Cookie", fmt.Sprintf("SESSION=%s", c.Session))
}

func (c *Client) GetPortfolio() (*GetPortfolioResponse, error) {
	url := "https://trading.vfs.com.vn/rest/front/api/findPortfolio?custNo=&subAccoNo=NONE&alloDate=%s&secCd="
	url = fmt.Sprintf(url, utils.GetTimeFormat(time.Now()))

	request, err := httputils.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	c.setHeaders(request)

	var response GetPortfolioResponse
	err = httputils.Execute(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetTodayTransaction(subAccoNo string) (*GetTodayTransactionResponse, error) {
	url := "https://trading.vfs.com.vn/rest/front/api/findPortTransaction?subAccoNo=%s&secCd=&fromDate=%s&toDate=%s&tradeType=-1"
	url = fmt.Sprintf(url, subAccoNo, utils.GetTimeFormat(time.Now()), utils.GetTimeFormat(time.Now()))

	request, err := httputils.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	c.setHeaders(request)

	var response GetTodayTransactionResponse
	err = httputils.Execute(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
