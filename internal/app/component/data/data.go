package data

import (
	klog "github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	zapgorm "moul.io/zapgorm2"
	"strings"
)

type Driver string

func (d Driver) String() string {
	return string(d)
}

const (
	MySQL       Driver = "mysql"
	PostgresSQL Driver = "postgres"
)

type Config struct {
	Driver          Driver
	Host            string
	Port            string
	Database        string
	Username        string
	Password        string
	Options         []string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxIdleTime int64
	ConnMaxLifeTime int64
}

// New 初始化 orm
func New(config *Config, kLogger klog.Logger, zLogger *zap.Logger) (db *gorm.DB, cleanup func(), err error) {
	if config == nil {
		return nil, func() {}, nil
	}

	gLogger := zapgorm.New(zLogger.WithOptions(zap.AddCallerSkip(3)))
	gLogger.SetAsDefault()

	switch config.Driver {
	case MySQL:

		if err != nil {
			return
		}
	case PostgresSQL:

		if err != nil {
			return
		}
	default:
		return nil, nil, nil
	}

	cleanup = func() {
		klog.NewHelper(kLogger).Info("closing the database resources")

		sqlDB, err := db.DB()
		if err != nil {
			klog.NewHelper(kLogger).Error(err)
		}

		if err := sqlDB.Close(); err != nil {
			klog.NewHelper(kLogger).Error(err)
		}
	}

	return db, cleanup, nil
}

func buildSource(c *Config) string {
	dsn := ""

	switch c.Driver {
	case PostgresSQL:
		options := strings.Join(c.Options, " ")
		dsn = "host=" + c.Host + " port=" + c.Port + " user=" + c.Username + " password=" + c.Password + " dbname=" + c.Database + " " + options
	case MySQL:
		fallthrough
	default:
		options := strings.Join(c.Options, "&")
		dsn = c.Username + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Database + "?" + options
	}

	return dsn
}
