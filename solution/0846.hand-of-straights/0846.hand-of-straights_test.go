package _846_hand_of_straights

import "testing"

func Test0846(t *testing.T) {
	hand := []int{1,2,3,6,2,3,4,7,8}
	w := 3
	res := isNStraightHand(hand,w)
	println(res)
}
