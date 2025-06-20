package errs

import (
	"fmt"
	"time"
)

// NewErrIndexOutOfRange creates an error representing an out-of-range index.
func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("go-tools: index out of range, length %d, index %d", length, index)
}

// NewErrInvalidType creates an error representing a failed type conversion.
func NewErrInvalidType(want string, got any) error {
	return fmt.Errorf("go-tools: type conversion failed, expected type:%s, actual value:%#v", want, got)
}

func NewErrInvalidIntervalValue(interval time.Duration) error {
	return fmt.Errorf("go-tools: invalid interval duration %d, expected value should be greater than 0", interval)
}

func NewErrInvalidMaxIntervalValue(maxInterval, initialInterval time.Duration) error {
	return fmt.Errorf("go-tools: max retry interval [%d] should be greater than or equal to the initial retry interval [%d]", maxInterval, initialInterval)
}

func NewErrRetryExhausted(lastErr error) error {
	return fmt.Errorf("go-tools: exceeded maximum retry attempts, last error returned by business logic: %w", lastErr)
}
