package main

import (
	"log/slog"
	"os"
	"fmt"
)

var logger *slog.Logger

func main() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger = slog.New(handler)
	argsParse()
	logger.Info("TG-commander launch")
	config()
	botHandler()
	logger.Info("TG-commander finish")
}

func argsParse() slog.Level {
	args := os.Args[1:]
	logLevel := slog.LevelInfo
	for i := range args {
		switch args[i] {
		case "-h", "--help":
			fmt.Println("Cross-platform project for remote computer control via a telegram bot")
			fmt.Printf("\t-h, --help\thelp summary\n")
			fmt.Printf("\t-s, --service\topen service interface\n")
			fmt.Printf("\t-d, --debug\tenable debug logger level\n")
			os.Exit(0)
		case "-d", "--debug": 
			handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			})
			logger = slog.New(handler)
			logger.Debug("Debug mode enabled")
		case "-s", "--service": 
			serviceHandler()
			os.Exit(0)
		default:
			fmt.Println("Unknown argument")
			fmt.Println("-h, --help for help summary")
			os.Exit(1)
		}
	}
	return logLevel
}