package utils

import (
	"github.com/indraprasetya154/golang-restful-api/helper"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	helper.PanicIfError(err)
}
