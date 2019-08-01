package main

import (
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"os"
	"strconv"
)

func main() {
	port := pflag.IntP("port", "p", 8080, "Http listening port")
	confFile := pflag.StringP("conf", "c", "conf.toml", "Service configuration file")
	pflag.Parse()

	args := pflag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "error: file serving directory missing\nusage: mockserver [flags] servingDir\n")
		os.Exit(1)
	}

	servingDir := args[0]

	//TODO validate servingDir

	conf, err := ParseConfig(*confFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading config file: %s\n", err)
		os.Exit(1)
	}

	pongo2.DefaultLoader.SetBaseDir(servingDir)

	//TODO better logging
	//TODO add gzip support
	//TODO add more utility filters to pongo2
	r := gin.Default()
	r.GET("/*path", NewHttpHandler(conf, servingDir))
	r.Run("0.0.0.0:" + strconv.Itoa(*port))
}
