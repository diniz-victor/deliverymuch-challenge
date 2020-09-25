package apigiphy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//GetGiphyLink get the gif link
func GetGiphyLink(query string) string {
	q := url.QueryEscape(query)
	l := 1
	k := "LyKGPGg76vk1mQ0jJM7mevHym3KM6NVC"
	url := fmt.Sprintf("http://api.giphy.com/v1/gifs/search?q=%s&limit=%v&api_key=%s", q, l, k)

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
