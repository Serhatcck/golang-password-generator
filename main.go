package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	days   = []string{"pazartesi", "sali", "carsamba", "persembe", "cuma", "cumartesi", "pazar"}
	months = []string{"ocak", "subat", "mart", "nisan", "mayis", "temmuz", "haziran", "agustos", "eylul", "ekim", "kasim", "aralik"}
)

type Pattern interface {
	getList([]string) []string
}

type Year struct {
	value int
}

func (y Year) getList(arr []string) []string {
	var list []string
	t := time.Now()
	nowYear := t.Year()

	for i := 0; i < y.value; i++ {
		list = append(list, strconv.Itoa(nowYear-i))
	}
	return matrix(arr, list)
}

type Month struct {
	value int
}

func (m Month) getList(arr []string) []string {
	return matrix(arr, months)
}

type Day struct {
	value string
}

func (d Day) getList(arr []string) []string {
	gunArr := strings.Split(d.value, "/")
	baslangic, _ := strconv.Atoi(gunArr[0])
	bitis, _ := strconv.Atoi(gunArr[1])
	return matrix(arr, days[baslangic:bitis])
}

type Special struct {
	value string
}

func (s Special) getList(arr []string) []string {
	special := []string{s.value}
	return matrix(arr, special)
}

type StringVar struct {
	value string
}

func (s StringVar) getList(arr []string) []string {
	stringVal := []string{s.value}
	return matrix(arr, stringVal)
}

type LeftPadding struct {
	paddingLength int
	paddingValue  string
}

func (l LeftPadding) getList(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		arr[i] = LeftPadd(arr[i], l.paddingValue, l.paddingLength)
	}

	return arr
}

func LeftPadd(s string, pad string, plength int) string {
	for i := len(s); i < plength; i++ {
		s = pad + s
	}
	return s
}

type RightPadding struct {
	paddingLength int
	paddingValue  string
}

func (r RightPadding) getList(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		arr[i] = RightPadd(arr[i], r.paddingValue, r.paddingLength)
	}

	return arr
}

func RightPadd(s string, pad string, plength int) string {
	for i := len(s); i < plength; i++ {
		s = s + pad
	}
	return s
}

func main() {
	payloads := []string{}
	patterns := parseArgs(os.Args[1:])
	for _, pattern := range patterns {
		payloads = pattern.getList(payloads)
	}

	for _, l := range payloads {
		fmt.Println(l)
	}

}

func parseArgs(args []string) []Pattern {

	patterns := []Pattern{}

	for _, str := range args {
		flag := strings.Split(str, "=")
		untiredFlag := strings.Split(flag[0], "-")[1]
		if untiredFlag == "yil" {
			year, _ := strconv.Atoi(flag[1])
			patterns = append(patterns, Year{year})
		} else if untiredFlag == "ay" {
			patterns = append(patterns, Month{})
		} else if untiredFlag == "gun" {
			patterns = append(patterns, Day{flag[1]})
		} else if untiredFlag == "ozel" {
			patterns = append(patterns, Special{flag[1]})
		} else if untiredFlag == "std" {
			patterns = append(patterns, StringVar{flag[1]})
		} else if untiredFlag == "lpad" {
			padd := strings.Split(flag[1], "/")
			paddingLengt, _ := strconv.Atoi(padd[1])
			patterns = append(patterns, LeftPadding{paddingLength: paddingLengt, paddingValue: padd[0]})
		} else if untiredFlag == "rpad" {
			padd := strings.Split(flag[1], "/")
			paddingLengt, _ := strconv.Atoi(padd[1])
			patterns = append(patterns, RightPadding{paddingLength: paddingLengt, paddingValue: padd[0]})
		}
	}
	return patterns

}

func matrix(arr1 []string, arr2 []string) []string {

	//eğer arr1 boş ise foreach in ilk indexi demektir o yüzden arr1 e arr2 değerini atarız
	if len(arr1) < 1 {
		arr1 = arr2
		return arr1
	}
	var matrix = []string{}
	for i := 0; i < len(arr1); i++ {
		for y := 0; y < len(arr2); y++ {
			matrix = append(matrix, arr1[i]+arr2[y])
		}
	}
	return matrix
}
