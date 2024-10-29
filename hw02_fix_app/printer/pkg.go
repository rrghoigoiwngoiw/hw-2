package printer

import (
	"fmt"

	"github.com/rrghoigoiwngoiw/hw-2/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	var str string
	for i := 0; i < len(staff); i++ {
		str = fmt.Sprintf("User ID: %d; Age: %d;", staff[i].UserID, staff[i].Age)
		str2 := fmt.Sprintf(`Name: %s; Department ID: %d; `, staff[i].Name, staff[i].DepartmentID)
		fmt.Println(str, str2)
	}

	fmt.Println(str)
}
