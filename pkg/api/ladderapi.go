package api

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func processData(dataChan <-chan *list.List) {
	for data := range dataChan {
		for e := data.Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value, " ")
		}
		fmt.Println()
	}
}

func parseLadderHTML(res *http.Response, dataChan chan<- *list.List) {
	doc := html.NewTokenizer(res.Body)
	defer res.Body.Close()

	for tokenType := doc.Next(); tokenType != html.ErrorToken; tokenType = doc.Next() {
		token := doc.Token()
		/* 		if tokenType == html.StartTagToken && token.Data == "td" {
		fmt.Println(token.Attr)
		} */
		if tokenType == html.StartTagToken && token.Data == "tr" {
			data := list.New()
			for tokenType := doc.Next(); tokenType != html.ErrorToken; tokenType = doc.Next() {
				token := doc.Token()
				if tokenType == html.EndTagToken && token.Data == "tr" {
					break
				}
				token.Data = strings.Trim(token.Data, "\n ")
				if tokenType == html.TextToken && token.Data != "" {
					data.PushBack(token.Data)
					// fmt.Println(token.Data)
				}
			}
			dataChan <- data
		}
	}
	defer close(dataChan)
}

// GetRankingByLeague gets PvP leaderboard for a given league
func GetRankingByLeague(queue queueType, season int, league leagueName) (ranking []byte, err error) {
	res, err := http.Get(baseURL + fmt.Sprintf(paramQueue, queue) +
		fmt.Sprintf(paramSeason, season) + fmt.Sprintf(paramLeague, league))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetRankingByClass gets PvP leaderboard for a given class
func GetRankingByClass(queue queueType, season int, class characterClass) (ranking []byte, err error) {
	res, err := http.Get(baseURL + fmt.Sprintf(paramQueue, queue) +
		fmt.Sprintf(paramSeason, season) + fmt.Sprintf(paramClass, class))
	if err != nil {
		return nil, err
	}

	dataChan := make(chan *list.List, 100)
	go parseLadderHTML(res, dataChan)
	processData(dataChan)

	/* 	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
		} */
	return nil, nil
}
