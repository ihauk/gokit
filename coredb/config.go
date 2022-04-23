package dbcore

import "fmt"

type DBConfig struct {
	// DSN string // data source name

	// MaxIdleConns int
	// MaxOpenConns int
	// AutoMigrate  bool // 自动建表，补全缺失字段，初始化数据
	// Debug        bool

	Addr string
	User string
	Pass string
	DB   string
}

func (cfg *DBConfig) buildDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Pass, cfg.Addr, cfg.DB)
}
