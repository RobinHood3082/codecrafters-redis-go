package main

import (
	"strconv"
	"time"
)

func pingHandler() string {
	return "+PONG\r\n"
}

func echoHandler(cmd string) string {
	return "$" + strconv.Itoa(len(cmd)) + "\r\n" + cmd + "\r\n"
}

func setHandler(cmd []string) string {
	db.Set(cmd[0], cmd[1])

	if len(cmd) > 2 {
		go expire(cmd[0], cmd[3])
	}

	return "+OK\r\n"
}

func expire(key, expiry string) {
	expiryMS, _ := strconv.Atoi(expiry)
	time.Sleep(time.Duration(expiryMS) * time.Millisecond)
	db.Del(key)
}

func getHandler(cmd string) string {
	val, ok := db.Get(cmd)
	if !ok {
		return "$-1\r\n"
	}
	return "$" + strconv.Itoa(len(val)) + "\r\n" + val + "\r\n"
}
