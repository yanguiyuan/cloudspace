package gorm

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yanguiyuan/yuan/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// 可选参数 gorm-gen.dal,gorm-gen.model 指定dal和model的生成路径
// 必需提供参数 mysql.dsn
func MysqlWithConfig(c *viper.Viper) {
	fmt.Println(c.GetString("mysql.dsn"))
	db, err := gorm.Open(mysql.Open(c.GetString("mysql.dsn")), &gorm.Config{})
	if err != nil {
		return
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:      c.GetString("gorm-gen.dal"),
		ModelPkgPath: c.GetString("gorm-gen.model"),
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
func MysqlGorm() {
	c := config.NewConfig()
	MysqlWithConfig(c)
}

// 必须提供 sqlite.dbname参数
func SqliteWithConfig(c *viper.Viper) {
	db, err := gorm.Open(sqlite.Open(c.GetString("sqlite.dbname")))
	if err != nil {
		return
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:      c.GetString("gorm-gen.dal"),
		ModelPkgPath: c.GetString("gorm-gen.model"),
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
func Sqlite() {
	c := config.NewConfig()
	SqliteWithConfig(c)
}
