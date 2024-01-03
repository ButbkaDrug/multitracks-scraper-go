/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/butbkadrug/multitracks-scraper-go/internal"
	"github.com/spf13/cobra"
)

var params *internal.SaveProjectParams
var url string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Provide url to a song on multitracks.com",
	Long: `Plese, provide all nessesery data for a scraper to run. First URL,
then path to a TEMPLATE file, and path to DESTenation folder. You can utilize
flags or pass arguments to the command.`,
	Run: func(cmd *cobra.Command, args []string) {

        if url == "" && len(args) < 1 {
            fmt.Fprintln(os.Stderr, "ERROR: Please, provide url using -u flag or as a first argument")
            return
        }

        fmt.Println(args)

        if url == "" {
            url = args[0]
            args = args[1:]
        }

        fmt.Println(args)

        if params.Template == "" && len(args) < 1{
            fmt.Fprintln(os.Stderr, "ERROR: Please, provide path to the template file. You can use -t flag or pass it as an argument")
            return
        }

        if params.Template == "" {
            params.Template = args[0]
            args = args[1:]
        }
        fmt.Println(args)

        if params.Dest == "" && len(args) < 1 {
            fmt.Fprintln(os.Stderr, "ERROR: Please, provide destination path. You can use -d flag or pass it as an argument")
            return
        }

        if params.Dest == "" {
            params.Dest = args[0]
        }

        song, err := internal.NewSong(url)

        if err != nil {
            fmt.Fprintln(os.Stderr, "ERROR: Can't get song... ", err)
            return
        }

        params.Song = song

        if err = internal.SaveProject(params); err != nil {
            fmt.Fprintln(os.Stderr, "ERROR: Can't save project... ", err)
            return
        }

        fmt.Println("Project succesfully created!")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

    params = &internal.SaveProjectParams{}

    getCmd.Flags().StringVarP(
        &params.Dest,
        "dest",
        "d",
        "",
        "Destination folder which will be used as a projects bay",
    )

    getCmd.Flags().StringVarP(
        &params.Template,
        "template",
        "t",
        "",
        "Path to a Reaper(DAW) project that will be used as a template",
    )

    getCmd.Flags().StringVarP(
        &url,
        "url",
        "u",
        "",
        "url to the page with the song on multitracks.com",
    )
}
