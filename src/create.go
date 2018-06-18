package src

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"

	uuid "github.com/satori/go.uuid"
)

func Create(number int) {

	var err error

	for i := 0; i < number; i++ {
		uuid := generateUUID()
		err = writeToFile(uuid)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = createNamespace(number)
	if err != nil {
		log.Printf("Error creating namespace ::: %s", err)
	}
	// err = createStorageClass()
	// if err != nil {
	// 	log.Printf("Error creating storage class ::: %s", err)
	// }
	err = deployWordpress(number)
	if err != nil {
		log.Printf("Error deploying wordpress ::: %s", err)
	}
}

func generateUUID() string {
	u1 := uuid.Must(uuid.NewV4())
	s := hex.EncodeToString(u1.Bytes()[:4])

	fmt.Println("uuid --- ", s)

	return s
}

func writeToFile(uuid string) error {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("/tmp/ns-test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.Write([]byte(uuid + "\n")); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func readFromFile() ([]string, error) {

	var uuids []string
	f, err := os.Open("/tmp/ns-test")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)

	for fileScanner.Scan() {
		uuid := fileScanner.Text()
		uuids = append(uuids, uuid)
	}

	return uuids, err
}

func createNamespace(number int) error {
	uuids, err := readFromFile()
	if err != nil {
		return err
	}

	for _, uuid := range uuids[len(uuids)-number:] {
		var stderr bytes.Buffer

		fmt.Println(uuid)
		cmd := exec.Command("kubectl", "create", "ns", "ns-"+uuid)
		//cmd := exec.Command("echo", "ns-"+uuid)
		cmd.Stderr = &stderr
		stdout, err := cmd.Output()

		fmt.Println("Stderr --- ", stderr.String())
		if err != nil {
			return err
		}
		fmt.Println("Printing output ::: ", string(stdout))
	}
	return nil
}

func createStorageClass() error {

	fileLocation := "src/storageclass/sc.yaml"
	cmd := exec.Command("kubectl", "create", "-f", fileLocation)
	//cmd := exec.Command("cat", fileLocation)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println("Printing output ::: ", string(stdout))

	return nil
}

func deployWordpress(number int) error {

	var out bytes.Buffer
	var stderr bytes.Buffer

	fileLocation := "src/wordpress/wp.yaml"

	uuids, err := readFromFile()
	if err != nil {
		return err
	}

	for _, uuid := range uuids[len(uuids)-number:] {

		command := "NAMESPACE=ns-" + uuid
		//+ " | kubectl apply -f -"
		fmt.Println("Printing command ::: ", command)
		cmd := exec.Command("kubetpl", "render", fileLocation, "-s", command)
		kubectl := exec.Command("kubectl", "apply", "-f", "-")

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
