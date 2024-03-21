package config

import "github.com/spf13/viper"

type Config struct{
	Port string `mapstructure:"PORT"`
	AuthService string `mapstructure:"AUTH_SRV"`
	Jwt string `mapstructure:"JWT"`
}

var envs =[]string{"AUTH_SRV","PORT"}

func LoadConfig() (config *Config,err error){
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	for _,env:=range envs{
		if err =viper.BindEnv(env);err != nil{
			return
		}
	}
	err=viper.Unmarshal(&config)
	return
}