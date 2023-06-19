package log

import (
	"log"
	"sync"
	"testing"
)

func TestPanic(t *testing.T) {

	log.Println("hello testpanic start")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				log.Println("defer func", i)
				// err := recover()
				// if err != nil {
				// 	log.Printf("defer recover err:%v", err)
				// }
				wg.Done()
			}()
			logPanic(i)
		}(i)
	}
	wg.Wait()
	log.Println("testpanic end")

}

func logPanic(i int) {
	log.Panicf("%v log panic!!", i)
}
