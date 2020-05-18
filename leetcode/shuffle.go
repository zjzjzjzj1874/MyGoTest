package leetcode

import (
	"fmt"
	"math/rand"
	"time"
)

//var ALL_CARDS = []byte{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
//var ALL_CARDS = []byte{0,1,2,3,4,5,6}
var ALL_CARDS = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, //红桃23456789 10 JQKA 共13张牌
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, //黑桃
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, //梅花
	0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, //方块
	0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4A, 0x4B, 0x4C, 0x4D, //方块
	0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5A, 0x5B, 0x5C, 0x5D, //方块
	0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, //方块
	0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7A, 0x7B, 0x7C, 0x7D, //方块
	0x8E, 0x8F, //小王
	0x9E, 0x9F}

func shuffle1() []byte {
	//结果数据
	var result_cards []byte
	result_cards = make([]byte, len(ALL_CARDS))
	//存储临时牌的变量
	tmp_cards := make([]byte, len(ALL_CARDS))
	copy(tmp_cards, ALL_CARDS)
	//记录牌的序号
	iCardIndex := 0
	//记录牌的位置
	iCardPosition := 0
	//生成随机数种子
	r_node := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < len(ALL_CARDS); i++ {
		iCardPosition = r_node.Intn(len(tmp_cards))
		result_cards[iCardIndex] = tmp_cards[iCardPosition]
		iCardIndex++
		tmp_cards = append(tmp_cards[:iCardPosition], tmp_cards[iCardPosition+1:]...)
	}
	return result_cards
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func shuffle2() []byte {
	//begin := time.Now()

	tmp_cards := make([]byte, len(ALL_CARDS))
	copy(tmp_cards, ALL_CARDS)
	for i := len(ALL_CARDS); i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)
		tmp_cards[lastIdx], tmp_cards[idx] = tmp_cards[idx], tmp_cards[lastIdx]
	}

	//fmt.Println(time.Now().Sub(begin))
	return tmp_cards
}

func transfer() {
	totalWeight := 3000

	for totalDis := 1000; totalDis > 0; totalDis-- {
		switch {
		case totalWeight >= 2000:
			totalWeight -= 5
		case totalWeight >= 1000:
			totalWeight -= 3
		default:
			totalWeight--
		}
	}

	fmt.Println("剩余 == ", totalWeight)

}
