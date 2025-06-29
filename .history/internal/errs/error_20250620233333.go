package errs

import (
	"fmt"
	"time"
)

// NewErrIndexOutOfRange 创建一个代表下标超出范围的错误
func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("go-tools: 下标超出范围，长度 %d, 下标 %d", length, index)
}

// NewErrInvalidType 创建一个代表类型转换失败的错误
func NewErrInvalidType(want string, got any) error {
	return fmt.Errorf("go-tools: 类型转换失败，预期类型:%s, 实际值:%#v", want, got)
}

func NewErrInvalidIntervalValue(interval time.Duration) error {
	return fmt.Errorf("go-tools: 无效的间隔时间 %d, 预期值应大于 0", interval)
}

func NewErrInvalidMaxIntervalValue(maxInterval, initialInterval time.Duration) error {
	return fmt.Errorf("go-tools: 最大重试间隔的时间 [%d] 应大于等于初始重试的间隔时间 [%d] ", maxInterval, initialInterval)
}

func NewErrRetryExhausted(lastErr error) error {
	return fmt.Errorf("go-tools: 超过最大重试次数，业务返回的最后一个 error %w", lastErr)
}
