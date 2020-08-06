package strutil

import "fmt"


func numTilePossibilities(tiles string) int {
	fmt.Printf("NumPerms for %s ", tiles)
    m := make(map[byte]int)

	for _,c := range tiles {
		cb := byte(c)
		if count,ok := m[cb] ; ok {
			count++
			m[cb] = count
		} else {
			m[cb] = 1
		}
	}

	count := numPerm(m)
	fmt.Printf("%v => %d\n",m, count)
    return count
}


func numPerm(m map[byte]int) int {
	count := 0
    for k,v := range m {
		if v > 0 {
			m[k]=v-1
			count = count + numPerm(m) + 1
			m[k]=v
		}
    }

	return count
}
