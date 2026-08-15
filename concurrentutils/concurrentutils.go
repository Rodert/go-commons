// Package concurrentutils 提供并发相关的工具函数
// Package concurrentutils provides concurrency utility functions
package concurrentutils

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

// WorkerPool 工作池，用于并发执行任务
// WorkerPool is a pool of workers for concurrent task execution
type WorkerPool struct {
	workers   int
	taskQueue chan func()
	workerWG  sync.WaitGroup
	taskWG    sync.WaitGroup
	submitWG  sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
	startOnce sync.Once
	stopOnce  sync.Once
	stateMu   sync.Mutex
	stopped   bool
	stopDone  chan struct{}
}

// ErrWorkerPoolStopped is returned when submitting to a stopped worker pool.
var ErrWorkerPoolStopped = errors.New("worker pool is stopped")

// NewWorkerPool 创建新的工作池
//
// 参数 / Parameters:
//   - workers: 工作协程数量 / number of worker goroutines
//
// 返回值 / Returns:
//   - *WorkerPool: 工作池实例 / worker pool instance
//
// 示例 / Example:
//
//	pool := NewWorkerPool(10)
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
		stopDone:  make(chan struct{}),
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
//
//	pool.Start()
//
// Start starts the worker pool
func (wp *WorkerPool) Start() {
	wp.stateMu.Lock()
	defer wp.stateMu.Unlock()
	if wp.stopped {
		return
	}
	wp.startOnce.Do(func() {
		for i := 0; i < wp.workers; i++ {
			wp.workerWG.Add(1)
			go wp.worker()
		}
	})
}

// worker 工作协程
// worker is a worker goroutine
func (wp *WorkerPool) worker() {
	defer wp.workerWG.Done()
	for task := range wp.taskQueue {
		func() {
			defer wp.taskWG.Done()
			if task != nil {
				task()
			}
		}()
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
//
//	err := pool.Submit(func() {
//	    // 执行任务
//	})
//
// Submit submits a task to the worker pool
func (wp *WorkerPool) Submit(task func()) error {
	wp.Start()

	wp.stateMu.Lock()
	if wp.stopped {
		wp.stateMu.Unlock()
		return ErrWorkerPoolStopped
	}
	wp.submitWG.Add(1)
	wp.taskWG.Add(1)
	wp.stateMu.Unlock()
	defer wp.submitWG.Done()

	select {
	case <-wp.ctx.Done():
		wp.taskWG.Done()
		return ErrWorkerPoolStopped
	case wp.taskQueue <- task:
		return nil
	}
}

// Stop 停止工作池，等待已入队的任务全部完成后再返回
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - 无 / none
//
// 示例 / Example:
//
//	pool.Stop()
//
// Stop stops the worker pool and waits for all queued tasks to finish
func (wp *WorkerPool) Stop() {
	wp.stopOnce.Do(func() {
		wp.stateMu.Lock()
		wp.stopped = true
		wp.cancel()
		wp.stateMu.Unlock()

		// Wait for blocked submitters to observe cancellation before closing the queue.
		wp.submitWG.Wait()
		close(wp.taskQueue)
		wp.taskWG.Wait()
		wp.workerWG.Wait()
		close(wp.stopDone)
	})
	<-wp.stopDone
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
//
//	pool.Wait()
//
// Wait waits for tasks accepted before the call to complete.
func (wp *WorkerPool) Wait() {
	wp.taskWG.Wait()
}

// RateLimiter 限流器，用于控制请求速率（令牌桶算法）
// RateLimiter limits the rate of requests using a token bucket algorithm
type RateLimiter struct {
	limit    int64 // 每秒允许的请求数 / requests per second
	tokens   int64 // 当前可用令牌数（单位：纳秒等价令牌） / available tokens in nanosecond units
	lastTime int64 // 上次更新时间（纳秒） / last update time in nanoseconds
	mu       sync.Mutex
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
//
//	limiter := NewRateLimiter(100) // 每秒100个请求
//
// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit int) *RateLimiter {
	if limit <= 0 {
		limit = 1
	}
	now := time.Now().UnixNano()
	return &RateLimiter{
		limit:    int64(limit),
		tokens:   int64(limit) * int64(time.Second), // 初始满令牌，单位为纳秒
		lastTime: now,
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
//
//	if limiter.Allow() {
//	    // 处理请求
//	}
//
// Allow checks if a request is allowed
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now().UnixNano()
	elapsed := now - rl.lastTime

	// tokens 以纳秒为单位累积，每纳秒产生 limit/1e9 个令牌。
	// 全部用整数运算，消除浮点截断误差。
	// cap = limit * 1e9（纳秒），即满桶上限
	capNs := rl.limit * int64(time.Second)
	rl.tokens += elapsed * rl.limit
	if rl.tokens > capNs {
		rl.tokens = capNs
	}
	rl.lastTime = now

	// 消耗一个令牌（等价于 1e9 纳秒令牌）
	if rl.tokens >= int64(time.Second) {
		rl.tokens -= int64(time.Second)
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
//
//	err := limiter.Wait(ctx)
//
// Wait waits until a request is allowed
func (rl *RateLimiter) Wait(ctx context.Context) error {
	// 每个令牌的间隔时间 = 1秒 / limit
	// interval per token = 1s / limit
	interval := time.Second / time.Duration(rl.limit)
	for {
		if rl.Allow() {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(interval):
		}
	}
}

// SafeCounter 并发安全的计数器，使用 atomic 操作
// SafeCounter is a thread-safe counter using atomic operations
type SafeCounter struct {
	value int64
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
//
//	counter := NewSafeCounter(0)
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
//
//	newValue := counter.Increment(1)
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
//
//	newValue := counter.Decrement(1)
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
//
//	value := counter.Get()
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
//
//	counter.Set(100)
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
//
//	oldValue := counter.Reset()
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
//
//	newValue := counter.Add(10)
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
//
//	cache := NewSafeCache()
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
//
//	cache.Set("key", "value")
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
//
//	value, exists := cache.Get("key")
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
//
//	cache.Delete("key")
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
//
//	if cache.Has("key") { ... }
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
//
//	cache.Clear()
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
//
//	size := cache.Size()
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
//
//	keys := cache.Keys()
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
//
//	value, existed := cache.GetOrSet("key", "default")
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
//
//	value := cache.GetOrCompute("key", func() interface{} {
//	    return expensiveComputation()
//	})
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
