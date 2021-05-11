package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var (
	title = flag.String("t", "", "Movie title to search for.\n eg: -t=The+Social+Network")
	key   = flag.String("key", "", "API Key to access.")
)

//eg: go run main.go -key=[your apikey] -t=The+Social+Network

const URL = "http://www.omdbapi.com/"

type resPoster struct {
	Poster string `json:"Poster"`
}

func main() {
	flag.Parse()
	finalUrl := URL + "?t=" + url.QueryEscape(*title) + "&apikey=" + url.QueryEscape(*key)
	resp, err := http.Get(finalUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http get url:%v failed. err: %v\n", finalUrl, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var res resPoster
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		fmt.Fprintf(os.Stderr, "Json decode failed. err: %v\n", err)
		os.Exit(1)
	}

	poster, err := http.Get(res.Poster)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http get poster:%v failed. err: %v\n", res.Poster, err)
		os.Exit(1)
	}

	defer poster.Body.Close()

	f, err := os.Create(*title + ".jpeg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "create file fail. err: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = io.Copy(f, poster.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write to file err:%v\n", err)
		os.Exit(1)
	}
	dir, err := os.Getwd()
	fmt.Printf("File was saved in %v\n", dir)
}
