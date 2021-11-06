package tools

import (
	"fmt"
	"math/rand"
	"time"
)

// 产生随机数
func GetRandomCode() string {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	return code
}
