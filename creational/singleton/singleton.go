package singleton

import (
	"sync"
	"time"
)

var singletonInstance *singleton
var singletonLock sync.Mutex

type singleton struct{}

func GetSingleton() (*singleton, string) {
	if singletonInstance != nil {
		return singletonInstance, "singleton instance has already been created!"
	}

	singletonLock.Lock()
	defer singletonLock.Unlock()

	time.Sleep(500 * time.Millisecond)

	if singletonInstance == nil {
		singletonInstance = &singleton{}
		return singletonInstance, "creating a singleton instance!"
	}

	return singletonInstance, "mutex lock acquired, but the singleton instance has already been created!"
}
