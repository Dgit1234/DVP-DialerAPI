package main

import (
	"encoding/json"
	"fmt"
	"github.com/fzzy/radix/extra/pubsub"
	"github.com/fzzy/radix/redis"
	"time"
)

var subChannelName string

// Redis String Methods
func RedisAdd(key, value string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisSet", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	isExists, _ := client.Cmd("EXISTS", key).Bool()

	if isExists {
		return "Key Already exists"
	} else {
		result, sErr := client.Cmd("set", key, value).Str()
		errHndlr(sErr)
		fmt.Println(result)
		return result
	}
}

func RedisSet(key, value string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisSet", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	result, sErr := client.Cmd("set", key, value).Str()
	errHndlr(sErr)
	fmt.Println(result)
	return result
}

func RedisSetNx(key, value string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisSet", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	result, sErr := client.Cmd("setnx", key, value).Str()
	errHndlr(sErr)
	fmt.Println(result)
	return result
}

func RedisGet(key string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisGet", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	strObj, _ := client.Cmd("get", key).Str()
	fmt.Println(strObj)
	return strObj
}

func RedisSearchKeys(pattern string) []string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisSearchKeys", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	strObj, _ := client.Cmd("keys", pattern).List()
	return strObj
}

func RedisIncr(key string) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisSet", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	result, sErr := client.Cmd("incr", key).Int()
	errHndlr(sErr)
	fmt.Println(result)
	return result
}

func RedisIncrBy(key string, value int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisSet", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	result, sErr := client.Cmd("incrby", key, value).Int()
	errHndlr(sErr)
	fmt.Println(result)
	return result
}

func RedisRemove(key string) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisRemove", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	result, sErr := client.Cmd("del", key).Bool()
	errHndlr(sErr)
	fmt.Println(result)
	return result
}

func RedisCheckKeyExist(key string) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in CheckKeyExist", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	result, sErr := client.Cmd("exists", key).Bool()
	errHndlr(sErr)
	fmt.Println(result)
	return result
}

// Redis Hashes Methods

func RedisHashGetAll(hkey string) map[string]string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisHashGetAll", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)
	strHash, _ := client.Cmd("hgetall", hkey).Hash()
	bytes, err := json.Marshal(strHash)
	if err != nil {
		fmt.Println(err)
	}
	text := string(bytes)
	fmt.Println(text)
	return strHash
}

func RedisHashGetField(hkey, field string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisHashGetAll", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)
	strValue, _ := client.Cmd("hget", hkey, field).Str()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strValue)
	return strValue
}

func RedisHashSetField(hkey, field, value string) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisHashSetField", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	result, _ := client.Cmd("hset", hkey, field, value).Bool()
	return result
}

func RedisHashSetNxField(hkey, field, value string) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisHashSetField", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	result, _ := client.Cmd("hsetnx", hkey, field, value).Bool()
	return result
}

func RedisHashSetMultipleField(hkey string, data map[string]string) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisHashSetField", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)
	fmt.Println(data)
	for key, value := range data {
		client.Cmd("hset", hkey, key, value).Bool()
	}
	fmt.Println(true)
	return true
}

func RedisRemoveHashField(hkey, field string) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisRemoveHashField", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	result, _ := client.Cmd("hdel", hkey, field).Bool()
	return result
}

// Redis List Methods

func RedisListLpop(lname string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisListLpop", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	lpopItem, _ := client.Cmd("lpop", lname).Str()
	fmt.Println(lpopItem)
	return lpopItem
}

func RedisListLpush(lname, value string) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisListLpush", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	result, _ := client.Cmd("lpush", lname, value).Bool()
	return result
}

func RedisListRpush(lname, value string) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisListLpush", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	result, _ := client.Cmd("rpush", lname, value).Bool()
	return result
}

func RedisListLlen(lname string) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in RedisListLlen", r)
		}
	}()
	client, err := redis.DialTimeout("tcp", redisIp, time.Duration(10)*time.Second)
	errHndlr(err)
	defer client.Close()

	// select database
	r := client.Cmd("select", redisDb)
	errHndlr(r.Err)

	result, _ := client.Cmd("llen", lname).Int()
	return result
}

// Redis PubSub

func PubSub() {
	subChannelName = fmt.Sprintf("dialer%s", dialerId)
	c2, err := redis.Dial("tcp", redisIp)
	errHndlr(err)
	defer c2.Close()
	psc := pubsub.NewSubClient(c2)
	psr := psc.Subscribe(subChannelName)
	//ppsr := psc.PSubscribe("*")

	//if ppsr.Err == nil {

	for {
		psr = psc.Receive()
		if psr.Err != nil {

			fmt.Println(psr.Err.Error())

			break
		}

		var subEvent = SubEvents{}
		json.Unmarshal([]byte(psr.Message), &subEvent)
		go OnEvent(subEvent)
	}
	//s := strings.Split("127.0.0.1:5432", ":")
	//}

	psc.Unsubscribe(subChannelName)

}
