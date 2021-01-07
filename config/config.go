package config

import (
	"fmt"
	"gtmhub-cli/output"
	"os"

	"github.com/spf13/viper"
)

var (
	refreshTokenKey = "refreshToken"
	userIDKey       = "userId"
	accountIDKey    = "accountId"
	tokenKey        = "token"
	gtmhubUrlKey    = "gtmhubUrl"
	gtmhubDCKey     = "dc"
)

func InitConfig() {
	homePath, _ := os.LookupEnv("HOME")
	homePath = joinPaths(homePath, ".gtmhub")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(homePath) // call multiple times to add many search paths

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if _, err := os.Stat(homePath); os.IsNotExist(err) {
				os.Mkdir(homePath, os.ModePerm)
			}

			emptyFile, err := os.Create(joinPaths(homePath, "config"))
			defer emptyFile.Close()
			if err != nil {
				panic(err.Error())
			}

		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}

func joinPaths(rootPath, addOn string) string {
	return rootPath + string(os.PathSeparator) + addOn
}

func SetToken(token string) {
	viper.Set(tokenKey, token)
	if err := viper.WriteConfig(); err != nil {
		output.PrintErrorAndExit(err)

	}
}

func SetAccountId(accountId string) {
	viper.Set(accountIDKey, accountId)
	if err := viper.WriteConfig(); err != nil {
		output.PrintErrorAndExit(err)
	}
}

func SetUserID(userID string) {
	viper.Set(userIDKey, userID)
	if err := viper.WriteConfig(); err != nil {
		output.PrintErrorAndExit(err)

	}
}

func SetRefreshToken(refreshToken string) {
	viper.Set(refreshTokenKey, refreshToken)
	if err := viper.WriteConfig(); err != nil {
		output.PrintErrorAndExit(err)
	}
}

func SetGtmhubUrl(url string) {
	viper.Set(gtmhubUrlKey, url)
	if err := viper.WriteConfig(); err != nil {
		output.PrintErrorAndExit(err)
	}
}

func SetGtmhubDC(dc string) {
	viper.Set(gtmhubDCKey, dc)
	if err := viper.WriteConfig(); err != nil {
		output.PrintErrorAndExit(err)
	}
}

func GetRefreshToken() string {
	return getStringKeyAndExitIfMissing(refreshTokenKey)
}

func GetGtmhubDC() string {
	return getStringKeyAndExitIfMissing(gtmhubDCKey)
}

func GetGtmhubUrl() string {
	return getStringKeyAndExitIfMissing(gtmhubUrlKey)
}

func GetToken() string {
	return getStringKeyAndExitIfMissing(tokenKey)
}

func GetAccountId() string {
	return getStringKeyAndExitIfMissing(accountIDKey)
}

func GetUserID() string {
	return getStringKeyAndExitIfMissing(userIDKey)
}

func getStringKeyAndExitIfMissing(key string) string {
	val := viper.GetString(key)
	if len(val) == 0 {
		output.PrintErrorAndExit(fmt.Errorf(authenticationFailedMsgFmt, key))
	}

	return val
}

func Clear() {

	viper.Set(refreshTokenKey, "")
	viper.Set(userIDKey, "")
	viper.Set(accountIDKey, "")
	viper.Set(tokenKey, "")

	viper.WriteConfig()
}
