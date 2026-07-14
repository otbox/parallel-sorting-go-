package service

import (
	"fmt"
	"sync"
	"time"
)

type groupSingleton struct {
	wg sync.WaitGroup
}

var (
	instance *groupSingleton
	once     sync.Once
)

func InitGroup(numThreads int) *sync.WaitGroup {
	once.Do(func() {
		instance = &groupSingleton{}
		instance.wg.Add(numThreads)
	})
	return &instance.wg
}

func GetGroup() *sync.WaitGroup {
	if instance == nil {
		panic("O grupo Singleton precisa ser inicializado primeiro usando InitGroup(numThreads)")
	}
	return &instance.wg
}

func Execute[T any](id int, f func() T) T {
	defer GetGroup().Done()
	beginTime := time.Now()
	defer func() {
		fmt.Printf("Process id %d ends. Duration: %s\n", id, time.Since(beginTime))
	}()
	return f()
}
