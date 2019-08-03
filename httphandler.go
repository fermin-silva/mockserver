package main

import (
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
		c.AbortWithError(500, err)
	}

	for k, v := range parsedfile.GetHeaders() {
		c.Header(k, v)
	}

	content, err := parsedfile.String(c)

	if err != nil {
		c.AbortWithError(500, err)
	}

	c.Writer.Write([]byte(content))
}
