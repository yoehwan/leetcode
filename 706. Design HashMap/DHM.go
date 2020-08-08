package main

import "fmt"


func main() {
	obj:=Constructor()
	obj.Put(1,2)
	obj.Put(2,2)
	par:=obj.Get(1)
	par_2:=obj.Get(3)
	obj.Put(2,2)

	fmt.Println(par)
	fmt.Println(par_2)
}

type MyHashMap struct {
	key []int
	value []int
}
/** Initialize your data structure here. */
func Constructor() MyHashMap {
   return MyHashMap{}
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	for idx,item:=range this.key{
		if item==key{
			this.value[idx]=value
			return
		}
	}
	this.key=append(this.key,key)
	this.value=append(this.value,value)

}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
 	for idx,item:=range this.key{
 		if item==key{
 			return this.value[idx]
		}
	}
	return -1

}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    for idx,item:=range this.key{
    	if item==key{
    		this.key=append(this.key[:idx],this.key[idx+1:]...)
			this.value=append(this.value[:idx],this.value[idx+1:]...)
		}
	}
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */