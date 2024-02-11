package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andrietri/guest-registration/database/database"
	"github.com/andrietri/guest-registration/http/router"
	"github.com/erajayatech/go-helper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("Cannot load file .env: ", err)
		panic(err)
	}
	logrus.Warn("ENV:", os.Getenv("MODE"))
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initializes database
	db := database.NewDBConnection()

	mode := helper.MustGetEnv("MODE")
	port := helper.MustGetEnv("PORT")
	if mode == "local" {
		port = helper.MustGetEnv("LOCAL_PORT")
	}

	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(database.Inject(db))

	if mode != "local" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// get default url request
	app.UseRawPath = true
	app.UnescapePathValues = true

	// cors configuration
	config := cors.DefaultConfig()
	config.AddAllowHeaders("Authorization", "x-source", "platform")
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"OPTIONS", "PUT", "POST", "GET", "DELETE"}
	app.Use(cors.New(config))

	// router
	// endpoint health check
	app.GET("/ping", ping)
	router.ApiV1(&app.RouterGroup, db)

	app.NoRoute(lostInSpce)

	// ----------------------------
	// handling gracefully shutdown
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: app,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	logrus.Warn("shutting down gracefully, press Ctrl+C again to force 🔴")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server forced to shutdown 🔴: ", err)
	}
	logrus.Warn("Server exiting 🔴")
}

func lostInSpce(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":        404,
		"data":          nil,
		"error_message": "Lost in space",
	})
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"service_name": os.Getenv("SERVICE_NAME"),
		"mode":         os.Getenv("MODE"),
	})
}
