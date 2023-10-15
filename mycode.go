// mypackage/mycode.go

package mycode

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

// AddNumbers is a simple function that adds two numbers. 
func Add测试(a, b int) int {
	return a + b
}

func getCurrentFunctionName() string {
	// Caller 的参数为 0 表示当前函数，1 表示调用当前函数的函数，依此类推
	pc, _, _, _ := runtime.Caller(1)
	funcObj := runtime.FuncForPC(pc)
	return funcObj.Name()
}

// 函数用时 用于检查执行时间是否超过限制，并输出警告信息。
// 参数：
//   - tag: 标签用于标识警告的来源
//   - detailed: 详细信息，描述执行的具体内容
//   - start: 开始时间，表示代码开始执行的时间点
//   - timeLimit: 时间限制，表示允许的最长执行时间
func FunTime(start time.Time, timeLimit time.Duration) {
	// 使用time.Since获取经过的时间
	elapsedTime := time.Since(start)

	// 如果经过的时间超过了限制，输出警告信息
	if elapsedTime > timeLimit {
		// elapsedTime = elapsedTime.Truncate(time.Millisecond)

		// strTime := fmt.Sprint(elapsedTime)

		minutes := int(elapsedTime.Minutes()) % 60
		seconds := int(elapsedTime.Seconds()) % 60
		milliseconds := elapsedTime.Milliseconds() % 1000

		// 格式化为字符串
		strTime := fmt.Sprintf("%02d:%02d.%02d", minutes, seconds, milliseconds)

		log.Println(getCurrentFunctionName(), ":", "已用时", strTime)
	}
}
