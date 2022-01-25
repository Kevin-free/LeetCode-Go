package _380_insert_delete_getrandom_o1

import (
	"math/rand"
)

// 380. O(1) 时间插入、删除和获取随机元素：中等
// tags：设计、数组、哈希表、随机化

type RandomizedSet struct {
	// 存储元素的值
	nums []int
	// 记录每个元素对应在 nums 中的索引
	valToIndex map[int]int
}

// 构造方法
func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:       []int{},
		valToIndex: map[int]int{},
	}
}

// 如果 val 不存在集合中，则插入并返回 true，否则直接返回 false
func (this *RandomizedSet) Insert(val int) bool {
	// 若 val 已存在，不用再插入
	if _, ok := this.valToIndex[val]; ok {
		return false
	}
	// 若 val 不存在，插入到 nums 尾部，
	// 并记录 val 对应的索引值
	this.nums = append(this.nums, val)
	this.valToIndex[val] = len(this.nums) - 1
	return true
}

// 如果 val 在集合中，则删除并返回 true，否则直接返回 false
func (this *RandomizedSet) Remove(val int) bool {
	// 若 val 不存在，不用再删除
	if _, ok := this.valToIndex[val]; !ok {
		return false
	}
	// 拿到待删除 val 的索引
	index := this.valToIndex[val]
	// 最后一个元素的索引
	lastIndex := len(this.nums) - 1
	// 将最后一个元素对应的索引修改为 index
	this.valToIndex[this.nums[lastIndex]] = index
	// 交换 val 和最后一个元素 （将 val 索引位置设为最后一个值也可）
	this.nums[index], this.nums[lastIndex] = this.nums[lastIndex], this.nums[index] //this.nums[index] = this.nums[lastIndex]
	// 将数组中最后一个值 val 删除
	this.nums = this.nums[0 : len(this.nums)-1]
	// 删除元素 val 对应的索引
	delete(this.valToIndex, val)
	return true
}

// 从集合中等概率地随机获得一个元素
func (this *RandomizedSet) GetRandom() int {
	// 随机获取 nums 中的一个元素
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
