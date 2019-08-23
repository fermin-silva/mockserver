package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//TODO return file object after parsing with config, template, etc
func resolveFile(servingdir, finalpath, url string, config *Config) (*ParsedFile, error) {
	//TODO if path is /user/create there could be a create.json file, so use "findFirst" also for the base case
	//TODO what is the precedence? index.*, exact match? directory? etc

	info, err := os.Stat(finalpath)

	if os.IsPermission(err) {
		return findFirst(servingdir, "403", url, config)
	}

	if os.IsNotExist(err) {
		fmt.Printf("%s does not exist\n", finalpath)

		parsed, err := findFirstGlob(finalpath, url, config)

		if err == nil && parsed != nil {
			return parsed, nil
		}

		//try 404 inside the requested path
		//ex: /users/john --> look in /users/john/404
		parsed, err = findFirst(finalpath, "404", url, config)

		if err == nil && parsed != nil {
			return parsed, nil
		}

		//remove last url element and look for 404 in there
		//ex: /users/john does not exist, so look for /users/404 and only then in /404
		parentpath := filepath.Dir(finalpath)

		parsed, err = findFirst(parentpath, "404", url, config)

		if err == nil && parsed != nil {
			return parsed, nil
		}

		//try the base dir to find a 404 file
		return findFirst(servingdir, "404", url, config)
	}

	if info.IsDir() {
		findFirst(servingdir, "index", url, config)
		//TODO if no index, return 404 inside there?
	} else {
		return ParseFile(finalpath, config) //TODO should we do a url matching at this point if its an exact match?
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

		//TODO this probably needs caching, for each request we are re-parsing everything
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
}
