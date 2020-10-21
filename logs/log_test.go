package logs

import (
    "fmt"
    "testing"
)

func TestInitLogger(t *testing.T) {
   InitLogger("../logs/test.log",200,5,30)
   GetLog().Error(fmt.Sprintf("FBI WARNING"))
}

func BenchmarkGetLog(b *testing.B) {
    InitLogger("../logs/test.log",200,5,30)
    for i := 0; i < b.N; i ++ {
        GetLog().Error(fmt.Sprintf("FBI WARNING"))
    }
}