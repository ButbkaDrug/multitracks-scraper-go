package internal

import(
    "fmt"
    "strings"
    "strconv"
    colly "github.com/gocolly/colly/v2"
)

type Song struct {
    Title string
    Album string
    Artist string
    Tempo string
    Key string
    Signature string
    Length string
    Songmap  []string
    Lyrics string
}

func NewSong(url string) (*Song, error) {
    var s = Song{}
    c := colly.NewCollector()

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting ", r.URL)
    })

    c.OnHTML("h1.song-banner--title", func(h *colly.HTMLElement) {
        s.Title = strings.Trim(h.Text, " \n\t")
    })

    c.OnHTML("h2.song-banner--artist a", func(h *colly.HTMLElement) {
        text := strings.Trim(h.Text," \n\t")

        if h.Index == 0 {
            s.Artist = text
            return
        }

        s.Album = text
    })

    c.OnHTML("div.song-banner--song-sections--list--item", func(h *colly.HTMLElement) {
        var repeat int
        var err error

        if repeat, err = strconv.Atoi(h.Attr("data-repeat")); err != nil {
            repeat = 1
        }

        for i:=0; i < repeat; i++ {
            s.Songmap = append(s.Songmap, h.ChildTexts("span")[0])
        }
    })

    c.OnHTML("dl.song-banner--meta-list", func(h *colly.HTMLElement) {
        h.ForEach("dd.song-banner--meta-list--desc", func(i int, h *colly.HTMLElement) {

            value := strings.Trim(h.Text, "\t\n")
            switch i {
            case 0:
                s.Key = value
            case 1:
                s.Tempo = value
            case 2:
                s.Signature = value
            case 3:
                s.Length = value

            }

        })

    })

    c.OnHTML("div.section-expand--block", func(h *colly.HTMLElement) {
        h.ForEach("p", func(i int, h *colly.HTMLElement) {


            var line string


            if i == 1 {

                if html, err := h.DOM.Html(); err == nil{
                    line = strings.ReplaceAll(html, "&#39;", "'")
                    line = strings.ReplaceAll(line, "<br/>", "\n")
                }

            } else {
                line = "\n" + h.Text + "\n"
            }

            s.Lyrics += line

        })
    })

    if err := c.Visit(url); err != nil {
        return &s, err
    }

    return &s, nil

}
