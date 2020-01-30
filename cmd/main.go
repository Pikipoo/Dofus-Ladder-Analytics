package main

import (
	"fmt"
	"log"

	"github.com/Pikipoo/Dofus-Ladder-Analytics/pkg/api"
)

func main() {
	res, err := api.GetRankingByClass(api.Solo, api.GetSeasonID(6), api.Roublard)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", res)
}
