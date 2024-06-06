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

func validateField(errSlice *ValidationErrors, tags []string, fieldName string, value reflect.Value) {
	tagData := strings.Split(tags[0], "|")
	for idx, tag := range tagData {
		fmt.Println(tag)
		switch strings.Split(tag, ":")[idx] {
		case "regexp":
			validateRegexp(errSlice, fieldName, tag, value)
		case "len":
			validateLengh(errSlice, fieldName, tag, value)
		case "in":
			validateIn(errSlice, fieldName, tag, value)
		case "max":
			validateMax(errSlice, fieldName, tag, value)
		case "min":
			validateMin(errSlice, fieldName, tag, value)
		}
	}
}

func validateMin(errSlice *ValidationErrors, fieldName, tag string, value reflect.Value) {
	min, err := strconv.Atoi(strings.Split(tag, ":")[1])
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

func validateMax(errSlice *ValidationErrors, fieldName, tag string, value reflect.Value) {
	max, err := strconv.Atoi(strings.Split(tag, ":")[1])
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
}

func validateRegexp(errSlice *ValidationErrors, fieldName, tag string, value reflect.Value) {
	re := regexp.MustCompile(strings.Split(tag, ":")[1])
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
}

func validateIn(errSlice *ValidationErrors, fieldName, tag string, value reflect.Value) {
	var inData string
	pattern := strings.ReplaceAll(strings.Split(tag, ":")[1], ",", "|")
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
}

func validateLengh(errSlice *ValidationErrors, fieldName, tag string, value reflect.Value) {
	lengh, err := strconv.Atoi(strings.Split(tag, ":")[1])
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
}
