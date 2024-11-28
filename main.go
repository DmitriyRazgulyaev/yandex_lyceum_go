package main

import (
	"fmt"
	"module3_1-lesson/internal/workingWithFiles"
)

func main() {
	result := workingWithFiles.LineByNum("/Users/dmitriyrazgulyaev/GolandProjects/lesson_testing/test", -1)
	fmt.Println(result)
}
