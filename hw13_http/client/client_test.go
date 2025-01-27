package client

import (
	"errors"
	"flag"
	"os"
	"os/exec"
	"testing"
)

func TestRunClient_MissingURL(t *testing.T) {
	// Командная строка для запуска текущего теста
	cmd := exec.Command(os.Args[0], "-test.run=TestRunClient_MissingURL_Helper")
	cmd.Env = append(os.Environ(), "TEST_MISSING_URL=1")

	// ошибка
	err := cmd.Run()
	var exitError *exec.ExitError
	if errors.As(err, &exitError) {
		if exitError.ExitCode() != 1 {
			t.Errorf("ожидался код выхода 1, получен: %d", exitError.ExitCode())
		}
	} else if err == nil {
		t.Error("ожидался код выхода 1, программа завершилась без ошибок")
	}
}

func TestRunClient_MissingURL_Helper(_ *testing.T) {
	if os.Getenv("TEST_MISSING_URL") != "1" {
		return
	}

	// некорректные флаги
	os.Args = []string{os.Args[0], "-method=GET"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	RunClient()
	os.Exit(0)
}
