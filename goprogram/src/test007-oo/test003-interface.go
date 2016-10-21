package main

import "fmt"

type USB interface {
	Name() (string, string)
	Connect()
}

type Computer struct {
	Brand string
}

func (m Computer) Name() (string, string) {
	return "zhangsan's mac ", m.Brand
}
func (m Computer) Connect() {
	fmt.Println(m, " ....Connect....")
}
func main() {
	var usdDevice USB
	usdDevice = Computer{Brand:"apple inc"}

	usdDevice.Connect()

	n,b:=usdDevice.Name()
	fmt.Println(n,b)


}
