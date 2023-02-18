package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// omdbapi.com API key
const OMDB_API_KEY = "2616241e"

const OMDB_URL = "http://www.omdbapi.com/?"

// MovieInfo Define Struct for response of omdbapi.com
// only necessary fields are defined
// `json:{field_name}` parse response JSON to golang struct member
type MovieInfo struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	ImdbRating string `json:"imdbRating"`
	ImdbID     string `json:"imdbID"`
}

func sendGetRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// close the stream of response body before exiting function
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return string(body), errors.New(resp.Status)
	}
	return string(body), nil
}

func SearchByTitle(title string) (*MovieInfo, error) {
	// http://www.omdbapi.com/?apikey=[yourkey]&t=[title]
	params := url.Values{}
	params.Set("apikey", OMDB_API_KEY)
	params.Set("t", title)

	requestUrl := OMDB_URL + params.Encode()
	body, err := sendGetRequest(requestUrl)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody: " + body)
	}

	ret := &MovieInfo{}
	return ret, json.Unmarshal([]byte(body), ret)
}

func SearchById(id string) (*MovieInfo, error) {
	// http://www.omdbapi.com/?apikey=[yourkey]&i=[id]
	params := url.Values{}
	params.Set("apikey", OMDB_API_KEY)
	params.Set("i", id)

	requestUrl := OMDB_URL + params.Encode()
	body, err := sendGetRequest(requestUrl)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody: " + body)
	}

	ret := &MovieInfo{}
	return ret, json.Unmarshal([]byte(body), ret)
}

func main() {
	fmt.Println("Search By Id")
	bodyFromId, err := SearchById("tt3896198")
	if err != nil {
		fmt.Println("Error in SearchById :", err.Error())
	} else {
		fmt.Printf("%+v\n", bodyFromId)
	}

	fmt.Println("Search By Title")
	bodyFromTitle, err := SearchByTitle("Game of")
	if err != nil {
		fmt.Println("Error in SearchById :", err.Error())
	} else {
		fmt.Printf("%+v\n", bodyFromTitle)
	}
}
