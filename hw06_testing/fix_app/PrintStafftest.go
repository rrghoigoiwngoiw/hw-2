package fixapp

import (
	"bytes"
	"os"
	"testing"

	"github.com/rrghoigoiwngoiw/hw-2/hw02_fix_app/printer"
	"github.com/rrghoigoiwngoiw/hw-2/hw02_fix_app/types"
)

func TestPrintStaff(_ *testing.T) {
	staff := []types.Employee{
		{UserID: 1, Name: "John Doe", Age: 30, DepartmentID: 101},
		{UserID: 2, Name: "Jane Doe", Age: 25, DepartmentID: 102},
	}

	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = stdout }()

	var buf bytes.Buffer
	done := make(chan struct{})

	go func() {
		_, _ = buf.ReadFrom(r)
		close(done)
	}()

	printer.PrintStaff(staff)

	_ = w.Close()
	<-done
}
