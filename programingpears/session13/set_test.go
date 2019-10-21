package pearls

import "testing"

func Test_Set(t *testing.T) {
	maxElements, maxVal := 1000, 10000
	// source :=
	// dst := []int{1, 2, 3, 4, 5, 6}
	suits := []struct {
		Source []int
		Dst    []int
	}{
		{Source: []int{1, 2, 3, 5, 6, 4, 4},
			Dst: []int{1, 2, 3, 4, 5, 6},
		},
	}

	sets := []struct {
		IntSet
		Name string
	}{
		{
			IntSet: NewIntSetArr(maxElements, maxVal),
			Name:   "IntSetArr",
		},
		{
			IntSet: NewIntSetSTL(maxElements, maxVal),
			Name:   "IntSetSTL",
		},
		{
			IntSet: NewIntSetList(maxElements, maxVal),
			Name:   "IntSetList",
		},
	}
	for _, set := range sets {
		for _, suit := range suits {
			for _, element := range suit.Source {
				set.Insert(element)
			}
			dist := set.Report()
			for i := 0; i < len(suit.Dst); i++ {
				if dist[i] != suit.Dst[i] {
					t.Errorf("test faied in source:%v,set type %s", suit.Source, set.Name)
					return
				}
			}
		}
		set.Report()
	}

}
