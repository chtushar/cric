package main

import (
	"bytes"
	"cric-score/model"
	"cric-score/utils"
	"encoding/json"
	"fmt"
	"os"

	xj "github.com/basgys/goxml2json"
	"github.com/ttacon/chalk"
)

const LIVE_SCORES = "http://static.espncricinfo.com/rss/livescores.xml"

func main() {
	xmlBytes, _ := utils.GetXML(LIVE_SCORES)
	data := bytes.NewReader(xmlBytes)
	js, err := xj.Convert(data)

	if err != nil {
		fmt.Println("Coudn't fetch the Live scores.")
		utils.ExitGracefully(err)
	}

	page := &model.Page{}
	err = json.Unmarshal(js.Bytes(), page)

	if err != nil {
		fmt.Println("Couldn't fetch the Live scores.")
	}

	items := &[]model.Item{}
	*items = utils.GetAllItems(page)

	if err != nil {
		fmt.Println("Couldn't fetch the Live scores.")
	} else {
		if len(os.Args) < 2 {
			fmt.Printf(chalk.Green.Color("🟩 Currently live: \n\n\n"))
			utils.PrintLiveMatches(*items)
			} else {
			team := os.Args[1]
			utils.GetScore(team, *items)
		}
	}
}
