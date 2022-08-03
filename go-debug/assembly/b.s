# command-line-arguments
"".add STEXT nosplit size=19 args=0x18 locals=0x0
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:3)	TEXT	"".add(SB), NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:3)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:4)	PCDATA	$2, $0
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:4)	PCDATA	$0, $0
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:4)	MOVQ	"".b+16(SP), AX
	0x0005 00005 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:4)	MOVQ	"".a+8(SP), CX
	0x000a 00010 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:4)	ADDQ	CX, AX
	0x000d 00013 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:4)	MOVQ	AX, "".~r2+24(SP)
	0x0012 00018 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:4)	RET
	0x0000 48 8b 44 24 10 48 8b 4c 24 08 48 01 c8 48 89 44  H.D$.H.L$.H..H.D
	0x0010 24 18 c3                                         $..
"".main STEXT nosplit size=2 args=0x0 locals=0x0
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:7)	TEXT	"".main(SB), NOSPLIT|ABIInternal, $0-0
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:7)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:8)	PCDATA	$2, $0
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:8)	PCDATA	$0, $0
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:8)	XCHGL	AX, AX
	0x0001 00001 (<unknown line number>)	RET
	0x0000 90 c3                                            ..
"".init.0 STEXT size=94 args=0x0 locals=0x8
	0x0000 00000 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	TEXT	"".init.0(SB), ABIInternal, $8-0
	0x0000 00000 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	MOVQ	(TLS), CX
	0x0009 00009 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	CMPQ	SP, 16(CX)
	0x000d 00013 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	JLS	87
	0x000f 00015 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	SUBQ	$8, SP
	0x0013 00019 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	MOVQ	BP, (SP)
	0x0017 00023 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	LEAQ	(SP), BP
	0x001b 00027 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x001b 00027 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	PCDATA	$2, $1
	0x001b 00027 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	PCDATA	$0, $0
	0x001b 00027 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	MOVQ	runtime/debug.modinfo(SB), AX
	0x0022 00034 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	MOVQ	runtime/debug.modinfo+8(SB), CX
	0x0029 00041 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build33# command-line-arguments
"".add STEXT nosplit size=19 args=0x18 locals=0x0
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:3)	TEXT	"".add(SB), NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:3)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:4)	PCDATA	$2, $0
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:4)	PCDATA	$0, $0
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:4)	MOVQ	"".b+16(SP), AX
	0x0005 00005 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:4)	MOVQ	"".a+8(SP), CX
	0x000a 00010 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:4)	ADDQ	CX, AX
	0x000d 00013 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:4)	MOVQ	AX, "".~r2+24(SP)
	0x0012 00018 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:4)	RET
	0x0000 48 8b 44 24 10 48 8b 4c 24 08 48 01 c8 48 89 44  H.D$.H.L$.H..H.D
	0x0010 24 18 c3                                         $..
"".main STEXT nosplit size=2 args=0x0 locals=0x0
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:7)	TEXT	"".main(SB), NOSPLIT|ABIInternal, $0-0
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:7)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:8)	PCDATA	$2, $0
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:8)	PCDATA	$0, $0
	0x0000 00000 (/Users/ahui/www/go-lib/go-debug/assembly/file.go:8)	XCHGL	AX, AX
	0x0001 00001 (<unknown line number>)	RET
	0x0000 90 c3                                            ..
