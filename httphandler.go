package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
)

func NewHttpHandler(config *Config, servingDir string) func(c *gin.Context) {
	return func(c *gin.Context) {
		serve(servingDir, c, config)
	}
}

func serve(servingDir string, c *gin.Context, config *Config) {
	finalPath := path.Join(servingDir, c.Request.URL.Path)

	//TODO this could have different backends: database, file, s3, etc
	parsedfile, err := resolveFile(servingDir, finalPath, c.Request.URL.Path, config)

	if err != nil {
		fmt.Println("error", err)
		c.AbortWithError(500, err)
		return
	}

	if parsedfile == nil {
		c.AbortWithError(500, fmt.Errorf("file resolver returned nil for path %s", finalPath))
		return
	}

	fmt.Println("Final file returned by resolveFile is", parsedfile.FilePath)

	for k, v := range parsedfile.GetHeaders() {
		c.Header(k, v)
	}

	content, err := parsedfile.String(c)

	if err != nil {
		c.AbortWithError(500, err)
	}

	c.Writer.Write([]byte(content))
}
