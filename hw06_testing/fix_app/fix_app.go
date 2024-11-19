package fix_app

import (
	"fmt"

	"github.com/rrghoigoiwngoiw/hw-2/hw02_fix_app/printer"
	"github.com/rrghoigoiwngoiw/hw-2/hw02_fix_app/reader"
	"github.com/rrghoigoiwngoiw/hw-2/hw02_fix_app/types"
)

func fix_app() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}
	staff, err = reader.ReadJSON(path)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
