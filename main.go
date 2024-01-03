package main

import (
	s "github.com/butbkadrug/multitracks-scraper-go/internal"
)


func main(){

    url := "https://www.multitracks.com/songs/Paul-Baloche/Paul-Baloche-Live/Hosanna---Praise-is-Rising-(Live)/"

    song, err := s.NewSong(url)

    if err != nil {
        panic(err)
    }

    params := &s.SaveProjectParams{
        Template: "/home/butbkadrug/empty.RPP",
        Dest: "/home/butbkadrug/",
        Song: song,

    }

    if err = s.SaveProject(params); err != nil {
        panic(err)
    }
}