"".init.0 STEXT size=94 args=0x0 locals=0x8
	0x0000 00000 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	TEXT	"".init.0(SB), ABIInternal, $8-0
	0x0000 00000 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	MOVQ	(TLS), CX
	0x0009 00009 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	CMPQ	SP, 16(CX)
	0x000d 00013 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	JLS	87
	0x000f 00015 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	SUBQ	$8, SP
	0x0013 00019 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	MOVQ	BP, (SP)
	0x0017 00023 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	LEAQ	(SP), BP
	0x001b 00027 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x001b 00027 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	PCDATA	$2, $1
	0x001b 00027 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	PCDATA	$0, $0
	0x001b 00027 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	MOVQ	runtime/debug.modinfo(SB), AX
	0x0022 00034 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	MOVQ	runtime/debug.modinfo+8(SB), CX
	0x0029 00041 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	MOVQ	CX, "".keepalive_modinfo+8(SB)
	0x0030 00048 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	PCDATA	$2, $-2
	0x0030 00048 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	PCDATA	$0, $-2
	0x0030 00048 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	CMPL	runtime.writeBarrier(SB), $0
	0x0037 00055 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	JNE	73
	0x0039 00057 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	MOVQ	AX, "".keepalive_modinfo(SB)
	0x0040 00064 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	MOVQ	(SP), BP
	0x0044 00068 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	ADDQ	$8, SP
	0x0048 00072 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	RET
	0x0049 00073 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	LEAQ	"".keepalive_modinfo(SB), DI
	0x0050 00080 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	CALL	runtime.gcWriteBarrier(SB)
	0x0055 00085 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	JMP	64
	0x0057 00087 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	NOP
	0x0057 00087 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	PCDATA	$0, $-1
	0x0057 00087 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	PCDATA	$2, $-1
	0x0057 00087 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	CALL	runtime.morestack_noctxt(SB)
	0x005c 00092 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 48 48  eH..%....H;a.vHH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 48 8b 05 00 00  ...H.,$H.,$H....
	0x0020 00 00 48 8b 0d 00 00 00 00 48 89 0d 00 00 00 00  ..H......H......
	0x0030 83 3d 00 00 00 00 00 75 10 48 89 05 00 00 00 00  .=.....u.H......
	0x0040 48 8b 2c 24 48 83 c4 08 c3 48 8d 3d 00 00 00 00  H.,$H....H.=....
	0x0050 e8 00 00 00 00 eb e9 e8 00 00 00 00 eb a2        ..............
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 runtime/debug.modinfo+0
	rel 37+4 t=15 runtime/debug.modinfo+8
	rel 44+4 t=15 "".keepalive_modinfo+8
	rel 50+4 t=15 runtime.writeBarrier+-1
	rel 60+4 t=15 "".keepalive_modinfo+0
	rel 76+4 t=15 "".keepalive_modinfo+0
	rel 81+4 t=8 runtime.gcWriteBarrier+0
	rel 88+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=92 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), ABIInternal, $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	85
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	PCDATA	$2, $0
	0x001b 00027 (<autogenerated>:1)	PCDATA	$0, $0
	0x001b 00027 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0022 00034 (<autogenerated>:1)	CMPB	AL, $1
	0x0025 00037 (<autogenerated>:1)	JLS	48
	0x0027 00039 (<autogenerated>:1)	PCDATA	$2, $-2
	0x0027 00039 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0027 00039 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002b 00043 (<autogenerated>:1)	ADDQ	$8, SP
	0x002f 00047 (<autogenerated>:1)	RET
	0x0030 00048 (<autogenerated>:1)	JNE	57
	0x0032 00050 (<autogenerated>:1)	PCDATA	$2, $0
	0x0032 00050 (<autogenerated>:1)	PCDATA	$0, $0
	0x0032 00050 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x00376288921/b001/_gomod_.go:6)	MOVQ	CX, "".keepalive_modinfo+8(SB)
	0x0030 00048 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	PCDATA	$2, $-2
	0x0030 00048 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	PCDATA	$0, $-2
	0x0030 00048 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	CMPL	runtime.writeBarrier(SB), $0
	0x0037 00055 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	JNE	73
	0x0039 00057 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	MOVQ	AX, "".keepalive_modinfo(SB)
	0x0040 00064 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	MOVQ	(SP), BP
	0x0044 00068 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	ADDQ	$8, SP
	0x0048 00072 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	RET
	0x0049 00073 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	LEAQ	"".keepalive_modinfo(SB), DI
	0x0050 00080 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	CALL	runtime.gcWriteBarrier(SB)
	0x0055 00085 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	JMP	64
	0x0057 00087 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	NOP
	0x0057 00087 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	PCDATA	$0, $-1
	0x0057 00087 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	PCDATA	$2, $-1
	0x0057 00087 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	CALL	runtime.morestack_noctxt(SB)
	0x005c 00092 (/var/folders/30/nz41pvvs50z9k63nhx_23dmw0000gn/T/go-build336288921/b001/_gomod_.go:6)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 48 48  eH..%....H;a.vHH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 48 8b 05 00 00  ...H.,$H.,$H....
	0x0020 00 00 48 8b 0d 00 00 00 00 48 89 0d 00 00 00 00  ..H......H......
	0x0030 83 3d 00 00 00 00 00 75 10 48 89 05 00 00 00 00  .=.....u.H......
	0x0040 48 8b 2c 24 48 83 c4 08 c3 48 8d 3d 00 00 00 00  H.,$H....H.=....
	0x0050 e8 00 00 00 00 eb e9 e8 00 00 00 00 eb a2        ..............
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 runtime/debug.modinfo+0
	rel 37+4 t=15 runtime/debug.modinfo+8
	rel 44+4 t=15 "".keepalive_modinfo+8
	rel 50+4 t=15 runtime.writeBarrier+-1
	rel 60+4 t=15 "".keepalive_modinfo+0
	rel 76+4 t=15 "".keepalive_modinfo+0
	rel 81+4 t=8 runtime.gcWriteBarrier+0
	rel 88+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=92 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), ABIInternal, $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	85
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	PCDATA	$2, $0
	0x001b 00027 (<autogenerated>:1)	PCDATA	$0, $0
	0x001b 00027 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0022 00034 (<autogenerated>:1)	CMPB	AL, $1
	0x0025 00037 (<autogenerated>:1)	JLS	48
	0x0027 00039 (<autogenerated>:1)	PCDATA	$2, $-2
	0x0027 00039 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0027 00039 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002b 00043 (<autogenerated>:1)	ADDQ	$8, SP
	0x002f 00047 (<autogenerated>:1)	RET
	0x0030 00048 (<autogenerated>:1)	JNE	57
	0x0032 00050 (<autogenerated>:1)	PCDATA	$2, $0
	0x0032 00050 (<autogenerated>:1)	PCDATA	$0, $0
	0x0032 00050 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x0037 00055 (<autogenerated>:1)	UNDEF
	0x0039 00057 (<autogenerated>:1)	MOVB	$1, "".initdone·(SB)
	0x0040 00064 (<autogenerated>:1)	CALL	"".init.0(SB)
	0x0045 00069 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x004c 00076 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0050 00080 (<autogenerated>:1)	ADDQ	$8, SP
	0x0054 00084 (<autogenerated>:1)	RET
	0x0055 00085 (<autogenerated>:1)	NOP
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0055 00085 (<autogenerated>:1)	PCDATA	$2, $-1
	0x0055 00085 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x005a 00090 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 46 48  eH..%....H;a.vFH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 80 f8 01 76 09 48 8b 2c 24 48 83 c4 08 c3  .....v.H.,$H....
	0x0030 75 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 01  u...............
	0x0040 e8 00 00 00 00 c6 05 00 00 00 00 02 48 8b 2c 24  ............H.,$
	0x0050 48 83 c4 08 c3 e8 00 00 00 00 eb a4              H...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 51+4 t=8 runtime.throwinit+0
	rel 59+4 t=15 "".initdone·+-1
	rel 65+4 t=8 "".init.0+0
	rel 71+4 t=15 "".initdone·+-1
	rel 86+4 t=8 runtime.morestack_noctxt+0
