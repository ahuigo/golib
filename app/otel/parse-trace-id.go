func GetTraceIDFromCtx(ctx *gin.Context) string {
        r := ctx.Request
        traceparent := r.Header.Get("traceparent")
        if traceparent == "" {
                return ""
        }

        // 在 W3C TraceContext 中，traceparent 的格式是:
        // "version-traceid-parentid-traceflags"
        // 例如: "00-0af7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01"
        // 还有一种格式的header, b3: 00-xxxx-00

        parts := strings.Split(traceparent, "-")
        if len(parts) < 2 {
                return ""
        }

        // 返回 trace_id 部分
        return parts[1]
}
