package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	"challenge/config"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gitlab.com/fillgoods-library/microservice/datasource/fsql"
	"gitlab.com/fillgoods-library/microservice/vault"
	"go.elastic.co/apm/module/apmgin"
)

func init() {
	initViper()
	config.Init()
	fmt.Println()
	runEnv := viper.GetString("run.env")
	vault.Init(vault.Config{
		Address:  viper.GetString("vault.baseurl"),
		Token:    viper.GetString("vault.token"),
		BasePath: "/" + runEnv,
	})

}

func createConnection(databaseName string) (*sql.DB, error) {

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		vault.GetString(config.VaultPath, "postgres_host"),
		vault.GetString(config.VaultPath, "postgres_port"),
		vault.GetString(config.VaultPath, "postgres_user"),
		vault.GetString(config.VaultPath, "postgres_password"),
		databaseName,
	)
	return fsql.NewFSQL(connString, fsql.POSTGRES)
}

func initViper() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read in viper config:%s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func main() {
	db, err := createConnection("prelive")

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)

	g := gin.Default()

	g.Use(apmgin.Middleware(g))

	g.GET("/", func(c *gin.Context) {

		wg := sync.WaitGroup{}

		done := make(chan struct{}, 1)
		done2 := make(chan struct{}, 1)

		defer func() {
			wg.Wait()
			close(done)
			close(done2)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			_, err := db.ExecContext(c.Request.Context(), "SELECT COUNT(*) FROM facebook_lives ")

			if err != nil {
				log.Println("ERROR COUNT(*) ", err.Error())
			}

			done <- struct{}{}
			log.Println("COUNT DONE")
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := db.QueryContext(c.Request.Context(), "SELECT * FROM facebook_lives"); err != nil {
				log.Println("ERROR", err.Error())
			}
			done2 <- struct{}{}
			log.Println("SELECT DONE")
		}()

		<-done
		<-done2

		c.String(200, "OKKK")

	})

	g.Run(":3003")

}