go.info."".add$abstract SDWARFINFO dupok size=29
	0x0000 04 6d 61 69 6e 2e 61 64 64 00 01 01 11 61 00 00  .main.add....a..
	0x0010 00 00 00 00 11 62 00 00 00 00 00 00 00           .....b.......
	rel 16+4 t=28 go.info.int+0
	rel 24+4 t=28 go.info.int+0
go.loc."".add SDWARFLOC size=103
	0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 13 00 00 00 00 00 00 00  ................
	0x0020 01 00 9c 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 ff ff ff ff ff ff ff ff 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 13 00 00 00 00  ................
	0x0050 00 00 00 02 00 91 08 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00                             .......
	rel 8+8 t=1 "".add+0
	rel 59+8 t=1 "".add+0
go.info."".add SDWARFINFO size=54
	0x0000 05 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 01 9c 13 00 00 00 00 00 00 00 00  ................
	0x0020 13 00 00 00 00 00 00 00 00 0f 7e 72 32 00 01 03  ..........~r2...
	0x0030 00 00 00 00 00 00                                ......
	rel 1+4 t=28 go.info."".add$abstract+0
	rel 5+8 t=1 "".add+0
	rel 13+8 t=1 "".add+19
	rel 24+4 t=28 go.info."".add$abstract+12
	rel 28+4 t=28 go.loc."".add+0
	rel 33+4 t=28 go.info."".add$abstract+20
	rel 37+4 t=28 go.loc."".add+51
	rel 48+4 t=28 go.info.int+0
