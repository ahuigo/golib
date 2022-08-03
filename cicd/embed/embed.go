package main
import (
    "embed"
    "html/template"
    "log"
    "net/http"
)
var (
    //go:embed resource
    res embed.FS
    pages = map[string]string{
        "/": "resource/index.html",
    }
    //go:embed resource/index.html
    s string
)
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        page, ok := pages[r.URL.Path]
        if !ok {
            w.WriteHeader(http.StatusNotFound)
            return
        }
        tpl, err := template.ParseFS(res, page)
        if err != nil {
            log.Printf("page %s not found in pages cache...", r.RequestURI)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "text/html")
        w.WriteHeader(http.StatusOK)
        data := map[string]interface{}{
            "userAgent": r.UserAgent(),
        }
        if err := tpl.Execute(w, data); err != nil {
            return
        }
    })
    http.FileServer(http.FS(res))
    log.Println("server started...")
    err := http.ListenAndServe(":8088", nil)
    if err != nil {
        panic(err)
    }
}
