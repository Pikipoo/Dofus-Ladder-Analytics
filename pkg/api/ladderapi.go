package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

func parseLadderHTML(res *http.Response) {
	doc, err := html.Parse(res.Body)
	if err != nil {
		return
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			fmt.Println(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

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
