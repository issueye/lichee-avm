package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
	licheeJs "github.com/issueye/lichee-js"
)

type AutoJsController struct {
	ScriptTimeoutSec int
}

func (autoJs *AutoJsController) AutoJsReceiveServer(ctx *gin.Context) {
	core := licheeJs.NewCore()
	vm := core.GetRts()
	licheeJs.NewRequest(vm, ctx.Request)
	licheeJs.NewResponse(vm, ctx.Writer)

	normalEndCh := make(chan bool)
	go func() {
		err := core.Run("AutoJsController", "index.js")
		if err != nil {
			switch err := err.(type) {
			case *goja.Exception:
				ctx.Writer.WriteHeader(http.StatusInternalServerError)
				fmt.Println("*goja.Exception:", err.String())
				if v := err.Value().ToObject(vm).Get("nativeType"); v != nil {
					fmt.Printf("%T:%[1]v\n", v.Export())
				}
			case *goja.InterruptedError:
				fmt.Println("*goja.InterruptedError:", err.String())
			default:
				fmt.Println("default err:", err)
			}
			normalEndCh <- false
			return
		}

		// if This.writeResultValue {
		// 	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 	w.Write([]byte((ret).String()))
		// }
		normalEndCh <- true
	}()

	if autoJs.ScriptTimeoutSec > 0 {
		timeoutCh := time.After(time.Duration(autoJs.ScriptTimeoutSec) * time.Second)
		select {
		case <-timeoutCh:
			vm.Interrupt("run code timeout, halt")
			ctx.Writer.WriteHeader(http.StatusInternalServerError)
		case <-normalEndCh:
		}
	} else {
		status := <-normalEndCh
		// 如果失败则直接返回失败
		if !status {
			ctx.JSON(http.StatusOK,
				map[string]any{
					"success": "false",
					"message": "失败",
				})
		}
	}
}
