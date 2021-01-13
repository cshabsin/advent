package tile

import "testing"

func TestReadEdge(t *testing.T) {
	testcases := []struct {
		lines     []string
		wantEdges []int
	}{
		{
			lines: []string{"Tile 2311:",
				"..##.#..#.", //4+8+32+256 = 300, 2+16+64+128=210
				"##..#.....",
				"#...##..#.",
				"####.#...#",
				"##.##.###.",
				"##...#.###",
				".#.#.#..##",
				"..#....#..",
				"###...#.#.",
				"..###..###",
			},
			wantEdges: []int{210, 318, 924, 89},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.lines[0], func(t *testing.T) {
			tile, err := parseLines(tc.lines)
			if err != nil {
				t.Fatal(err)
			}
			for i := 0; i < 4; i++ {
				if got := tile.ReadEdge(i); got != tc.wantEdges[i] {
					t.Errorf("tile.ReadEdge(%d) got %d, want %d", i, got, tc.wantEdges[i])
				}
			}
		})
	}
}
