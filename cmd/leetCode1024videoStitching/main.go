package main

import (
	"fmt"
	"sort"
)

var (
	clips = [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	time  = 10
)

func videoStitching(clips [][]int, time int) int {
	sort.Slice(clips, func(i int, j int) bool {
		return clips[i][0] < clips[j][0] || (clips[i][0] == clips[j][0] && clips[i][1] > clips[j][1])
	})
	clipCount := 0
	clipTime := 0
	for clipTime < time {

		maxClipTime := 0
		i := 0
		for i < len(clips) && clips[i][0] <= clipTime {
			if clips[i][1]-clipTime > maxClipTime {
				maxClipTime = clips[i][1] - clipTime
			}
			i++
		}
		if maxClipTime == 0 {
			return -1
		} else {
			clipTime += maxClipTime
			clipCount++
		}
	}

	//fmt.Println(clips)
	return clipCount
}

func main() {
	result := videoStitching(clips, time)
	fmt.Println(result)

}
