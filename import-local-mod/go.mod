module c/go_cli_demo

go 1.15

require (
	ahui1 v0.0.0
	github.com/ahuigo/go-hello v0.0.0-20190325051759-913dff133b48
	github.com/ahuigo/requests v1.0.8
)

replace ahui1 v0.0.0 => ./ahui1

// 外部依赖也行
//replace ahui1 v0.0.0 => ../ahui1
