package types

import "time"

type IRetryable interface {
	Retry() error
}
type Retryable[T any] struct{}

func (r *Retryable[T]) Retry(fn func() (T, error), retries int) (T, error) {
	var result T
	var err error
	for i := 0; i < retries; i++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}
	}
	return result, err
}

type IRetryableVoid interface {
	RetryVoid() error
}
type RetryableVoid struct{}

func (rv *RetryableVoid) Retry(fn func() error, retries int) error {
	for i := 0; i < retries; i++ {
		if err := fn(); err == nil {
			return nil
		}
	}
	return nil
}

type IRetryableWithDelay interface {
	RetryWithDelay(delayMs int) error
}
type RetryableWithDelay[T any] struct{}

func (r *RetryableWithDelay[T]) Retry(fn func() (T, error), retries int, delayMs int) (T, error) {
	var result T
	var err error
	for i := 0; i < retries; i++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}
		if delayMs > 0 {
			time.Sleep(time.Duration(delayMs) * time.Millisecond)
		}
	}
	return result, err
}

type IRetryableWithTimeout interface {
	RetryWithTimeout(timeout time.Duration) error
}
type RetryableWithTimeout[T any] struct{}

func (r *RetryableWithTimeout[T]) Retry(fn func() (T, error), retries int, timeout time.Duration) (T, error) {
	startTime := time.Now()
	for i := 0; i < retries; i++ {
		if time.Since(startTime) > timeout {
			return *new(T), nil // Timeout reached, return zero value without error
		}
		if result, err := fn(); err == nil {
			return result, nil
		}
	}
	return *new(T), nil
}

type IRetryableWithDelayAndTimeout interface {
	RetryWithDelayAndTimeout(delayMs int, timeout time.Duration) error
}
type RetryableWithDelayAndTimeout[T any] struct{}

func (r *RetryableWithDelayAndTimeout[T]) Retry(fn func() (T, error), retries int, delayMs int, timeout time.Duration) (T, error) {
	startTime := time.Now()
	var result T
	var err error
	for i := 0; i < retries; i++ {
		if time.Since(startTime) > timeout {
			return result, nil // Timeout reached, return zero value without error
		}
		result, err = fn()
		if err == nil {
			return result, nil
		}
		if delayMs > 0 {
			time.Sleep(time.Duration(delayMs) * time.Millisecond)
		}
	}
	return result, err
}

type IRetryableWithDelayAndTimeoutAndRetries interface {
	RetryWithDelayAndTimeoutAndRetries(delayMs int, timeout time.Duration, retries int) error
}
type RetryableWithDelayAndTimeoutAndRetries[T any] struct{}

func (r *RetryableWithDelayAndTimeoutAndRetries[T]) Retry(fn func() (T, error), retries int, delayMs int, timeout time.Duration) (T, error) {
	startTime := time.Now()
	var result T
	var err error
	for i := 0; i < retries; i++ {
		if time.Since(startTime) > timeout {
			return result, nil // Timeout reached, return zero value without error
		}
		result, err = fn()
		if err == nil {
			return result, nil
		}
		if delayMs > 0 {
			time.Sleep(time.Duration(delayMs) * time.Millisecond)
		}
	}
	return result, err
}

type IRetryableWithRetries interface {
	RetryWithRetries(retries int) error
}
type RetryableWithRetries[T any] struct{}

func (r *RetryableWithRetries[T]) Retry(fn func() (T, error), retries int) (T, error) {
	var result T
	var err error
	for i := 0; i < retries; i++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}
	}
	return result, err
}

type IRetryableVoidWithDelay interface {
	RetryVoidWithDelay(delayMs int) error
}
type RetryableVoidWithDelay struct{}

func (rv *RetryableVoidWithDelay) Retry(fn func() error, retries int, delayMs int) error {
	for i := 0; i < retries; i++ {
		if err := fn(); err == nil {
			return nil
		}
		if delayMs > 0 {
			time.Sleep(time.Duration(delayMs) * time.Millisecond)
		}
	}
	return nil
}

type IRetryableVoidWithTimeout interface {
	RetryVoidWithTimeout(timeout time.Duration) error
}
type RetryableVoidWithTimeout struct{}

func (rv *RetryableVoidWithTimeout) Retry(fn func() error, retries int, timeout time.Duration) error {
	startTime := time.Now()
	for i := 0; i < retries; i++ {
		if time.Since(startTime) > timeout {
			return nil // Timeout reached, return without error
		}
		if err := fn(); err == nil {
			return nil
		}
	}
	return nil
}

type IRetryableVoidWithDelayAndTimeout interface {
	RetryVoidWithDelayAndTimeout(delayMs int, timeout time.Duration) error
}
type RetryableVoidWithDelayAndTimeout struct{}

func (rv *RetryableVoidWithDelayAndTimeout) Retry(fn func() error, retries int, delayMs int, timeout time.Duration) error {
	startTime := time.Now()
	for i := 0; i < retries; i++ {
		if time.Since(startTime) > timeout {
			return nil // Timeout reached, return without error
		}
		if err := fn(); err == nil {
			return nil
		}
		if delayMs > 0 {
			time.Sleep(time.Duration(delayMs) * time.Millisecond)
		}
	}
	return nil
}
