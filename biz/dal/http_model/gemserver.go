package http_model

import (
	"fmt"
	"github.com/goccy/go-json"
	"ttidl1/conf"
)

type GemserverModel struct {
	Host      string
	Appid     string
	Username2 string
	Username3 string
	Username4 string
	Username  string
	valid     bool
}

var gemserverClient GemserverModel

type GemserverTask struct {
	url  string
	data QueryInterface
}

func NewGemserverModel(server conf.ConfigServer) GemserverModel {
	if gemserverClient.valid == true {
		return gemserverClient
	}
	gemserverClient.Host = server.GemserverHost
	gemserverClient.Appid = server.GemserverAppid
	gemserverClient.Username2 = "" // 暂无小学
	gemserverClient.Username3 = server.Gemserver3Username
	gemserverClient.Username4 = server.Gemserver4Username
	gemserverClient.Username = server.GemserverUsername
	gemserverClient.valid = true
	return gemserverClient
}

func (s GemserverTask) GetUrl() string {
	furl := fmt.Sprintf("%s%s", gemserverClient.Host, s.url)
	return furl
}

func (s GemserverTask) GetData() QueryInterface {
	return s.data
}

type ItemDetailsQuery struct {
	Appid    string
	Username string
	Subject  string
	ItemIds  []string
}

func (s ItemDetailsQuery) BuildParams() string {
	mp := map[string]interface{}{}
	mp["appid"] = s.Appid
	mp["username"] = s.Username
	mp["subject"] = s.Subject
	mp["item_ids"] = s.ItemIds
	mpJson, _ := json.Marshal(mp)
	ret := fmt.Sprintf("data=%s", mpJson)
	return ret
}

type ItemDetail struct {
	Id   string `json:"_id"`
	Data struct {
		Desc string `josn:"desc"`
		Stem string `json:"stem"`
		Html string `json:"html"`
		Qs   []struct {
			Desc string      `json:"desc"`
			Exp  string      `json:"exp"`
			Opts []string    `json:"opts"`
			Ans  interface{} `json:"ans"`
		} `json:"qs"`
	} `json:"data"`
}

type ItemDetailsResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Item []ItemDetail `json:"items"`
	} `json:"data"`
}

func (s *GemserverModel) GetItemDetails(subject string, item_id string) (ItemDetailsResp, error) {
	query := ItemDetailsQuery{
		ItemIds:  []string{item_id},
		Subject:  subject,
		Appid:    s.Appid,
		Username: s.Username, // 非必要不指定学段
	}
	url := fmt.Sprintf("/%s/pub/item/batch/details", subject)
	httpTask := GemserverTask{url: url, data: query}
	retBody, err := requestGet(httpTask)
	if err != nil {
		return ItemDetailsResp{}, err
	}
	var ret ItemDetailsResp
	_ = json.Unmarshal(retBody, &ret)
	return ret, nil
}
