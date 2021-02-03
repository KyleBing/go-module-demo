package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	// 如果没有名字，返回错误
	if name == "" {
		return "", errors.New("name wasn't giving")
	}
	msg := fmt.Sprintf(randomFormat(), name)
	// 如果成功，error 返回 nil
	return msg, nil
}

// 批量问候
func Hellos(name []string) (map[string]string, error) {
	msgs := make(map[string]string)
	for _, name := range name {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		msgs[name] = msg
	}
	return msgs, nil
}

// init 方法初始化随机数的种子
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Welcome, %v",
		"Hello, %v",
		"Hey, %v how are you",
	}
	return formats[rand.Intn(len(formats))]
}
