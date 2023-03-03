package retries

import (
	"errors"
	"fmt"
	retry "github.com/avast/retry-go"
	"time"
)

const RETRYABLE = "retryable"
const NOT_RETRYABLE = "not-retryable"

func willPassOnThirdTry(n int) error {
	var errMessage string
	if n < 3 {
		errMessage = RETRYABLE
	} else if n < 5 {
		errMessage = NOT_RETRYABLE
	}
	if len(errMessage) == 0 {
		return nil
	}
	fmt.Println(errMessage)
	return errors.New(errMessage)
}

func isRetryable(err error) bool {
	if err.Error() == RETRYABLE {
		return true
	}
	return false
}

func RetryFailingFunction() {
	val := 0
	err := retry.Do(
		func() error {
			val += 2
			err := willPassOnThirdTry(val)
			return err
		},
		retry.DelayType(func(n uint, config *retry.Config) time.Duration {
			return retry.BackOffDelay(n, config)
		}),
		retry.Attempts(3),
		retry.RetryIf(isRetryable),
	)
	if err == nil {
		fmt.Printf("Final value of val: %v", val)
	} else {
		fmt.Printf("Final value of err: %v", err)
	}
}
