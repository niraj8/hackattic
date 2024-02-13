package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"niraj8/hackattic/challenges/helper"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/forPelevin/gomoji"
	"github.com/redis/go-redis/v9"
)

func main() {
	var problem struct {
		RDB          string `json:"rdb"`
		Requirements struct {
			CheckTypeOf string `json:"check_type_of"`
		} `json:"requirements"`
	}
	err := helper.GetChallenge("the_redis_one", &problem)
	if err != nil {
		log.Fatalf("error fetching challenge: %v", err)
	}

	log.Println(problem)

	solution := make(map[string]interface{}, 0)
	solution[problem.Requirements.CheckTypeOf] = nil

	// type Solution struct {
	// 	DBCount       int    `json:"db_count"`
	// 	EmojiKeyValue string `json:"emoji_key_value"`
	// 	ExpiryMillis  int    `json:"expiry_millis"`
	// }

	// dump the rdb file to disk
	decodedRDB, err := base64.StdEncoding.DecodeString(problem.RDB)
	if err != nil {
		log.Fatalf("error decoding rdb contents: %v", err)
	}

	// remove the first 5 characters
	decodedRDB = decodedRDB[5:]
	decodedRDB = append([]byte("REDIS"), decodedRDB...)
	err = os.WriteFile("snapshot.rdb", decodedRDB, 0644)
	if err != nil {
		log.Fatalf("error writing the rdb snapshot file to disk: %v", err)
	}

	ctx := context.Background()

	// start the redis server
	err = exec.CommandContext(ctx, "redis-server", "--dbfilename", "snapshot.rdb").Start()
	if err != nil {
		log.Fatalf("error starting the redis server: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	keyspaces, err := rdb.Info(ctx, "keyspace").Result()
	if err != nil {
		log.Fatalf("error fetching keyspace info: %v", err)
	}

	dbIndices := make([]int, 0)

	dbs := strings.Split(keyspaces, "\n")
	for _, db := range dbs {
		if strings.HasPrefix(db, "db") {
			dbIndex, err := strconv.Atoi(strings.Replace(strings.Split(db, ":")[0], "db", "", 1))
			if err != nil {
				log.Fatalf("error converting db index to int:%v", err)
			}
			dbIndices = append(dbIndices, dbIndex)
		}
	}
	solution["db_count"] = len(dbIndices)

	fmt.Println(keyspaces)

	fmt.Println("db indices", dbIndices)

	for _, dbIndex := range dbIndices {
		rdb = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   dbIndex,
		})
		keys, err := rdb.Keys(ctx, "*").Result()
		if err != nil {
			log.Fatalf("error fetching keys for db: %v", err)
		}

		for _, key := range keys {
			expiryDuration, err := rdb.PExpireTime(ctx, key).Result()
			if err != nil {
				log.Fatalf("error fetching expirytime: %v", err)
			}
			valueType, err := rdb.Type(ctx, key).Result()
			if err != nil {
				log.Fatalf("error fetching type: %v", err)
			}

			if problem.Requirements.CheckTypeOf == key {
				solution[key] = valueType
			}

			if gomoji.ContainsEmoji(key) {
				value, err := rdb.Get(ctx, key).Result()
				if err != nil {
					log.Fatalf("error getting value: %v", err)
				}
				solution["emoji_key_value"] = value
			}

			fmt.Println(key, expiryDuration, valueType)
		}

	}

	resp, err := helper.SubmitChallengeSolution("the_redis_one", solution)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}
