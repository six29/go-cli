package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	ext  string
	size int64
	list bool
	del  bool
	wLog io.Writer
}

func main() {
	root := flag.String("root", ".", "Root directory to start")
	logFile := flag.String("log", "", "Log deletes to this file")
	ext := flag.String("ext", "", "File extension to filter")
	size := flag.Int64("size", 0, "Minimum file size to filter")
	list := flag.Bool("list", false, "List files only")
	del := flag.Bool("del", false, "Delete found files")
	flag.Parse()

	var (
		f = os.Stdout
	)

	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
		del:  *del,
		wLog: f,
	}

	if *logFile != "" {
		f, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
		c.wLog = f
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run(root string, out io.Writer, cfg config) error {
	delLogger := log.New(cfg.wLog, "DELETED FILE: ", log.LstdFlags)

	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filterOut(path, cfg.ext, cfg.size, info) {
			return nil
		}

		if cfg.list {
			return listFile(path, out)
		}

		if cfg.del {
			return delFile(path, delLogger)
		}

		return listFile(path, out)

	})
}
