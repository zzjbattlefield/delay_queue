package queue

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

type JobItem struct {
	ID    string `json:"id"`    //jobID job的唯一标识
	Topic string `json:"topic"` //job的类别
	Delay int64  `json:"delay"` //job的延迟时间
	TTR   int64  `json:"ttr"`   //job执行的最长时间
	Body  string `json:"body"`  //jobBody json格式存储
}

func getJob(jobID string) (*JobItem, error) {
	val, err := RedisDB.Get(context.Background(), jobID).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		log.Printf("获取job详情时发生错误 jobID:%s # %s", jobID, err.Error())
		return nil, err
	}
	jobItem := &JobItem{}
	json.Unmarshal([]byte(val), jobItem)
	return jobItem, nil
}

func setJob(jobID string, jobInfo JobItem) error {
	jobJson, err := json.Marshal(jobInfo)
	if err != nil {
		log.Printf("job转码时发生错误 jobID:%s # %s", jobID, err.Error())
		return err
	}
	err = RedisDB.Set(context.Background(), jobID, jobJson, 0).Err()
	if err != nil {
		log.Printf("设置job时发生错误 jobID:%s # %s", jobID, err.Error())
	}
	return nil
}
