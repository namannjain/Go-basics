package main

type Subject interface {
	register()
	deRegister()
	notifyAll()
}
