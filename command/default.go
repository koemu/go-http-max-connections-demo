package command

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/codegangsta/cli"
)

func CmdDefault(c *cli.Context) {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		time.Sleep(time.Millisecond * 500) // for watch netstat
		io.WriteString(w, "Default")
	})

	http_config := &http.Server{
		Addr: fmt.Sprintf(":%d", c.Int("port")),
	}
	err := http_config.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
