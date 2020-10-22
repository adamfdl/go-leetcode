package main

// TAG: Stack

func main() {

}

type MinStack struct {
	items []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		items: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.items = append(this.items, x)
}

func (this *MinStack) Pop() {
	toRemoveIndex := len(this.items) - 1
	this.items = this.items[:toRemoveIndex]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	min := this.items[0]
	for _, val := range this.items {
		if val < min {
			min = val
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
