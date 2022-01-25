package _846_hand_of_straights

func isNStraightHand(hand []int, W int) bool {
	if len(hand) % W != 0 {
		return false
	}

	count := make(map[int]int, 0) // 记录牌-数量map
	first := 1 << 31 - 1 //第一张牌：初始化 MAX_VALUE
	for _, card := range hand {
		// 计数
		// if count[card] == 0 {
		//     count[card] = 1
		// }else {
		//     count[card] = count[card] + 1
		// }
		// 等价于下面一行
		count[card] += 1
		// 更新第一张牌
		if card < first {
			first = card
		}
	}
	//for i := 0; i < len(hand); i++ {
	//	// 计数
	//	count[hand[i]] += 1
	//	// 更新第一张牌
	//	if hand[i] < first {
	//		first = hand[i]
	//	}
	//}

	for len(count) > 0 {
		for count[first] == 0 {
			first += 1
		}
		for card := first; card < first + W; card++ {
			if count[card] == 0 {
				return false
			}
			if count[card] == 1 {
				delete(count, card)
			}else {
				count[card] -= 1
			}
		}
	}
	return true
}
