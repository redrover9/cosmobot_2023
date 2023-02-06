package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/kosa3/pexels-go"
)

func getWord(path string) ([]string, error) {
	verbFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer verbFile.Close()

	var lines []string
	scanner := bufio.NewScanner(verbFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()

}

func getCaption() string {
	verbLines, err := getWord("sexVerbs.txt")
	if err != nil {
		log.Fatal(err)
	}
	partLines, err := getWord("bodyParts.txt")
	if err != nil {
		log.Fatal(err)
	}
	objectLines, err := getWord("householdObjects.txt")
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	min := 1
	verbMax := len(verbLines) - 1
	partMax := len(partLines) - 1
	objectMax := len(objectLines) - 1

	verb := verbLines[rand.Intn(verbMax-min)+min]
	part := partLines[rand.Intn(partMax-min)+min]
	object := objectLines[rand.Intn(objectMax-min)+min]

	caption := verb + " his " + part + " with a " + object
	return caption

}

func getPhotoURL() string {
	cli := pexels.NewClient("abc-123")
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

func getFont() string {
	fonts := []string{"Cedarville_Cursive/CedarvilleCursive-Regular.ttf", "Crafty_Girls/CraftyGirls-Regular.ttf", "Dancing_Script/DancingScript-VariableFont_wght.ttf", "Delius_Swash_Caps/DeliusSwashCaps-Regular.ttf", "Indie_Flower/IndieFlower-Regular.ttf", "Sassy_Frass/SassyFrass-Regular.ttf"}
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 4
	font := fonts[rand.Intn(max-min)+min]
	return font
}

func main() {

}
