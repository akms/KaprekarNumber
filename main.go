package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Scan() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func Square(int_number int) (square_num int) {
	square_num = int_number * int_number
	return
}

func MakeStrSlice(square_str string) []string {
	var square_slice []string = make([]string, 0)
	for _, c := range square_str {
		square_slice = append(square_slice, string([]rune{c}))
	}
	return square_slice
}

func MakeCheckNumOne(square_slice []string, digit int) (check_num int) {
	str_num_head := square_slice[0:(digit / 2)]
	str_num_tail := square_slice[(digit / 2):digit]
	int_num_head, _ := strconv.Atoi(strings.Join(str_num_head, ""))
	int_num_tail, _ := strconv.Atoi(strings.Join(str_num_tail, ""))
	check_num = int_num_head + int_num_tail
	return
}

func CheckKaprekarNumberOne(int_number int, check_num int) {
	if int_number == check_num {
		fmt.Printf("%d is Kaprekar Number. (Definition 1)\n", int_number)
	} else {
		fmt.Printf("%d is not Kaprekar Number. (Definition 1)\n", int_number)
	}
}

func KaprekarNumberOne(int_number int) {
	var square_str string
	var digit, square_num, check_num int

	square_num = Square(int_number)
	square_str = strconv.Itoa(square_num)
	digit = utf8.RuneCountInString(square_str)
	square_slice := MakeStrSlice(square_str)
	check_num = MakeCheckNumOne(square_slice, digit)
	CheckKaprekarNumberOne(int_number, check_num)
}

func AscendingSortStrSlice(ascending_str_slice []string) []string {
	sort.Sort(sort.StringSlice(ascending_str_slice))
	return ascending_str_slice
}

func DescendingSortStrSlice(descending_str_slice []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(descending_str_slice)))
	return descending_str_slice
}

func MakeCheckNumTw(str_slice []string) (check_num int) {
	check_num, _ = strconv.Atoi(strings.Join(str_slice, ""))
	return
}

func CheckKaprekarNumberTw(int_number, descending_num, ascending_num int) {
	var ans_num int
	ans_num = descending_num - ascending_num
	if ans_num == int_number {
		fmt.Printf("%d is Kaprekar Number. (Definition 2)\n", int_number)
	} else {
		fmt.Printf("%d is not Kaprekar Number. (Definition 2)\n", int_number)
	}
}

func KaprekarNumberTw(str_number string, int_number int) {
	str_slice := MakeStrSlice(str_number)
	ascending_order := AscendingSortStrSlice(str_slice)
	ascending_num := MakeCheckNumTw(ascending_order)
	descending_order := DescendingSortStrSlice(str_slice)
	descending_num := MakeCheckNumTw(descending_order)

	CheckKaprekarNumberTw(int_number, descending_num, ascending_num)
}

func main() {
	var str_number string
	var int_number int
	fmt.Print("Please enter number \n> ")

	str_number = Scan()
	int_number, _ = strconv.Atoi(str_number)

	KaprekarNumberOne(int_number)
	KaprekarNumberTw(str_number, int_number)

	return
}