go.range."".add SDWARFRANGE size=0
go.isstmt."".add SDWARFMISC size=0
	0x0000 04 05 01 0e 00                                   .....
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=33
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+2
	rel 27+4 t=29 gofile../Users/ahui/www/go-lib/go-debug/assembly/file.go+0
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
	0x0000 02 02 00                                         ...
go.loc."".init.0 SDWARFLOC size=0
go.info."".init.0 SDWARFINFO size=35
	0x0000 03 22 22 2e 69 6e 69 74 2e 30 00 00 00 00 00 00  ."".init.0......
	0x0010 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00  ................
	0x0020 00 01 00                                         ...
	rel 11+8 t=1 "".init.0+0
	rel 19+8 t=1 "".init.0+94
	rel 29+4 t=29 gofile.._gomod_.go+0
go.range."".init.0 SDWARFRANGE size=0
go.isstmt."".init.0 SDWARFMISC size=0
	0x0000 04 0f 04 0c 03 07 01 1e 02 09 01 0e 02 07 00     ...............
go.string."0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nmod\tgithub.com/ahuigo/repo\t(devel)\t\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2" SRODATA dupok size=96
	0x0000 30 77 af 0c 92 74 08 02 41 e1 c1 07 e6 d6 18 e6  0 00055 (<autogenerated>:1)	UNDEF
	0x0039 00057 (<autogenerated>:1)	MOVB	$1, "".initdone·(SB)
	0x0040 00064 (<autogenerated>:1)	CALL	"".init.0(SB)
	0x0045 00069 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x004c 00076 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0050 00080 (<autogenerated>:1)	ADDQ	$8, SP
	0x0054 00084 (<autogenerated>:1)	RET
	0x0055 00085 (<autogenerated>:1)	NOP
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0055 00085 (<autogenerated>:1)	PCDATA	$2, $-1
	0x0055 00085 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x005a 00090 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 46 48  eH..%....H;a.vFH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 80 f8 01 76 09 48 8b 2c 24 48 83 c4 08 c3  .....v.H.,$H....
	0x0030 75 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 01  u...............
	0x0040 e8 00 00 00 00 c6 05 00 00 00 00 02 48 8b 2c 24  ............H.,$
	0x0050 48 83 c4 08 c3 e8 00 00 00 00 eb a4              H...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 51+4 t=8 runtime.throwinit+0
	rel 59+4 t=15 "".initdone·+-1
	rel 65+4 t=8 "".init.0+0
	rel 71+4 t=15 "".initdone·+-1
	rel 86+4 t=8 runtime.morestack_noctxt+0
go.info."".add$abstract SDWARFINFO dupok size=29
	0x0000 04 6d 61 69 6e 2e 61 64 64 00 01 01 11 61 00 00  .main.add....a..
	0x0010 00 00 00 00 11 62 00 00 00 00 00 00 00           .....b.......
	rel 16+4 t=28 go.info.int+0
	rel 24+4 t=28 go.info.int+0
go.loc."".add SDWARFLOC size=103
	0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 13 00 00 00 00 00 00 00  ................
	0x0020 01 00 9c 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 ff ff ff ff ff ff ff ff 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 13 00 00 00 00  ................
	0x0050 00 00 00 02 00 91 08 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00                             .......
	rel 8+8 t=1 "".add+0
	rel 59+8 t=1 "".add+0
go.info."".add SDWARFINFO size=54
	0x0000 05 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 01 9c 13 00 00 00 00 00 00 00 00  ................
	0x0020 13 00 00 00 00 00 00 00 00 0f 7e 72 32 00 01 03  ..........~r2...
	0x0030 00 00 00 00 00 00                                ......
	rel 1+4 t=28 go.info."".add$abstract+0
	rel 5+8 t=1 "".add+0
	rel 13+8 t=1 "".add+19
	rel 24+4 t=28 go.info."".add$abstract+12
	rel 28+4 t=28 go.loc."".add+0
	rel 33+4 t=28 go.info."".add$abstract+20
	rel 37+4 t=28 go.loc."".add+51
	rel 48+4 t=28 go.info.int+0
go.range."".add SDWARFRANGE size=0
go.isstmt."".add SDWARFMISC size=0
	0x0000 04 05 01 0e 00                                   .....
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=33
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+2
	rel 27+4 t=29 gofile../Users/ahui/www/go-lib/go-debug/assembly/file.go+0
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
	0x0000 02 02 00                                         ...
