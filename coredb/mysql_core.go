package dbcore

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/ihauk/gokit/logger"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

// interface ?

var (
	globalDB  *gorm.DB
	globalCfg *DBConfig
	injectors []func(db *gorm.DB)
)

func Connect(cfg *DBConfig) {

	CreateDatabase(cfg)

	dsn := cfg.buildDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: gormlog.Default.LogMode(gormlog.Info)})
	if err != nil {
		panic(err)
	}

	sdb, _ := db.DB()
	sdb.SetMaxOpenConns(50)
	sdb.SetMaxIdleConns(10)
	sdb.SetConnMaxLifetime(2 * time.Minute)

	callInjector(db)

	globalDB = db
	globalCfg = cfg

	go PingDBConnection()
}

type ctxTransactionKey struct{}

// 如果使用跨模型事务则传参
func GetDB(ctx context.Context) *gorm.DB {
	iface := ctx.Value(ctxTransactionKey{})

	if iface != nil {
		tx, ok := iface.(*gorm.DB)
		if !ok {
			logger.ErrorLog.Println("unexpect context value type: %s", reflect.TypeOf(tx))
			return nil
		}

		return tx
	}

	return globalDB.WithContext(ctx)
}

func RegisterInjector(f func(*gorm.DB)) {
	injectors = append(injectors, f)
}

func callInjector(db *gorm.DB) {
	for _, v := range injectors {
		v(db)
	}
}

func CreateDatabase(cfg *DBConfig) {
	gdsn := cfg.buildDSN()
	slashIndex := strings.LastIndex(gdsn, "/")
	dsn := gdsn[:slashIndex+1]

	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		logger.ErrorLog.Println("open db failed", err)
	}

	createSQL := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4;",
		cfg.DB,
	)

	err = db.Exec(createSQL).Error
	if err != nil {
		logger.ErrorLog.Println("create db failed", err)
	}
}

// 自动初始化表结构
func SetupTableModel(db *gorm.DB, model interface{}) {
	// if GetDBConfig().AutoMigrate {
	err := db.AutoMigrate(model)
	if err != nil {
		logger.ErrorLog.Println("setup table failed", err)
	}
	// }
}

func PingDBConnection() {

	t := time.NewTicker(time.Second * 60)

	for _ = range t.C {
		sdb, err := globalDB.DB()
		err = sdb.Ping()
		if err != nil {
			logger.ErrorLog.Println("DB connection error ", err)
		}
	}
}
