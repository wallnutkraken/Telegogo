package TeleGogo

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"reflect"
)

func toMultiPart(args interface{}) (*multipart.Writer, *bytes.Buffer, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	t := reflect.TypeOf(args)
	v := reflect.ValueOf(args)

	fieldC := t.NumField()
	if fieldC == 0 {
		return nil, nil, errors.New("No fields present in interface")
	}

	for x := 0; x < fieldC; x++ {
		value := v.Field(x)
		fieldType := t.Field(x)

		teleTag := fieldType.Tag.Get("tele")

		if teleTag != "-" {
			if teleTag[0] == '!' {
				/* Field is a filepath */
				path := valueString(value)
				uploader, err := writer.CreateFormFile(teleTag[1:], path)
				if err != nil {
					return nil, nil, err
				}
				actualFile, err := os.Open(path)
				if err != nil {
					return nil, nil, err
				}
				if _, err = io.Copy(uploader, actualFile); err != nil {
					return nil, nil, err
				}
				actualFile.Close()
			} else if !isZeroOfUnderlyingType(value.Interface()) {
				/* Field is not empty; let's go. */
				writer.WriteField(teleTag, valueString(value))
			}
		}
	}
	if err := writer.Close(); err != nil {
		return nil, nil, err
	}
	return writer, &buffer, nil
}

func valueString(val reflect.Value) string {
	return fmt.Sprint(val.Interface())
}

func isZeroOfUnderlyingType(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
