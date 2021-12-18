package main

import "testing"

func mustParse(t *testing.T, line string) *snailfish {
	s, err := parse(line)
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func TestExplode(t *testing.T) {
	testcases := []struct {
		in          string
		want        string
		wantChanged bool
	}{
		{
			in:          "[[[[[9,8],1],2],3],4]",
			want:        "[[[[0,9],2],3],4]",
			wantChanged: true,
		},
		{
			in:          "[7,[6,[5,[4,[3,2]]]]]",
			want:        "[7,[6,[5,[7,0]]]]",
			wantChanged: true,
		},
		{
			in:          "[[6,[5,[4,[3,2]]]],1]",
			want:        "[[6,[5,[7,0]]],3]",
			wantChanged: true,
		},
		{
			in:          "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			want:        "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			wantChanged: true,
		},
		{
			in:          "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			want:        "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
			wantChanged: true,
		},
		{
			in:          "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
			want:        "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			wantChanged: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.in, func(t *testing.T) {
			lf := mustParse(t, tc.in)
			got, changed := explode(lf, 0, true)
			if got.String() != tc.want {
				t.Errorf("explode got %v, want %v", got, tc.want)
			}
			if changed != tc.wantChanged {
				t.Errorf("explode got changed %v, want %v", changed, tc.wantChanged)
			}
		})
	}
}

func checkEq(t *testing.T, s *snailfish, val string) *snailfish {
	if s.String() != val {
		t.Fatalf("test val %v didn't match expected val %q", s, val)
	}
	return s
}

func TestAddToLeft(t *testing.T) {
	testcases := []struct {
		fish string
		node func(s *snailfish) *snailfish
		val  int
		want string
	}{
		{
			fish: "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			node: func(s *snailfish) *snailfish {
				return checkEq(t, s.second.first.first, "[3,7]")
			},
			val:  1,
			want: "[8,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
		},
	}
	for _, tc := range testcases {
		fish := mustParse(t, tc.fish)
		tc.node(fish).addToLeft(tc.val)
		if got := fish.String(); got != tc.want {
			t.Errorf("addToLeft got %v, want %v", got, tc.want)
		}
	}
}

func TestAdd(t *testing.T) {
	testcases := []struct {
		desc string
		a, b string
		want string
	}{
		{
			desc: "sample",
			a:    "[[[[4,3],4],4],[7,[[8,4],9]]]",
			b:    "[1,1]",
			want: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
		{
			desc: "cascade1",
			a:    "[1,1]",
			b:    "[2,2]",
			want: "[[1,1],[2,2]]",
		},
		{
			desc: "cascade2",
			a:    "[[[1,1],[2,2]],[3,3]]",
			b:    "[4,4]",
			want: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			desc: "cascade3",
			a:    "[[[[1,1],[2,2]],[3,3]],[4,4]]",
			b:    "[5,5]",
			want: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			desc: "cascade4",
			a:    "[[[[3,0],[5,3]],[4,4]],[5,5]]",
			b:    "[6,6]",
			want: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			"1",
			"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
			"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		},
		{
			"2",
			"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
			"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
			"[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
		},
		{
			"3",
			"[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
			"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
			"[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
		},
		{
			"4",
			"[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
			"[7,[5,[[3,8],[1,4]]]]",
			"[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]",
		},
		{
			"5",
			"[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]",
			"[[2,[2,2]],[8,[8,1]]]",
			"[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
		},
		{
			"6",
			"[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
			"[2,9]",
			"[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
		},
		{
			"7",
			"[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
			"[1,[[[9,3],9],[[9,0],[0,7]]]]",
			"[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]",
		},
		{
			"8",
			"[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]",
			"[[[5,[7,4]],7],1]",
			"[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]",
		},
		{
			"9",
			"[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]",
			"[[[[4,2],2],6],[8,7]]",
			"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			a := mustParse(t, tc.a)
			b := mustParse(t, tc.b)
			if got := a.add(b); got.String() != tc.want {
				t.Errorf("add %q + %q, got\n%v\n    , want\n%v", tc.a, tc.b, got, tc.want)
				a := mustParse(t, tc.a)
				b := mustParse(t, tc.b)
				doDebug = true
				a.add(b)
				doDebug = false
			}
		})
	}
}
