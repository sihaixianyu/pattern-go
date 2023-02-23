package singleton

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

var printLock sync.Mutex

func TestSingleton(t *testing.T) {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			_, info := GetSingleton()

			log := strings.Builder{}
			log.WriteString(fmt.Sprintf("[INFO] Goroutine[%2d]: ", i))
			log.WriteString(info)

			printLock.Lock()
			fmt.Println(log.String())
			printLock.Unlock()
		}(i)

		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()
}
