package util

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/avp365/arch-pat/internal/pkg/task"
)

func RunTest() {
	files := GetTestFilesFromDirectory()
	dataAssert := AdapterInt(GetDataFromFiles(files))

	for _, v := range dataAssert {
		fmt.Println("test ", v[0], v[0])

		if task.Task(v[0]) == v[1] {
			fmt.Println("test passed")
			continue
		}
		fmt.Println("failed")
	}

}

func GetTestFilesFromDirectory() [][]string {
	files := [][]string{}
	entries, err := os.ReadDir(getPath())
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {

		match, _ := regexp.MatchString(".in", e.Name())

		if match {
			name := strings.Split(e.Name(), ".in")[0]
			files = append(files, []string{getPath() + name + ".in", getPath() + name + ".out"})
		}
	}

	return files

}
func GetDataFromFiles(files [][]string) [][]string {
	dataAssert := [][]string{}
	for _, paths := range files {
		in, _ := os.ReadFile(paths[0])
		out, _ := os.ReadFile(paths[1])
		dataAssert = append(dataAssert, []string{string(in), string(out)})
	}

	return dataAssert

}
func AdapterInt(data [][]string) [][]int {
	fmt.Println(data)
	dataAssert := [][]int{}

	for _, v := range data {
		n, _ := strconv.Atoi(v[0])
		c, _ := strconv.Atoi(v[1])
		dataAssert = append(dataAssert, []int{n, c})
	}

	return dataAssert

}
func getPath() string {
	path, _ := os.Getwd()
	return path + "/data/"
}
