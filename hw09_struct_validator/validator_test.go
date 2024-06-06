package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte `validate:"len:5"`
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	testsSuccess := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "1234567890abcdef1234567890abcdef",
				Name:   "John Doe",
				Age:    30,
				Email:  "johndoe@example.com",
				Role:   "admin",
				Phones: []string{"1234567890"},
			},
			expectedErr: nil,
		},
		{
			in: App{
				Version: "1.2.3",
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 200,
				Body: "Success message",
			},
			expectedErr: nil,
		},

		{
			in: Token{
				Header:    []byte("header data"),
				Payload:   []byte("12345"),
				Signature: []byte("signature data"),
			},
			expectedErr: nil,
		},
	}

	testsFail := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "12345678",
				Name:   "John Doe",
				Age:    30,
				Email:  "john.doe@example.com",
				Role:   "admin",
				Phones: []string{"1234567890"},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Email",
					Err:   fmt.Errorf("validation error: field 'Email'"),
				},
			},
		},
		{
			in: Token{
				Header:    []byte("header data"),
				Payload:   []byte("123456"),
				Signature: []byte("signature data"),
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Payload",
					Err:   fmt.Errorf("validation error: field 'Payload'"),
				},
			},
		},
		{
			in: User{
				ID:     "1234567890abcdef1234567890abcdef",
				Name:   "John Doe",
				Age:    15,
				Email:  "john.doe@example.com",
				Role:   "admin",
				Phones: []string{"1234567890"},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Age",
					Err:   fmt.Errorf("validation error: field 'Age'"),
				},
				ValidationError{
					Field: "Email",
					Err:   fmt.Errorf("validation error: field 'Email'"),
				},
			},
		},
	}

	for i, tt := range testsSuccess {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			require.Equal(t, err, tt.expectedErr)
		})
	}

	for i, tt := range testsFail {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			if !errors.As(err, &ValidationErrors{}) {
				t.Errorf("unexpected error: got %v, expected %v", err, tt.expectedErr)
			} else {
				require.Equal(t, tt.expectedErr, err)
				// fmt.Println(err)
			}
		})
	}
}
