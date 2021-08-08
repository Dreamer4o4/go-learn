package src

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func recoverHandler() handlerfunc {
	return func(ctxt *Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("-----------------------\r\n")
				log.Printf("%s\n\n", fmt.Sprintf("%s", trace(fmt.Sprintf("%s", err))))
				ctxt.Resw.WriteHeader(http.StatusInternalServerError)
				log.Printf("-----------------------\r\n")
			}
		}()

		ctxt.NextStep()
	}
}

func logger() handlerfunc {
	return func(ctxt *Context) {
		// Start timer
		t := time.Now()
		// Process request
		ctxt.NextStep()
		// Calculate resolution time
		log.Printf("[%s] %s in %v", ctxt.Req.URL, ctxt.Req.Method, time.Since(t))
	}
}
