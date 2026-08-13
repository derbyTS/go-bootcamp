package main

import (
	"fmt"
	// "io"
	"os/exec"
	// "strings"
	// "time"
)

func main() {
	//===========Sample ls
	cmd := exec.Command("ls", "-l")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Output: ", string(output))

	//===========Sample pipe
	// This sample is like echo -e "foo\nbar\nbaz" | grep "foo"
	// pr, pw := io.Pipe()
	//
	// cmd := exec.Command("grep", "foo")
	// cmd.Stdin = pr
	//
	// go func() {
	// 	defer pw.Close()
	// 	pw.Write([]byte("foo\nbar\n\baz"))
	// }()
	//
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// fmt.Println("Output: ", string(output))

	//===========Sample print env
	// cmd := exec.Command("printenv", "SHELL")
	//
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// fmt.Println("Output: ", string(output))

	//===========Sample with time process
	// // cmd := exec.Command("sleep", "5")
	// cmd := exec.Command("sleep", "60") // For process kill sample
	//
	// // cmd.Start() // Uncomment this to see the error
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	//
	// time.Sleep(2 * time.Second)
	//
	// // Uncomment to test kill process
	// err = cmd.Process.Kill()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Process kill")

	// Uncomment to test wait
	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Process complete")

	//===========Sample with grep
	// cmd := exec.Command("grep", "foo")
	//
	// cmd.Stdin = strings.NewReader("foo\nbar\nbaz\n")
	//
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// fmt.Println("Output ", string(output))

	//===========Sample simple echo

	// cmd := exec.Command("echo", "Test")
	//
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// fmt.Println("Output: ", string(output))
}
