package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/kosa3/pexels-go"
)

func getPhotoURL() string {
	cli := pexels.NewClient(os.Getenv("PEXELS_API_KEY"))
	ctx := context.Background()
	ps, err := cli.PhotoService.Search(ctx, &pexels.PhotoParams{
		Query:   "men",
		Page:    1,
		PerPage: 80,
	})
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 80
	sourceURL := ps.Photos[rand.Intn(max-min)+min].Src.Medium
	return sourceURL
}

func getPhoto() int64 {
	resp, err := http.Get(getPhotoURL())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	out, err := os.Create("dude.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	photo, _ := io.Copy(out, resp.Body)
	return photo
}

func main() {
	getPhoto()
}
