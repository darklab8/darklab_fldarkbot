package utils

import (
	"darkbot/app/settings/logus"
	"fmt"
	"os"
	"os/exec"
)

func ShellRunArgs(program string, args ...string) {
	logus.Debug(fmt.Sprintf("OK attempting to run: %s", program), logus.Args(args))
	executable, _ := exec.LookPath(program)

	args = append([]string{""}, args...)
	command := exec.Cmd{
		Path:   executable,
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	err := command.Run()

	logus.CheckFatal(err, "failed to run shell command")
}