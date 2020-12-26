package config

import "github.com/spf13/viper"

// VaultPath _
var VaultPath string

func Init() {
	VaultPath = viper.GetString("vault.path")
}
