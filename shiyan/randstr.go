// 产生随机字符串的方法

package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// 随机字符串
func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := r.Intn(4)
	switch x {
	case 0:
		fmt.Println(toSha1(string(Krand(32, KC_RAND_KIND_NUM))))
	case 1:
		fmt.Println(toSha1(string(Krand(32, KC_RAND_KIND_LOWER))))
	case 2:
		fmt.Println(toSha1(string(Krand(32, KC_RAND_KIND_UPPER))))
	case 3:
		fmt.Println(toSha1(string(Krand(32, KC_RAND_KIND_ALL))))
	default:
		fmt.Println("default")

		fmt.Println(toSha1(string(Krand(32, KC_RAND_KIND_ALL))))
	}
}

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

func toSha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
