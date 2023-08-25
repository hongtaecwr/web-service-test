package config

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	Database struct {
		User     string
		Password string
		Net      string
		Addr     string
		DBName   string
	}
	Server struct {
		Address string
	}
}

var (
	C config
)

func ReadConfig() {
	Config := &C
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("$GOPATH", "src", "web-service-test", "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

func GetKey() (*rsa.PrivateKey, *rsa.PublicKey) {
	keyPrivatePath, _ := filepath.Abs("config/keys/key.rsa")
	key, err := os.ReadFile(keyPrivatePath)
	if err != nil {
		panic(err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		panic(err)
	}
	keyPublicPath, _ := filepath.Abs("config/keys/key.rsa.pub")
	key, err = os.ReadFile(keyPublicPath)
	if err != nil {
		panic(err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
	if err != nil {
		panic(err)
	}
	return privateKey, publicKey
}
