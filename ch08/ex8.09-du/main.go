// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 250.

// The du3 command computes the disk usage of the files in a directory.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type FileSizeResp struct {
	rootid int
	size   int64
}

//!+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//!+
	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan FileSizeResp)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(root, i, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	//!-

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

	nfiles := make(map[string]int64)
	nbytes := make(map[string]int64)
loop:
	for {
		select {
		case sizeResp, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles[roots[sizeResp.rootid]]++
			nbytes[roots[sizeResp.rootid]] += sizeResp.size
		case <-tick:
			printDiskUsage(nfiles, nbytes, roots)
		}
	}

	printDiskUsage(nfiles, nbytes, roots) // final totals
	//!+
	// ...select loop...
}

//!-

func printDiskUsage(nfiles, nbytes map[string]int64, roots []string) {
	fmt.Println()
	for _, r := range roots {
		fmt.Printf("%s : %d files  %.1f GB\n", r, nfiles[r], float64(nbytes[r])/1e9)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+walkDir
func walkDir(dir string, rootid int, n *sync.WaitGroup, fileSizes chan<- FileSizeResp) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, rootid, n, fileSizes)
		} else {
			fileSizes <- FileSizeResp{rootid, entry.Size()}
		}
	}
}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
