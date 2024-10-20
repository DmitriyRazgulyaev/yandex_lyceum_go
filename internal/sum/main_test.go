package sum

import "testing"

//type Test struct {
//	Name     string
//	Value1   int
//	Value2   int
//	Expected int
//}
//
//func TestSum(t *testing.T) {
//	var testes = []Test{
//		{"Test_1", 1, 2, 3},
//		{"Test_2", -1, 2, 1},
//		{"Test_3", 0, 0, 0},
//	}
//
//	for _, tc := range testes {
//		tc := tc
//		t.Run(tc.Name, func(t *testing.T) {
//			got := Sum(tc.Value1, tc.Value2)
//			if got != tc.Expected {
//				t.Errorf("Sum(%v) = %v; expected %v", tc.Value1, tc.Value2, tc.Expected)
//			}
//		})
//	}
//}
//
//func TestMultiply(t *testing.T) {
//	tests := []struct {
//		Name        string
//		FirstValue  int
//		SecondValue int
//		Expect      int
//	}{
//		{
//			Name:        "Positive values",
//			FirstValue:  2,
//			SecondValue: 3,
//			Expect:      6,
//		},
//
//		{
//			Name:        "Zero-value",
//			FirstValue:  0,
//			SecondValue: 5,
//			Expect:      0,
//		},
//
//		{
//			Name:        "Negative value",
//			FirstValue:  5,
//			SecondValue: -4,
//			Expect:      -20,
//		},
//
//		{
//			Name:        "Negative values",
//			FirstValue:  -6,
//			SecondValue: -8,
//			Expect:      48,
//		},
//
//		{
//			Name:        "Negative and Zero values",
//			SecondValue: 0,
//			FirstValue:  -3,
//			Expect:      0,
//		},
//	}
//
//	for _, tc := range tests {
//		tc := tc
//		t.Run(tc.Name, func(t *testing.T) {
//			got := Multipy(tc.FirstValue, tc.SecondValue)
//			if got != tc.Expect {
//				t.Errorf("Sum(%v, %v) = %v; want %v", tc.FirstValue, tc.SecondValue, got, tc.Expect)
//			}
//
//		})
//	}
//}

func TestGetUTFLength(t *testing.T) {
	cases := []struct {
		Name string
		In []byte
		Out int
	}{
		{
			Name: "test-1",
			In: []byte{115, 116, 114},
			Out: 3,
		},

		{
			Name: "test-2",
			In: []byte,
		},
	}
}
