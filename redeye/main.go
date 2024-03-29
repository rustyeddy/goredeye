package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/rustyeddy/redeye"
)

type Config struct {
	Addr     string
	Pipeline string
	Device   interface{}
}

var config Config

func init() {
	flag.StringVar(&config.Addr, "addr", ":1234", "Listent to address")
	flag.StringVar(&config.Pipeline, "pipeline", "", "Pipeline to apply")
}

// go:embed index.html
// var content embed.FS

func main() {
	flag.Parse()

	srv := redeye.NewWebServer(config.Addr)
	srv.Handle("/", http.FileServer(http.Dir("./html")))

	if len(os.Args) < 1 {
		log.Println("No video capture devices specified")
		return
	}

	vidsrcs := getVideoSrcs(os.Args[1:])
	for i, vsrc := range vidsrcs {
		mjpg := redeye.NewMJPEGPlayer(i)
		srv.Handle(mjpg.URL(), mjpg)

		imgQ := vsrc.Play()
		mjpg.Play(imgQ)
	}
	srv.Listen()
}

func getVideoSrcs(args []string) []*redeye.VideoSource {

	var capdevs []*redeye.VideoSource
	for _, capstr := range os.Args[1:] {

		// Open up the video capture device
		cap := redeye.GetVideoSource(capstr)
		if cap == nil {
			log.Println("Failed to get capture device", capstr)
			os.Exit(1)
		}

		if config.Pipeline != "" {
			log.Println("Looking up pipcap")
		}
	}
	return capdevs

}
