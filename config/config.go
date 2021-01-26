package config

func GetFromConfig() string {
	banana := GetEnv("banana")

	return banana
}
