package gid-wrapper

import (
	"fmt"
	"os"
	"syscall"
)

func SetGid() {
	// get the current process id
	_ = os.Getpid()

	_ = syscall.Getgid()

	// set the new gid to 1101
	err := syscall.Setgid(1101)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// verify that the gid has been updated
	newGid := syscall.Getgid()
	if newGid != 1101 {
		fmt.Println("Error: gid not updated")
		return
	}

	fmt.Println("GID updated successfully!")
}
