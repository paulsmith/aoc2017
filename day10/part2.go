package main

import "fmt"

func main() {
	input := "230,1,2,221,97,252,168,169,57,99,0,254,181,255,235,167"
	var lengths []byte
	for _, c := range input {
		lengths = append(lengths, byte(c))
	}
	for _, n := range []byte{17, 31, 73, 47, 23} {
		lengths = append(lengths, n)
	}
	const size = 256
	list := make([]byte, size)
	for i := 0; i < size; i++ {
		list[i] = byte(i)
	}
	const rounds = 64
	pos := 0
	skip := 0
	for i := 0; i < rounds; i++ {
		for _, length := range lengths {
			sublist := sel(list, pos, length)
			rev := reverse(sublist)
			update(list, rev, pos)
			pos += int(length) + skip
			skip++
		}
	}
	const hashsize = 16
	var dense [hashsize]byte
	for i := 0; i < (size / hashsize); i++ {
		dense[i] = list[(i*hashsize)+0] ^
			list[(i*hashsize)+1] ^
			list[(i*hashsize)+2] ^
			list[(i*hashsize)+3] ^
			list[(i*hashsize)+4] ^
			list[(i*hashsize)+5] ^
			list[(i*hashsize)+6] ^
			list[(i*hashsize)+7] ^
			list[(i*hashsize)+8] ^
			list[(i*hashsize)+9] ^
			list[(i*hashsize)+10] ^
			list[(i*hashsize)+11] ^
			list[(i*hashsize)+12] ^
			list[(i*hashsize)+13] ^
			list[(i*hashsize)+14] ^
			list[(i*hashsize)+15]
	}
	var hash string
	for _, b := range dense {
		hash += fmt.Sprintf("%02x", b)
	}
	fmt.Println(hash)
}

func update(s []byte, t []byte, pos int) {
	start := pos
	for i := 0; i < len(t); i++ {
		pos = (start + i) % len(s)
		s[pos] = t[i]
	}
}

func sel(s []byte, pos int, length byte) (res []byte) {
	start := pos
	for i := 0; i < int(length); i++ {
		pos = (start + i) % len(s)
		res = append(res, s[pos])
	}
	return
}

func reverse(s []byte) (res []byte) {
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return
}
