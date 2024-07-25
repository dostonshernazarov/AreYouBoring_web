package boredapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Bored struct {
	Ok string
}

type BoredRes struct {
	Resultt string
}

type BoredBody struct {
	Activity      string
	Availability  float64
	Type          string
	Participants  int64
	Price         float32
	Accessibility string
	Duration      string
	KidFriendly   bool
	Link          string
	Key           string
}

type BoredServiceClient interface {
	GetTip(ctx context.Context, ok *Bored) (*BoredRes, error)
}

type boredServiceClient struct {
}

func NewBoredClient() BoredServiceClient {
	return &boredServiceClient{}
}

func (c *boredServiceClient) GetTip(ctx context.Context, ok *Bored) (*BoredRes, error) {

	url := "https://bored-api.appbrewery.com/random"
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}
	var rest BoredBody

	err = json.Unmarshal(body, &rest)
	if err != nil {
		log.Println("Error unmarshaling body:", err)
		return nil, err
	}

	return &BoredRes{
		Resultt: rest.Activity,
	}, nil
}
