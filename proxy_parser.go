package proxylist

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetProxyList is
func GetProxyList() ([]Proxy, error) {
	client := &http.Client{}
	r, _ := http.NewRequest("GET", "https://www.proxynova.com/proxy-server-list/", nil)

	response, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status : %v", response.StatusCode)
	}

	proxyList := make([]Proxy, 0)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("#tbl_proxy_list > tbody > tr").Each(func(i int, s *goquery.Selection) {
		// document.write('169.57.1.84');
		ip := s.Find("td:nth-child(1) > abbr > script").Text()
		ip = strings.TrimSpace(ip)

		// 8123
		port := s.Find("td:nth-child(2)").Text()
		port = strings.TrimSpace(port)

		// 1649 ms
		speedStr := s.Find("td:nth-child(4) > small").Text()
		speedStr = strings.TrimSpace(speedStr)

		// elite, transparent, anonymous
		anonymity := s.Find("td:nth-child(7) > span").Text()
		anonymity = strings.TrimSpace(anonymity)
		anonymity = strings.ToLower(anonymity)

		if len(ip) == 0 || len(port) == 0 || len(speedStr) == 0 || len(anonymity) == 0 {
			return
		}

		if len(strings.Split(ip, "'")) < 1 {
			return
		}

		ip = strings.Split(ip, "'")[1]
		speedStr = strings.Split(speedStr, " ")[0]
		speed, err := strconv.Atoi(speedStr)
		if err != nil {
			return
		}

		proxyList = append(proxyList, newProxy(ip, port, anonymity, speed))
	})

	sort.Slice(proxyList, func(i, j int) bool {
		return proxyList[i].speed < proxyList[j].speed
	})

	return proxyList, nil
}
