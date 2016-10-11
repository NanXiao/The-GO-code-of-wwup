package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"log"
	"os"
	"os/exec"
)

func main() {
	/* The arguments of process */
	fmt.Printf("There are %v arguments:\n", len(os.Args)-1)
	for i, v := range os.Args[1:] {
		fmt.Printf("Argument %v is %v\n", i, v)
	}
	fmt.Println()

	/* Process and parent process ids */
	fmt.Printf("The process id is %v.\n", os.Getpid())
	fmt.Printf("The parent process id is %v.\n", os.Getppid())
	fmt.Println()

	/* File descriptors */
	fmt.Printf("File descriptor of stdin is %v.\n", os.Stdin.Fd())
	fmt.Printf("File descriptor of stdin is %v.\n", os.Stdout.Fd())
	fmt.Printf("File descriptor of stdin is %v.\n", os.Stderr.Fd())
	file, err := os.Open("./wwup.go")
	if err == nil {
		fmt.Printf("File descriptor of file is %v.\n", file.Fd())
		file.Close()
	}
	fmt.Println()

	/* Resource limits */
	var olim unix.Rlimit
	err = unix.Getrlimit(unix.RLIMIT_NOFILE, &olim)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Now Current File Number Limit is %v.\n", olim.Cur)
	fmt.Printf("Now Maximum File Number Limit is %v.\n", olim.Max)

	nlim := unix.Rlimit{4096, 4096}
	err = unix.Setrlimit(unix.RLIMIT_NOFILE, &nlim)
	if err != nil {
		log.Fatal(err)
	}
	var rlim unix.Rlimit
	err = unix.Getrlimit(unix.RLIMIT_NOFILE, &rlim)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New current File Number Limit is %v.\n", rlim.Cur)
	fmt.Printf("New maximum File Number Limit is %v.\n", rlim.Max)

	err = unix.Setrlimit(unix.RLIMIT_NOFILE, &olim)
	if err != nil {
		log.Fatal(err)
	}
	err = unix.Getrlimit(unix.RLIMIT_NOFILE, &rlim)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Now current File Number Limit is %v.\n", rlim.Cur)
	fmt.Printf("Now maximum File Number Limit is %v.\n", rlim.Max)
	fmt.Println()

	/* Environment variables */
	err = os.Setenv("FOO", "foo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The value of FOO is %v.\n", os.Getenv("FOO"))
	os.Unsetenv("FOO")
	fmt.Println()

	/* Create new process */
	cmd := exec.Command("sleep", "5")
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Waiting for sleep command to finish.")
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
