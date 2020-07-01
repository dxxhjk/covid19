package main

import "github.com/dxxhjk/covid19/frontend"

func main() {
	fe := frontend.FrontEnd{}
	fe.StartListening()
}
