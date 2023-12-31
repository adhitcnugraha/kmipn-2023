package main

import (
	"embed"
	"fmt"
	"kmipn-2023/client"
	"kmipn-2023/db"
	"kmipn-2023/handler/api"
	"kmipn-2023/handler/web"
	"kmipn-2023/middleware"
	"kmipn-2023/model"
	repo "kmipn-2023/repository"
	"kmipn-2023/service"

	"github.com/gin-contrib/cors"

	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIHandler struct {
	// struct api handler here
	UserAPIHandler api.UserAPI
	// SellerAPIHandler api.SellerAPI
}

type ClientHandler struct {
	// struct Client Handler here
	AuthWeb      web.AuthWeb
	HomeWeb      web.HomeWeb
	DashboardWeb web.DashboardWeb
	ModalWeb     web.ModalWeb
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
			Host:         "0.0.0.0",
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
	// CORS Middleware
	gin.Use(cors.Default())

	userRepo := repo.NewUserRepo(db)
	// sellerRepo := repo.NewSellerRepo(db)
	sessionRepo := repo.NewSessionsRepo(db)
	// adminRepo := repo.NewAdminRepo(db)
	// productRepo := repo.NewProductRepo(db)
	// orderRepo := repo.NewOrderRepo(db)

	userService := service.NewUserService(userRepo, sessionRepo)
	// sellerService := service.NewSellerService(sellerRepo, sessionRepo)
	// productService := service.NewProductService(productRepo)
	// orderService := service.NewOrderService(orderRepo)

	userAPIHandler := api.NewUserAPI(userService)
	// sellerAPIHandler := api.NewSellerAPI(sellerService)
	// adminAPIHandler := api.NewAdminAPI(adminService)
	// productAPIHandler := api.NewProductAPI(productService)
	// orderAPIHandler := api.NewOrderAPI(orderService)

	apiHandler := APIHandler{
		UserAPIHandler: userAPIHandler,
		// SellerAPIHandler: sellerAPIHandler,
		// AdminAPIHandler: adminAPIHandler,
		// ProductAPIHandler: productAPIHandler,
		// OrderAPIHandler: orderAPIHandler,
	}

	version := gin.Group("/api/v1")
	{
		user := version.Group("/user")
		{
			user.POST("/login", apiHandler.UserAPIHandler.Login)
			user.POST("/register", apiHandler.UserAPIHandler.Register)

			user.Use(middleware.Auth())
			user.GET("/product", apiHandler.UserAPIHandler.GetUserProductCategory)
		}
		// seller := version.Group("/seller")
		// {
		// 	seller.POST("/login", apiHandler.SellerAPIHandler.Login)
		// 	seller.POST("/register", apiHandler.SellerAPIHandler.Register)

		// 	seller.Use(middleware.Auth())
		// }
	}
	return gin
}

func RunClient(db *gorm.DB, gin *gin.Engine, embed embed.FS) *gin.Engine {
	// Bagian 1
	sessionRepo := repo.NewSessionsRepo(db)
	sessionService := service.NewSessionService(sessionRepo)

	// Bagian 2
	userClient := client.NewUserClient()

	// Bagian 3
	authWeb := web.NewAuthWeb(userClient, sessionService, embed)
	modalWeb := web.NewModalWeb(embed)
	homeWeb := web.NewHomeWeb(embed)
	dashboardWeb := web.NewDashboardWeb(userClient, sessionService, embed)

	client := ClientHandler{
		authWeb, homeWeb, dashboardWeb, modalWeb,
	}

	gin.StaticFS("/static", http.Dir("frontend/public"))

	gin.GET("/", client.HomeWeb.Index)

	user := gin.Group("/client")
	{

		user.GET("/login", client.AuthWeb.Login)
		user.POST("/login/process", client.AuthWeb.LoginProcess)
		user.GET("/register", client.AuthWeb.Register)
		user.POST("/register/process", client.AuthWeb.RegisterProcess)
		user.Use(middleware.Auth())
		user.GET("/logout", client.AuthWeb.Logout)
	}

	main := gin.Group("/client")
	{
		main.Use(middleware.Auth())
		main.GET("/dashboard", client.DashboardWeb.Dashboard)
	}

	modal := gin.Group("/client")
	{
		modal.GET("/modal", client.ModalWeb.Modal)
	}

	return gin
}
