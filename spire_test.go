package spire

import (
	"testing"
	"os"
	"github.com/spf13/viper"
	"fmt"
	"time"
)

func getConfig(filename string){
	viper.SetConfigName(filename)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s \n", err))
	}
}

func TestMain(m *testing.M) {

	getConfig("config")
	code := m.Run()

	os.Exit(code)
}

func Test(t *testing.T) {

	accessToken := viper.GetString("access_token")
	client := NewClient(accessToken)
	spires := client.fetch(WithTime(time.Date(2018,8,30,0,0,0,0,time.UTC)))
	for key, value := range *spires {
		fmt.Println(key)
		fmt.Println(value)
	}
}