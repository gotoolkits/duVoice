package main

import (
	"flag"
	"fmt"
	"os"

	auth "github.com/gotoolkits/duVoice/auth"
	tts "github.com/gotoolkits/duVoice/tts"
	"github.com/spf13/viper"
)

var defaultPath = "/tmp/tts/notify.mp3"

func main() {

	var context, outpath string
	flag.StringVar(&outpath, "o", defaultPath, "语音输出文件与路径")
	flag.StringVar(&context, "msg", "百度文件语音转换功能", "需转换的文本内容")
	flag.Parse()

	au, err := LoadJsonConf()

	if err != nil {
		fmt.Println("Get auth config file failed:", err)
		os.Exit(1)
	}

	duApi := tts.NewAPI_Util(au.Client_id, au.Client_secret)

	duApi.Text2AudioFile(outpath, context)

	fmt.Println("转换的文本内容为:", context)
	fmt.Println("转换语音的位置:", outpath)

}

func LoadJsonConf() (auth.Credentials_Request, error) {

	var au auth.Credentials_Request = auth.Credentials_Request{}

	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/duVoice/")
	err := viper.ReadInConfig()
	if err != nil {
		return auth.Credentials_Request{}, err
	} else {
		au.Client_id = viper.GetString("ApiKey")
		au.Client_secret = viper.GetString("SecretKey")
	}
	return au, nil
}
