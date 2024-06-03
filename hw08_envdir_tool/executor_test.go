package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Set env", func(t *testing.T) {
		cmd := []string{"/bin/echo", "-e", "test1"}
		env := Environment{
			"ARG1": {Value: "arg1-Done"},
			"ARG2": {Value: "arg2-Done"},
		}
		returnCode := RunCmd(cmd, env)
		require.Equal(t, returnCode, 0)
		require.Contains(t, os.Environ(), "ARG1=arg1-Done")
		require.Contains(t, os.Environ(), "ARG2=arg2-Done")
	})
	t.Run("ls fail", func(t *testing.T) {
		cmd := []string{"/bin/ls", "-j"}
		env := Environment{
			"ENV_EX_1": EnvValue{Value: "NOTSET"},
		}

		returnCode := RunCmd(cmd, env)
		require.Equal(t, 2, returnCode)
	})
	t.Run("pwd fail", func(t *testing.T) {
		cmd := []string{"/bin/pwd", "-R"}
		env := Environment{
			"ENV_EX_1": EnvValue{Value: "NOTSET"},
		}
		returnCode := RunCmd(cmd, env)
		require.Equal(t, 1, returnCode)
	})
}
