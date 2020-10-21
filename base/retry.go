package base

import "errors"

var (
    MaxRetries = 10
    errMaxRetriesReached = errors.New("exceeded retry limit")
)

type Func func(attempt int) (retry bool, err error)

func Retry(fn Func) error {
    var (
        err error
        cont bool
    )
    attempt := 1
    for {
        cont,err = fn(attempt)
        if !cont || err != nil {
            break
        }
        attempt++
        if attempt > MaxRetries {
            return errMaxRetriesReached
        }
    }
    return err
}

func IsMaxRetries(err error) bool {
    return err == errMaxRetriesReached
}