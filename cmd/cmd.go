package cmd

import (
	"fmt"
	"net/http"

	"github.com/zzjbattlefield/delay_queue/config"
	"github.com/zzjbattlefield/delay_queue/queue"
	"github.com/zzjbattlefield/delay_queue/route"
)

type Cmd struct {
}

func (cmd *Cmd) Run() {
	config.Init() //配置初始化
	queue.Init()  //延迟队列的初始化
	cmd.runWeb()
}

func (cmd *Cmd) runWeb() {
	http.HandleFunc("/push", route.PushJob)
	http.HandleFunc("/delete", route.DeleteJob)
	fmt.Printf("开始监听%s \n", config.Setting.BindAddress)
	http.ListenAndServe(config.DefaultBindAddress, nil)
}
