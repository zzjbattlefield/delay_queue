package queue

import (
	"fmt"
	"log"
	"time"

	"github.com/zzjbattlefield/delay_queue/config"
)

func Init() {
	//初始化队列相关
	//初始化redis
	initRedis()
	initTimer()
}

func initTimer() {
	//按照配置的bucket大小创建对应的携程去扫描bucket
	times := make([]*time.Ticker, config.Setting.BucketSize)
	for i := 0; i < config.Setting.BucketSize; i++ {
		times[i] = time.NewTicker(1 * time.Second)
		bucketName := fmt.Sprintf(config.Setting.BucketName, i)
		go waitTicker(*times[i], bucketName)
	}

}

func waitTicker(ticker time.Ticker, bucketName string) {
	for {
		select {
		case t := <-ticker.C:
			tickerHander(t, bucketName)
		}
	}
}

func tickerHander(t time.Time, bucketName string) {
	for {
		//从bucketName获取job列表判断时间是否已经可以执行
		//job时间到的时候在去查询一下jobID是否存在
		//如果存在放到readQueue让服务端消费
		bucketItem, err := getBucketItem(bucketName)
		if err != nil {
			log.Printf("扫描bucket错误#bucket-%s # %s", bucketName, err.Error())
			return
		}

		//bucket为空
		if bucketItem == nil {
			return
		}

		if bucketItem.Timestamp > t.Unix() {
			return
		}
		//item的延迟时间比当前时间小 通过JobID获取job的详情
		jobInfo, err := getJob(bucketItem.JobID)
		if err != nil {
			return
		}

		if jobInfo == nil {
			//当前job已被删除 把job从bucket中删除不再执行后续逻辑
			if err = removeItem(bucketItem.JobID); err != nil {
				fmt.Printf("删除bucketItem时出错 jobID=%s # %s", bucketItem.JobID, err)
			}
			continue
		}

		//再次确定job中的delay是否小于当前时间
		if jobInfo.Delay <= t.Unix() {
			if err = removeItem(bucketItem.JobID); err != nil {
				fmt.Printf("删除bucketItem时出错 jobID=%s # %s", bucketItem.JobID, err)
			}
		}

		if err := pushJobToReadyQueue(jobInfo.Topic, *jobInfo); err != nil {
			fmt.Printf("添加job到ready队列时出错 jobID=%s # queueName=%s # %s")
		}
	}

}
