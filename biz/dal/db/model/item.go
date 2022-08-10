package model

import (
	"gopkg.in/mgo.v2/bson"
	dal "ttidl1/biz/dal/db"
)

type DbTaggedItem struct {
	Id      bson.ObjectId `bson:"_id"`
	Simhash string        `bson:"simhash"`
	Data    struct {
		Type       int    `bson:"type"`
		Subtype    int    `bson:"subtype"`
		Difficulty int    `bson:"difficulty"`
		Stem       string `bson:"stem"`
		Qs         []struct {
			TagIds        []bson.ObjectId   `bson:"tag_ids"`
			SubTagIds     []bson.ObjectId   `bson:"tag_sub_ids"`
			OptsTagIds    [][]bson.ObjectId `bson:"opts_tag_ids"`
			OptsSubTagIds [][]bson.ObjectId `bson:"opts_sub_tag_ids"`
			Desc          string            `bson:"desc"`
			Ans           interface{}       `bson:"ans"`
			Qs            []struct {
				Desc          string            `bson:"desc"`
				TagIds        []bson.ObjectId   `bson:"tag_ids"`
				SubTagIds     []bson.ObjectId   `bson:"tag_sub_ids"`
				OptsTagIds    [][]bson.ObjectId `bson:"opts_tag_ids"`
				OptsSubTagIds [][]bson.ObjectId `bson:"opts_sub_tag_ids"`
			}
		}
		TagIds    []bson.ObjectId `bson:"tag_ids"`
		SubTagIds []bson.ObjectId `bson:"tag_sub_ids"`
	} `bson:"data"`
	UserAdd   bool          `bson:"user_add"`
	RefItemId bson.ObjectId `bson:"ref_item_id"`
}

var DbTaggedFileds = bson.M{
	"_id":             1,
	"user_add":        1,
	"ref_item_id":     1,
	"data.difficulty": 1, "data.tag_ids": 1,
	"data.stem":        1,
	"data.qs.desc":     1,
	"data.qs.qs.desc":  1,
	"data.qs.ans":      1,
	"data.tag_sub_ids": 1, "data.qs.tag_ids": 1,
	"data.qs.tag_sub_ids": 1, "data.qs.opts_tag_ids": 1,
	"data.qs.opts_sub_tag_ids": 1, "data.type": 1,
	"data.subtype": 1,
}

type DbItemModel struct {
	mdc dal.MongoClient
}

func NewDbItemModel(mdc dal.MongoClient) DbItemModel {
	ret := DbItemModel{mdc: mdc}
	return ret
}

func (s *DbItemModel) FindOneById(subject string, item_id bson.ObjectId) DbTaggedItem {
	db := s.mdc.GetSubjectDb(subject)
	var ret DbTaggedItem
	_ = db.C("items").FindId(item_id).Select(DbTaggedFileds).One(&ret)
	return ret
}
