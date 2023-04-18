package core

import (
	"encoding/json"
	"fmt"
	"hostscan/elog"
	"hostscan/models"
	"hostscan/utils"
	"hostscan/vars"
	"strings"
	"sync"
)

type Task struct {
	Uri  string
	Host string
}

func goScan(taskChan chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-taskChan:
			if !ok {
				return
			} else {
				vars.ProcessBar.Add(1)
				flag := *vars.Flag
				body := utils.GetHttpBody(task.Uri, task.Host)
				if strings.Contains(body, flag) {
					var result models.Result
					result.Uri = task.Uri
					result.Host = task.Host
					result.Title = body
					resultStr, _ := json.Marshal(result)
					elog.Notice(fmt.Sprintf("Uri: %s, Host: %s --> %s", task.Uri, task.Host, body))
					utils.WriteLine(string(resultStr), *vars.OutFile)
				}
			}
		}
	}
}
