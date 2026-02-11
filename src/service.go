package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/kardianos/service"
)

type program struct{}

func serviceHandler() {
	var action string
	
	fmt.Println("Welcome to service interface")
	fmt.Println("[C]reate | [D]elete | [R]estart")
	fmt.Print("Choose action: ")

	_, err := fmt.Scan(&action)
	if err != nil {logger.Error("Error with get user action", "err", err)}
	logger.Debug("User choosed action", "action", action)

	switch strings.ToLower(action) {
	case "c": serviceWork("c")
	case "d": serviceWork("d")
	case "r": serviceWork("r")
	default: logger.Warn("Unknown action")
	}
}

func serviceWork(action string) {
	execPath, err := os.Getwd()
	if err != nil {logger.Error("Failed get a execPath", "err", err.Error()); os.Exit(1)}
	logger.Info(execPath)
	svcConfig := &service.Config{
	Name:        "TGCommanderService",
	DisplayName: "TG-Commander Service",
	Description: "Cross-platform project for remote computer control via a telegram bot\ngithub.com/Konare1ka/TG-commander",
	WorkingDirectory: execPath,
	Option: service.KeyValue{"StartType": "automatic"},
	}

    prg := &program{}
    s, err := service.New(prg, svcConfig)
    if err != nil { logger.Error("Can't create service", "err", err) }
	switch action {
	case "c":
		if err := s.Install(); err != nil { logger.Error("Can't install service", "err", err) 
		} else {logger.Info("Successfully service installed")}
		if err := s.Start(); err != nil { logger.Error("Can't start service", "err", err)
		} else {logger.Info("Successfully service started")}
	case "d":
		if err := s.Uninstall(); err != nil {logger.Error("Can't delete service", "err", err)
		} else { logger.Info("Successfully service deleted") }
	case "r":
		if err := s.Restart(); err != nil {logger.Error("Can't restart service", "err", err)
		} else {logger.Info("Successfully service restarted")}
	}

}
func (p *program) Start(s service.Service) error {
	logger.Info("Service TG-commander launch")
    go p.run()
    return nil
}

func (p *program) run() {
	config()
	botHandler()
}

func (p *program) Stop(s service.Service) error {
    logger.Info("Service TG-commander stop")
    return nil
}