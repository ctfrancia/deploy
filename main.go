package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("/bin/bash", "deployment.sh").Output()
	if err != nil {
		fmt.Println("in err")
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
