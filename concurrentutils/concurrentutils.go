// Package concurrentutils 提供并发相关的工具函数
// Package concurrentutils provides concurrency utility functions
package concurrentutils

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

// WorkerPool 工作池，用于并发执行任务
// WorkerPool is a pool of workers for concurrent task execution
type WorkerPool struct {
	workers    int
	taskQueue  chan func()
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
	once       sync.Once
}

// NewWorkerPool 创建新的工作池
//
// 参数 / Parameters:
//   - workers: 工作协程数量 / number of worker goroutines
//
// 返回值 / Returns:
//   - *WorkerPool: 工作池实例 / worker pool instance
//
// 示例 / Example:
//   pool := NewWorkerPool(10)
//
// NewWorkerPool creates a new worker pool
func NewWorkerPool(workers int) *WorkerPool {
	if workers <= 0 {
		workers = 1
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		workers:   workers,
		taskQueue: make(chan func(), workers*2),
		ctx:       ctx,
		cancel:    cancel,
	}
}

// Start 启动工作池
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - 无 / none
//
// 示例 / Example:
//   pool.Start()
//
// Start starts the worker pool
func (wp *WorkerPool) Start() {
	wp.once.Do(func() {
		for i := 0; i < wp.workers; i++ {
			wp.wg.Add(1)
			go wp.worker()
		}
	})
}

// worker 工作协程
// worker is a worker goroutine
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for {
		select {
		case <-wp.ctx.Done():
			return
		case task, ok := <-wp.taskQueue:
			if !ok {
				return
			}
			if task != nil {
				task()
			}
		}
	}
}

// Submit 提交任务到工作池
//
// 参数 / Parameters:
//   - task: 要执行的任务函数 / task function to execute
//
// 返回值 / Returns:
//   - error: 如果工作池已关闭则返回错误 / error if pool is closed
//
// 示例 / Example:
//   err := pool.Submit(func() {
//       // 执行任务
//   })
//
// Submit submits a task to the worker pool
func (wp *WorkerPool) Submit(task func()) error {
	select {
	case <-wp.ctx.Done():
		return wp.ctx.Err()
	case wp.taskQueue <- task:
		return nil
	default:
		// 如果队列已满，检查是否已关闭
		// If queue is full, check if closed
		select {
		case <-wp.ctx.Done():
			return wp.ctx.Err()
		case wp.taskQueue <- task:
			return nil
		}
	}
}

// Stop 停止工作池
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - 无 / none
//
// 示例 / Example:
//   pool.Stop()
//
// Stop stops the worker pool
func (wp *WorkerPool) Stop() {
	wp.cancel()
	// 等待所有任务完成后再关闭channel
	// Wait for all tasks to complete before closing channel
	wp.wg.Wait()
	close(wp.taskQueue)
}

// Wait 等待所有任务完成
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - 无 / none
//
// 示例 / Example:
//   pool.Wait()
//
// Wait waits for all tasks to complete
func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

// RateLimiter 限流器，用于控制请求速率
// RateLimiter limits the rate of requests
type RateLimiter struct {
	limit     int64         // 每秒允许的请求数 / requests per second
	interval  time.Duration // 时间窗口 / time window
	tokens    int64         // 当前可用令牌数 / current available tokens
	lastTime  int64         // 上次更新时间（纳秒） / last update time in nanoseconds
	mu        sync.Mutex
}

// NewRateLimiter 创建新的限流器
//
// 参数 / Parameters:
//   - limit: 每秒允许的请求数 / requests per second
//
// 返回值 / Returns:
//   - *RateLimiter: 限流器实例 / rate limiter instance
//
// 示例 / Example:
//   limiter := NewRateLimiter(100) // 每秒100个请求
//
// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit int) *RateLimiter {
	if limit <= 0 {
		limit = 1
	}
	return &RateLimiter{
		limit:    int64(limit),
		interval: time.Second,
		tokens:   int64(limit),
		lastTime: time.Now().UnixNano(),
	}
}

