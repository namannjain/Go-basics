package main

import "fmt"

type Item struct {
	name         string
	inStock      bool
	observerList []Observer
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) updateAvailability() {
	fmt.Println("Ïtem is available now:  ", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *Item) register(o Observer){
  i.observerList = append(i.observerList, o)
}

func (i *Item) deRegister(o Observer){
  i.observerList = removeFromSlice(i.observerList, o)
}

func removeFromSlice(o []Observer, obser1 Observer) []Observer{
  n := len(o)
  for i,observer := o{
    if observer.getId() == obser1.getId(){
      o[n-1], o[i] = o[i], o[n-1]
      return o[:n-1]
    }
  }
  return o
}

func (i *Item) notifyAll(){
  for _, o := range i.observerList{
    o.update(i.name)
  }
}