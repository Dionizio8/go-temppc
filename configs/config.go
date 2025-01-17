package configs

import "github.com/spf13/viper"

type conf struct {
	WebServerPort          string `mapstructure:"WEB_SERVER_PORT"`
	ViaCEPClientURL        string `mapstructure:"VIA_CEP_CLIENT_URL"`
	WeatherAPIClientURL    string `mapstructure:"WEATHER_API_CLIENT_URL"`
	WeatherAPIClientAPIKey string `mapstructure:"WEATHER_API_CLIENT_API_KEY"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
