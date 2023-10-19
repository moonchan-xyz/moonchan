package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/moonchan-xyz/moonchan/back/controller"
	"github.com/moonchan-xyz/moonchan/back/controller/mastodon"
	"github.com/moonchan-xyz/moonchan/back/controller/webfinger"
	docs "github.com/moonchan-xyz/moonchan/back/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /
// @Schemes http

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// utils.MetaData["HOST"] = "fedi.moonchan.xyz"

	docs.SwaggerInfo.BasePath = "/"

	r := gin.Default()
	// ping pong
	r.GET("/ping", controller.Ping)
	r.Any("/echo", controller.Echo)
	r.Any("/test", controller.Test)

	// webfinger
	r.GET(".well-known/webfinger", webfinger.Webfinger)

	// mastodon
	v1 := r.Group("/api/v1")
	{
		// v1.GET("/preferences", mastodon.Preferences)
		accounts := v1.Group("/accounts")
		{
			accounts.GET("/lookup", mastodon.LookUp)
			accounts.GET("/:id", mastodon.GetAccount)
			accounts.GET("/verify_credentials", mastodon.VerifyCredentials)
			accounts.PATCH("/update_credentials", mastodon.UpdateCredentials)
			accounts.POST("/:id/block", mastodon.Block)
			accounts.POST("/:id/unblock", mastodon.Unblock)
			accounts.POST("/:id/follow", mastodon.Follow)
			accounts.POST("/:id/remove_from_followers", mastodon.RemoveFromFollowers)
			accounts.POST("/:id/unfollow", mastodon.Unfollow)
			accounts.POST("/:id/mute", mastodon.Mute)
			accounts.POST("/:id/unmute", mastodon.Unmute)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run("0.0.0.0:3000") // listen and serve on 0.0.0.0:3000
}
