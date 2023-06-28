package main

import (
	"embed"
	"fmt"
	"kmipn-2023/db"
	"kmipn-2023/model"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIHandler struct {
	// struct api handler here
}

type ClientHandler struct {
	// struct Client Handler here
}

var Resources embed.FS

func main() {
	gin.SetMode(gin.ReleaseMode) // release

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		router := gin.New()
		db := db.NewDB()
		router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("[%s] \"%s %s %s\"\n",
				param.TimeStamp.Format(time.RFC822),
				param.Method,
				param.Path,
				param.ErrorMessage,
			)
		}))
		router.Use(gin.Recovery())
		dbCredential := model.Credential{
			Host:         "localhost",
			Username:     "postgres",
			Password:     "rastafara86",
			DatabaseName: "postgres",
			Port:         5432,
			Schema:       "public",
		}
		conn, err := db.Connect(&dbCredential)
		if err != nil {
			panic(err)
		}

		// comment
		conn.AutoMigrate(
			&model.User{}, &model.Seller{}, &model.Admin{},
			&model.Product{}, &model.Order{}, &model.Session{},
		)

		router = RunServer(conn, router)
		router = RunClient(conn, router, Resources)

		fmt.Println("Server is running on port 8080")
		err = router.Run(":8080")
		if err != nil {
			panic(err)
		}
	}()
	wg.Wait()

}

func RunServer(db *gorm.DB, gin *gin.Engine) *gin.Engine {
	return nil
}

func RunClient(db *gorm.DB, gin *gin.Engine, embed embed.FS) *gin.Engine {
	return nil
}
