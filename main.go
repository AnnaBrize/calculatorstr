package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	strbool1, strbool2     bool
	strok1, strok2, result string
	chislo                 int
)

func prov(stroka string) (str1 string, str2 string, str3 string, bool1 bool, bool2 bool) {
	lenslise := strings.Count(stroka, "\"")
	x := strings.Index(stroka, "\"")
	switch {
	case lenslise == 4:
		strslice := strings.SplitAfter(stroka, "\"")

		str1 = strings.Trim(strslice[1], "\" ")
		str2 = strings.Trim(strslice[3], "\" \n")
		str3 = strings.Trim(strslice[2], "\" ")
		bool1 = true
		bool2 = true
	case lenslise == 2:
		if x == 0 {
			strslice := strings.SplitAfter(stroka, "\"")
			ss := strings.SplitAfter(strslice[2], " ")

			str1 = strings.Trim(strslice[1], "\" ")
			str2 = strings.TrimSpace(ss[2])
			str3 = strings.TrimSpace(ss[1])
			bool1 = true
			bool2 = false
		} else {
			str1, str2, str3 = "", "", ""
			bool1 = false
			bool2 = false
		}
	default:
		str1, str2, str3 = "", "", ""
		bool1 = false
		bool2 = false
	}
	return
}

// деление строки на число
func del(stroka string, ch int) (strok string) {
	strclice := strings.Split(stroka, "")
	arr := []string{}
	j := len(strclice) / ch
	for i := 0; i < j; i++ {
		arr = append(arr, strclice[i])
	}
	strok = strings.Join(arr, "")
	return
}

// уьножение строки на число
func um(stroka string, ch int) (strok string) {
	resstr := strings.Repeat(stroka, ch)
	if len(resstr) > 40 {
		strclice := strings.Split(resstr, "")
		arr := []string{}
		for i := 0; i < 40; i++ {
			arr = append(arr, strclice[i])
		}
		strok = strings.Join(arr, "") + "..."
	} else {
		strok = resstr
	}
	return
}

func main() {
	read := bufio.NewReader(os.Stdin)
	text, _ := read.ReadString('\n')

	strok1, strok2, z, strbool1, strbool2 := prov(text)

	if strbool1 == true && len(strok1) <= 10 {
		if strbool2 == true && len(strok2) <= 10 {
			switch {
			case z == "+":
				result = strok1 + strok2
			case z == "-":
				result = strings.Replace(strok1, strok2, "", 1)
			default:
				fmt.Println("error")
				return
			}
		} else if strbool2 == false {
			chislo, _ := strconv.Atoi(strok2)
			if chislo <= 10 && chislo >= 1 {
				switch {
				case z == "*":
					result = um(strok1, chislo)
				case z == "/":
					result = del(strok1, chislo)
				default:
					fmt.Println("error")
					return
				}
			} else {
				fmt.Println("error")
				return
			}
		} else {
			fmt.Println("error")
			return
		}
	} else {
		fmt.Println("error")
		return
	}

	result = "\"" + result + "\""
	fmt.Println(result)
}
