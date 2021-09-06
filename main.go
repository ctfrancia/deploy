package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("/bin/bash", "/opt/adamo/deployment/bin/deployment.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
