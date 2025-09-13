/*
Processes spawning refers to creating and managing separate operating system processes from within a goprogram
This involves starting new processes to run tasks concurrently or in isolation from the main program
We use process spanning for Concurrency,Isolation and Resource Management
Concurrency - processes spawning runs tasks in parallel to utilize multiple cpu cores
Isolation - process spawning runs tasks in separate environmets to avoid interference and improve stability
resource management - process spawning offloads resource intensive tasks to separate processes to manage system resources more effectively

we use os/exec package for process spawning
exec.Command
cmd.Stdin/cmd.Stdout
cmd.Start / cmd.Wait
cmd.Output
*/

package advanced

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output:", string(output))
}

// ============pipe write and pipe write
// func main() {
// 	pr, pw := io.Pipe()

// 	cmd := exec.Command("grep", "foo")
// 	cmd.Stdin = pr

// 	go func() {
// 		defer pw.Close()
// 		pw.Write([]byte("food is good\nbar\nbaz\n"))
// 	}()

// 	output, err := cmd.Output()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("Output:", string(output))
// }

// =====DEALING WITH Environmet Variables==========
// func main() {
// 	cmd := exec.Command("printenv", "SHELL")

// 	output, err := cmd.Output()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}

// 	fmt.Println("Output:", string(output))
// }

//============WAIT Command , Kill Command
// func main() {
// 	cmd := exec.Command("sleep", "5")

// 	err := cmd.Start()

// 	if err != nil {
// 		fmt.Println("Error starting command:", err)
// 		return
// 	}

// 	time.Sleep(2 * time.Second)

// 	err = cmd.Process.Kill()
// 	if err != nil {
// 		fmt.Println("Error killing process:", err)
// 		return
// 	}

// 	//========wait command
// 	//Waiting
// 	// err = cmd.Wait()
// 	// if err != nil {
// 	// 	fmt.Println("Error waiting:", err)
// 	// 	return
// 	// }

// 	fmt.Println("Process is completed")
// }

//================grep============

// func main() {
// 	cmd := exec.Command("grep", "foo")

// 	//Set input for the command
// 	cmd.Stdin = strings.NewReader("food is good\nbar\nbaz\n")

// 	output, err := cmd.Output()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("Output:", string(output))
// }

// ==============ECHO==================
// func main() {

// 	//cmd := exec.Command("echo", "Hello World!")
// 	cmd := exec.Command("echo", "Hello World!")
// 	//after we ran the above cmd we need to extract values we receive from this which are output and error
// 	output, err := cmd.Output()
// 	if err != nil {
// 		fmt.Println("error:", err)
// 		return
// 	}

// 	fmt.Println("Output:", string(output))

// }
