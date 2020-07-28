package config

import (
	"fmt"
	"golearn1/services/users/structs"

	"github.com/spf13/viper"
)

// S struct containing unmarshalled viper config
var S structs.DbCred

// Init configuring DB connection credentials
func Init() *viper.Viper {
	var V = viper.New()
	V.SetConfigName("dbconfig")
	V.SetConfigType("yaml")
	V.AddConfigPath(".")
	err := V.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	V.BindEnv("dbuser")
	V.BindEnv("dbpassword")
	V.Unmarshal(&S)
	fmt.Println("Current viper config: \n", S)
	return V
}
