package utils

import (
    "strconv"
    "math"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
)

/**
 *  62进制
 *   '0'-'9' --- 0-9
 *   A-Z     --- 10-35
 *   a-z     --- 36-62
 */

/**
 * 单个字符转数字，传入ascii
 */
func Str2int(i rune) int64{
    var res  int32
     if i >= 48 && i <= 57 {
        res = i - 48
    } else if i >= 65 && i <= 90 {
        res = i - 65 + 10
    } else {
        res = i - 97 + 36
    }
    return int64(res)
}

/**
 * 数字转字符
 */
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

/**
 * 10进制转62进制
 */
func Convert_10_to_62(num int) string {
    var res string
    var ys int
    s := num
    for s >= 1 {
        tmp := s/62
        ys = s - tmp * 62
        s = tmp
        res = Int2str(ys) + res
    }
    return res
}

/**
 * 62进制转10进制
 */
func Convert_62_to_10(str string) int64 {
    var res int64 = 0
    len := len(str)
    for k, v := range str {
        res = res + Str2int(v) * int64(math.Pow(float64(62), float64(len - 1 - k)))
    }
    return res
}

func GetSortUrl(db *sql.DB, url string) string {
    stmt, _ := db.Prepare("INSERT url SET url = ?")
    res, _ := stmt.Exec(url)
    id, _ := res.LastInsertId()
    return Convert_10_to_62(int(id))
}

func GetOriUrl(db *sql.DB, url string) string {
    id := Convert_62_to_10(url)
    rows, _ := db.Query("SELECT url FROM url WHERE id = ?", id)
    var oriurl string
    for rows.Next() {
        err := rows.Scan(&oriurl)
        if err != nil {
            panic(err)
        }
    }
    defer rows.Close()
    return oriurl
}
