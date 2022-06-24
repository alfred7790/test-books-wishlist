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

func (v *Handle) SearchBooks(terms string) (*entity.GoogleBooksResponse, error) {
	url := fmt.Sprintf(booksURL, v.CurrencyHost, v.AccessKey, terms, maxResults)
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

	if result.StatusCode != 200 {
		str := fmt.Sprintf("Attempt to query failed with status code %d", result.StatusCode)
		fmt.Println(str)
		return nil, errors.New(str)
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	var booksBody entity.GoogleBooksResponse
	var booksBodyErr entity.GoogleBooksError
	if err = json.Unmarshal(body, &booksBody); err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	if len(booksBody.Kind) == 0 {
		if err = json.Unmarshal(body, &booksBodyErr); err != nil {
			fmt.Printf("Error: %v\n", err)
			return nil, err
		}

		fmt.Println(booksBodyErr.Error)
		return nil, errors.New(booksBodyErr.Error.Message)
	}

	return &booksBody, nil
}
