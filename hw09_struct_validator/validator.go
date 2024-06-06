package main

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	msgErr := make([]string, 0, len(v))
	for _, msg := range v {
		msgErr = append(msgErr, msg.Err.Error())
	}
	return strings.Join(msgErr, "; ")
}

func Validate(v interface{}) error {
	var (
		dataType  = reflect.TypeOf(v)
		valueType = reflect.ValueOf(v)
		errSlice  ValidationErrors
	)
	if dataType.Kind() != reflect.Struct {
		log.Fatalf("invalid type: expected struct, got %v", dataType.Kind())
	}
	for i := 0; i < dataType.NumField(); i++ {
		if alias, ok := dataType.Field(i).Tag.Lookup("validate"); ok {
			if len(alias) == 0 {
				continue
			}
			validateField(
				&errSlice,
				strings.Fields(alias),
				dataType.Field(i).Name,
				valueType.FieldByName(dataType.Field(i).Name),
			)
		}
	}
	if len(errSlice) != 0 {
		return errSlice
	}
	return nil
}

func validateField(errSlice *ValidationErrors, tags []string, fieldName string, value reflect.Value) *ValidationErrors {
	tagData := strings.Split(tags[0], "|")
	for idx, vle := range tagData {
		switch strings.Split(vle, ":")[idx] {
		case "regexp":
			re := regexp.MustCompile(strings.Split(vle, ":")[1])
			if len(re.Find([]byte(value.String()))) != 0 {
				log.Printf("validation field %v success", fieldName)
			} else {
				*errSlice = append(
					*errSlice,
					ValidationError{
						Field: fieldName,
						Err:   fmt.Errorf("validation error: field '%s'", fieldName),
					},
				)
			}
		case "len":
			lengh, err := strconv.Atoi(strings.Split(vle, ":")[1])
			if err != nil {
				log.Fatalln(err)
			}
			if value.Len() <= lengh {
				log.Printf("validation field %v success", fieldName)
			} else {
				*errSlice = append(
					*errSlice,
					ValidationError{
						Field: fieldName,
						Err:   fmt.Errorf("validation error: field '%s'", fieldName),
					},
				)
			}
		case "in":
			var inData string
			pattern := strings.ReplaceAll(strings.Split(vle, ":")[1], ",", "|")
			re := regexp.MustCompile(pattern)
			if value.Kind() == reflect.String {
				inData = value.String()
			}
			if value.Kind() == reflect.Int {
				inData = strconv.Itoa(int(value.Int()))
			}

			if find := re.FindStringSubmatch(inData); len(find) != 0 {
				log.Printf("validation field %v success", fieldName)
			} else {
				*errSlice = append(
					*errSlice,
					ValidationError{
						Field: fieldName,
						Err:   fmt.Errorf("validation error: field '%s'", fieldName),
					},
				)
			}
		case "max":
			max, err := strconv.Atoi(strings.Split(vle, ":")[1])
			if err != nil {
				log.Fatalf("can't get max value, return err: %v", err)
			}
			if max >= int(value.Int()) {
				log.Printf("validation field %v success", fieldName)
			} else {
				*errSlice = append(
					*errSlice,
					ValidationError{
						Field: fieldName,
						Err:   fmt.Errorf("validation error: field '%s'", fieldName),
					},
				)
			}
		case "min":
			min, err := strconv.Atoi(strings.Split(vle, ":")[1])
			if err != nil {
				log.Fatalf("can't get min value, return err: %v", err)
			}
			if min <= int(value.Int()) {
				log.Printf("validation field %v success", fieldName)
			} else {
				*errSlice = append(
					*errSlice,
					ValidationError{
						Field: fieldName,
						Err:   fmt.Errorf("validation error: field '%s'", fieldName),
					},
				)
			}
		}
	}
	return errSlice
}