go.loc."".init.0 SDWARFLOC size=0
go.info."".init.0 SDWARFINFO size=35
	0x0000 03 22 22 2e 69 6e 69 74 2e 30 00 00 00 00 00 00  ."".init.0......
	0x0010 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00  ................
	0x0020 00 01 00                                         ...
	rel 11+8 t=1 "".init.0+0
	rel 19+8 t=1 "".init.0+94
	rel 29+4 t=29 gofile.._gomod_.go+0
go.range."".init.0 SDWARFRANGE size=0
go.isstmt."".init.0 SDWARFMISC size=0
	0x0000 04 0f 04 0c 03 07 01 1e 02 09 01 0e 02 07 00     ...............
go.string."0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nmod\tgithub.com/ahuigo/repo\t(devel)\t\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2" SRODATA dupok size=96
	0x0000 30 77 af 0c 92 74 08 02 41 e1 c1 07 e6 d6 18 e6  0w...t..A.......
	0x0010 70 61 74 68 09 63 6f 6d 6d 61 6e 64 2d 6c 69 6e  path.command-lin
	0x0020 65 2d 61 72 67 75 6d 65 6e 74 73 0a 6d 6f 64 09  e-arguments.mod.
	0x0030 67 69 74 68 75 62 2e 63 6f 6d 2f 61 68 75 69 67  github.com/ahuig
	0x0040 6f 2f 72 65 70 6f 09 28 64 65 76 65 6c 29 09 0a  o/repo.(devel)..
	0x0050 f9 32 43 31 86 18 20 72 00 82 42 10 41 16 d8 f2  .2C1.. r..B.A...
go.loc."".init SDWARFLOC size=0
go.info."".init SDWARFINFO size=33
	0x0000 03 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+92
	rel 27+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
go.isstmt."".init SDWARFMISC size=0
	0x0000 04 0f 04 0c 03 07 01 05 02 09 01 02 02 13 01 10  ................
	0x0010 02 07 00                                         ...
runtime/debug.modinfo SDATA size=16
	0x0000 00 00 00 00 00 00 00 00 60 00 00 00 00 00 00 00  ........`.......
	rel 0+8 t=1 go.string."0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nmod\tgithub.com/ahuigo/repo\t(devel)\t\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"+0
"".keepalive_modinfo SDATA size=16
	0x0000 00 00 00 00 00 00 00 00 60 00 00 00 00 00 00 00  ........`.......
	rel 0+8 t=1 go.string."0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nmod\tgithub.com/ahuigo/repo\t(devel)\t\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"+0
"".initdone· SNOPTRBSS size=1
type..importpath.unsafe. SRODATA dupok size=9
	0x0000 00 00 06 75 6e 73 61 66 65                       ...unsafe
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
w...t..A.......
	0x0010 70 61 74 68 09 63 6f 6d 6d 61 6e 64 2d 6c 69 6e  path.command-lin
	0x0020 65 2d 61 72 67 75 6d 65 6e 74 73 0a 6d 6f 64 09  e-arguments.mod.
	0x0030 67 69 74 68 75 62 2e 63 6f 6d 2f 61 68 75 69 67  github.com/ahuig
	0x0040 6f 2f 72 65 70 6f 09 28 64 65 76 65 6c 29 09 0a  o/repo.(devel)..
	0x0050 f9 32 43 31 86 18 20 72 00 82 42 10 41 16 d8 f2  .2C1.. r..B.A...
go.loc."".init SDWARFLOC size=0
go.info."".init SDWARFINFO size=33
	0x0000 03 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+92
	rel 27+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
go.isstmt."".init SDWARFMISC size=0
	0x0000 04 0f 04 0c 03 07 01 05 02 09 01 02 02 13 01 10  ................
	0x0010 02 07 00                                         ...
runtime/debug.modinfo SDATA size=16
	0x0000 00 00 00 00 00 00 00 00 60 00 00 00 00 00 00 00  ........`.......
	rel 0+8 t=1 go.string."0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nmod\tgithub.com/ahuigo/repo\t(devel)\t\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"+0
"".keepalive_modinfo SDATA size=16
	0x0000 00 00 00 00 00 00 00 00 60 00 00 00 00 00 00 00  ........`.......
	rel 0+8 t=1 go.string."0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nmod\tgithub.com/ahuigo/repo\t(devel)\t\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"+0
"".initdone· SNOPTRBSS size=1
type..importpath.unsafe. SRODATA dupok size=9
	0x0000 00 00 06 75 6e 73 61 66 65                       ...unsafe
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
