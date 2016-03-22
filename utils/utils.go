package utils

import (
    "strconv"
    "strings"
    "fmt"
)

func Str2int(char rune) int64{
    var res  int32
    if char >= 0 && char <= 9 {
        res = char - 0
    } else if char >= 'A' && char <= 'Z' {
        res = char - 'A' + 10
    } else {
        res = char - 'a' + 36
    }
    return int64(res)
}

func Int2str(i int) string {
    var res string
    if i >=0 && i <= 9 {
        res =  strconv.Itoa(int(i))
    } else if i >= 10 && i <= 35 {
        res = string(i - 10 + 'A')
    } else {
        res = string(i - 36 + 'a')
    }
    return res
}

func Convert_10_to_62(num int) string {
    var res []string
    var ys int
    s := num
    for s >= 1 {
        tmp := s/62
        ys = s - tmp * 62
        s = tmp
        res = append(res, Int2str(ys))
        fmt.Println("s =", s, ", ys =", ys)
    }
    return strings.Join(res, "")
}

func Convert_62_to_10(str string) int64 {
    var res int64 = 0
    for _, s := range str {
        fmt.Println("s =", s)
        res = res + Str2int(s) * 62
    }
    return res
}
