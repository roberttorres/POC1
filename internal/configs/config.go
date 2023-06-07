package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DBName string
}

func init() {
	//viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "pocdb.cluster-cwturhaakaag.sa-east-1.rds.amazonaws.com")
	viper.SetDefault("database.port", "3306")
	//viper.SetDefault("database.user", "admin")
	//viper.SetDefault("database.pass", "admin1234")
	//viper.SetDefault("database.dbname", "pocDb")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:   viper.GetString("database.host"),
		Port:   viper.GetString("database.port"),
		User:   viper.GetString("database.user"),
		Pass:   viper.GetString("database.pass"),
		DBName: viper.GetString("database.dbname"),
	}

	return nil

}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
