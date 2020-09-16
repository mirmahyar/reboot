/*
Author : Mirmahyar Mirsamadi

Description : A simple CLI tool to reboot the Unix system after a
predefined time set by user.

*/

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	fmt.Println("How much time do you need before reboot (seconds): ")

	var seconds string

	fmt.Scanln(&seconds)

	fmt.Println("Rebooting the system in ", seconds, "seconds ...")

	cmd := exec.Command("sleep", seconds)
	cmd.Run()
	cmd = exec.Command("reboot")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
