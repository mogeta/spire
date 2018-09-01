package spire

import (
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
	"time"
	"encoding/json"
)

const URL = "https://app.spire.io/api/v2/streaks?"

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

type Query struct {
	date string
}

type queryOption func(*Query)

func NewClient(accessToken string) Client {

	return Client{
		accessToken:accessToken,
	}
}

func WithTime(t time.Time) queryOption {
	return func(q *Query) {
		q.date =  t.Format("20060102")
	}
}

func (c Client)fetch(opt ...queryOption) *Spires{
	query := Query{}
	for _, o := range opt {
		o(&query)
	}

	values := url.Values{}
	values.Add("access_token", c.accessToken)

	if query.date != "" {
		values.Add("date",query.date)
	}

	resp, err := http.Get(URL + values.Encode())
	if err != nil {
		fmt.Println(err)

	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	spireData := new(Spires)
	err = json.Unmarshal(body,spireData)
	if err != nil {
		fmt.Errorf("%s",err)
	}

	return spireData
}


