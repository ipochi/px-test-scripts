package src

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Cleanup does the cleanup of storageclasses,namespaces
func Cleanup() {

	var err error

	err = deleteWordpress()
	if err != nil {
		log.Printf("Error deploying wordpress ::: %s", err)
	}
	// err = deleteStorageClass()
	// if err != nil {
	// 	log.Printf("Error deleting storage class ::: %s", err)
	// }
	err = deleteNameSpaces()
	if err != nil {
		log.Printf("Error deleting namespace ::: %s", err)
	}
}

func deleteNameSpaces() error {
	uuids, err := readFromFile()
	if err != nil {
		return err
	}

	for _, uuid := range uuids {
		cmd := exec.Command("kubectl", "delete", "ns", "ns-"+uuid)
		//cmd := exec.Command("echo", "ns-"+uuid)
		stdout, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println("Printing output ::: ", string(stdout))
	}

	err = os.Remove("/tmp/ns-test")

	if err != nil {
		return err
	}

	return nil

}
func deleteStorageClass() error {

	fileLocation := "src/storageclass/sc.yaml"
	cmd := exec.Command("kubectl", "delete", "-f", fileLocation)
	//cmd := exec.Command("cat", fileLocation)

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println("Printing output ::: ", string(stdout))

	return nil
}

func deleteWordpress() error {

	var out bytes.Buffer
	var stderr bytes.Buffer

	fileLocation := "src/wordpress/wp.yaml"

	uuids, err := readFromFile()
	if err != nil {
		return err
	}

	for _, uuid := range uuids {

		command := "NAMESPACE=ns-" + uuid
		//+ " | kubectl apply -f -"
		fmt.Println("Printing command ::: ", command)
		cmd := exec.Command("kubetpl", "render", fileLocation, "-s", command)
		kubectl := exec.Command("kubectl", "delete", "-f", "-")

		pipe, err := cmd.StdoutPipe()
		defer pipe.Close()
		if err != nil {
			return err
		}

		kubectl.Stdin = pipe
		err = cmd.Start()
		if err != nil {
			return err
		}

		kubectl.Stdout = &out
		kubectl.Stderr = &stderr

		//cmd := exec.Command("echo", "ns-"+uuid)
		err = kubectl.Run()
		if err != nil {
			fmt.Println("Printing output ::: ", stderr.String())
			return err
		}
		fmt.Println("Printing output ::: ", out.String())
	}
	return nil
}
