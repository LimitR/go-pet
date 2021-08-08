package apiserver

type Congig struct {
	BingAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Congig {
	return &Congig{
		BingAddr: ":8080",
		LogLevel: "debug",
	}
}
