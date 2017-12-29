package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/prestonp/boop/deploy"
	"github.com/prestonp/boop/server"
)

var (
	host       = flag.String("host", "0.0.0.0", "server host")
	port       = flag.String("port", "8080", "server port")
	scriptPath = flag.String("script.path", "./deploy.sh", "path to deploy script")
	logPath    = flag.String("log.path", "/tmp/deploy-logs", "path to log files (note this will get wiped on startup)")
	ghSecret   = flag.String("github_secret", os.Getenv("GITHUB_SECRET"), "the github webook secret")
)

func main() {
	flag.Parse()
	deployer, err := deploy.New(*scriptPath, *logPath)
	if err != nil {
		panic(err)
	}

	addr := net.JoinHostPort(*host, *port)
	svc := server.New(deployer, *ghSecret)
	fmt.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, svc))
}
