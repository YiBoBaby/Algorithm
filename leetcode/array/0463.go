package main

func islandPerimeter(grid [][]int) int {
	x := len(grid)
	y := len(grid[0])
	var t int
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if grid[i][j] == 1 {
				// 上
				if j+1 < y {
					if grid[i][j+1] == 0 {
						t++
					}
				} else {
					t++
				}
				// 下
				if j-1 >= 0 {
					if grid[i][j-1] == 0 {
						t++
					}
				} else {
					t++
				}
				// 左
				if i-1 >= 0 {
					if grid[i-1][j] == 0 {
						t++
					}
				} else {
					t++
				}
				// 右
				if i+1 < x {
					if grid[i+1][j] == 0 {
						t++
					}
				} else {
					t++
				}
			}
		}
	}
	return t
}
