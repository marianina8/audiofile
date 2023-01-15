package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/pterm/pterm"
)

var header = []string{
	"ID",
	"Path",
	"Status",
	"Title",
	"Album",
	"Album Artist",
	"Composer",
	"Genre",
	"Artist",
	"Lyrics",
	"Year",
	"Comment",
}

var IdColor = color.New(color.FgGreen).SprintFunc()

func row(audio Audio) []string {
	return []string{
		IdColor(audio.Id),
		audio.Path,
		audio.Status,
		audio.Metadata.Tags.Title,
		audio.Metadata.Tags.Album,
		audio.Metadata.Tags.AlbumArtist,
		audio.Metadata.Tags.Composer,
		audio.Metadata.Tags.Genre,
		audio.Metadata.Tags.Artist,
		audio.Metadata.Tags.Lyrics,
		strconv.Itoa(audio.Metadata.Tags.Year),
		audio.Metadata.Tags.Comment,
	}
}

// AudioList is a slice of Audio structs
type AudioList []Audio

func (list *AudioList) Table() (string, error) {
	data := pterm.TableData{header}
	for _, audio := range *list {
		data = append(
			data,
			row(audio),
		)
	}
	return pterm.DefaultTable.WithHasHeader().WithData(data).Srender()
}

func (audio *Audio) Plain() string {
	return fmt.Sprintf("Id,Path,Tags,Transcript\n%s,%s,%v,%s\n", audio.Id, audio.Path, audio.Metadata.Tags, audio.Metadata.Transcript)
}

func (list *AudioList) Plain() string {
	plaintext := ""
	for _, audio := range *list {
		plaintext += audio.Plain()
	}
	return fmt.Sprintf("Id,Path,Tags,Transcript\n%s\n", plaintext)
}

func (audio *Audio) Table() (string, error) {
	data := pterm.TableData{header, row(*audio)}
	return pterm.DefaultTable.WithHasHeader().WithData(data).Srender()
}

type Audio struct {
	Id       string
	Path     string
	Metadata Metadata
	Status   string
	Error    []error
}

func (audio *Audio) JSON() (string, error) {
	audioJSON, err := json.Marshal(audio)
	if err != nil {
		return "", err
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(audioJSON), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
