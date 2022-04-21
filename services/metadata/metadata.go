package metadata

import (
	"audiofile/internal/interfaces"
	"audiofile/storage"

	"fmt"
	"net/http"
)

type MetadataService struct {
	Server  *http.Server
	Storage interfaces.Storage
}

func CreateMetadataServer(port int, storage interfaces.Storage) *http.Server {
	mux := http.NewServeMux()
	metadataService := &MetadataService{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%v", port),
			Handler: mux,
		},
		Storage: storage,
	}
	mux.HandleFunc("/upload", metadataService.uploadHandler)
	mux.HandleFunc("/request", metadataService.getByIDHandler)
	// mux.HandleFunc("/list", metadataService.listHandler)
	return metadataService.Server
}

func Run(port int) {
	flatfileStorage := storage.FlatFile{}
	server := CreateMetadataServer(port, flatfileStorage)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("error starting api: ", err)
	}
}
