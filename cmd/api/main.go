package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/gorilla/mux"
)

type GiphyResponse struct {
	Data []struct {
		Type             string        `json:"type"`
		ID               string        `json:"id"`
		URL              string        `json:"url"`
		Slug             string        `json:"slug"`
		BitlyGifURL      string        `json:"bitly_gif_url"`
		BitlyURL         string        `json:"bitly_url"`
		EmbedURL         string        `json:"embed_url"`
		Username         string        `json:"username"`
		Source           string        `json:"source"`
		Title            string        `json:"title"`
		Rating           string        `json:"rating"`
		ContentURL       string        `json:"content_url"`
		Tags             []interface{} `json:"tags"`
		FeaturedTags     []interface{} `json:"featured_tags"`
		UserTags         []interface{} `json:"user_tags"`
		SourceTld        string        `json:"source_tld"`
		SourcePostURL    string        `json:"source_post_url"`
		IsSticker        int           `json:"is_sticker"`
		ImportDatetime   string        `json:"import_datetime"`
		TrendingDatetime string        `json:"trending_datetime"`
		Images           struct {
			Original struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
				Frames   string `json:"frames"`
				Hash     string `json:"hash"`
			} `json:"original"`
			Downsized struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized"`
			DownsizedLarge struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized_large"`
			DownsizedMedium struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized_medium"`
			DownsizedSmall struct {
				Height  string `json:"height"`
				Width   string `json:"width"`
				Mp4Size string `json:"mp4_size"`
				Mp4     string `json:"mp4"`
			} `json:"downsized_small"`
			DownsizedStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized_still"`
			FixedHeight struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_height"`
			FixedHeightDownsampled struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_height_downsampled"`
			FixedHeightSmall struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_height_small"`
			FixedHeightSmallStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"fixed_height_small_still"`
			FixedHeightStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"fixed_height_still"`
			FixedWidth struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_width"`
			FixedWidthDownsampled struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_width_downsampled"`
			FixedWidthSmall struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_width_small"`
			FixedWidthSmallStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"fixed_width_small_still"`
			FixedWidthStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"fixed_width_still"`
			Looping struct {
				Mp4Size string `json:"mp4_size"`
				Mp4     string `json:"mp4"`
			} `json:"looping"`
			OriginalStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"original_still"`
			OriginalMp4 struct {
				Height  string `json:"height"`
				Width   string `json:"width"`
				Mp4Size string `json:"mp4_size"`
				Mp4     string `json:"mp4"`
			} `json:"original_mp4"`
			Preview struct {
				Height  string `json:"height"`
				Width   string `json:"width"`
				Mp4Size string `json:"mp4_size"`
				Mp4     string `json:"mp4"`
			} `json:"preview"`
			PreviewGif struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"preview_gif"`
			PreviewWebp struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"preview_webp"`
			Four80WStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"480w_still"`
		} `json:"images"`
		User struct {
			AvatarURL   string `json:"avatar_url"`
			BannerImage string `json:"banner_image"`
			BannerURL   string `json:"banner_url"`
			ProfileURL  string `json:"profile_url"`
			Username    string `json:"username"`
			DisplayName string `json:"display_name"`
			IsVerified  bool   `json:"is_verified"`
		} `json:"user"`
		AnalyticsResponsePayload string `json:"analytics_response_payload"`
		Analytics                struct {
			Onload struct {
				URL string `json:"url"`
			} `json:"onload"`
			Onclick struct {
				URL string `json:"url"`
			} `json:"onclick"`
			Onsent struct {
				URL string `json:"url"`
			} `json:"onsent"`
		} `json:"analytics"`
	} `json:"data"`
	Pagination struct {
		TotalCount int `json:"total_count"`
		Count      int `json:"count"`
		Offset     int `json:"offset"`
	} `json:"pagination"`
	Meta struct {
		Status     int    `json:"status"`
		Msg        string `json:"msg"`
		ResponseID string `json:"response_id"`
	} `json:"meta"`
}

type RecipesResponseData struct {
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Link        string   `json:"link"`
	Gif         string   `json:"gif"`
}
type RecipesResponse struct {
	Keywords []string              `json:"keywords"`
	Recipes  []RecipesResponseData `json:"recipes"`
}

type RecipePuppyResponse struct {
	Title   string  `json:"title"`
	Version float64 `json:"version"`
	Href    string  `json:"href"`
	Results []struct {
		Title       string `json:"title"`
		Href        string `json:"href"`
		Ingredients string `json:"ingredients"`
		Thumbnail   string `json:"thumbnail"`
	} `json:"results"`
}

func getRecipePuppy(p string) RecipePuppyResponse {
	resp, err := http.Get(fmt.Sprintf("http://www.recipepuppy.com/api/?i=%s", p))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var recipePuppyResponse RecipePuppyResponse
	err = json.Unmarshal(bodyBytes, &recipePuppyResponse)
	if err != nil {
		log.Fatalln(err)
	}

	return recipePuppyResponse
}

func getGiphyLink(query string) string {
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

func joinResponses(pr RecipePuppyResponse, s []string) RecipesResponse {
	var recipes []RecipesResponseData
	for _, puppyResponse := range pr.Results {
		responseData := RecipesResponseData{
			Title:       puppyResponse.Title,
			Ingredients: strings.Split(puppyResponse.Ingredients, ","),
			Link:        puppyResponse.Href,
			Gif:         getGiphyLink(puppyResponse.Title),
		}

		recipes = append(recipes, responseData)
	}

	return RecipesResponse{
		Keywords: s,
		Recipes:  recipes,
	}
}
func main() {
	router := mux.NewRouter().StrictSlash(true)

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func setupRouter(router *mux.Router) {
	router.
		Methods("GET").
		Path("/recipes").
		HandlerFunc(postFunction)
}

func postFunction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := r.URL.Query()["i"][0]

	s := strings.Split(p, ",")
	if len(s) > 3 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Please try 3 ingredients at maximum."})
		return
	}

	rp := getRecipePuppy(p)

	sort.Slice(rp.Results, func(i, j int) bool {
		return rp.Results[i].Title < rp.Results[j].Title
	})

	response := joinResponses(rp, s)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

type ErrorResponse struct {
	Error string `json:"error"`
}
