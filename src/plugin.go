package main

import (
	"os"
	"bufio"
	"os/exec"
	"path/filepath"
	"strings"
)

var plst map[string]string
var pluginPath string

func pluginExecute() []string {
	launchCom, args := commandMaker()
	cmd := exec.Command(launchCom, args...)
	logger.Debug(cmd.String())
	stdout, err := cmd.StdoutPipe()
	if err != nil {logger.Error("Can't create stdout pipe", "err", err)}
	if err := cmd.Start(); err != nil {logger.Error("Can't launch plugin", "err", err)}

	scanner := bufio.NewScanner(stdout) //Dynamically monitor all stdout

	var output []string

	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		output = append(output, line)
	}

	if err := cmd.Wait(); err != nil {
		output = append(output, err.Error())
		return output
	}
	return output
}

func commandMaker() (string, []string) {
	pluginFullPath := filepath.Join(pluginPath, message[0] + plst[message[0]])	//pluginPath + plugin name + plugin extension
	args := append([]string{pluginFullPath}, message[1:]...)
	logger.Debug(pluginFullPath)

	var launchCom string

	switch plst[message[0]] {
	case ".sh": launchCom = "bash"
	case ".py": launchCom = "python"
	case ".go": 
		launchCom = "go"
		args = append([]string{"run"}, args...)	//inserting "run" at beginning of args array
	default: 
		launchCom = pluginFullPath				//if plugin is binary(ex. exe) or haven't extension
		args = message[1:]
	}

	return launchCom, args
}

func pluginsListMaker() {
	plst = make(map[string]string)
	if cfg.PluginPath != "" {
		pluginPath = cfg.PluginPath
	} else {
	execPath, err := os.Executable()
	if err != nil {logger.Error("Failed get a execPath", "err", err.Error()); os.Exit(1)}
	pluginPath = filepath.Join(filepath.Dir(execPath), "plugins")
	logger.Debug("Plugin", "path", pluginPath)
	}

	_, err := os.Stat(pluginPath)
	if err != nil {
		logger.Warn("Plugins direcotry not exists")
		os.Exit(1)
	}

	entries, err := os.ReadDir(pluginPath)
	if err != nil { logger.Error("Error when getting a list of files in the plugins directory", "error", err.Error())}

	for _, entry := range entries {
		if !entry.IsDir() {
			text := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))	//without extension
			ext := filepath.Ext(entry.Name())										//extension only
			plst[text] = ext
			logger.Debug("Check file", "entry", entry.Name())
		}
	}

	if len(plst) == 0 {logger.Warn("No plugins")}
}
