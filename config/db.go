package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type Database struct {
	Driver              string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Host                string `mapstructure:"host" json:"host" yaml:"host"`
	Port                int    `mapstructure:"port" json:"port" yaml:"port"`
	Database            string `mapstructure:"database" json:"database" yaml:"database"`
	UserName            string `mapstructure:"username" json:"username" yaml:"username"`
	Password            string `mapstructure:"password" json:"password" yaml:"password"`
	Charset             string `mapstructure:"charset" json:"charset" yaml:"charset"`
	MaxIdleConns        int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns        int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`
}

var DB *gorm.DB

func InitDB(cfg *viper.Viper) *gorm.DB {
	// 根据驱动配置进行初始化
	switch cfg.GetString("driver") {
	case "mysql":
		return initMySqlGorm(cfg)
	default:
		return initMySqlGorm(cfg)
	}
}

// 初始化 mysql gorm.DB
func initMySqlGorm(cfg *viper.Viper) *gorm.DB {
	dbCfg := &Database{}
	if err := cfg.Unmarshal(dbCfg); err != nil {
		fmt.Println("failed to unmarshal db config")
		return nil
	}
	if dbCfg.Database == "" {
		return nil
	}
	dsn := dbCfg.UserName + ":" + dbCfg.Password + "@tcp(" + dbCfg.Host + ":" + strconv.Itoa(dbCfg.Port) + ")/" +
		dbCfg.Database + "?charset=" + dbCfg.Charset + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
	}); err != nil {
		fmt.Println("mysql connect failed, err:", err)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)
		return db
	}
}
