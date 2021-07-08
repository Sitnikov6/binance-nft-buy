package main

import (
	"github.com/1makarov/binance-nft-buy/pkg/binance"
	binanceapi "github.com/1makarov/binance-nft-buy/pkg/binance-api"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalln(err)
	}

	headers := binanceapi.Headers{
		ClientType: viper.GetString("headers.clientType"),
		Cookie:     viper.GetString("headers.cookie"),
		CsrfToken:  viper.GetString("headers.csrfToken"),
		UserAgent:  viper.GetString("headers.userAgent"),
	}

	box := binance.MysteryBox{
		Productid: viper.GetString("mysterybox.id"),
		Volume:    viper.GetInt("mysterybox.amount"),
	}

	time := viper.GetString("mysterybox.time")

	proxy := viper.GetString("proxy")

	api := binanceapi.CreateApi(proxy, headers)
	client := binance.CreateClient(api, box, time)

	client.Start()
}

func initConfigs() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}