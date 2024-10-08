package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/log"
	"github.com/dsc-sgu/shawty/internal/server/middleware"
	"github.com/dsc-sgu/shawty/internal/server/routes"
	apiroutes "github.com/dsc-sgu/shawty/internal/server/routes/api"
	webroutes "github.com/dsc-sgu/shawty/internal/server/routes/webui"
	"github.com/gin-gonic/gin"
)

func Launch() {
	// debug or release
	gin.SetMode(config.C.Server.Mode)

	// server will run using this context
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	// gin server engine
	e := gin.New()

	e.Use(
		middleware.AccessLogMiddleware(),
		middleware.AuthMiddleware([]string{
			`^/home$`,
			`^/home/.*$`,
			`^/links$`,
			`^/links/.*$`,
			`^/api/links$`,
			`^/api/links/.*$`,
		}),
	)

	e.Static("/static", "./static")
	e.GET("/ping", routes.GetPing)
	e.GET("/s/:name", routes.Redirect)
	e.GET("/", webroutes.GetIndex)
	e.POST("/auth", webroutes.PostAuth)
	e.GET("/logout", webroutes.GetLogout)
	e.GET("/home", webroutes.GetHome)
	e.GET("/links", webroutes.GetLinks)
	e.POST("/links", webroutes.PostLink)
	e.GET("/links/new", webroutes.NewLink)
	e.DELETE("/links/:id", webroutes.DeleteLink)

	g := e.Group("/api")
	{
		g.GET("/version", apiroutes.GetVersion)
		g.POST("/auth", apiroutes.PostAuth)
		g.GET("/links", apiroutes.GetLinks)
		g.POST("/links", apiroutes.PostLink)
		g.DELETE("/links", apiroutes.DeleteLink)
	}

	// disable trusted proxy warning
	if err := e.SetTrustedProxies(nil); err != nil {
		log.S.Fatalw(
			"Failed to configure trusted proxies settings",
			"error", err,
		)
	}

	// create a new server
	srv := &http.Server{
		Handler: e,
	}
	// setting onShutdown logic
	srv.RegisterOnShutdown(onShutdown)

	// create listener
	listener, err := net.Listen("tcp", fmt.Sprintf(
		"%s:%d",
		config.C.Server.Host,
		config.C.Server.Port,
	))
	if err != nil {
		log.S.Fatalw("Failed to create a TCP listener", "error", err)
	}
	defer listener.Close()

	// perform startup logic
	err = onStartup(ctx)
	if err != nil {
		log.S.Fatalw("Failed to start the service", "error", err)
	}

	// server runs in a goroutine
	go func() {
		if err := srv.Serve(listener); err != nil &&
			err != http.ErrServerClosed {
			log.S.Fatalw(
				"An error occurred, cannot listen for requests anymore",
				"error", err,
			)
		}
	}()

	// listen for the interrupt signal
	<-ctx.Done()

	// restore default behavior on the interrupt signal and notify user of shutdown
	cancel()
	log.S.Info("Shutting down gracefully, press Ctrl+C to force")

	ctx, cancel = context.WithTimeout(
		context.Background(),
		time.Duration(5)*time.Second,
	)
	defer cancel()

	// perform shutdown logic
	if err := srv.Shutdown(ctx); err != nil {
		log.S.Errorw(
			"Server forced to shutdown",
			"error", err,
		)
	}
}
