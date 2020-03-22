package main

type MaxQueue struct {
	A []int
	B []int
}

func Constructor() MaxQueue {
	return MaxQueue{[]int{}, []int{}}
}

func (this *MaxQueue) Max_value() int {
	if len(this.B) != 0 {
		return this.B[0]
	}
	return -1
}

func (this *MaxQueue) Push_back(value int) {
	this.A = append(this.A, value)
	//  找到位置
	for len(this.B) != 0 && value > this.B[len(this.B)-1] {
		this.B = this.B[:len(this.B)-1]
	}
	this.B = append(this.B, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.A) > 0 {
		val := this.A[0]
		this.A = this.A[1:]
		if len(this.B) > 0 && val == this.B[0] {
			this.B = this.B[1:]
		}
		return val
	}

	return -1
}

func main() {

}
