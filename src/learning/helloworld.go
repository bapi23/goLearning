package main

import ("fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "content"
	filename := "./fromString.txt"
  file, err := os.Create(filename)
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("All done with file of %v charactes", ln)

	bytes := []byte(content)
	ioutil.WriteFile("./fromString2.txt", bytes, 0644)

	content2, err := ioutil.ReadFile(filename)
	checkError(err)
	result := string(content2)
	fmt.Println(result)
}

func checkError(err error) {
	if err != nil{
		panic(err)
	}
}
