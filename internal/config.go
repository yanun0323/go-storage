package internal

type Config struct {
	Upload struct {
		Token string `yaml:"token"`
	} `yaml:"upload"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	SQLite bool `yaml:"sqlite"`
	MySQL  struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`
}
