package ezreb

import (
	"github.com/nlopes/slack"
	"strings"
	"fmt"
	"time"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

func handleMessage(e slack.MessageEvent, rtm *slack.RTM) {
	//	rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("%+v", e), "C0T7WHNHW"))
	//channels := rtm.GetChannels()
	channel, _ := rtm.GetChannelInfo(e.Channel)
	if strings.Contains(e.Text, "say") {
		if strings.Contains(e.Text, "hello to") {
			var greeting string
			greeting = strings.Trim(strings.Split(e.Text, "hello to ")[1], "!?.,'\"")
			greeting = strings.Replace(greeting, "your", "東東", -1)
			greeting = strings.Replace(greeting, "my", "your", -1)
			greeting = strings.Replace(greeting, "東東", "my", -1)
			rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("Hello %s!", greeting), e.Channel))
		}
	}
	if strings.Contains(e.Text, "price") && channel.Name == "tf2-data" {
		rtm.SendMessage(rtm.NewOutgoingMessage("Please give me a moment...", channel.ID))
		CurrencyData := getCurrencyData()
		t := time.Now()
		params := slack.NewPostMessageParameters()
		//		outgoingMessage := rtm.NewOutgoingMessage("Backpack.tf Data "+t.Format(time.UnixDate), e.Channel)
		keys := slack.Attachment{
			Title:     "Mann Co. Supply Crate Keys",
			TitleLink: "http://backpack.tf/stats/Unique/Mann%20Co.%20Supply%20Crate%20Key/Tradable/Craftable",
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "Min Price",
					Value: fmt.Sprintf("%f", CurrencyData.Response.Currencies.Keys.Price.Value) + " " + CurrencyData.Response.Currencies.Keys.Price.Currency,
				},
				slack.AttachmentField{
					Title: "Max Price",
					Value: fmt.Sprintf("%f", CurrencyData.Response.Currencies.Keys.Price.ValueHigh) + " " + CurrencyData.Response.Currencies.Keys.Price.Currency,
				},
			},
		}
		buds := slack.Attachment{
			Title:     "Mann Co. Army Grade Earbuds",
			TitleLink: "http://backpack.tf/stats/Unique/Earbuds/Tradable/Craftable",
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "Min Price",
					Value: fmt.Sprintf("%f", CurrencyData.Response.Currencies.Earbuds.Price.Value) + " " + CurrencyData.Response.Currencies.Earbuds.Price.Currency,
				},
				slack.AttachmentField{
					Title: "Max Price",
					Value: fmt.Sprintf("%f", CurrencyData.Response.Currencies.Earbuds.Price.ValueHigh) + " " + CurrencyData.Response.Currencies.Earbuds.Price.Currency,
				},
			},
		}
		params.Attachments = []slack.Attachment{keys, buds}
		//		fmt.Printf("Sent to channel %s at %s", channelID, timestamp)
		rtm.PostMessage(channel.ID, "Backpack.tf Data "+t.Format(time.UnixDate), params)
	}
}

type webapiJSON struct {
	status     string
	statuscode int
	body       string
}

func getWebJSON(url string) (webapiJSON, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return *new(webapiJSON), err
	}
	robots, err := ioutil.ReadAll(res.Body)
	//fmt.Printf("%+v\n\n\n\n\n", res)
	apiJSON := webapiJSON{res.Status, res.StatusCode, string(robots[:])}
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return *new(webapiJSON), err
	}
	//fmt.Printf("%s", robots)
	return apiJSON, nil
}

type Currency struct {
	Response struct {
		Currencies struct {
			Earbuds struct {
				Blanket   int    `json:"blanket"`
				Craftable string `json:"craftable"`
				Defindex  int    `json:"defindex"`
				Plural    string `json:"plural"`
				Price     struct {
					Currency   string  `json:"currency"`
					Difference float64 `json:"difference"`
					Value      float64 `json:"value"`
					ValueHigh  float64 `json:"value_high"`
				} `json:"price"`
				Priceindex int    `json:"priceindex"`
				Quality    int    `json:"quality"`
				Round      int    `json:"round"`
				Single     string `json:"single"`
				Tradable   string `json:"tradable"`
			} `json:"earbuds"`
			Hat struct {
				Blanket   int    `json:"blanket"`
				Craftable string `json:"craftable"`
				Defindex  int    `json:"defindex"`
				Plural    string `json:"plural"`
				Price     struct {
					Currency   string  `json:"currency"`
					Difference float64 `json:"difference"`
					Value      float64 `json:"value"`
					ValueHigh  float64 `json:"value_high"`
				} `json:"price"`
				Priceindex int    `json:"priceindex"`
				Quality    int    `json:"quality"`
				Round      int    `json:"round"`
				Single     string `json:"single"`
				Tradable   string `json:"tradable"`
			} `json:"hat"`
			Keys struct {
				Blanket   int    `json:"blanket"`
				Craftable string `json:"craftable"`
				Defindex  int    `json:"defindex"`
				Plural    string `json:"plural"`
				Price     struct {
					Currency   string  `json:"currency"`
					Difference float64 `json:"difference"`
					Value      float64 `json:"value"`
					ValueHigh  float64 `json:"value_high"`
				} `json:"price"`
				Priceindex int    `json:"priceindex"`
				Quality    int    `json:"quality"`
				Round      int    `json:"round"`
				Single     string `json:"single"`
				Tradable   string `json:"tradable"`
			} `json:"keys"`
			Metal struct {
				Blanket   int    `json:"blanket"`
				Craftable string `json:"craftable"`
				Defindex  int    `json:"defindex"`
				Plural    string `json:"plural"`
				Price     struct {
					Currency   string  `json:"currency"`
					Difference float64 `json:"difference"`
					Value      float64 `json:"value"`
					ValueHigh  float64 `json:"value_high"`
				} `json:"price"`
				Priceindex int    `json:"priceindex"`
				Quality    int    `json:"quality"`
				Round      int    `json:"round"`
				Single     string `json:"single"`
				Tradable   string `json:"tradable"`
			} `json:"metal"`
		} `json:"currencies"`
		Name    string `json:"name"`
		Success int    `json:"success"`
		URL     string `json:"url"`
	} `json:"response"`
}

func getCurrencyData() Currency {
	var c Currency
	currencyData, _ := getWebJSON("http://backpack.tf/api/IGetCurrencies/v1/?key=56eaba18dea9e90b24c42c59")
	if currencyData.statuscode != 200 {
		panic(currencyData.statuscode)
	}
	err := json.Unmarshal([]byte(currencyData.body), &c)
	if err != nil {
		log.Fatal(err)
		return c
	}
	return c
}
