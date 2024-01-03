package main

import (
	"fmt"

    s "github.com/butbkadrug/multitracks-scraper-go/internal"
)


func main(){

    url := "https://www.multitracks.com/songs/Jesus-Culture/Your-Love-Never-Fails-(Live)/Your-Love-Never-Fails/"

    song, err := s.NewSong(url)

    if err != nil {
        panic(err)
    }

    fmt.Println(song)
}
