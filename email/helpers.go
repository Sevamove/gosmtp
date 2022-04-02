package email

import (
	"errors"
	"log"
	"net/smtp"

	"github.com/spf13/viper"
)

/*
   @notice Load .env file into the programme.
   @param _path Path to config file (ex: . or ./config).
   @param _name Name of config file (ex: secret).
   @param _type Extension of config file (ex: env).
*/
func loadEnv(_path, _name, _type string) (env Env, err error) {
	viper.AddConfigPath(_path)
	viper.SetConfigName(_name)
	viper.SetConfigType(_type)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err.Error())
	}

	err = viper.Unmarshal(&env)
	return
}

func getEnv() Env {
	path := "."
	name := "secret"
	extn := "env"

	env, _ := loadEnv(path, name, extn)

	return env
}

// @dev We implement the next 2 functions in order to authenticate us.
// @dev By default, Go does not support LOGIN auth method. See more:
// @dev https://stackoverflow.com/questions/57783841/how-to-send-email-using-outlooks-smtp-servers

func (a *LoginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *LoginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}
