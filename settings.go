package settings

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/yookoala/realpath"
	"gopkg.in/ini.v1"
	"os"
)

var Logger *logrus.Logger
var Config *ini.File

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
	Config, err = ini.Load(filePath)
	if err != nil {
		Logger.WithFields(logrus.Fields{"err": err}).Fatal("load env file failed")
		os.Exit(1)
	}
	Logger.WithFields(logrus.Fields{"file_path": filePath}).Info("load config file success")
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
