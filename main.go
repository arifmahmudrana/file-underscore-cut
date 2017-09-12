package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	zglob "github.com/mattn/go-zglob"
)

func main()  {
	filesToBeRenamed := make(map[string]string)
	matches, _ := zglob.Glob(os.Args[1] + string(os.PathSeparator) + `**` + string(os.PathSeparator) + `video*.mp4`)
	re := regexp.MustCompile(`(^video[1-9][0-9]*_[1-9][0-9]*)_[1-9][0-9]*(\.mp4$)`)
	for _, path := range matches {
		dir := filepath.Dir(path)
		file := filepath.Base(path)
		if re.Match([]byte(file)) {
			filesToBeRenamed[path] = dir + string(os.PathSeparator) + re.ReplaceAllString(file, `$1$2`)
		}
	}

	for o, n := range filesToBeRenamed {
		if e := os.Rename(o, n); e != nil {
			fmt.Println("Can't rename file " + o + " to " + n)
			os.Exit(1)
		}
	}
}
