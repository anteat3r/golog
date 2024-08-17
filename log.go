package golog

import (
  "runtime"
  "strings"
  "fmt"
  "time"
)

func logMsg(msg string, v ...any) {
  fileres := ""
  pc := make([]uintptr, 10)
  n := runtime.Callers(3, pc)
  var frames *runtime.Frames
  if n == 0 {
    goto fileend
  }
  pc = pc[:n]
  frames = runtime.CallersFrames(pc)
  for {
    frame, more := frames.Next()
    if strings.Contains(frame.File, "asm_amd64") { break }
    if strings.Contains(frame.File, "go/pkg/mod") { break }
    fileres += fmt.Sprintf("%v:%v -> ", frame.File, frame.Line)
    if !more { break }
  }

  fileend:
  res := ""
  for _, e := range v {
    res += fmt.Sprintf("%v %T, ", e, e)
  }
  fmt.Printf(
    "%v %v%v: %v\n\n",
    strings.Split(time.Now().String(), ".")[0],
    fileres,
    msg,
    res,
  )
}

func Error(v ...any) {
  logMsg("ERROR", v...)
}

func Info(v ...any) {
  logMsg("INFO", v...)
}
