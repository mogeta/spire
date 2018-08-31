package spire

import (
	"testing"
	"os"
	"github.com/spf13/viper"
	"fmt"
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
	spire_io.NewClient(accessToken)
}