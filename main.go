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

func scan() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func square(int_number int) (square_num int) {
	square_num = int_number * int_number
	return
}

func make_str_slice(square_str string) []string {
	var square_slice []string = make([]string, 0)
	for _, c := range square_str {
		square_slice = append(square_slice, string([]rune{c}))
	}
	return square_slice
}

func make_check_num_one(square_slice []string, digit int) (check_num int) {
	str_num_head := square_slice[0:(digit / 2)]
	str_num_tail := square_slice[(digit / 2):digit]
	int_num_head, _ := strconv.Atoi(strings.Join(str_num_head, ""))
	int_num_tail, _ := strconv.Atoi(strings.Join(str_num_tail, ""))
	check_num = int_num_head + int_num_tail
	return
}

func check_kaprekar_number_one(int_number int, check_num int) {
	if int_number == check_num {
		fmt.Printf("%d is Kaprekar Number. (Definition 1)\n", int_number)
	} else {
		fmt.Printf("%d is not Kaprekar Number. (Definition 1)\n", int_number)
	}
}

func kaprekar_number_one(int_number int) {
	var square_str string
	var digit, square_num, check_num int

	square_num = square(int_number)
	square_str = strconv.Itoa(square_num)
	digit = utf8.RuneCountInString(square_str)
	square_slice := make_str_slice(square_str)
	check_num = make_check_num_one(square_slice, digit)
	check_kaprekar_number_one(int_number, check_num)
}

func ascending_sort_str_slice(ascending_str_slice []string) []string {
	sort.Sort(sort.StringSlice(ascending_str_slice))
	return ascending_str_slice
}

func descending_sort_str_slice(descending_str_slice []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(descending_str_slice)))
	return descending_str_slice
}

func make_check_num_tw(str_slice []string) (check_num int) {
	check_num, _ = strconv.Atoi(strings.Join(str_slice, ""))
	return
}

func check_kaprekar_number_tw(int_number, descending_num, ascending_num int) {
	var ans_num int
	ans_num = descending_num - ascending_num
	if ans_num == int_number {
		fmt.Printf("%d is Kaprekar Number. (Definition 2)\n", int_number)
	} else {
		fmt.Printf("%d is not Kaprekar Number. (Definition 2)\n", int_number)
	}
}

func kaprekar_number_tw(str_number string, int_number int) {
	str_slice := make_str_slice(str_number)
	ascending_order := ascending_sort_str_slice(str_slice)
	ascending_num := make_check_num_tw(ascending_order)
	descending_order := descending_sort_str_slice(str_slice)
	descending_num := make_check_num_tw(descending_order)

	check_kaprekar_number_tw(int_number, descending_num, ascending_num)
}

func main() {
	var str_number string
	var int_number int
	fmt.Print("Please enter number \n> ")

	str_number = scan()
	int_number, _ = strconv.Atoi(str_number)

	kaprekar_number_one(int_number)
	kaprekar_number_tw(str_number, int_number)

	return
}
