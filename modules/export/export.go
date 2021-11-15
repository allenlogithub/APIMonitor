package export

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
)

type csvStruct struct {
	Data    []interface{}
	Headers []string
}

func (wrap csvStruct) ToCSV(path string) error {
	// create a file
	f, createErr := os.Create(path + ".csv")
	if createErr != nil {
		fmt.Println(createErr)
		return errors.New("Create CSV file failed.")
	}
	defer f.Close()

	str := ""
	// get headers
	headers := wrap.Headers
	str = str + headers[0]
	for i := 1; i < len(headers); i++ {
		str = str + "," + headers[i]
	}
	str = str + "\n"

	// get data
	for i := range wrap.Data {
		vals := getValues(wrap.Data[i])
		str = str + fmt.Sprintf("%#v", vals[0])
		for j := 1; j < len(vals); j++ {
			str = str + "," + fmt.Sprintf("%#v", vals[j])
		}
		str = str + "\n"
	}

	// write lines
	writeString(f, str)

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

func getHeaders(data interface{}) []string {
	val := reflect.ValueOf(data)
	var headers []string
	for i := 0; i < val.NumField(); i++ {
		headers = append(headers, val.Type().Field(i).Name)
	}

	return headers
}

func getValues(data interface{}) []interface{} {
	val := reflect.ValueOf(data)
	var values []interface{}
	for i := 0; i < val.NumField(); i++ {
		values = append(values, val.Field(i).Interface())
	}

	return values
}

func CSVWrapper(data []interface{}) csvStruct {
	return csvStruct{
		Data:    data,
		Headers: getHeaders(data[0]),
	}
}
