package utils

import (
	"cricli/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

func GetAllItems(page *model.Page) (items []model.Item) {
	return page.Rss.Channel.Item
}

func ExitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func GetScore(team string, items []model.Item) {
	for _, match := range items {
		fmt.Println(match.Title)
	}
}
