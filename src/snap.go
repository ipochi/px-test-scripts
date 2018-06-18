package src

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

func CreateSnapshots(wpnumber int) error {

	var out bytes.Buffer
	var stderr bytes.Buffer

	fileLocation := "src/snapshot/snap.yaml"

	uuids, err := readFromFile()
	if err != nil {
		return err
	}

	for _, uuid := range uuids {

		for i := 0; i < wpnumber; i++ {

			snapuuid := generateUUID()
			ns := "NAMESPACE=ns-" + uuid
			snap := "RANDOM_UUID=" + snapuuid
			wp := "WP_NUMBER=" + strconv.Itoa(i)

			cmd := exec.Command("kubetpl", "render", fileLocation, "-s", ns, "-s", snap, "-s", wp)
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

			err = kubectl.Run()
			if err != nil {
				fmt.Println("Printing output ::: ", stderr.String())
				return err
			}
			fmt.Println("Printing output ::: ", out.String())
		}
	}
	return nil
}

func CreateGroupSnapshot() error {

	var out bytes.Buffer
	var stderr bytes.Buffer
	fileLocation := "src/snapshot/groupsnap.yaml"

	uuids, err := readFromFile()
	if err != nil {
		return err
	}

	for _, uuid := range uuids {

		gsuuid := generateUUID()
		ns := "NAMESPACE=ns-" + uuid
		snapgroup := "SNAP_GROUP=ns-" + uuid

		snapuuid := "RANDOM_UUID=" + gsuuid
		cmd := exec.Command("kubetpl", "render", fileLocation, "-s", ns, "-s", snapuuid, "=s", snapgroup)
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

		err = kubectl.Run()
		if err != nil {
			fmt.Println("Printing output ::: ", stderr.String())
			return err
		}
		fmt.Println("Printing output ::: ", out.String())
	}
	return nil
}
