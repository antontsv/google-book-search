package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	books "google.golang.org/api/books/v1"
)

func fatal(msg string) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", msg))
	os.Exit(1)
}

func fatalIfErr(err error) {
	if err != nil {
		fatal(err.Error())
	}
}

func main() {
	if len(os.Args) < 2 {
		fatal("Need Book search string an argument")
	}
	svc, err := books.New(http.DefaultClient)
	fatalIfErr(err)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	searchStr := strings.Join(os.Args[1:], " ")
	volumes, err := svc.Volumes.List(searchStr).Context(ctx).Do()
	fatalIfErr(err)

	fmt.Println("Search results:")
	for i, v := range volumes.Items {
		fmt.Printf("%d. %s\n", i+1, v.VolumeInfo.Title)
	}

}
