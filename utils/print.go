package utils

import (
	"encoding/json"
	"fmt"

	"github.com/marianina8/audiofile/models"
)

func Print(b []byte, jsonFormat bool) ([]byte, error) {
	var err error
	if jsonFormat {
		if IsAtty() {
			err = Pager(string(b))
			if err != nil {
				return b, fmt.Errorf("\n  paging: %v\n  ", err)
			}
		} else {
			fmt.Println(string(b))
		}
	} else {
		var audios models.AudioList
		err := json.Unmarshal(b, &audios)
		if err != nil {
			return b, fmt.Errorf("\n  unmarshalling: %v\n  ", err)
		}
		tableData, err := audios.Table()
		if err != nil {
			return b, fmt.Errorf("\n  printing table: %v\n  ", err)
		}
		if IsAtty() {
			err = Pager(tableData)
			if err != nil {
				return b, fmt.Errorf("\n  paging: %v\n  ", err)
			}
		} else {
			fmt.Println(tableData)
		}
	}
	return b, nil
}
