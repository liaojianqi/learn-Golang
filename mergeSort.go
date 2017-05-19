package main

import (
    "fmt"
)

func _sort(num []int, a, b int) {
    if b - a <= 1 {
        return
    }
    t := (a + b) /2
    _sort(num, a, t)
    _sort(num, t, b)
    newNum := []int{}
    i, j := a, t
    for i < t && j < b {
        if num[i] < num[j] {
            newNum = append(newNum, num[i])
            i++
        } else {
            newNum = append(newNum, num[j])
            j++
        }
    }
    for ; i < t; i++ {
        newNum = append(newNum, num[i])
    }
    for ; j < b; j++ {
        newNum = append(newNum, num[j])
    }
    for i := 0; i < b - a; i++ {
        num[a + i] = newNum[i]
    }
}
func sort(num []int) {
    _sort(num, 0, len(num))
}
func main() {
    num := []int{9,0,1,4,3,2,7,6,8,5}
    sort(num)
    fmt.Println(num)
}