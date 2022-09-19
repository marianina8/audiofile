package metadata

import (
	"fmt"
	"io"
	"net/http"
)

func (m *MetadataService) listHandler(res http.ResponseWriter, req *http.Request) {
	audioFiles, err := m.Storage.List()
	if err != nil {
		res.WriteHeader(500)
		return
	}
	jsonData, err := json.Marshal(audioFiles)
	if err != nil {
			res.WriteHeader(500)
			return
		}
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, []byte(jsonData), "", "    ")
	if err != nil {
		res.WriteHeader(500)
		return
	}
	io.WriteString(res, prettyJSON.String())
}
