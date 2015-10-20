package command

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/netutil"

	"github.com/codegangsta/cli"
)

func CmdLimited(c *cli.Context) {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		time.Sleep(time.Millisecond * 500) // for watch netstat
		io.WriteString(w, "Limited")
	})

	port := fmt.Sprintf(":%d", c.Int("port"))
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	limit_listener := netutil.LimitListener(listener, c.Int("max-connections"))
	http_config := &http.Server{
		Addr: port,
	}

	defer limit_listener.Close()
	err = http_config.Serve(limit_listener)
	if err != nil {
		log.Fatalln(err)
	}
}
