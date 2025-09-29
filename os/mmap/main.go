package main

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

func main() {
	// 1. 创建或打开一个文件
	fileName := "mapped_file.log"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// 确保文件有足够的空间（Mmap 通常需要文件具有一定大小）
	// 这里写入一些占位符，确保文件大小至少为 1024000 字节
	if _, err := file.Seek(1023999, 0); err != nil {
		fmt.Fprintf(os.Stderr, "Error seeking file: %v\n", err)
		os.Exit(1)
	}
	if _, err := file.Write([]byte{0}); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
		os.Exit(1)
	}

	// 2. 获取文件描述符
	fd := int(file.Fd())

	// 3. 内存映射
	// 映射 1024000 字节，允许读写，共享映射
	mappedData, err := syscall.Mmap(fd, 0, 1024000, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error mmaping file: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		// 5. 解除映射
		if err := syscall.Munmap(mappedData); err != nil {
			fmt.Fprintf(os.Stderr, "Error munmaping file: %v\n", err)
		}
	}()

	// 4. 读写内存
	// 写入数据
	copy(mappedData[0:5], []byte("Hello"))
	copy(mappedData[5:10], []byte(" World"))

	// 读取数据
	fmt.Printf("Data read from mapped memory: %s\n", mappedData[0:10])

	// 6. 同步到文件 (对于 MAP_SHARED，虽然修改会写回，但显式同步更保险)
	// if err := syscall.Msync(mappedData, syscall.MS_SYNC); err != nil {
	// 把该区间对应的页写回磁盘：但只有被标记为脏（modified）的页才会产生实际写回 I/O(前10个字节修改的就是脏页)
	// 或更明确地用：unix.Msync(mappedData[:N], unix.MS_SYNC)
	if err := unix.Msync(mappedData, unix.MS_SYNC); err != nil {
		fmt.Fprintf(os.Stderr, "Error msyncing file: %v\n", err)
	}

	fmt.Println("Data written and synced to file.")

}
