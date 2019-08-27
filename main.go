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
	run(os.Args[1:])
}

func run(osargs []string) {
	port := pflag.IntP("port", "p", 8080, "Http listening port")
	confFile := pflag.StringP("conf", "c", "", "Service configuration file")
	pflag.CommandLine.Parse(osargs)

	args := pflag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "error: file serving directory missing\nusage: mockserver [flags] servingDir\n")
		os.Exit(1)
	}

	servingDir := args[0]

	//TODO validate servingDir

	var conf *Config
	var err error

	if confFile == nil || *confFile == "" {
		conf = &DefaultConfig
	} else {
		conf, err = ParseConfig(*confFile)

		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading config file: %s\n", err)
			os.Exit(1)
		}
	}

	pongo2.DefaultLoader.SetBaseDir(servingDir)

	//TODO better logging
	//TODO add gzip support
	//TODO add more utility filters to pongo2
	r := gin.Default()
	r.GET("/*path", NewHttpHandler(conf, servingDir))
	r.Run("0.0.0.0:" + strconv.Itoa(*port))
}
