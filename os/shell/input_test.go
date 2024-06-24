package demo

func TestInput(t *testing.T){
    // 读取命令行输入
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Read from console failed,err:", err)
			return
		}
        println("input:",input)
    }
}
