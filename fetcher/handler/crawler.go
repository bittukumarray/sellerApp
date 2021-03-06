package handler

import (
	"bytes"
	"encoding/json"
	"fetcher/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func Fetch(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var u model.Url
	err = json.Unmarshal(body, &u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid data format"))
		return
	}
	if u.URL == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Url is empty"))
		return
	}

	c := colly.NewCollector()

	var p model.Product

	c.OnHTML("#productTitle", func(e *colly.HTMLElement) {
		data := strings.ReplaceAll(e.Text, "\n", "")
		p.Name = data
	})

	c.OnHTML("#imgTagWrapperId", func(e *colly.HTMLElement) {
		p.ImageUrl = e.ChildAttr("img", "data-old-hires")
	})

	c.OnHTML("#feature-bullets", func(e *colly.HTMLElement) {
		p.Description = e.ChildText("li > span.a-list-item")

	})

	c.OnHTML("#priceblock_ourprice", func(e *colly.HTMLElement) {
		p.Price = e.Text
	})

	c.OnHTML("#acrCustomerReviewText", func(e *colly.HTMLElement) {
		data := strings.Split(e.Text, " ")
		data[0] = strings.ReplaceAll(data[0], ",", "")
		val, err := strconv.Atoi(data[0])
		if err != nil {
			fmt.Errorf("error %v\n", err)
		}
		p.TotalReviews = val
	})

	c.Visit(u.URL)

	var res model.ResponseModel
	res.URL = u.URL
	res.Product = p

	d, err := json.Marshal(res)
	if err != nil {
		fmt.Println("err")
	}

	req, err := http.NewRequest(http.MethodPost, "http://writer-api:9000/product", bytes.NewBuffer(d))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while sending request"))
		return
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while getting response"))
		return
	}

	if resp.StatusCode == http.StatusCreated {
		w.WriteHeader(http.StatusCreated)
		w.Write(d)
	}

}
