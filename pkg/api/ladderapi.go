package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func processDataWithClassFiltering(dataChan <-chan []string, class characterClass) []character {
	rankings := []character{}
	for data := range dataChan {
		characterData, err := newCharaterFromArray(data)
		if err == nil && characterData.class == class {
			rankings = append(rankings, *characterData)
		} else {
			log.Println(err)
		}
	}
	return rankings
}

func parseLadderHTML(res *http.Response, dataChan chan<- []string) {
	doc := html.NewTokenizer(res.Body)
	defer res.Body.Close()

	for tokenType := doc.Next(); tokenType != html.ErrorToken; tokenType = doc.Next() {
		token := doc.Token()
		if tokenType == html.StartTagToken && token.Data == "tr" {
			data := []string{}
			for tokenType := doc.Next(); tokenType != html.ErrorToken; tokenType = doc.Next() {
				token := doc.Token()
				if tokenType == html.EndTagToken && token.Data == "tr" {
					break
				}
				token.Data = strings.Trim(token.Data, "\n ")
				if tokenType == html.TextToken && token.Data != "" {
					data = append(data, token.Data)
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
func GetRankingByClass(queue queueType, season int, class characterClass) (rankings []character, err error) {
	res, err := http.Get(baseURL + fmt.Sprintf(paramQueue, queue) +
		fmt.Sprintf(paramSeason, season) + fmt.Sprintf(paramClass, CharacterClassMap[class]))
	if err != nil {
		return nil, err
	}

	dataChan := make(chan []string, 100)
	go parseLadderHTML(res, dataChan)
	rankings = processDataWithClassFiltering(dataChan, class)

	return rankings, nil
}
