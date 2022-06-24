package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"test-books-wishlist/internal/entity"
	"time"
)

const (
	booksURL   = "%s/volumes?key=%s&q=%s&maxResults=%v"
	maxResults = 20
)

// Service interface contains main functions
type Service interface {
	SetHostAndKey(host, accessKey string)
	SearchBooks(terms string) (*entity.GoogleBooksResponse, error)
}

//Handle struct for vcc handle for url
type Handle struct {
	CurrencyHost           string
	AccessKey              string
	lastSuccessfulSyncTime time.Time
}

//NewHandle function to create new Handle object
func NewHandle() *Handle {
	return &Handle{}
}

func (v *Handle) SetHostAndKey(host, accessKey string) {
	v.CurrencyHost = host
	v.AccessKey = accessKey
}

func (v *Handle) SearchBooks(terms, apiKey string) (*entity.GoogleBooksResponse, error) {
	url := fmt.Sprintf(booksURL, v.CurrencyHost, apiKey, terms, maxResults)
	return v.requestBooks(url)
}

func (v *Handle) requestBooks(url string) (*entity.GoogleBooksResponse, error) {
	options := &Options{
		Method: http.MethodGet,
		URL:    url,
	}

	result, err := RequestJSON(options)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	var booksBody entity.GoogleBooksResponse
	var booksBodyErr entity.GoogleBooksError

	if result.StatusCode != 200 {
		if err = json.Unmarshal(body, &booksBodyErr); err != nil {
			fmt.Printf("Error: %v\n", err)
			return nil, err
		}

		return nil, errors.New(booksBodyErr.Error.Message)
	}

	if err = json.Unmarshal(body, &booksBody); err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return &booksBody, nil
}
