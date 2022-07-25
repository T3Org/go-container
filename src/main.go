package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main(){
	if len(os.Args) <= 1 {
		Usage()
	}
	switch os.Args[1] {
	case "run":
		run()
	case "initContainer":
		initContainer()
	default:
		Usage()
}
}

func run(){
	fmt.Printf("Running %v as user id %d in process %d", os.Args[2:], os.Getuid(), os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"initContainer"}, os.Args[2:]...)...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNS | syscall.CLONE_NEWPID,
        UidMappings: []syscall.SysProcIDMap{
            { 
                ContainerID: 0, 
                HostID: 1000, 
                Size: 1,
            },
        },
        GidMappings: []syscall.SysProcIDMap{
            { 
                ContainerID: 0, 
                HostID: 1000, 
                Size: 1,
            },
        },
    }
    must(cmd.Run())
}

func Usage(){
	fmt.Printf("Usage: run <command>\n")
	os.Exit(1)
}

func initContainer() {
    fmt.Printf("Running %v as user id %d in process %d\n", os.Args[2:], os.Getuid(), os.Getpid())

    must(syscall.Chroot("/"))
    must(os.Chdir("/"))
    must(syscall.Mount("proc", "proc", "proc", 0, ""))

    cmd := exec.Command(os.Args[2], os.Args[3:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    must(cmd.Run())
    must(syscall.Unmount("proc", 0))
}

func must (err error) {
    if err != nil {
        panic(err)
    }
}
