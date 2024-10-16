package nknovh_engine

type configuration struct {
	Version       string
	Db            string `env:"DB_CONNECTION_STRING"`
	LogLevel      string `env:"LOG_LEVEL"`
	Port          string `env:"PORT"`
	DbType        string `env:"DB_TYPE"`
	WalletPath    string `env:"WALLET_PATH"`
	WebPath       string `env:"WEB_PATH"`
	TemplatesPath string `env:"TEMPLATE_PATH"`
	NodesPath     string `env:"NODES_PATH"`

	NeighborPoll struct {
		ConnTimeout    int `json:"ConnTimeout"`
		Interval       int `json:"Interval"`
		RemoveInterval int `json:"RemoveInterval"`
	} `json:"NeighborPoll"`
	MainPoll struct {
		ConnTimeout    int `json:"ConnTimeout"`
		Interval       int `json:"Interval"`
		EntriesPerNode int `json:"EntriesPerNode"`
	} `json:"MainPoll"`
	DirtyPoll struct {
		ConnTimeout int `json:"ConnTimeout"`
		Interval    int `json:"Interval"`
	} `json:"DirtyPoll"`
	Threads struct {
		Neighbors int `json:"Neighbors"`
		Main      int `json:"Main"`
		Dirty     int `json:"Dirty"`
	} `json:"Threads"`
	Wallets struct {
		Interval int `json:"Interval"`
	} `json:"Wallets"`
	Messengers struct {
		Telegram struct {
			Use   bool   `json:"Use"`
			Token string `json:"Token"`
		}
	} `json:"Messengers"`
	TrustedProxies []string `json:"TrustedProxies"`
	SeedList       []string `json:"SeedList"`
}
