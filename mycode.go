// mypackage/mycode.go

备份我的代码

导入（
	“FMMT”
	“日志”
	“运行”
	“时间”
）

// AddNumbers1 是一个简单的函数，用于将两个数字相加。
func Add测试(a, b int)整数{
	返回a+b
}

func 测试(a, b int )整数{
	返回a+b
}


func getCurrentFunctionName()字符串{
	// Caller 的参数为 0 表示当前函数，1 表示调用当前函数的函数，此类依推
	pc, _, _, _ := 运行时。来电者( 1 )
	funcObj := 运行时。FuncForPC(pc)
	返回funcObj 。（姓名）
}

// 函数限制时用于检查执行时间是否超过，并输出警告信息。
// 参数：
// - tag：用于标识警告的来源的标签
//-detailed：详细信息，描述执行的具体内容
// - start: 开始时间，表示代码开始执行的时间点
// - timeLimit: 时间限制，表示允许的最后执行时间
func FunTime(开始时间.时间, timeLimit 时间.持续时间) {
	// 使用time.Since获取经过的时间
	经过时间：=时间。自从（开始）

	// 如果经过的时间超过了限制，则输出警告信息
	如果经过时间 > 时间限制 {
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
