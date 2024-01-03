package internal

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type SaveProjectParams struct {
    Template string
    Dest string
    Song *Song
}

func SaveProject(p *SaveProjectParams) error {
    project, err := getTemplate(p.Template)

    if err != nil {
        return err
    }

    project = updateTemplate(project, p.Song)

    path, err := makeDir(p)

    if err != nil {
        return err
    }

    return saveFile(path, []byte(project))
}

func saveFile(p string, data []byte) error {
    filename := fmt.Sprintf("%s.RPP", path.Base(p))
    file := filepath.Join(p, filename)


    if _, err := os.Open(file); err == nil {
        return err
    }

    if err := os.WriteFile(file, data, 0777); err != nil {
        return err
    }

    return nil
}

// Attempts to create a dir for a project. Returns a path where dir were created
func makeDir(p *SaveProjectParams) (string, error) {

    dirName := fmt.Sprintf("%s - %s - %s - %s BPM",
        p.Song.Title,
        p.Song.Artist,
        p.Song.Key,
        p.Song.Tempo,
    )

    dir := filepath.Join(p.Dest, dirName)

    if err := os.Mkdir(dir, 0777); err != nil {
        return dir, err
    }

    return dir, nil
}

func updateTemplate(t string, s *Song) string {
    tempo := fmt.Sprintf("TEMPO %s %s", s.Tempo, s.Signature)

    t = strings.Replace(t, "TEMPO 120 4 4", tempo, 1)
    t = strings.Replace(t, "TITLE TITLE", "TITLE "+s.Title, 1)
    t = strings.Replace(t, "|MAP", "|"+strings.Join(s.Songmap, ""), 1)

    return t
}

func getTemplate(p string) (string, error) {

    f, err := os.ReadFile(p)

    if err != nil {
        return "", err
    }

    return string(f), nil
}
