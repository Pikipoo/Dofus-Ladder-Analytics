package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

func parseLadderHTML(res *http.Response) {
	doc := html.NewTokenizer(res.Body)
	for tokenType := doc.Next(); tokenType != html.ErrorToken; {
		token := doc.Token()
		if tokenType == html.TextToken && len(token.Data) > 1 {
			fmt.Println(token.Data)
		}
		tokenType = doc.Next()
	}

	defer res.Body.Close()
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
	parseLadderHTML(res)
	/* 	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
		} */
	return nil, nil
}
