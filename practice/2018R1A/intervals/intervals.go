package intervals

import "sort"

type Intervals struct {
	a []Interval
}

type Interval struct {
	L, H float64
}

func (i *Intervals) Add(l, h float64) {
	i.a = append(i.a, Interval{l, h})

	i.SortAndMerge()
}

func (i *Intervals) SortAndMerge() {
	sort.Slice(i.a, func(a, b int) bool {
		return i.a[a].L < i.a[b].L
	})

	for k := 0; k < len(i.a)-1; k++ {
		if i.a[k].H >= i.a[k+1].L {
			i.a[k].H = i.a[k+1].H
			i.a = append(i.a[:k+1], i.a[k+2:]...)
			k--
		}
	}
}
