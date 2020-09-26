package recipes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type giphy struct {
	endpoint string
	apiKey   string
}

//NewGiphy creates a new giphy
func NewGiphy(endpoint, apiKey string) GiphyGateway {
	return &giphy{
		endpoint: endpoint,
		apiKey:   apiKey,
	}
}

//Get get the gif link
func (g giphy) GetGiphy(query string) string {
	q := url.QueryEscape(query)
	l := 1
	url := fmt.Sprintf("%s?q=%s&limit=%v&api_key=%s", g.endpoint, q, l, g.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var giphyResponse GiphyResponse
	err = json.Unmarshal(bodyBytes, &giphyResponse)
	if err != nil {
		log.Fatalln(err)
	}
	if len(giphyResponse.Data) == 0 {
		return ""
	}
	return giphyResponse.Data[0].BitlyGifURL
}
