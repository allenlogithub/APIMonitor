package export

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"

	"modules/requests"
)

func ToCSV(data []requests.ResponseData, path string) error {
	// create a file
	f, createErr := os.Create(path + ".csv")
	if createErr != nil {
		fmt.Println(createErr)
		return errors.New("Create CSV file failed.")
	}
	defer f.Close()

	// get headers and write line
	headers, gHErr := getHeaders(data[0])
	if gHErr != nil {
		return errors.New("Get Headers failed.")
	}
	str := headers[0]
	for i := 1; i < len(headers); i++ {
		str = str + "," + headers[i]
	}
	str = str + "\n"
	writeString(f, str)

	// get data and write line
	for i := range data {
		v := reflect.ValueOf(data[i])

		str := fmt.Sprintf("%#v", v.Field(0))
		for j := 1; j < v.NumField(); j++ {
			str = str + "," + fmt.Sprintf("%#v", v.Field(j))
		}
		str = str + "\n"
		writeString(f, str)
	}

	return nil
}

func writeString(f *os.File, s string) error {
	buf := bytes.NewBufferString(s)
	_, nBSErr := buf.WriteTo(f)
	if nBSErr != nil {
		return errors.New("Error in w.WriteString in ToCSV")
	}

	return nil
}

func getHeaders(data requests.ResponseData) ([]string, error) {
	val := reflect.ValueOf(data)
	var headers []string
	for i := 0; i < val.NumField(); i++ {
		headers = append(headers, val.Type().Field(i).Name)
	}
	return headers, nil
}
