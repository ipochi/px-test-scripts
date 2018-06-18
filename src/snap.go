package src

import (
	"bytes"
	"fmt"
	"os/exec"
)

func CreateSnapshots(number int) error {

	var out bytes.Buffer
	var stderr bytes.Buffer

	fileLocation := "src/snapshot/snap.yaml"

	uuids, err := readFromFile()
	if err != nil {
		return err
	}

	for _, uuid := range uuids[len(uuids)-number:] {

		replace := "NAMESPACE=ns-" + uuid
		cmd := exec.Command("kubetpl", "render", fileLocation, "-s", replace)
		kubectl := exec.Command("kubectl", "apply", "-f", "-")

		pipe, err := cmd.StdoutPipe()
		fmt.Println("pipe ---, ", pipe)
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

		err = kubectl.Run()
		if err != nil {
			fmt.Println("Printing output ::: ", stderr.String())
			return err
		}
		fmt.Println("Printing output ::: ", out.String())
	}
	return nil

}
