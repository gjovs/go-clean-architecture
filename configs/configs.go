package configs

import "github.com/spf13/viper"

type config struct {
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	WebServerPort  string `mapstructure:"WWW_SERVER_PORT"`
	GRPCServerPort string `mapstructure:"GRPC_SERVER_PORT"`
}

func LoadConfig(path string) (*config, error) {
	var configuration *config
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	
	if err != nil {
		panic(err)
	}
	
	err = viper.Unmarshal(&configuration)
	
	if err != nil {
		panic(err)
	}
	
	return configuration, err
}
