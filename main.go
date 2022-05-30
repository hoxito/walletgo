package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/hoxito/walletgo/rest/middlewares"
	"github.com/hoxito/walletgo/rest/routes"
	"github.com/hoxito/walletgo/tools/env"
	"github.com/joho/godotenv"
	ginprometheus "github.com/zsais/go-gin-prometheus"

	cors "github.com/itsjamie/gin-cors"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading environment file, setting defult environment variables")
	}
	router := gin.Default()

	store, err := redis.NewStore(10, "tcp", env.Get().RedisURL, "redispw", []byte(env.Get().SessionSecret))
	if err != nil {
		fmt.Print("error al conectar a instancia de redis:", err)
	}
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(sessions.Sessions("mysession", store))
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

	p := ginprometheus.NewPrometheus("gin")
	p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL.Path
		for _, p := range c.Params {
			if p.Key == "walletid" {
				url = strings.Replace(url, p.Value, ":wallet", 1)
				break
			}
			if p.Key == "transactionid" {
				url = strings.Replace(url, p.Value, ":wallet", 1)
				break
			}
		}
		return url
	}
	p.Use(router)

	routes.WalletRoute(router)
	routes.LoginRoute(router)
	routes.UserRoute(router)
	router.Run(fmt.Sprintf(":%d", env.Get().Port))
}
