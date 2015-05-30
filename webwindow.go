package webwindow

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"text/template"
)

const scriptName = "win.js"

func createTempfile() *os.File {
	f, err := ioutil.TempFile("", scriptName)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func renderTemplate(f *os.File, config Config) {
	t := template.Must(template.New(scriptName).Parse(WinJS))
	t.Execute(f, config)
}

func runScript(f *os.File) {
	osascript := "osascript -l JavaScript " + f.Name()
	cmd := exec.Command("/bin/sh", "-c", osascript)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func handleOpen(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	open := "open " + r.Form["href"][0]
	cmd := exec.Command("/bin/sh", "-c", open)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func handleWWJS(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, WWJS)
}

func initHTTP(config Config) {
	http.HandleFunc("/ww/open", handleOpen)
	http.HandleFunc("/ww.js", handleWWJS)
	go http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
}

func Open(config Config) {
	f := createTempfile()
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	renderTemplate(f, config)
	initHTTP(config)
	runScript(f)
}
