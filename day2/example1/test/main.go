package test1


import(
	"fmt"
)

func init() {
	fmt.Println("shawanyi")
}

func List(n int) string{

	for i := 0; i <= n; i++ {
		fmt.Printf("%d+%d=%d\n", i, n - i, n)
	}
	return "test"
}


