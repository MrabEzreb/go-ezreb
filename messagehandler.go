package ezreb

import (
	"github.com/nlopes/slack"
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
