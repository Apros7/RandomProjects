package main

import (
	"fmt"
)

func array() {
	var intArr [3]int32 // 3 x 4 bytes = 12 bytes of memory
	intArr2 := [...]int32{1, 2, 3}
	intArr[1] = 123
	fmt.Println(intArr2)
	fmt.Println(&intArr[0]) // &: Gives memory localtion, which is continuous
	fmt.Println(intArr[1:3])

	var intSlice []int32 // 3 x 4 bytes = 12 bytes of memory
	var intSlice2 []int32 = []int32{1, 2, 3}
	fmt.Println(len(intSlice), cap(intSlice))
	intArrLong := append(intSlice, intSlice2...) // Will allocate more memory (not 4 x 4 bytes = 16 bytes, although only 3 items)
	fmt.Println(len(intArrLong), cap(intArrLong))
	var intArrLonger []int32 = make([]int32, 3, 8) // 3 items to begin with, but allocate memory in heap for 8 items = 32 bytes
	fmt.Println(intArrLonger)
}

func maps() {
	var strIntMap map[string]uint32 = make(map[string]uint32)
	fmt.Println(strIntMap)
	strIntMap["Lucas"] = 21
	strIntMap["Zilas"] = 17
	fmt.Println(strIntMap)
	delete(strIntMap, "Lucas")
	fmt.Println(strIntMap)
	var LucasValue, LucasOK = strIntMap["Lucas"]
	fmt.Println(LucasValue, LucasOK) // Will return default value of value type here uint32 default value = 0
}
