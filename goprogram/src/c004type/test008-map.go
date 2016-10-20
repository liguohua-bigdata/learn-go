package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	//声明一个map
	var personDB map[string]PersonInfo
	fmt.Println(personDB)
	//创建一个map
	personDB = make(map[string]PersonInfo)
	fmt.Println(personDB)
	// 往这个map里插入几条数据
	personDB["p001"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDB["p002"] = PersonInfo{"1", "Jack", "Room 101,..."}
	fmt.Println(personDB)
	//从map中取建伟 p001的元素
	person, ok := personDB["p001"]
	if ok {
		fmt.Println("Found person", person, "with key p001.")
	} else {
		fmt.Println("Did not find person with key p001.")
	}
	//删除map中的元素
	delete(personDB, "p001")
	fmt.Println(personDB)
}
