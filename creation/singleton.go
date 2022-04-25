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
	// Initialize singleton when theis package is imported for the first time
	eagerSingleton = &Singleton{}
}

func EagerSingleton() *Singleton {
	return eagerSingleton
}

func LazySingleton() *Singleton {
	if lazySingleton == nil {
		// Only the first invoke will be executed
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}

	return lazySingleton
}