// Allow 检查是否允许请求
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - bool: 如果允许则返回true / true if request is allowed
//
// 示例 / Example:
//   if limiter.Allow() {
//       // 处理请求
//   }
//
// Allow checks if a request is allowed
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now().UnixNano()
	elapsed := now - rl.lastTime

	// 计算应该补充的令牌数（每秒补充limit个令牌）
	// Calculate tokens to add (add limit tokens per second)
	// elapsed是纳秒，interval是秒，所以需要转换
	// elapsed is in nanoseconds, interval is in seconds, so need conversion
	elapsedSeconds := float64(elapsed) / float64(time.Second)
	tokensToAdd := int64(elapsedSeconds * float64(rl.limit))
	
	if tokensToAdd > 0 {
		rl.tokens = min(rl.tokens+tokensToAdd, rl.limit)
		rl.lastTime = now
	}

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}

// Wait 等待直到允许请求
//
// 参数 / Parameters:
//   - ctx: 上下文，用于取消等待 / context for cancellation
//
// 返回值 / Returns:
//   - error: 如果上下文被取消则返回错误 / error if context is cancelled
//
// 示例 / Example:
//   err := limiter.Wait(ctx)
//
// Wait waits until a request is allowed
func (rl *RateLimiter) Wait(ctx context.Context) error {
	for {
		if rl.Allow() {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(rl.interval / time.Duration(rl.limit)):
			// 等待一小段时间后重试
			// Wait a short time before retrying
		}
	}
}

// min 返回两个整数中的较小值
// min returns the smaller of two integers
func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// SafeCounter 并发安全的计数器
// SafeCounter is a thread-safe counter
type SafeCounter struct {
	value int64
	mu    sync.RWMutex
}

// NewSafeCounter 创建新的安全计数器
//
// 参数 / Parameters:
//   - initialValue: 初始值 / initial value
//
// 返回值 / Returns:
//   - *SafeCounter: 计数器实例 / counter instance
//
// 示例 / Example:
//   counter := NewSafeCounter(0)
//
// NewSafeCounter creates a new safe counter
func NewSafeCounter(initialValue int64) *SafeCounter {
	return &SafeCounter{
		value: initialValue,
	}
}

// Increment 增加计数器的值
//
// 参数 / Parameters:
//   - delta: 增加的值，默认为1 / value to add, default is 1
//
// 返回值 / Returns:
//   - int64: 增加后的值 / value after increment
//
// 示例 / Example:
//   newValue := counter.Increment(1)
//
// Increment increments the counter value
func (sc *SafeCounter) Increment(delta int64) int64 {
	return atomic.AddInt64(&sc.value, delta)
}

// Decrement 减少计数器的值
//
// 参数 / Parameters:
//   - delta: 减少的值，默认为1 / value to subtract, default is 1
//
// 返回值 / Returns:
//   - int64: 减少后的值 / value after decrement
//
// 示例 / Example:
//   newValue := counter.Decrement(1)
//
// Decrement decrements the counter value
func (sc *SafeCounter) Decrement(delta int64) int64 {
	return atomic.AddInt64(&sc.value, -delta)
}

// Get 获取当前值
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - int64: 当前值 / current value
//
// 示例 / Example:
//   value := counter.Get()
//
// Get gets the current value
func (sc *SafeCounter) Get() int64 {
	return atomic.LoadInt64(&sc.value)
}

// Set 设置计数器的值
//
// 参数 / Parameters:
//   - value: 要设置的值 / value to set
//
// 返回值 / Returns:
//   - 无 / none
//
// 示例 / Example:
//   counter.Set(100)
//
// Set sets the counter value
func (sc *SafeCounter) Set(value int64) {
	atomic.StoreInt64(&sc.value, value)
}

// Reset 重置计数器为0
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - int64: 重置前的值 / value before reset
//
// 示例 / Example:
//   oldValue := counter.Reset()
//
// Reset resets the counter to 0
func (sc *SafeCounter) Reset() int64 {
	return atomic.SwapInt64(&sc.value, 0)
}

// Add 添加值并返回新值
//
// 参数 / Parameters:
//   - delta: 要添加的值 / value to add
//
// 返回值 / Returns:
//   - int64: 添加后的值 / value after adding
//
// 示例 / Example:
//   newValue := counter.Add(10)
//
// Add adds a value and returns the new value
func (sc *SafeCounter) Add(delta int64) int64 {
	return atomic.AddInt64(&sc.value, delta)
}

// SafeCache 并发安全的缓存
// SafeCache is a thread-safe cache
type SafeCache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

