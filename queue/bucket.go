package queue

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/zzjbattlefield/delay_queue/config"
)

type BucketItem struct {
	Timestamp int64
	JobID     string
}

func getBucketItem(bucketName string) (*BucketItem, error) {
	val, err := RedisDB.ZRangeWithScores(context.Background(), bucketName, 0, 0).Result()
	if err != nil {
		return nil, err
	}
	if len(val) == 0 {
		return nil, nil
	}

	buckInfo := &BucketItem{JobID: val[0].Member.(string), Timestamp: int64(val[0].Score)}
	return buckInfo, nil
}

func removeItem(bucketName string, jobID string) error {
	if err := RedisDB.ZRem(context.Background(), bucketName, jobID).Err(); err != nil {
		return err
	}
	return nil

}

func setBucketItem(bucketName string, delay int64, jobID string) error {
	return RedisDB.ZAdd(context.Background(), bucketName, &redis.Z{Score: float64(delay), Member: jobID}).Err()
}

//轮询获取bucket的名称
func generateBucketName() chan string {
	ch := make(chan string)
	go func() {
		i := 1
		for {
			ch <- fmt.Sprintf(config.Setting.BucketName, i)
			if i >= config.Setting.BucketSize {
				i = 1
			} else {
				i++
			}
		}
	}()
	return ch
}
