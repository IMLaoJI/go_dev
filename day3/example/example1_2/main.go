package main

import (
	"fmt"
	"strings"
)

func urlProcess(url string) string {

	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}

	return url
}

/*strings.TrimSpace(" sksk ") =>"sksk"
strings.Trim("abbacba", "ab") =>"c"
strings.TrimLeft("a","b")
strings.TrimRight("b","c")
"heheheworld", "he", "wo", 0

strings.Fields("abc cde edk") ["abc", "cde", "edk]

strings.Split("abc,cde,edk", ",") ["abc", "cde", "edk]
strings.Join(["abc", "cde", "edk], ",") "abc,cde,edk"

strings.Replace("str", 3) "strstrstr"*/
func pathProcess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}

	return path
}

func Stringdex(str string, substr string)(int, int) {
	index1 := strings.Index(str, substr)
	index2 := strings.LastIndex(str, substr)
	return index1,index2
}

func main() {
	//var (
	//	url  string
	//	path string
	//)
	//a := strings.TrimLeft("bcbbbabbbv","b")
	fmt.Println(strings.Fields("abc cde edk")[0])
	fmt.Println(strings.Replace("sssssstr", "s","o",3))
	a := []int{1,2,3,3,4,5,2}
	fmt.Println(cap(a))
	fmt.Printf("%p\n",a)
	//a = a[1:len(a)]
	a = append(a[:0], a[1:]...)
	fmt.Println(a)
	fmt.Println(cap(a))
	fmt.Printf("%p",a)
	//fmt.Scanf("%s%s", &url, &path)
	//url = urlProcess(url)
	//path = pathProcess(path)
	//
	//fmt.Println(url)
	//fmt.Println(path)
}
