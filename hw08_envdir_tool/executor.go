package main

import (
	"errors"
	"os"
	"os/exec"
)

const (
	codeOk           = 0
	codeFail         = 1
	envCodeUnsetFail = 4
	envCodeSetFail   = 7
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	commandBin, args := cmd[0], cmd[1:]
	for k, v := range env {
		if v.NeedRemove {
			if err := os.Unsetenv(k); err != nil {
				return envCodeUnsetFail
			}
			continue
		}
		if err := os.Setenv(k, v.Value); err != nil {
			return envCodeSetFail
		}
	}
	objCmd := exec.Command(commandBin, args...)
	objCmd.Stdin = os.Stdin
	objCmd.Stderr = os.Stderr
	objCmd.Stdout = os.Stdout
	if err := objCmd.Run(); err != nil {
		var execErr *exec.ExitError
		if errors.As(err, &execErr) {
			return execErr.ExitCode()
		}
		return codeFail
	}
	return codeOk
}
