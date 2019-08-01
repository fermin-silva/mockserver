package main

import (
	"os"
)

//TODO return file object after parsing with config, template, etc
func resolveFile(servingdir, filepath string, config *Config) (*ParsedFile, error) {
	//TODO if path is /user/create there could be a create.json file, so use "findFirst" also for the base case
	//TODO what is the precedence? index.*, exact match? directory? etc

	info, err := os.Stat(filepath)

	if os.IsPermission(err) {
		//TODO find if there is 403.whatever file
		findFirst(servingdir, "403")
	}

	if os.IsNotExist(err) {
		//TODO find if there is 404.whatever file
		findFirst(servingdir, "404")
	}

	if info.IsDir() {
		findFirst(servingdir, "index")
	} else {
		return ParseFile(filepath, config)
	}

	//TODO throw error at this point
	return nil, nil
}

func findFirst(dir, name string) {
	//TODO implement return first file name that matches name (and not the extension)
	//TODO or do here file url matching based on file config 'matches' ?
}
