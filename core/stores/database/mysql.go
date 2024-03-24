/*
 * @Author: yujiajie
 * @Date: 2024-03-22 14:47:30
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-22 14:52:27
 * @FilePath: /gateway/core/stores/database/mysql.go
 * @Description:
 */
package database

import (
	"fmt"
	"gateway/options"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

func NewDB(c *options.MysqlConfig) error {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: c.Default,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		PrepareStmt: true,
	})
	if err != nil {
		fmt.Println("mysql init failed")
		return err
	}
	if c.Cluster {
		initCluster(db, c)
	}
	return nil
}

func initCluster(db *gorm.DB, c *options.MysqlConfig) {
	if len(c.Sources) == 0 && len(c.Replicas) == 0 {
		return
	}

	config := dbresolver.Config{}
	config.Sources = []gorm.Dialector{}
	config.Replicas = []gorm.Dialector{}

	for _, sourceConfig := range c.Sources {
		config.Sources = append(config.Sources, mysql.Open(sourceConfig))
	}
	for _, replicaConfig := range c.Replicas {
		config.Replicas = append(config.Replicas, mysql.Open(replicaConfig))
	}
	db.Use(dbresolver.Register(config))
}
