package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Rdb         string
	RdbHost     string
	RdbPort     string
	RdbUser     string
	RdbPassWord string
	RdbName     int
	RdbPoolSize int

	TENCENTCLOUD_SECRET_ID  string
	TENCENTCLOUD_SECRET_KEY string
	SdkAppId                string
	SignName                string
	SmsId                   string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadRedis(file)
	LoadContent(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("123qwe456asd789zxc")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("1234")
	DbName = file.Section("database").Key("DbName").MustString("tracer")
}

func LoadRedis(file *ini.File) {
	Rdb = file.Section("redis").Key("Rdb").MustString("redis")
	RdbHost = file.Section("redis").Key("RdbHost").MustString("localhost")
	RdbPort = file.Section("redis").Key("RdbPort").MustString("6379")
	RdbUser = file.Section("redis").Key("RdbUser").MustString("")
	RdbPassWord = file.Section("redis").Key("RdbPassWord").MustString("")
	RdbName = file.Section("redis").Key("RdbName").MustInt(0)
	RdbPoolSize = file.Section("redis").Key("RdbPoolSize").MustInt(100)
}

func LoadContent(file *ini.File) {
	TENCENTCLOUD_SECRET_ID = file.Section("tencent").Key("TENCENTCLOUD_SECRET_ID").MustString("")
	TENCENTCLOUD_SECRET_KEY = file.Section("tencent").Key("TENCENTCLOUD_SECRET_KEY").MustString("")
	SdkAppId = file.Section("tencent").Key("SdkAppId").MustString("")
	SignName = file.Section("tencent").Key("SignName").MustString("")
	SmsId = file.Section("tencent").Key("SmsId").MustString("")
}
