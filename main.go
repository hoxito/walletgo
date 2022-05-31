package main

import (
	"fmt"
	"time"

	"github.com/hoxito/walletgo/rest/middlewares"
	"github.com/hoxito/walletgo/rest/routes"
	"github.com/hoxito/walletgo/tools/env"
	"github.com/joho/godotenv"
	"github.com/penglongli/gin-metrics/ginmetrics"

	docs "github.com/hoxito/walletgo/docs"
	cors "github.com/itsjamie/gin-cors"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Wallet API
// @version         1.0
// @description     This is a simple wallet API. you can handle users, wallets and transactions between them
// @termsOfService  http://swagger.io/terms/

// @contact.name   jose aranciba
// @contact.email  josearanciba09@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3010

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading environment file, setting defult environment variables")
	}
	router := gin.Default()

	if err != nil {
		fmt.Print("error al conectar a instancia de redis:", err)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Use(static.Serve("/", static.LocalFile("www", true)))
	router.Use(middlewares.ErrorHandler)
	router.Use(cors.Middleware(cors.Config{
		Origins:         "http://localhost",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, Size",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	//prometheus instrumentation

	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(router)

	//routing

	routes.WalletRoute(router)
	routes.LoginRoute(router)
	routes.UserRoute(router)
	router.Run(fmt.Sprintf(":%d", env.Get().Port))
}
