package fix_app

import (
	"bytes"
	"os"
	"testing"

	"github.com/rrghoigoiwngoiw/hw-2/hw02_fix_app/printer"
	"github.com/rrghoigoiwngoiw/hw-2/hw02_fix_app/types"
)

func Test_PrintStaff(t *testing.T) {
	staff := []types.Employee{
		{UserID: 1, Name: "John Doe", Age: 30, DepartmentID: 101},
		{UserID: 2, Name: "Jane Doe", Age: 25, DepartmentID: 102},
	}

	// Перенаправляем стандартный вывод в буфер
	var buf bytes.Buffer
	stdout := os.Stdout
	defer func() { os.Stdout = stdout }() // Восстанавливаем стандартный вывод после теста

	printer.PrintStaff(staff)

	// Ожидаемое значение
	expected := `User ID: 1, Age: 30; Name: John Doe, Department ID: 101
User ID: 2, Age: 25; Name: Jane Doe, Department ID: 102
`

	if buf.String() != expected {
		t.Errorf("Unexpected output. Got:\n%s\nExpected:\n%s", buf.String(), expected)
	}
}
