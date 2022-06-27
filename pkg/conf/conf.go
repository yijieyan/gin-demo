package conf

import "github.com/spf13/viper"

var v = viper.New()

func Init(filePath string) error {
	v.SetConfigFile(filePath)
	return v.ReadInConfig()
}

func Get(key string) interface{} {
	return v.Get(key)
}

func GetBool(key string) bool {
	return v.GetBool(key)
}

func GetInt(key string) int {
	return v.GetInt(key)
}

func GetString(key string) string {
	return v.GetString(key)
}

func Unmarshal(rawVal interface{}) error {
	return v.Unmarshal(rawVal)
}

func All() map[string]interface{} {
	return v.AllSettings()
}
