package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	fmt.Println("Setting up Air and dependencies...")

	// Step 1: Run `go mod tidy`
	fmt.Println("Running 'go mod tidy'...")
	if err := runCommand("go", "mod", "tidy"); err != nil {
		fmt.Printf("Error running 'go mod tidy': %v\n", err)
		os.Exit(1)
	}

	// Step 2: Install Air
	fmt.Println("Installing Air...")
	if err := runCommand("go", "install", "go install github.com/air-verse/air@latest"); err != nil {
		fmt.Printf("Error installing Air: %v\n", err)
		os.Exit(1)
	}

	// Step 3: Copy Air binary to the current directory
	fmt.Println("Copying Air binary to the current directory...")
	airBinaryPath, err := exec.LookPath("air")
	if err != nil {
		fmt.Printf("Error locating Air binary: %v\n", err)
		os.Exit(1)
	}

	destPath := filepath.Join(".", "air")
	if err := copyFile(airBinaryPath, destPath); err != nil {
		fmt.Printf("Error copying Air binary: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Air binary successfully copied to the current directory as 'air'")
	fmt.Println("Setup complete!")
}

// runCommand runs a shell command and waits for it to finish
func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = destFile.ReadFrom(sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}
