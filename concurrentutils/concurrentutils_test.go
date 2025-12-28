package concurrentutils

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewWorkerPool(t *testing.T) {
	tests := []struct {
		name    string
		workers int
		want    int
	}{
		{"normal", 10, 10},
		{"zero", 0, 1},
		{"negative", -1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pool := NewWorkerPool(tt.workers)
			if pool.workers != tt.want {
				t.Errorf("NewWorkerPool() workers = %v, want %v", pool.workers, tt.want)
			}
		})
	}
}

func TestWorkerPool_Submit(t *testing.T) {
	pool := NewWorkerPool(2)
	pool.Start()
	defer pool.Stop()

	var counter int64
	var wg sync.WaitGroup
	tasks := 10

	for i := 0; i < tasks; i++ {
		wg.Add(1)
		err := pool.Submit(func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		})
		if err != nil {
			t.Errorf("Submit() error = %v", err)
		}
	}

	wg.Wait()

	if counter != int64(tasks) {
		t.Errorf("WorkerPool executed %v tasks, want %v", counter, tasks)
	}
}

func TestWorkerPool_Stop(t *testing.T) {
	pool := NewWorkerPool(2)
	pool.Start()

	// 提交一些任务
	// Submit some tasks
	var counter int64
	for i := 0; i < 5; i++ {
		pool.Submit(func() {
			atomic.AddInt64(&counter, 1)
		})
	}

	// 停止工作池
	// Stop the pool
	pool.Stop()

	// 尝试提交新任务应该失败
	// Attempting to submit new tasks should fail
	err := pool.Submit(func() {})
	if err == nil {
		t.Errorf("Submit() after Stop() should return error")
	}
}

func TestNewRateLimiter(t *testing.T) {
	tests := []struct {
		name  string
		limit int
		want  int64
	}{
		{"normal", 100, 100},
		{"zero", 0, 1},
		{"negative", -1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limiter := NewRateLimiter(tt.limit)
			if limiter.limit != tt.want {
				t.Errorf("NewRateLimiter() limit = %v, want %v", limiter.limit, tt.want)
			}
		})
	}
}

func TestRateLimiter_Allow(t *testing.T) {
	limiter := NewRateLimiter(10) // 每秒10个请求
	// 10 requests per second

	// 前10个请求应该被允许
	// First 10 requests should be allowed
	allowed := 0
	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			allowed++
		}
	}

	if allowed != 10 {
		t.Errorf("RateLimiter.Allow() allowed %v requests, want 10", allowed)
	}

	// 第11个请求应该被拒绝（在没有等待的情况下）
	// 11th request should be denied (without waiting)
	if limiter.Allow() {
		t.Errorf("RateLimiter.Allow() should deny 11th request immediately")
	}

	// 等待一段时间后应该允许
	// After waiting, should allow
	time.Sleep(200 * time.Millisecond)
	if !limiter.Allow() {
		t.Errorf("RateLimiter.Allow() should allow after waiting")
	}
}

func TestRateLimiter_Wait(t *testing.T) {
	limiter := NewRateLimiter(5) // 每秒5个请求
	// 5 requests per second

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 消耗所有令牌
	// Consume all tokens
	for i := 0; i < 5; i++ {
		limiter.Allow()
	}

	// 等待应该成功（在超时之前）
	// Wait should succeed (before timeout)
	err := limiter.Wait(ctx)
	if err != nil {
		t.Errorf("RateLimiter.Wait() error = %v, want nil", err)
	}
}

func TestRateLimiter_Wait_Timeout(t *testing.T) {
	limiter := NewRateLimiter(1) // 每秒1个请求
	// 1 request per second

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// 消耗令牌
	// Consume token
	limiter.Allow()

	// 等待应该超时
	// Wait should timeout
	err := limiter.Wait(ctx)
	if err == nil {
		t.Errorf("RateLimiter.Wait() should timeout")
	}
}

func TestNewSafeCounter(t *testing.T) {
	counter := NewSafeCounter(10)
	if counter.Get() != 10 {
		t.Errorf("NewSafeCounter() initial value = %v, want 10", counter.Get())
	}
}

