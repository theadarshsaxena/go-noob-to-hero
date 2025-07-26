package main

import (
	"slices"
)

/*Problem Statement:
A monkey is given ‘n’ piles of bananas, whereas the 'ith' pile has ‘a[i]’ bananas.
An integer ‘h’ is also given, which denotes the time (in hours) for all the bananas to be eaten.

Each hour, the monkey chooses a non-empty pile of bananas and eats ‘k’ bananas.
If the pile contains less than ‘k’ bananas, then the monkey consumes all the bananas
and won’t eat any more bananas in that hour.

Find the minimum number of bananas ‘k’ to eat per hour so that the monkey can eat all the
bananas within ‘h’ hours
*/

func kokoEatingBananas(a []int, h int) int {
	l,r:=1, slices.Max(a)  // TODO: check other function in slices https://pkg.go.dev/slices
	ans:=-1
	for l<=r {
		m:=r-(r-l)/2
		th := findTotalHoursTaken(a,m)
		if th <= h {
			ans = m
			r = m-1
		} else {
			l = m+1
		}
	}
	return ans
}

func findTotalHoursTaken(a []int, num int) int { //num is bananas eaten per hour
	hours := 0
	for _, v := range a {
		hours +=v/num
		if v%num != 0 {
			hours++
		}
	}
	return hours
}
// func main() {
// 	piles := []int{30,11,23,4,20}
// 	h := 5
// 	fmt.Println(kokoEatingBananas(piles, h))
// }

// TODO: Check the following program, it was taking lower time

// func minEatingSpeed(piles []int, h int) int {
//   hours_required := func (speed int) int {
//     total := 0;
//     for _, pile := range piles {
//       total += int(math.Ceil(float64(pile)/ float64(speed)));
//     }
//     return total;
//   }

//   low := 1;
//   high := slices.Max(piles);

//   for low < high {
//     mid := (low + high) / 2;
//     if hours_required(mid) > h {
//       low = mid + 1;
//     } else {
//       high = mid;
//     }
//   }

//   return low;
// }

// TODO: Check this now
// Alternative to  total += int(math.Ceil(float64(pile)/ float64(speed)));
// You can write total += (pile + speed -1)/speed

// TODO: Check this whole example, this is computationally efficient but less readable
// func minEatingSpeed(a []int, h int) int {
//    l,r:=1, slices.Max(a)
//    th := 0
// 	  ans:=-1
//    m := 0
// 	  for l<=r {
// 		m=r-(r-l)/2
// 		th = 0
//         for _, v := range a {
//             th +=(v+m-1)/m
//         }
// 		if th <= h {
// 			ans = m
// 			r = m-1
// 		} else {
// 			l = m+1
// 		}
// 	  }
// 	return ans
// }