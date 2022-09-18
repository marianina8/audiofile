package metadata

import (
	"audiofile/extractors/tags"
	"audiofile/extractors/transcript"
	"audiofile/models"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (m *MetadataService) uploadHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("req: ", req)
	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println("error creating formfile: ", err)
		res.WriteHeader(500)
		return
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("error opening file: ", err)
		res.WriteHeader(500)
		return
	}
	defer f.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		fmt.Println("error copying file to buffer: ", err)
		res.WriteHeader(500)
		return
	}

	id, audioFilePath, err := m.Storage.Upload(buf.Bytes(), handler.Filename)
	audio := &models.Audio{
		Id:   id,
		Path: audioFilePath,
	}
	err = m.Storage.SaveMetadata(audio)
	if err != nil {
		fmt.Println("error saving metadata: ", err)
		res.WriteHeader(500)
		return
	}
	audio.Status = "Initiating"

	io.WriteString(res, id)
	go func() {
		var errors []error

		audio.Status = "Complete"

		// tags
		err = tags.Extract(audio)
		if err != nil {
			fmt.Println("error extracting tags metadata: ", err)
			errors = append(errors, err)
		}
		err = m.Storage.SaveMetadata(audio)
		if err != nil {
			fmt.Println("error saving metadata: ", err)
			errors = append(errors, err)
		}

		// transcript
		err = transcript.Extract(audio)
		if err != nil {
			fmt.Println("error extracting transcript metadata: ", err)
			errors = append(errors, err)
		}

		audio.Error = errors
		audio.Status = "Complete"
		err = m.Storage.SaveMetadata(audio)
		if err != nil {
			fmt.Println("error saving metadata: ", err)
			errors = append(errors, err)
		}

		if len(errors) > 0 {
			fmt.Println("errors occurred extracting metadata: ")
			for i := 0; i < len(errors); i++ {
				fmt.Printf("\terror[%d]: %s\n", i, errors[i].Error())
			}
		} else {
			fmt.Println("successfully extracted and saved audio metadata: ", audio)
		}
	}()
}
