package metadata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (m *MetadataService) listHandler(res http.ResponseWriter, req *http.Request) {
	audioFiles, err := m.Storage.List()
	if err != nil {
		fmt.Println("error saving metadata: ", err)
		res.WriteHeader(500)
		return
	}
	jsonData, err := json.Marshal(audioFiles)
	if err != nil {
		fmt.Println("error saving metadata: ", err)
		res.WriteHeader(500)
	}
	io.WriteString(res, string(jsonData))
}
