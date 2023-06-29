package main

import (
	"embed"
	"fmt"
	"kmipn-2023/db"
	"kmipn-2023/model"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
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

		// router = RunServer(conn, router)
		// router = RunClient(conn, router, Resources)

		fmt.Println("Server is running on port 8080")
		err = router.Run(":8080")
		if err != nil {
			panic(err)
		}
	}()
	wg.Wait()

}

// func RunServer(db *gorm.DB, gin *gin.Engine) *gin.Engine {
// userRepo := repo.NewUserRepo(db)
// sessionRepo := repo.NewSessionsRepo(db)
// adminRepo := repo.NewAdminRepo(db)
// productRepo := repo.NewProductRepo(db)
// orderRepo := repo.NewOrderRepo(db)

// userService := service.NewUserService(userRepo, sessionRepo)
// adminService := service.NewAdminService(adminRepo, sessionRepo)
// productService := service.NewProductService(productRepo)
// orderService := service.NewOrderService(orderRepo)

// userAPIHandler := api.NewUserAPI(userService)
// adminAPIHandler := api.NewAdminAPI(adminService)
// productAPIHandler := api.NewProductAPI(productService)
// orderAPIHandler := api.NewOrderAPI(orderService)

// apiHandler := APIHandler {
// UserAPIHandler: userAPIHandler,
// AdminAPIHandler: adminAPIHandler,
// ProductAPIHandler: productAPIHandler,
// OrderAPIHandler: orderAPIHandler,
// }

// version := gin.Group("api/v1")
// {
// user := version.Group("/user")
// {
// user.POST("/login"), apiHandler.UserAPIHandler.Login
// user.POST("/register"), apiHandler.UserAPIHandler.Register

// user.Use(middleware.Auth())
// user.GET("/product", apiHandler.ProductAPIHandler. --> kurang satu param lagi)
// }
// }
// 	return nil
// }

// func RunClient(db *gorm.DB, gin *gin.Engine, embed embed.FS) *gin.Engine {
// 	return nil
// }
