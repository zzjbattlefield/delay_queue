package cmd

import (
	"github.com/zzjbattlefield/delay_queue/config"
	"github.com/zzjbattlefield/delay_queue/queue"
)

func Run() {
	config.Init() //配置初始化
	queue.Init()  //延迟队列的初始化
}
