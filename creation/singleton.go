package creation

import (
	"sync"
)

type Singleton struct{}

var eagerSingleton *Singleton

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func init() {
	eagerSingleton = &Singleton{}
}

func GetEagerInstance() *Singleton {
	return eagerSingleton
}

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}

	return lazySingleton
}
