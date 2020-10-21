package base

import (
    "errors"
    "fmt"
    "log"
    "testing"
)

func TestRetryDoSuccessful(t *testing.T)  {
    SomeFunction := func() (string, error) {
        panic("something went badly wrong")
    }
    var value string
    err := Retry(func(attempt int) (retry bool, err error) {
        retry = attempt < 5 // try 5 times
        defer func() {
            if r := recover(); r != nil {
                err = errors.New(fmt.Sprintf("panic: %v", r))
            }
        }()
        value, err = SomeFunction()
        log.Println(value)
        return
    })
    if err != nil {
        //log.Fatalln("error:", err)
    }
}