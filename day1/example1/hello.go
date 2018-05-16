package main

import (
	"fmt"
	"strconv"
	//"strings"
	"strings"
)

func main() {
	//fmt.Println("hello world")
	//s := make([]byte, 5)
	//fmt.Println(cap(s),"----",len(s))
	//fmt.Printf("%p\n",s)
	//s1 := s[2:4]
	//fmt.Println(cap(s),"----",len(s))
	//fmt.Printf("%p\n",s)
	//fmt.Printf("%p\n",s1)
	//s2 := []byte("")
	//fmt.Println(cap(s2),"----",len(s2))
	//fmt.Printf("%p\n",s2)
	//
	//s3 := append(s2, 'a')
	//fmt.Printf("%p\n",s2)
	//fmt.Println(cap(s3),"----",len(s3))
	//fmt.Printf("%p\n",s3)
	//s4 :=[]int{1,2,3,4,5,6,7}
	//fmt.Println(cap(s4),"----",len(s4))
	//fmt.Printf("%p\n",s4)
	//s5 := append(s4,2)
	//fmt.Println(cap(s4),"----",len(s4))
	//fmt.Printf("%p\n",s4)
	//fmt.Println(cap(s5),"----",len(s5))
	//fmt.Printf("%p\n",s5)
	//s := []byte{}
	//fmt.Printf("%p\n",s)
	//s1 := append(s, 'a') // 等同于 arr[0] = 'a'
	//fmt.Printf("%p\n",s1)
	//s2 := append(s, 'b') // 等同于 arr[0] = 'b'
	//fmt.Printf("%p\n",s2)
	//fmt.Println(string(s1), "==========", string(s2)) // 只是把同一份数组打印出来了
	//fmt.Println(cap(s), len(s))

	//s3 := []byte("")
	//fmt.Printf("%p\n",s3)
	//s4 := append(s3, 'a')
	//fmt.Printf("%p\n",s4)
	//s5 := append(s3, 'b')
	//fmt.Printf("%p\n",s5)
	//fmt.Println(s4, "==========", s5)
	//fmt.Println(string(s4), "==========", string(s5))
	//fmt.Println(cap(s3), len(s3))
	//s := []byte{}
	//s1 := append(s, 'a')
	//s2 := append(s, 'b')
	//fmt.Println(string(s1), ",", string(s2))
	//fmt.Println(cap(s), len(s))

	//在make([]byte,0,0)这样情况下，s容量肯定不够用，所以s1，s2使用的都是各自从s复制出来的数组，结果也自然符合预期a，b了。
	//s := make([]byte, 0, 0)
	//fmt.Printf("%p\n",s)
	//s = append(s, 'a')
	//fmt.Printf("%p\n",s)
	//s = append(s, 'b')
	//fmt.Printf("%p\n",s)

	var c rune = '3'
	//var i int =98
	i1, _ := strconv.Atoi(string(c))
	fmt.Println("'a' convert to", i1)
	//fmt.Println(string(s1), ",", string(s2))
	//s := make([]byte, 0, 1)
	//fmt.Printf("%p\n",s)
	//fmt.Println(cap(s), len(s))
	//s1 := append(s, 'a')
	//fmt.Printf("%p\n",s1)
	//fmt.Println(cap(s1), len(s1))
	//s2 := append(s, 'b')
	//fmt.Printf("%p\n",s2)
	//fmt.Println(string(s1), ",", string(s2))

	var i int = 104
	var j int = 84
	fmt.Println(string(rune(i)))
	fmt.Println(string(rune(j)))
	str1 := "abca"
	var  a rune = 'a'
	var n rune = 'E'
	fmt.Println(n,"sss")
	str1 += "123"
	index := strings.Index(str1, "a")
	tem := []rune(str1)
	for i := 0; i < len(tem); i++ {
		if i==index {
			tem = append(tem[:i], tem[i+1:len(tem)]...)
		}

	}
	fmt.Println(a==n)
	fmt.Println(len(str1))
	s := string(tem)
	//contains := strings.Contains(str1, "b")
	fmt.Println(str1)
	fmt.Println(index)
	fmt.Println(s)
	ab := []int{1,2,3,4,5}
	test(ab)
	fmt.Println(ab)
	aa :="a"
	fmt.Println(int(aa[0]))

}

func test(test []int)  {
	test[1] =0
}
