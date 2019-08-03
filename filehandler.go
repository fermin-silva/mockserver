package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//TODO return file object after parsing with config, template, etc
func resolveFile(servingdir, filepath, url string, config *Config) (*ParsedFile, error) {
	//TODO if path is /user/create there could be a create.json file, so use "findFirst" also for the base case
	//TODO what is the precedence? index.*, exact match? directory? etc

	info, err := os.Stat(filepath)

	if os.IsPermission(err) {
		return findFirst(servingdir, "403", url, config)
	}

	if os.IsNotExist(err) {
		fmt.Printf("%s does not exist\n", filepath)

		parsed, err := findFirstGlob(filepath, url, config)

		if err == nil && parsed != nil {
			return parsed, nil
		}

		//try 404 inside the requested path
		parsed, err = findFirst(filepath, "404", url, config)

		if err == nil && parsed != nil {
			return parsed, nil
		}

		//try the base dir to find a 404 file
		return findFirst(servingdir, "404", url, config)
	}

	if info.IsDir() {
		findFirst(servingdir, "index", url, config)
	} else {
		return ParseFile(filepath, config)
	}

	//TODO throw error at this point
	return nil, nil
}

func findFirstGlob(findpath, url string, config *Config) (*ParsedFile, error) {
	fmt.Printf("Finding first glob file name %s\n", findpath)

	matches, err := filepath.Glob(findpath + "*")

	if err != nil {
		return nil, err
	}

	fmt.Printf("Glob matches: %v\n", matches)

	for _, name := range matches {
		f, err := os.Stat(name)

		if err != nil {
			return nil, err
		}

		if f.IsDir() {
			continue
		}

		if parsed, err := ParseFile(name, config); err != nil {
			return nil, err
		} else if parsed.Matches(url) {
			return parsed, nil
		}
	}

	return nil, nil
}

func findFirst(dir, name, url string, config *Config) (*ParsedFile, error) {
	fmt.Printf("Finding first file named %s in dir %s\n", name, dir)

	return findFirstGlob(filepath.Join(dir, name), url, config)

	//TODO implement return first file name that matches name (and not the extension)
	//TODO or do here file url matching based on file config 'matches' ?
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		fmt.Printf("Trying to find first with file %s\n", f.Name())

		//TODO caching files? regex compilation? etc?

		if strings.HasPrefix(f.Name(), name) {
			if parsed, err := ParseFile(path.Join(dir, f.Name()), config); err != nil {
				return nil, err
			} else if parsed.Matches(url) {
				return parsed, nil
			}
		}
	}

	return nil, nil
}