func TestSafeCounter_Increment(t *testing.T) {
	counter := NewSafeCounter(0)

	result := counter.Increment(1)
	if result != 1 {
		t.Errorf("Increment() = %v, want 1", result)
	}

	result = counter.Increment(5)
	if result != 6 {
		t.Errorf("Increment() = %v, want 6", result)
	}
}

func TestSafeCounter_Decrement(t *testing.T) {
	counter := NewSafeCounter(10)

	result := counter.Decrement(1)
	if result != 9 {
		t.Errorf("Decrement() = %v, want 9", result)
	}

	result = counter.Decrement(5)
	if result != 4 {
		t.Errorf("Decrement() = %v, want 4", result)
	}
}

func TestSafeCounter_Get(t *testing.T) {
	counter := NewSafeCounter(42)
	if counter.Get() != 42 {
		t.Errorf("Get() = %v, want 42", counter.Get())
	}
}

func TestSafeCounter_Set(t *testing.T) {
	counter := NewSafeCounter(0)
	counter.Set(100)
	if counter.Get() != 100 {
		t.Errorf("Set() value = %v, want 100", counter.Get())
	}
}

func TestSafeCounter_Reset(t *testing.T) {
	counter := NewSafeCounter(50)
	oldValue := counter.Reset()
	if oldValue != 50 {
		t.Errorf("Reset() old value = %v, want 50", oldValue)
	}
	if counter.Get() != 0 {
		t.Errorf("Reset() new value = %v, want 0", counter.Get())
	}
}

func TestSafeCounter_Add(t *testing.T) {
	counter := NewSafeCounter(10)
	result := counter.Add(5)
	if result != 15 {
		t.Errorf("Add() = %v, want 15", result)
	}
}

func TestSafeCounter_Concurrent(t *testing.T) {
	counter := NewSafeCounter(0)
	var wg sync.WaitGroup
	goroutines := 100
	opsPerGoroutine := 100

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				counter.Increment(1)
			}
		}()
	}

	wg.Wait()
	expected := int64(goroutines * opsPerGoroutine)
	if counter.Get() != expected {
		t.Errorf("Concurrent operations: counter = %v, want %v", counter.Get(), expected)
	}
}

func TestNewSafeCache(t *testing.T) {
	cache := NewSafeCache()
	if cache == nil {
		t.Errorf("NewSafeCache() = nil, want non-nil")
	}
	if cache.Size() != 0 {
		t.Errorf("NewSafeCache() size = %v, want 0", cache.Size())
	}
}

func TestSafeCache_SetAndGet(t *testing.T) {
	cache := NewSafeCache()

	cache.Set("key1", "value1")
	cache.Set("key2", 42)

	value1, exists1 := cache.Get("key1")
	if !exists1 {
		t.Errorf("Get() exists = false, want true")
	}
	if value1 != "value1" {
		t.Errorf("Get() = %v, want 'value1'", value1)
	}

	value2, exists2 := cache.Get("key2")
	if !exists2 {
		t.Errorf("Get() exists = false, want true")
	}
	if value2 != 42 {
		t.Errorf("Get() = %v, want 42", value2)
	}
}

func TestSafeCache_Delete(t *testing.T) {
	cache := NewSafeCache()
	cache.Set("key", "value")

	cache.Delete("key")

	_, exists := cache.Get("key")
	if exists {
		t.Errorf("Delete() key still exists")
	}
}

func TestSafeCache_Has(t *testing.T) {
	cache := NewSafeCache()

	if cache.Has("key") {
		t.Errorf("Has() = true, want false")
	}

	cache.Set("key", "value")
	if !cache.Has("key") {
		t.Errorf("Has() = false, want true")
	}
}

func TestSafeCache_Clear(t *testing.T) {
	cache := NewSafeCache()
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	cache.Clear()

	if cache.Size() != 0 {
		t.Errorf("Clear() size = %v, want 0", cache.Size())
	}
}

