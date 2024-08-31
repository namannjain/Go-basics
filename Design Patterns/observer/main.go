package main

func main() {
	//adding new product
	shirtItem := newItem("tommy")

	//normal users
	obser1 := &Customer{id: "abc@gmail.com"}
	obser2 := &Customer{id: "xyz@gmai.com"}

	//register
	shirtItem.register(obser1)
	shirtItem.register(obser2)

	//update product available
	shirtItem.updateAvailability()
}
