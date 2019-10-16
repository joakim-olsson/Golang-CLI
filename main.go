package main

import (
	"bufio"
	"fmt"
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
		if err != nil {
			fmt.Fprintln(os.Stderr, "Something went wrong")
		}
		runCommand(cmdString)
	}
}

func sum(nums []string) int {
	sum := 0
	for _, n := range nums {
		j, _ := strconv.Atoi(n)
		sum += j
	}
	return sum
}
func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "add":
		fmt.Println(sum(arrCommandStr))
	}
	cmd := exec.Command(arrCommandStr[0])
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
