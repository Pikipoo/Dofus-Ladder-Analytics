package main

import (
	"fmt"
	"log"

	"github.com/Pikipoo/Dofus-Ladder-Analytics/pkg/api"
)

func main() {
	res, err := api.GetRankingByClass(api.GROUP, api.GetSeasonID(6), api.ELIOTROPE)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
