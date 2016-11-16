package patzer

import (
	"flag"
)

type Config struct {
	Debug int
}

var Cfg Config

func init() {
	flag.IntVar(&Cfg.Debug, "debug", 0, "enable debug level: 0-9")
}

func (c *Config) Parse() {
	flag.Parse()
}
