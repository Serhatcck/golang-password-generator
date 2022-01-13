package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	gunler = []string{"pazartesi", "sali", "carsamba", "persembe", "cuma", "cumartesi", "pazar"}
	aylar  = []string{"ocak", "subat", "mart", "nisan", "mayis", "temmuz", "haziran", "agustos", "eylul", "ekim", "kasim", "aralik"}
)

type payloads interface {
	executePayloads() []string
}
type Yil struct {
	value int
}
type Ay struct {
	value bool
}
type Gun struct {
	value string
}

type Pattern struct {
	ay           Ay
	yil          Yil
	gun          Gun
	ozelkarakter []string
	siralama     []string
	padding      string
}

func (yil Yil) executePayloads() []string {
	var payloads []string
	t := time.Now()
	year := t.Year()

	for i := 0; i < yil.value; i++ {
		payloads = append(payloads, strconv.Itoa(year-i))
	}
	return payloads
}

func (ay Ay) executePayloads() []string {
	return aylar
}

func (gun Gun) executePayloads() []string {
	gunArr := strings.Split(gun.value, "-")
	baslangic, _ := strconv.Atoi(gunArr[0])
	bitis, _ := strconv.Atoi(gunArr[1])
	return gunler[baslangic:bitis]
}

func (pattern Pattern) getField(key string) reflect.Value {
	r := reflect.ValueOf(pattern)
	f := reflect.Indirect(r).FieldByName(key)
	return f
}

func (pattern Pattern) payloads(key string, val string) []string {
	if key == "yil" {
		return pattern.yilPayloads(val)
	} else {
		return []string{}
	}
}

func (pattern Pattern) yilPayloads(val string) []string {

	return []string{}
}

func (pattern Pattern) generate() {
	list := []string{}
	var appendCount = 0
	for _, key := range pattern.siralama {
		if appendCount == 0 {
			list = append(list, getPayloads(pattern, key)...)
		} else {
			fmt.Println(key)
			var matrix = []string{}
			payloads := getPayloads(pattern, key)
			for i := 0; i < len(list); i++ {
				for y := 0; y < len(payloads); y++ {
					matrix = append(matrix, list[i]+payloads[y])
				}
			}
			list = matrix
		}
		appendCount++
	}
	for _, l := range list {
		fmt.Println(l)
	}
}

func getPayloads(pattern Pattern, field string) []string {
	if field == "yil" {
		return pattern.yil.executePayloads()
	} else if field == "ay" {
		return pattern.ay.executePayloads()
	} else if field == "gun" {
		return pattern.gun.executePayloads()
	}
	return []string{}
}

func main() {
	var pattern Pattern

	flag.BoolVar(&pattern.ay.value, "ay", false, "Desene Ay eklensin mi")
	flag.IntVar(&pattern.yil.value, "yil", 0, "Desende şimdiki yıldan kaç sene önceki yıllar eklensin")
	flag.StringVar(&pattern.gun.value, "gun", "", "Desen haftanın kaç günü ile dahil edilsin")
	flag.Parse()

	pattern.siralama = parseArgs(os.Args[1:])
	pattern.generate()

}

func parseArgs(args []string) []string {
	var parsedArgs []string
	for _, str := range args {
		flag := strings.Split(str, "=")[0]
		untiredFlag := strings.Split(flag, "-")[1]
		parsedArgs = append(parsedArgs, untiredFlag)
	}
	return parsedArgs
}
