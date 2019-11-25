package main

import "C"

func main() {
	//println("a")

	C.puts(C.CString("Hello, World\n"))
}
