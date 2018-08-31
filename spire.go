package spire_io

import (
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Client struct {
	accessToken string
}


type Spires []*SpireData

type SpireData struct {
	Type         string  `json:"type"`
	StartAt      int     `json:"start_at"`
	StopAt       int     `json:"stop_at"`
	Value        float64 `json:"value"`
	SubValue     float64 `json:"sub_value"`
	OriginalType string  `json:"original_type"`
	Comment      string  `json:"comment"`
	ModifiedType string  `json:"modified_type"`
	Modified     bool    `json:"modified"`
}


func NewClient(accessToken string) Client {

	return Client{
		accessToken:accessToken,
	}
}

func (c Client)fetchData(date string)[]byte{
	values := url.Values{}
	values.Add("access_token", c.accessToken)
	values.Add("date",date)

	resp, err := http.Get("https://app.spire.io/api/v2/streaks" + "?" + values.Encode())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	//println(string(body))
	return body
}

func (c Client)fetch(data []byte){
	spireData := new(Spires)
	err := json.Unmarshal(data,spireData)
	if err != nil {
		fmt.Errorf("%s",err)
	}

	for key, value := range *spireData {
		fmt.Println(key)
		fmt.Println(value)
	}
}
