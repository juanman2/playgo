package attendance

import "fmt"

func checkRecord(s string) bool {
	a := false
	l := 0
	for _, n := range s {
		fmt.Printf("%v..", n)
		switch n {
		case 'L':
			fmt.Printf("late..")
			if l == 2 {
				return false
			}
			l = l + 1
		case 'A':
			fmt.Printf("absent..")
			if a == true {
				return false
			}
			a = true
		case 'P':
			fmt.Printf("present..")
			l = 0
		}
	}
	return true
}
