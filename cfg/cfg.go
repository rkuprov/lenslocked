package cfg

import (
	"encoding/json"
	"io"
	"os"
)

type Cfg struct {
	Postgres Postgres `json:"postgres"`
}

type Postgres struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DBName       string `json:"dbname"`
	SSLMode      string `json:"sslmode"`
	PoolMaxConns string `json:"pool_max_conns"`
}

func (c *Cfg) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	bts, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	return json.Unmarshal(bts, c)
}
