package settings

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/yookoala/realpath"
	"gopkg.in/ini.v1"
	"os"
)

var Logger *logrus.Logger
var Config *ini.Section

func init() {
	initLogger()

	loadConfig()
}

func loadConfig() {
	file := flag.String("c", "etc/app.ini", "config file")
	flag.Parse()
	filePath, err := realpath.Realpath(*file)
	if err != nil || !fileExists(filePath) {
		Logger.WithFields(logrus.Fields{"file_path": filePath}).Fatal("config file not found")
		os.Exit(1)
	}
	iniConf, err := ini.Load(filePath)
	if err != nil {
		Logger.WithFields(logrus.Fields{"err": err}).Fatal("load env file failed")
		os.Exit(1)
	}
	env := getEnv("ENV", "dev")
	Config = iniConf.Section(env)
	Logger.WithFields(logrus.Fields{"file_path": filePath, "env": env}).Info("load config file success")
}

func initLogger() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.Out = os.Stdout
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Val(key string) string {
	return Config.Key(key).Value()
}

func Int(key string) int {
	res, _ := Config.Key(key).Int()
	return res
}

func Int64(key string) int64 {
	res, _ := Config.Key(key).Int64()
	return res
}
