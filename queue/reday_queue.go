package queue

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zzjbattlefield/delay_queue/config"
)

func pushJobToReadyQueue(queueName string, jobInfo JobItem) error {
	readyQueueName := fmt.Sprintf(config.Setting.QueueName, queueName)
	jobJson, err := json.Marshal(jobInfo)
	if err != nil {
		return err
	}
	return RedisDB.RPush(context.Background(), readyQueueName, jobJson).Err()
}
