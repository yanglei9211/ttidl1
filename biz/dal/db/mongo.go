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
	subject2dbname map[string]string
}

func (s *MongoClient) getSession() *mgo.Session {
	return s.session
}

func (s *MongoClient) GetSubjectDb(subject string) *mgo.Database {
	ses := s.getSession()
	db_name := s.subject2dbname[subject]
	return ses.DB(db_name)
}

var mongoClient MongoClient

func NewMongoClient(server conf.ConfigServer) MongoClient {
	if mongoClient.valid == true {
		return mongoClient
	}
	db_url := server.DbUrl
	hlog.Infof("connect mongodb: %s", db_url)
	mdc, err := mgo.Dial(db_url)
	if err == nil {
		hlog.Info("connect mongodb success")
		mongoClient.session = mdc
		mongoClient.valid = true
	} else {
		hlog.Error("connect mongodb fail")
		panic(fmt.Sprintf("connect mongodb: %s fail", db_url))
	}
	mongoClient.subject2dbname = make(map[string]string)
	for k, v := range server.Subjects {
		mongoClient.subject2dbname[k] = v.DbName
	}
	return mongoClient
}
