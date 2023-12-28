package main

type RecentCounter struct {
	count []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.count = append(this.count, t)
	for this.count[0] < t-3000 {
		this.count = this.count[1:]
	}
	return len(this.count)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
