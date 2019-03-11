package config

type Config struct {
	KeyApp  string
	Token   string
	IDBoard string
}

func Setup() (c Config) {
	c.KeyApp = "fa6b1a601cfd6559cc134d0055507cc2"
	c.Token = "e0a44f959cde9dfb0883e2865d5632232b0b3ac93900263d22d6e7f84a1d0017"
	c.IDBoard = "iCBtQXmr"
	return
}
