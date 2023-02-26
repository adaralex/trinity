package main

import (
	"fmt"
	"github.com/adaralex/trinity/graph/model"
	"github.com/go-cmd/cmd"
	"gorm.io/gorm"
	"os"
	"strings"
	"time"
)

type ScannerDef struct {
	gorm.Model
	model.Scanner
}

func (sd *ScannerDef) PrepareCommand(params map[string]string, timeout int) []int {
	result := make([]int, 0)
	for _, c := range sd.Install {
		// Watch out, problematic split here
		split := strings.Split(c, " ")
		if len(split) > 1 {
			result = append(result, sd.exec(split[0], split[1:], params, timeout))
		} else {
			result = append(result, sd.exec(split[0], nil, params, timeout))
		}
	}

	return result
}

func (sd *ScannerDef) RunCommand(params map[string]string, timeout int) []int {
	result := make([]int, 0)
	for _, c := range sd.Run {
		// Watch out, problematic split here
		split := strings.Split(c, " ")
		if len(split) > 1 {
			result = append(result, sd.exec(split[0], split[1:], params, timeout))
		} else {
			result = append(result, sd.exec(split[0], nil, params, timeout))
		}
	}

	return result
}

func (sd *ScannerDef) ReportCommand(params map[string]string, timeout int) []int {
	result := make([]int, 0)
	for _, c := range sd.Report {
		// Watch out, problematic split here
		split := strings.Split(c, " ")
		if len(split) > 1 {
			result = append(result, sd.exec(split[0], split[1:], params, timeout))
		} else {
			result = append(result, sd.exec(split[0], nil, params, timeout))
		}
	}

	return result
}

func (sd *ScannerDef) exec(executeCmd string, commandParams []string, projectParams map[string]string, timeout int) int {
	//if runtime.GOOS == "windows" {
	//	fmt.Printf("can't run on Windows machines\n")
	//	return
	//}

	commandParameters := make([]string, 0)
	for _, s := range commandParams {
		tmp := s
		if projectParams != nil {
			for k, v := range projectParams {
				tmp = strings.ReplaceAll(tmp, k, v)
			}
		}
		commandParameters = append(commandParameters, tmp)
	}

	command := cmd.NewCmd(executeCmd, commandParameters...)
	command.Env = os.Environ()

	statusChan := command.Start()

	ticker := time.NewTicker(2 * time.Second)
	var previousSize int
	go func() {
		for range ticker.C {
			status := command.Status()
			n := len(status.Stdout)
			if n > 0 && previousSize < n {
				for _, s := range status.Stdout {
					fmt.Println(s)
				}
				previousSize = n
			}
		}
	}()

	go func() {
		<-time.After(time.Second * time.Duration(timeout)) // Scan Timeout
		_ = command.Stop()
	}()

	// Check if command is done
	select {
	case finalStatus := <-statusChan:
		if finalStatus.Error != nil {
			fmt.Printf("error %+v\n", finalStatus.Error)
		} else {
			fmt.Printf("finished! %d\n", finalStatus.Exit)
		}
	default:
		// no, still running
	}

	// Block waiting for command to exit, be stopped, or be killed
	finalStatus := <-statusChan
	if finalStatus.Error != nil {
		fmt.Printf("error %+v\n", finalStatus.Error)
	}

	if len(finalStatus.Stdout) > 0 {
		fmt.Println("-- Stdout --")
		for _, s := range finalStatus.Stdout {
			fmt.Println(s)
		}
	}
	if len(finalStatus.Stderr) > 0 {
		fmt.Println("-- Stderr --")
		for _, s := range finalStatus.Stderr {
			fmt.Println(s)
		}
	}
	fmt.Printf("return code %d\n", finalStatus.Exit)
	return finalStatus.Exit
}
