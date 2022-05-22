package main

import (
	"fmt"
	"math"
	"time"
)

type Su struct {
	self  uint
	next  uint
	cache uint
}

var n uint = 3

var SuList = struct {
	iSu  int
	list []*Su
}{
	0,
	[]*Su{
		&Su{
			3,
			3,
			9,
		},
	},
}

func procSu(e uint) {
	for n < e {
		//获取目前自然数序列增长最小合数
		m := compSu()
		//遍历合数
		iteratorSu(m)
	}
	SuList.list = append(SuList.list, &Su{
		2,
		2,
		4,
	})
}

func compSu() uint {
	var n = uint(math.MaxInt32)
	var idx []int
	var lastSu int
	for i := 0; i <= SuList.iSu; i++ {
		var tmpN uint
		if SuList.list[i].cache > 0 {
			tmpN = SuList.list[i].cache
		} else {
			tmpN = SuList.list[i].self * SuList.list[i].next
			SuList.list[i].cache = tmpN
		}
		if tmpN <= n {
			if tmpN < n {
				idx = append([]int(nil), i)
			} else {
				idx = append(idx, i)
			}
			n = tmpN
		}
	}
	for _, i := range idx {
		SuList.list[i].next++
		SuList.list[i].cache = 0
		lastSu = i
	}
	if SuList.list[lastSu].next-SuList.list[lastSu].self == 1 {
		SuList.iSu++
	}
	return n
}

func appendSu(s uint) {
	SuList.list = append(SuList.list, &Su{
		s,
		s,
		s * s,
	})
}

func updateN(number uint) {
	n = number
}

func iteratorSu(number uint) {
	var i uint = n + 1
	for {
		if i < number {
			if i&1 == 1 {
				appendSu(i)
			}
		} else {
			break
		}
		i++
	}
	updateN(number)
}

func main() {
	s := time.Now()
	procSu(10000000)
	for _, v := range SuList.list {
		fmt.Print(v.self, ",")
	}
	fmt.Println()
	fmt.Println(`cost:`, time.Since(s).String())
}
