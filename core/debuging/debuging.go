package debuging

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
)

/**
debuging 模块是开发期和生产期使用调试程序模块,
debuging 以Trace开头的标记显示在控制台上的,
debuging 以Log开头的标记为显示在控制台上以及记录在日志中


Trace 显示当前 GoroutineID 堆栈 信息
*/
func Trace(v ...interface{}) {
	fmt.Fprintf(os.Stderr, "[TRACE]:(G:%s) %+#v\n%v", GoroutineID(), GetStackFrame(1), DumpValue(v...))
}
func Tracef(msg string, v ...interface{}) {
	if len(v) == 0 {
		fmt.Fprintf(os.Stderr, "[TRACE]:(G:%s) %+#v\n[Msg]:%s", GoroutineID(), GetStackFrame(1), msg)
	} else {
		fmt.Fprintf(os.Stderr, "[TRACE]:(G:%s) %+#v\n[Msg]:%s\n%v", GoroutineID(), GetStackFrame(1), msg, DumpValue(v...))
	}

}

//TracePanic 出现panic的时候,打印堆栈信息
func TracePanic(v ...interface{}) {
	if x := recover(); x != nil {
		i := 0
		funcName, file, line, ok := runtime.Caller(i)
		for ok {
			fmt.Fprintf(os.Stderr, "frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
			i++
			funcName, file, line, ok = runtime.Caller(i)
		}

		for k := range v {
			fmt.Fprintf(os.Stderr, "EXRAS#%v DATA:%v\n", k, DumpValue(v[k]))
		}
	}
}

//================================================================//
//Pause 暂停程序
//控制台模式
func Pause(isPause bool) {
	if isPause {
		fmt.Fprintf(os.Stderr, "[TRACE]:(G:%s) %+#v", GoroutineID(), GetStack())
		fmt.Fprint(os.Stderr, "press ENTER to continue\n")
		fmt.Scanln()
	}
}

//GoroutineID 获取当前运行的的 Goroutine id
func GoroutineID() string {
	buf := make([]byte, 20)
	buf = buf[:runtime.Stack(buf, false)]
	return string(bytes.Split(buf, []byte(" "))[1])

}

//======================================================//
