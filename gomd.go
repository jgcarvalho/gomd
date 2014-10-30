package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func prepare(pdb string) (string, error) {
	gro := strings.Replace(pdb, ".pdb", ".gro", 1)
	pdb2gmx := exec.Command("gmx_mpi", "pdb2gmx", "-f", pdb, "-o", gro)
	pdb2gmx.Stdin = strings.NewReader("6\n1\n")
	err := pdb2gmx.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for pdb2gmx to finish...")
	err = pdb2gmx.Wait()
	if err != nil {
		log.Printf("Command pdb2gmx finished with error: %v", err)
		return "", err
	}
	return gro, nil
}

func run() {}

func main() {
	fmt.Println("Go MD!!!")
	pdb := "./CYP_noSS.pdb"

	prepare(pdb)
}
