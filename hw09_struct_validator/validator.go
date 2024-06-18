package hw09structvalidator

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type User1 struct {
	ID     string   `json:"id" validate:"len:36"`
	Age    int      `validate:"min:18|max:50"`
	Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
	Phones []string `validate:"len:11"`
	Some   []int    `validate:"in:15"`
}

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

var hashType = map[reflect.Kind]bool{
	reflect.Int:    true,
	reflect.Int8:   true,
	reflect.Int16:  true,
	reflect.Int32:  true,
	reflect.Int64:  true,
	reflect.Uint:   true,
	reflect.Uint8:  true,
	reflect.Uint16: true,
	reflect.Uint32: true,
	reflect.Uint64: true,
}

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
			RunValidateField(
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

func RunValidateField(errSlice *ValidationErrors, tags []string, fieldName string, value reflect.Value) {
	tagData := strings.Split(tags[0], "|")
	valid := func(tag string, value reflect.Value) {
		validateField(
			errSlice,
			tag,
			fieldName,
			value,
		)
	}
	for _, tag := range tagData {
		if value.Kind() == reflect.Slice {
			if typeEl := value.Type().Elem(); hashType[typeEl.Kind()] {
				valid(tag, value)
			} else {
				for j := 0; j < value.Len(); j++ {
					valid(tag, value.Index(j))
				}
			}
		} else {
			valid(tag, value)
		}
	}
}

func validateField(errSlice *ValidationErrors, tag string, fieldName string, value reflect.Value) {
	switch strings.Split(tag, ":")[0] {
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
