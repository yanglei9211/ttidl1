package conf

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/ini.v1"
	"os"
	"strings"
)

type subject struct {
	DbName string
}

type ConfigServer struct {
	Address  string
	Port     int
	DbUrl    string
	RedisUrl string
	Amqp     string
	Subjects map[string]*subject
	source   *ini.File
}

var serverConfig ConfigServer

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (s *ConfigServer) Load(path string) *ConfigServer {
	exists, err := PathExists(path)
	if !exists {
		return s
	}
	s.source, err = ini.Load(path)
	if err != nil {
		panic(err)
	}
	return s
}

func (s *ConfigServer) Init() *ConfigServer {
	//判断配置是否加载成功
	if s.source == nil {
		return s
	}
	s.Address = s.source.Section("server").Key("address").MustString("0.0.0.0")
	s.Port = s.source.Section("server").Key("port").MustInt(8080)
	s.DbUrl = s.source.Section("server").Key("db_url").MustString("")
	s.RedisUrl = s.source.Section("server").Key("redis_url").MustString("")
	s.Amqp = s.source.Section("server").Key("ampq").MustString("")
	subject_str := s.source.Section("server").Key("subjects").MustString("")
	subject_names := strings.Split(subject_str, ",")
	s.Subjects = make(map[string]*subject)
	for _, subject_name := range subject_names {
		db_name := s.source.Section(subject_name).Key("db_name").MustString("")
		curSubject := subject{db_name}
		s.Subjects[subject_name] = &curSubject
	}
	return s
}

func NewServerConf() ConfigServer {
	if serverConfig.source != nil {
		// 已加载,直接返回
		return serverConfig
	}
	hlog.Info("begin to load config")
	serverConfig = *(&ConfigServer{}).Load("conf/app.ini").Init()
	fmt.Println(serverConfig)
	return serverConfig
}

//func init() {
//	fmt.Println("init server config")
//	serverConfig = (&ConfigServer{}).Load("conf/app.ini").Init()
//}
