package main

import "fmt"

type Languege struct{
	Name string
	Count int
}

func main(){
	zalupka:=[]string{
    "python",
    "java",
    "go",
    "python",
    "go",
}

mapes := make(map[string]int)

for _, v := range zalupka{
	mapes[v]++
}


for i,v := range mapes{
	fmt.Print(i)
	fmt.Println(v)
}

}
