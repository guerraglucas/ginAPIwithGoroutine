package handlers

import (
	"sync"

	"github.com/gin-gonic/gin"
)

func GetSomething(c *gin.Context) {
	var something []string
	var lock sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			something = append(something, "Something")
		}()
	}
	wg.Wait()
	c.JSON(200, gin.H{
		"message": something,
	})
}
