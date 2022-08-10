package db

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/mgo.v2"
	"ttidl1/conf"
)

type MongoClient struct {
	session *mgo.Session
	valid   bool
	// map[string]string: subject2dbname
}

func (s *MongoClient) getSession() *mgo.Session {
	return s.session
}

func (s *MongoClient) GetSubjectDb(subject string) *mgo.Database {
	ses := s.getSession()
	return ses.DB("klx_xmath")
}

var mongoClient MongoClient

func NewMongoClient(server conf.ConfigServer) MongoClient {
	if mongoClient.valid == true {
		return mongoClient
	}
	db_url := server.DbHost
	hlog.Infof("connect mongodb: %s", db_url)
	mdc, err := mgo.Dial(db_url)
	if err == nil {
		hlog.Info("connect mongodb success")
		mongoClient.session = mdc
		mongoClient.valid = true
	} else {
		hlog.Info("connect mongodb fail")
		panic(fmt.Sprintf("connect mongodb: %s fail", db_url))
	}
	return mongoClient
}
