package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Rodert/go-commons/concurrentutils"
)

func main() {
	fmt.Println("=== Concurrent Utils Examples ===\n")

	// 示例1: Worker Pool
	// Example 1: Worker Pool
	fmt.Println("1. Worker Pool:")
	pool := concurrentutils.NewWorkerPool(3)
	pool.Start()
	defer pool.Stop()

	var counter int64
	var wg sync.WaitGroup
	tasks := 10

	for i := 0; i < tasks; i++ {
		wg.Add(1)
		taskID := i
		pool.Submit(func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
			fmt.Printf("   Task %d executed\n", taskID)
		})
	}

	wg.Wait()
	fmt.Printf("   Total tasks executed: %d\n\n", counter)

	// 示例2: Rate Limiter
	// Example 2: Rate Limiter
	fmt.Println("2. Rate Limiter:")
	limiter := concurrentutils.NewRateLimiter(5) // 每秒5个请求
	// 5 requests per second

	allowed := 0
	denied := 0
	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			allowed++
			fmt.Printf("   Request %d: Allowed\n", i+1)
		} else {
			denied++
			fmt.Printf("   Request %d: Denied\n", i+1)
		}
	}
	fmt.Printf("   Allowed: %d, Denied: %d\n\n", allowed, denied)

	// 示例3: Rate Limiter with Wait
	// Example 3: Rate Limiter with Wait
	fmt.Println("3. Rate Limiter with Wait:")
	limiter2 := concurrentutils.NewRateLimiter(2) // 每秒2个请求
	// 2 requests per second

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println("   Making 5 requests with rate limiting...")
	for i := 0; i < 5; i++ {
		if err := limiter2.Wait(ctx); err != nil {
			fmt.Printf("   Request %d: Error - %v\n", i+1, err)
			break
		}
		fmt.Printf("   Request %d: Allowed\n", i+1)
	}
	fmt.Println()

	// 示例4: Safe Counter
	// Example 4: Safe Counter
	fmt.Println("4. Safe Counter:")
	counter2 := concurrentutils.NewSafeCounter(0)

	// 并发增加
	// Concurrent increment
	var wg2 sync.WaitGroup
	goroutines := 10
	opsPerGoroutine := 100

	for i := 0; i < goroutines; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				counter2.Increment(1)
			}
		}()
	}

	wg2.Wait()
	fmt.Printf("   Initial value: 0\n")
	fmt.Printf("   Increments: %d goroutines × %d ops = %d\n", goroutines, opsPerGoroutine, goroutines*opsPerGoroutine)
	fmt.Printf("   Final value: %d\n", counter2.Get())

	// 测试其他操作
	// Test other operations
	counter2.Add(50)
	fmt.Printf("   After Add(50): %d\n", counter2.Get())

	counter2.Decrement(25)
	fmt.Printf("   After Decrement(25): %d\n", counter2.Get())

	oldValue := counter2.Reset()
	fmt.Printf("   Reset() returned: %d, new value: %d\n\n", oldValue, counter2.Get())

	// 示例5: Safe Cache
	// Example 5: Safe Cache
	fmt.Println("5. Safe Cache:")
	cache := concurrentutils.NewSafeCache()

	// 基本操作
	// Basic operations
	cache.Set("name", "Go Commons")
	cache.Set("version", "1.0.0")
	cache.Set("count", 42)

	if name, exists := cache.Get("name"); exists {
		fmt.Printf("   name: %v\n", name)
	}

	if version, exists := cache.Get("version"); exists {
		fmt.Printf("   version: %v\n", version)
	}

	fmt.Printf("   Cache size: %d\n", cache.Size())
	fmt.Printf("   Has 'count': %v\n", cache.Has("count"))

	// GetOrSet
	// GetOrSet
	value1, existed1 := cache.GetOrSet("new_key", "new_value")
	fmt.Printf("   GetOrSet('new_key', 'new_value'): value=%v, existed=%v\n", value1, existed1)

	value2, existed2 := cache.GetOrSet("new_key", "another_value")
	fmt.Printf("   GetOrSet('new_key', 'another_value'): value=%v, existed=%v\n", value2, existed2)

	// GetOrCompute
	// GetOrCompute
	computeCount := 0
	computedValue := cache.GetOrCompute("computed", func() interface{} {
		computeCount++
		return fmt.Sprintf("computed_value_%d", computeCount)
	})
	fmt.Printf("   GetOrCompute('computed'): %v (computed %d times)\n", computedValue, computeCount)

	computedValue2 := cache.GetOrCompute("computed", func() interface{} {
		computeCount++
		return "should_not_compute"
	})
	fmt.Printf("   GetOrCompute('computed') again: %v (computed %d times)\n", computedValue2, computeCount)

	// 获取所有键
	// Get all keys
	keys := cache.Keys()
	fmt.Printf("   All keys: %v\n", keys)

	// 删除
	// Delete
	cache.Delete("count")
	fmt.Printf("   After deleting 'count', size: %d\n", cache.Size())

	// 清空
	// Clear
	cache.Clear()
	fmt.Printf("   After Clear(), size: %d\n\n", cache.Size())

	// 示例6: 并发安全的缓存操作
	// Example 6: Concurrent safe cache operations
	fmt.Println("6. Concurrent Cache Operations:")
	concurrentCache := concurrentutils.NewSafeCache()
	var wg3 sync.WaitGroup
	concurrentOps := 100

	for i := 0; i < concurrentOps; i++ {
		wg3.Add(1)
		go func(id int) {
			defer wg3.Done()
			key := fmt.Sprintf("key_%d", id%10) // 只有10个不同的键
			// Only 10 different keys
			concurrentCache.Set(key, id)
			concurrentCache.Get(key)
			concurrentCache.Has(key)
		}(i)
	}

	wg3.Wait()
	fmt.Printf("   Completed %d concurrent operations\n", concurrentOps)
	fmt.Printf("   Final cache size: %d\n", concurrentCache.Size())
}

