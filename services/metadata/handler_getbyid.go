package metadata

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (m *MetadataService) getByIDHandler(res http.ResponseWriter, req *http.Request) {
	value, ok := req.URL.Query()["id"]
	if !ok || len(value[0]) < 1 {
		fmt.Println("Url Param 'id' is missing")
		res.WriteHeader(500)
		return
	}
	id := string(value[0])
	fmt.Println("requesting audio by id: ", id)

	audio, err := m.Storage.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "no such file or directory") {
			io.WriteString(res, "id not found")
			res.WriteHeader(200)
			return
		}
		res.WriteHeader(500)
		return
	}
	audioString, err := audio.JSON()
	if err != nil {
		res.WriteHeader(500)
		return
	}
	io.WriteString(res, audioString)
}