// NewSafeCache 创建新的安全缓存
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - *SafeCache: 缓存实例 / cache instance
//
// 示例 / Example:
//   cache := NewSafeCache()
//
// NewSafeCache creates a new safe cache
func NewSafeCache() *SafeCache {
	return &SafeCache{
		data: make(map[string]interface{}),
	}
}

// Set 设置缓存值
//
// 参数 / Parameters:
//   - key: 缓存键 / cache key
//   - value: 缓存值 / cache value
//
// 返回值 / Returns:
//   - 无 / none
//
// 示例 / Example:
//   cache.Set("key", "value")
//
// Set sets a cache value
func (sc *SafeCache) Set(key string, value interface{}) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.data[key] = value
}

// Get 获取缓存值
//
// 参数 / Parameters:
//   - key: 缓存键 / cache key
//
// 返回值 / Returns:
//   - interface{}: 缓存值，如果不存在则返回nil / cache value, nil if not exists
//   - bool: 是否存在 / whether the key exists
//
// 示例 / Example:
//   value, exists := cache.Get("key")
//
// Get gets a cache value
func (sc *SafeCache) Get(key string) (interface{}, bool) {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	value, exists := sc.data[key]
	return value, exists
}

// Delete 删除缓存值
//
// 参数 / Parameters:
//   - key: 缓存键 / cache key
//
// 返回值 / Returns:
//   - 无 / none
//
// 示例 / Example:
//   cache.Delete("key")
//
// Delete deletes a cache value
func (sc *SafeCache) Delete(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	delete(sc.data, key)
}

// Has 检查键是否存在
//
// 参数 / Parameters:
//   - key: 缓存键 / cache key
//
// 返回值 / Returns:
//   - bool: 如果存在则返回true / true if key exists
//
// 示例 / Example:
//   if cache.Has("key") { ... }
//
// Has checks if a key exists
func (sc *SafeCache) Has(key string) bool {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	_, exists := sc.data[key]
	return exists
}

// Clear 清空所有缓存
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - 无 / none
//
// 示例 / Example:
//   cache.Clear()
//
// Clear clears all cache
func (sc *SafeCache) Clear() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.data = make(map[string]interface{})
}

// Size 获取缓存大小
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - int: 缓存中的键值对数量 / number of key-value pairs
//
// 示例 / Example:
//   size := cache.Size()
//
// Size gets the cache size
func (sc *SafeCache) Size() int {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return len(sc.data)
}

// Keys 获取所有键
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - []string: 所有键的列表 / list of all keys
//
// 示例 / Example:
//   keys := cache.Keys()
//
// Keys returns all keys
func (sc *SafeCache) Keys() []string {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	keys := make([]string, 0, len(sc.data))
	for k := range sc.data {
		keys = append(keys, k)
	}
	return keys
}

// GetOrSet 获取值，如果不存在则设置并返回
//
// 参数 / Parameters:
//   - key: 缓存键 / cache key
//   - value: 如果不存在则设置的值 / value to set if not exists
//
// 返回值 / Returns:
//   - interface{}: 缓存值 / cache value
//   - bool: 是否是已存在的值（true表示已存在，false表示新设置） / whether value existed (true if existed, false if newly set)
//
// 示例 / Example:
//   value, existed := cache.GetOrSet("key", "default")
//
// GetOrSet gets a value, or sets it if not exists
func (sc *SafeCache) GetOrSet(key string, value interface{}) (interface{}, bool) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	if v, exists := sc.data[key]; exists {
		return v, true
	}

	sc.data[key] = value
	return value, false
}

// GetOrCompute 获取值，如果不存在则通过函数计算并设置
//
// 参数 / Parameters:
//   - key: 缓存键 / cache key
//   - compute: 计算值的函数 / function to compute value
//
// 返回值 / Returns:
//   - interface{}: 缓存值 / cache value
//
// 示例 / Example:
//   value := cache.GetOrCompute("key", func() interface{} {
//       return expensiveComputation()
//   })
//
// GetOrCompute gets a value, or computes and sets it if not exists
func (sc *SafeCache) GetOrCompute(key string, compute func() interface{}) interface{} {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	if v, exists := sc.data[key]; exists {
		return v
	}

	value := compute()
	sc.data[key] = value
	return value
}

