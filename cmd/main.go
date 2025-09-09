package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/haribabu/cardano-node-monitor/internal/api"
	"github.com/haribabu/cardano-node-monitor/internal/config"
	"github.com/haribabu/cardano-node-monitor/internal/monitor"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cardano-node-monitor",
	Short: "A monitoring tool for Cardano nodes",
	Long:  `A comprehensive monitoring tool that tracks Cardano node health, performance metrics, and network connectivity.`,
	Run:   runServer,
}

func init() {
	rootCmd.PersistentFlags().String("config", "", "config file (default is $HOME/.cardano-monitor.yaml)")
	rootCmd.PersistentFlags().String("node-url", "http://localhost:12798", "Cardano node API URL")
	rootCmd.PersistentFlags().String("port", "8080", "Server port")
	rootCmd.PersistentFlags().Duration("interval", 30*time.Second, "Monitoring interval")
}

func runServer(cmd *cobra.Command, args []string) {
	cfg, err := config.Load(cmd)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	nodeMonitor := monitor.New(cfg)
	
	go func() {
		if err := nodeMonitor.Start(context.Background()); err != nil {
			log.Printf("Monitor error: %v", err)
		}
	}()

	router := gin.Default()
	apiHandler := api.NewHandler(nodeMonitor)
	apiHandler.RegisterRoutes(router)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		fmt.Printf("Starting Cardano Node Monitor on port %s\n", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	nodeMonitor.Stop()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}