package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
)

const (
	serverHost = "192.168.10.216" // 服务端地址
	serverPort = "8080"      // 服务端端口
)

func main() {
	if runtime.GOOS == "linux" {
		runServer() // 在 Linux 上运行服务端
	} else {
		runClient() // 在 Windows 或 Mac 上运行客户端
	}
}

// runServer 启动服务端监听客户端连接
func runServer() {
	fmt.Println("Running server on", serverHost+":"+serverPort)
	listener, err := net.Listen("tcp", serverHost+":"+serverPort)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn) // 处理每个客户端连接
	}
}

// handleConn 处理客户端连接
func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connected to", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	for {
		cmd, err := reader.ReadString('\n') // 读取客户端发送的命令
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println("Received command:", cmd)
		output, err := exec.Command("/bin/bash", "-c", cmd).CombinedOutput() // 在 bash 中执行命令并获取输出
		if err != nil {
			log.Println(err)
			output = []byte(err.Error()) // 如果出错，返回错误信息
		}
		conn.Write(output) // 将输出结果发送给客户端
	}
}

// runClient 启动客户端连接服务端并发送命令
func runClient() {
	fmt.Println("Connecting to", serverHost+":"+serverPort)
	conn, err := net.Dial("tcp", serverHost+":"+serverPort)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter command: ") // 输入要执行的命令
		scanner.Scan()
		cmd := scanner.Text()
		if cmd == "exit" { // 如果输入 exit，退出程序
			break
		}
		conn.Write([]byte(cmd + "\n"))     // 将命令发送给服务端
		output, err := reader.ReadString('\n') // 读取服务端返回的结果
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println("Output:", output) // 打印输出结果
	}
}