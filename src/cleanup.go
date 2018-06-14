package src

import (
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
	err = deleteStorageClass()
	if err != nil {
		log.Printf("Error deleting storage class ::: %s", err)
	}
	err = deleteNameSpaces()
	if err != nil {
		log.Printf("Error deleting namespace ::: %s", err)
	}
}

func deleteNameSpaces() error {
	uuids, err := readUUIDsFromFile()
	if err != nil {
		return err
	}

	for _, uuid := range uuids {
		//cmd := exec.Command("kubectl delete ns", "ns-"+uuid)
		cmd := exec.Command("echo", "ns-"+uuid)
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
	//cmd := exec.Command("kubectl delete -f", fileLocation)
	cmd := exec.Command("cat", fileLocation)

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println("Printing output ::: ", string(stdout))

	return nil
}

func deleteWordpress() error {
	fileLocation := "src/wordpress/wp.yaml"

	uuids, err := readUUIDsFromFile()
	if err != nil {
		return err
	}

	for _, uuid := range uuids {

		command := "kubetpl render " + fileLocation + " -s NAMESPACE=" + "ns-" + uuid + " | kubectl delete -f -"
		cmd := exec.Command(command)

		stdout, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println("Printing output ::: ", string(stdout))
	}

	return nil
}
