package main

import (
	"fmt"
	"strings"
)

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0]) // must be "apple"

	// perbedaan penulisan array dan slice
	var fruitsA = []string{"apple", "grape"}   // slice
	var fruitsB = [2]string{"banana", "melon"} // array
	var fruitsC = [...]string{fruitsA[0], fruitsB[1]}
	fmt.Println(fruitsA)
	fmt.Println(fruitsB)
	fmt.Println(fruitsC)

	fmt.Println(strings.Repeat("=>", 10))

	// []string{"apple", "grape", "banana", "melon"}
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(aFruits)  // [apple, grape, banana]
	fmt.Println(bFruits)  // [grape, banana, melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	fmt.Println(strings.Repeat("=>", 10))
	baFruits[0] = "pineaple"
	fmt.Println(aFruits)  // [apple, pineaple, banana]
	fmt.Println(bFruits)  // [pineaple, banana, melon]
	fmt.Println(aaFruits) // [pineaple]
	fmt.Println(baFruits) // [pineaple]

	fmt.Println(strings.Repeat("==>", 5), "Fungsi Len", strings.Repeat("<==", 5))

	var fruity = []string{"anggur", "jambu", "nangka", "salak", "timun"}
	fmt.Println("Panjang dari slice fruity adalah =>", len(fruity))

	fmt.Println(strings.Repeat("==>", 5), "Fungsi Cap", strings.Repeat("<==", 5))

	var aFruity = fruity[0:3]
	fmt.Println("aFruity len =>", len(aFruity)) // 3
	fmt.Println("aFruity cap =>", cap(aFruity)) // 5

	var bFruity = fruity[2:4]
	fmt.Println("bFruity len =>", len(bFruity)) // 2
	fmt.Println("bFruity cap =>", cap(bFruity)) // 3

	fmt.Println(strings.Repeat("==>", 5), "Fungsi Append", strings.Repeat("<==", 5))

	var Append = []int{1, 2, 3}
	var Aappend = append(Append, 4)
	fmt.Println("before append =>", Append)
	fmt.Println("After append =>", Aappend)
	Aappend[0] = 2
	fmt.Println(Append)
	fmt.Println(Aappend)

	var newSlice = make([]string, 3, 4)
	newSlice[0] = "a"
	newSlice[1] = "b"
	newSlice[2] = "c"
	fmt.Println("before append newSlice =", newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	var newSlice2 = newSlice[0:3]
	fmt.Println("newSlice2 =", newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	var newSlice3 = append(newSlice2, "d")
	fmt.Println("newSlice2 = ", newSlice2) // [a,b,c,d] ??
	fmt.Println("newSlice3 =", newSlice3)  // [a,b,c,d]
	newSlice3[0] = "sd"
	fmt.Println("newSlice2 = ", newSlice2) // [a,b,c,d] ??
	fmt.Println("newSlice3 =", newSlice3)  // [a,b,c,d]

	fmt.Println(strings.Repeat("==>", 5), "Fungsi Copy", strings.Repeat("<==", 5))

	var dst = make([]string, 3)
	var src = []string{"watermelon", "pinnaple", "apple", "orange"}
	var n = copy(dst, src)
	fmt.Println(len(src))
	fmt.Println(n)

}
