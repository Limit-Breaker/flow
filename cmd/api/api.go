package api

import (
	"context"
	"flow/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"flow/config"
)

const (
	ModeDev  = "dev"  //开发模式
	ModeTest = "test" //测试模式
	ModeProd = "prod" //生产模式
)

var (
	settings string
	port     string
	mode     string
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start API server",
		Example: "flow server config/settings.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&settings, "settings", "c", "C:\\MyWorkspace\\code\\flow\\config\\api-settings.yaml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8002", "Tcp port server listening on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}

func usage() {
	usageStr := `starting api server`
	slog.Info("%s\n", usageStr)
}

func setup() {
	config.MustLoad(settings)
}

func run() error {
	if config.ApplicationConfig.Mode == ModeProd {
		gin.SetMode(gin.ReleaseMode)
	}
	r := routers.InitRouter()
	srv := &http.Server{
		Addr:    config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port,
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	fmt.Printf("Server %s Run http://%s:%s/ \r\n",
		config.ApplicationConfig.Name,
		config.ApplicationConfig.Host,
		config.ApplicationConfig.Port)

	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", time.Now())
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", time.Now())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	return nil

	return nil
}
