package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var testDataSet = `Package os provides a platform-independent interface 
to operating system functionality. The design is Unix-like, 
although the error handling is Go-like; failing calls 
return  values of type error rathere than error numbers. 
Often, more  information is available within the error. 
For example, if  a call that takes a file name fails, such as Open or Stat, 
the error will include the failing file name when printed and  will be of 
type *PathError, which may be unpacked for more information`

func TestCopy(t *testing.T) {
	fileName := "testFile"
	copyFileName := "testFileCopy"
	caseOffsetLimits := []struct {
		ofs int64
		lmt int64
		exp []byte
	}{
		{
			ofs: 201, lmt: 0,
			exp: []byte(`there than error numbers. 
Often, more  information is available within the error. 
For example, if  a call that takes a file name fails, such as Open or Stat, 
the error will include the failing file name when printed and  will be of 
type *PathError, which may be unpacked for more information`),
		},
		{ofs: 100, lmt: 12, exp: []byte("is Unix-like")},
		{ofs: 304, lmt: 15, exp: []byte("call that takes")},
		{
			ofs: 404, lmt: 0,
			exp: []byte(`me when printed and  will be of 
type *PathError, which may be unpacked for more information`),
		},
	}
	err := os.WriteFile(fileName, []byte(testDataSet), 0644)
	require.NoError(t, err)
	for _, tc := range caseOffsetLimits {
		tc := tc
		t.Run("copy with offset and limit", func(t *testing.T) {
			err = Copy(fileName, copyFileName, tc.ofs, tc.lmt)
			require.NoError(t, err)
			data, err := os.ReadFile(copyFileName)
			require.NoError(t, err)
			require.Equal(t, tc.exp, data)
		})
	}
}
