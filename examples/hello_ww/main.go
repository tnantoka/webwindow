package main

import (
	ww "github.com/tnantoka/webwindow"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		html := `<body>
      Hello, <a href="https://github.com/tnantoka/webwindow" target="_blank">WebWindow</a>.
      <script src="/ww.js"></script>
    </body>`
		io.WriteString(w, html)
	})

	config := ww.NewConfig(33333, "Example")
	ww.Open(config)
}
