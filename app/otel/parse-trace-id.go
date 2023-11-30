package otel

import (
	"strings"

	"github.com/gin-gonic/gin"
)
func GetTraceIDFromCtx(ctx *gin.Context) string {
        r := ctx.Request
        traceparent := r.Header.Get("traceparent")
        if traceparent == "" {
                return ""
        }

        /* 在 W3C TraceContext 中，有两种header传trace：
        	traceparent: "version-traceid-parentid-traceflags"
				例如: "00-0af7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01"
			b3: 00-xxxx-00
		gin框架中, middleware/otel.go 会解析这个header:
			"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
			e.Use(otelgin.Middleware(AppName))
				ctx := cfg.Propagators.Extract(savedCtx, propagation.HeaderCarrier(c.Request.Header))
		*/


        parts := strings.Split(traceparent, "-")
        if len(parts) < 2 {
                return ""
        }

        // 返回 trace_id 部分
        return parts[1]
}
