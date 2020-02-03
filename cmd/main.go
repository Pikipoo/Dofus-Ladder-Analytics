package main

import (
	"log"

	"github.com/Pikipoo/Dofus-Ladder-Analytics/pkg/api"
)

func main() {
	_, err := api.GetRankingByClass(api.SOLO, api.GetSeasonID(6), api.ROUBLARD)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s", res)
}
