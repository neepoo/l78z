package functional

type Config struct {
	Addr    string
	MaxConn uint32
}

type Option func(Config) Config

func WithAddr(addr string) Option {
	return func(config Config) Config {
		config.Addr = addr
		return config
	}
}

func WithMaxConn(cnt uint32) Option {
	return func(config Config) Config {
		config.MaxConn = cnt
		return config
	}
}

func NewConfig(opts ...Option) Config {
	c := Config{}
	for _, opt := range opts {
		c = opt(c)
	}
	return c
}
