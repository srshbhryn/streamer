package webserver

type config_t struct {
	port  string
	debug bool
}

var config config_t

func init() {
	port := "8080"
	debug := true
	config = config_t{
		port:  port,
		debug: debug,
	}
}
