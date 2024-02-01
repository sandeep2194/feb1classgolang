/*
In this example, we've developed an API that connects to Google and performs searches. This API can be integrated into any application, enabling searches without opening a browser. Our emphasis is on building real-time APIs for diverse projects.

Before running this code, ensure you've executed the following command to install the required package:
go get -u github.com/PuerkitoBio/goquery
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

// SearchResult struct represents a single search result
type SearchResult struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"URL"`
}

// SearchResponse struct represents the response containing search results
type SearchResponse struct {
	Results []SearchResult `json:"results"`
}

// searchGoogle function performs a search on Google
func searchGoogle(query string) ([]SearchResult, error) {
	// Construct the Google URL
	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", strings.ReplaceAll(query, " ", "+"))
	// Fetch the Google page
	doc, err := goquery.NewDocument(searchURL)
	if err != nil {
		return nil, err
	}
	// Extract search results
	var results []SearchResult
	doc.Find(".tF2Cxc").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".BVGONb").Text()
		description := s.Find(".aC0pRe").Text()
		url, _ := s.Find(".BVGONb a").Attr("href")

		result := SearchResult{
			Title:       title,
			Description: description,
			URL:         url,
		}
		results = append(results, result)
	})
	return results, nil
}

// searchHandler method processes search requests
func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing 'q' parameter", http.StatusBadRequest)
		return
	}

	results, err := searchGoogle(query)
	if err != nil {
		log.Printf("Error searching Google: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := SearchResponse{Results: results}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func example3ApiDev() {
	router := mux.NewRouter()
	router.HandleFunc("/search", searchHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
