package words

import (
	"fmt"
	"strings"
)

func WordCount(s string) (mm map[string]int) {

	fmt.Println("Thitiphong")
	// แสดง word ที่ซ้ำกัน แสดงจำนวนด้วย

	m := strings.Replace(s, ",", " ", -1)
	n := strings.Replace(m, ".", " ", -1)
	fmt.Println(m)
	fmt.Println(n)

	s1 := strings.Fields(s)
	// fmt.Println(len(s1))
	mm = make(map[string]int)
	for _, x := range s1 {
		mm[x]++
	}

	return mm
}
