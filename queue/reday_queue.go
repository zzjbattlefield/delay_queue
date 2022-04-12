package queue

import (
	"context"
	"fmt"

	"github.com/zzjbattlefield/delay_queue/config"
)

func pushJobToReadyQueue(queueName string, jobInfo JobItem) error {
	readyQueueName := fmt.Sprintf(config.Setting.QueueName, queueName)
	return RedisDB.RPush(context.Background(), readyQueueName, jobInfo).Err()
}
