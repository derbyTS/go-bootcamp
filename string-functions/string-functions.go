package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func init() {
	fmt.Println("It go first")
	const a = "asd"
}

func main() {
	str := "Hello Go!"
	fmt.Println(len(str))

	for _, v := range str {
		fmt.Printf("%c ", v)
		fmt.Println(v == 'e')
	}
	fmt.Println(str[1])
	r := []rune(str)
	fmt.Println(r)

	num := 16
	var str1 string = strconv.Itoa(num)
	fmt.Printf("%T\n", str1)
	fmt.Println(reflect.TypeOf(str1))

	fruits := "apple banana orange"
	fruitsArray := strings.Split(fruits, " ")
	fmt.Println(fruitsArray[0])

	str2 := "apple"
	str2Slice := strings.Split(str2, "")
	fmt.Println(reflect.TypeOf(str2Slice))
	str2Slice = append(str2Slice, "s")
	fmt.Println(str2Slice)
	fmt.Printf("capacity is %d and len is %d\n", cap(str2Slice), len(str2Slice))
	m := make(map[string]string, 0)
	m["s"] = "s"
	m["t"] = "t"
	m["u"] = "u"
	fmt.Println(m)

	if v, ok := m["s"]; ok {
		fmt.Println(v)
	}
	// fmt.Println(v) // this is invalid cause the v is inside live inside the if condition only

	sValue, isExist := m["s"]
	fmt.Printf("Is it exist: %v\n", isExist)
	fmt.Printf("The value is %v\n", sValue)

	str3 := []string{"h", "e", "l", "l", "o"}
	str3Join := strings.Join(str3, "")
	fmt.Println(str3Join)
	fmt.Println(strings.Contains(str3Join, "lo"))
	fmt.Println(strings.Contains(str3Join, "x"))

	str4 := " go go go go go"
	strReplace := strings.Replace(str4, "go", "lang", 33)
	fmt.Println(strReplace)
	fmt.Println(strings.TrimSpace(strReplace))
	fmt.Println(strings.ReplaceAll(strReplace, " ", ""))
	fmt.Println(strings.ToUpper(strReplace))
	fmt.Println(strings.ToLower(strReplace))
	fmt.Println(strings.Repeat(strReplace, 3))
	fmt.Println(strings.Count(strReplace, "lang"))
	fmt.Println(strings.HasPrefix(strReplace, " l"))
	fmt.Println(strings.HasPrefix(strReplace, "La"))
	fmt.Println(strings.HasSuffix(strReplace, "ng"))
	fmt.Println(strings.HasSuffix(strReplace, "lanG"))

	fmt.Println("Start of regex ....................")
	strReg1 := "Hell0, 123 Go 99"
	regexPattern1 := regexp.MustCompile(`\d+`)
	regexMatches := regexPattern1.FindAllString(
		strReg1,
		-1,
	) // It will return a slice of string and -1 means all matches
	fmt.Println(regexMatches)
	fmt.Println("Start of Unicode Package ....................")
	strUnicode := "愛してます Gracie"
	fmt.Println(utf8.RuneCountInString(strUnicode))

	fmt.Println("Start of StringBuilder ....................")
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")

	strResult := builder.String()
	fmt.Println(strResult)

	builder.WriteRune(' ')
	builder.WriteString("Again")
	fmt.Println(builder.String())

	builder.Reset()
	builder.WriteString("New")

	fmt.Println(builder.String())
	fmt.Println(builder.Len())
}

func testSnipp() (int, error) {
	// return 1, errors.New("Err")
	return 1, fmt.Errorf("Error %s", "err")
}
