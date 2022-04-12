package queue

import (
	"context"

	"github.com/go-redis/redis/v8"
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
	if val == nil {
		return nil, nil
	}
	buckInfo := &BucketItem{JobID: val[0].Member.(string), Timestamp: int64(val[0].Score)}
	return buckInfo, nil
}

func removeItem(jobID string) error {
	if err := RedisDB.Del(context.Background(), jobID).Err(); err != nil {
		return err
	}
	return nil

}

func setBucketItem(bucketName string, delay int64, jobID string) error {
	return RedisDB.ZAdd(context.Background(), bucketName, &redis.Z{Score: float64(delay), Member: jobID}).Err()
}
