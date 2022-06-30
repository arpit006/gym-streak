package config

import "github.com/spf13/viper"

func readViperString(key string) string {
	// type assertion T.(v)
	return viper.Get(key).(string)
	// type casting v.(T)
}

func readViperInteger(key string) int {
	return viper.Get(key).(int)
}