func TestSafeCache_Size(t *testing.T) {
	cache := NewSafeCache()

	if cache.Size() != 0 {
		t.Errorf("Size() = %v, want 0", cache.Size())
	}

	cache.Set("key1", "value1")
	if cache.Size() != 1 {
		t.Errorf("Size() = %v, want 1", cache.Size())
	}

	cache.Set("key2", "value2")
	if cache.Size() != 2 {
		t.Errorf("Size() = %v, want 2", cache.Size())
	}
}

func TestSafeCache_Keys(t *testing.T) {
	cache := NewSafeCache()
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	keys := cache.Keys()
	if len(keys) != 3 {
		t.Errorf("Keys() length = %v, want 3", len(keys))
	}

	keyMap := make(map[string]bool)
	for _, k := range keys {
		keyMap[k] = true
	}

	if !keyMap["key1"] || !keyMap["key2"] || !keyMap["key3"] {
		t.Errorf("Keys() missing expected keys")
	}
}

func TestSafeCache_GetOrSet(t *testing.T) {
	cache := NewSafeCache()

	// 第一次调用应该设置值
	// First call should set value
	value1, existed1 := cache.GetOrSet("key", "default")
	if existed1 {
		t.Errorf("GetOrSet() existed = true, want false")
	}
	if value1 != "default" {
		t.Errorf("GetOrSet() = %v, want 'default'", value1)
	}

	// 第二次调用应该返回已存在的值
	// Second call should return existing value
	value2, existed2 := cache.GetOrSet("key", "new")
	if !existed2 {
		t.Errorf("GetOrSet() existed = false, want true")
	}
	if value2 != "default" {
		t.Errorf("GetOrSet() = %v, want 'default'", value2)
	}
}

func TestSafeCache_GetOrCompute(t *testing.T) {
	cache := NewSafeCache()
	callCount := 0

	compute := func() interface{} {
		callCount++
		return "computed"
	}

	// 第一次调用应该计算值
	// First call should compute value
	value1 := cache.GetOrCompute("key", compute)
	if value1 != "computed" {
		t.Errorf("GetOrCompute() = %v, want 'computed'", value1)
	}
	if callCount != 1 {
		t.Errorf("GetOrCompute() compute called %v times, want 1", callCount)
	}

	// 第二次调用应该返回缓存的值
	// Second call should return cached value
	value2 := cache.GetOrCompute("key", compute)
	if value2 != "computed" {
		t.Errorf("GetOrCompute() = %v, want 'computed'", value2)
	}
	if callCount != 1 {
		t.Errorf("GetOrCompute() compute called %v times, want 1", callCount)
	}
}

func TestSafeCache_Concurrent(t *testing.T) {
	cache := NewSafeCache()
	var wg sync.WaitGroup
	goroutines := 100
	opsPerGoroutine := 100

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				key := "key"
				cache.Set(key, id*opsPerGoroutine+j)
				cache.Get(key)
				cache.Has(key)
			}
		}(i)
	}

	wg.Wait()

	// 验证最终状态
	// Verify final state
	if cache.Size() != 1 {
		t.Errorf("Concurrent operations: size = %v, want 1", cache.Size())
	}
}

func TestRateLimiter_Allow_WithTime(t *testing.T) {
	limiter := NewRateLimiter(2) // 每秒2个请求
	// 2 requests per second

	// 前2个请求应该被允许
	// First 2 requests should be allowed
	if !limiter.Allow() {
		t.Errorf("RateLimiter.Allow() should allow 1st request")
	}
	if !limiter.Allow() {
		t.Errorf("RateLimiter.Allow() should allow 2nd request")
	}

	// 第3个请求应该被拒绝
	// 3rd request should be denied
	if limiter.Allow() {
		t.Errorf("RateLimiter.Allow() should deny 3rd request")
	}

	// 等待半秒后应该允许
	// After waiting half second, should allow
	time.Sleep(600 * time.Millisecond)
	if !limiter.Allow() {
		t.Errorf("RateLimiter.Allow() should allow after waiting")
	}
}

