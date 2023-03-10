package main

import (
	"bufio"
	"context"
	"fmt"
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

	caption := verb + " his " + part + " with a " + object + "!"
	return caption

}

func getPhotoURL() string {
	cli := pexels.NewClient(os.Getenv("PEXELS_API_KEY"))
	ctx := context.Background()
	ps, err := cli.PhotoService.Search(ctx, &pexels.PhotoParams{
		Query:       "handsome man",
		Orientation: "landscape",
		Page:        1,
		PerPage:     80,
		Size:        "large",
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

func getPhoto() {
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

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func getFont() string {
	fonts := []string{"Cedarville_Cursive/CedarvilleCursive-Regular.ttf", "Crafty_Girls/CraftyGirls-Regular.ttf", "Dancing_Script/DancingScript-VariableFont_wght.ttf", "Delius_Swash_Caps/DeliusSwashCaps-Regular.ttf", "Indie_Flower/IndieFlower-Regular.ttf", "Sassy_Frass/SassyFrass-Regular.ttf"}
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	font := fonts[rand.Intn(max-min)+min]
	return font
}

func main() {
	fmt.Println(getCaption())
	getPhoto()
	fmt.Println(getFont())
}
