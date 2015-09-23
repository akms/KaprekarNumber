package main

import (
	"bufio"
	"fmt"
	"os"
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

func make_check_num(square_slice []string, digit int) (check_num int) {
	str_num_a := square_slice[0:(digit / 2)]
	str_num_b := square_slice[(digit / 2):digit]
	int_num_a, _ := strconv.Atoi(strings.Join(str_num_a, ""))
	int_num_b, _ := strconv.Atoi(strings.Join(str_num_b, ""))
	check_num = int_num_a + int_num_b
	return
}

func check_kaprekar_number(int_number int, check_num int) {
	if int_number == check_num {
		fmt.Printf("%d is Kaprekar Number\n", int_number)
	} else {
		fmt.Printf("%d is not Kaprekar Number\n", int_number)
	}
}

func main() {
	var str_number, square_str string
	var digit, int_number, square_num, check_num int
	fmt.Print("Please enter number \n> ")

	str_number = scan()
	int_number, _ = strconv.Atoi(str_number)
	square_num = square(int_number)
	fmt.Println(square_num)
	square_str = strconv.Itoa(square_num)
	digit = utf8.RuneCountInString(square_str)
	square_slice := make_str_slice(square_str)
	check_num = make_check_num(square_slice, digit)
	check_kaprekar_number(int_number, check_num)

	return
}
