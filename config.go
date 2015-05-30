package webwindow

type Config struct {
	Port   int
	Host   string
	Root   string
	Width  int
	Height int
	Title  string
}

func NewConfig(port int, title string) Config {
	config := Config{
		Host:   "localhost",
		Port:   port,
		Root:   "index.html",
		Width:  640,
		Height: 480,
		Title:  title,
	}
	return config
}
