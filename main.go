package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		cmdString, err:= reader.ReadString('\n')
		if len(cmdString) > 2 {
			runCommand(cmdString)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Something went wrong")
		}
	}
}

func sum(nums []string) int {
	res := 0
	for _, n := range nums {
		j, _ := strconv.Atoi(n)
		res += j
	}
	return res
}

func sub(nums []string) int {
	if len(nums) < 2 {
		main()
	}

	res, _ := strconv.Atoi(nums[1])
	for i := 2; i < len(nums); i++ {
	j, _ := strconv.Atoi(nums[i])
		res -= j
	}
	return res
}

func mul(nums []string) int {
	if len(nums) < 2 {
		main()
	}
	res, _ := strconv.Atoi(nums[1])

	for i := 2; i < len(nums); i++ {
		j, _ := strconv.Atoi(nums[i])
		res *= j
	}
	return res
}

func div(nums []string) float64 { // TODO
	if len(nums) < 2 {
		main()
	}

	res, _ := strconv.ParseFloat(nums[1], 64)
	for i := 2; i < len(nums); i++ {
		j, _ := strconv.ParseFloat(nums[i], 64)
		res /= j
	}
	return res
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println("Could not find source file")
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		fmt.Println("Could not find destination file")
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func WriteToFile(args []string) {
	fileName := args[1]
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	for i := 2; i < len(args); i++ {
		src.WriteString(args[i] + " ")
	}
	fmt.Printf("Wrote %d words to file \n", len(args)-2)
}

func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "add":
		fmt.Println(sum(arrCommandStr))
	case "sub":
		fmt.Println(sub(arrCommandStr))
	case "mul":
		fmt.Println(mul(arrCommandStr))
	case "div":
		fmt.Println(div(arrCommandStr))
	case "copy":
		CopyFile(arrCommandStr[1], arrCommandStr[2])
	case "write":
		WriteToFile(arrCommandStr)
	}
	cmd := exec.Command(arrCommandStr[0])
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
