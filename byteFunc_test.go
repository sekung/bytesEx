package bytesEx

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

func TestReversed(t *testing.T) {
	tests := []struct {
		src []byte
		res []byte
	}{
		{[]byte("123"), []byte("321")},
		{[]byte{0x11, 0x22, 0x33}, []byte{0x33, 0x22, 0x11}},
	}
	for _, test := range tests {
		old := test.src
		Reversed(test.src)
		if !reflect.DeepEqual(test.src, test.res) {
			t.Errorf("got %v for input %v; expected %v", test.src, old, test.res)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		src []byte
		res []byte
	}{
		{[]byte("123"), []byte("321")},
		{[]byte{0x11, 0x22, 0x33}, []byte{0x33, 0x22, 0x11}},
	}
	for _, test := range tests {
		actual := Reverse(test.src)
		if !reflect.DeepEqual(actual, test.res) {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		src []byte
		inx int
		val []byte
		res []byte
	}{
		{[]byte("123"), 1, []byte("321"), []byte("132123")},
		{[]byte{0x11, 0x22, 0x33}, -1, []byte{0x33, 0x22, 0x11}, nil},
		{[]byte{0x11, 0x22, 0x33}, 0, []byte{0x33, 0x22, 0x11}, []byte{0x33, 0x22, 0x11, 0x11, 0x22, 0x33}},
		{[]byte{0x11, 0x22, 0x33}, 1, []byte{0x33, 0x22, 0x11}, []byte{0x11, 0x33, 0x22, 0x11, 0x22, 0x33}},
		{[]byte{0x11, 0x22, 0x33}, 2, []byte{0x33, 0x22, 0x11}, []byte{0x11, 0x22, 0x33, 0x22, 0x11, 0x33}},
		{[]byte{0x11, 0x22, 0x33}, 3, []byte{0x33, 0x22, 0x11}, []byte{0x11, 0x22, 0x33, 0x33, 0x22, 0x11}},
		{[]byte{0x11, 0x22, 0x33}, 4, []byte{0x33, 0x22, 0x11}, nil},
	}
	for _, test := range tests {
		actual, _ := Insert(test.src, test.inx, test.val)
		if !reflect.DeepEqual(actual, test.res) {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestPopInx(t *testing.T) {
	tests := []struct {
		src []byte
		inx int
		res []byte
	}{
		{[]byte("123"), 1, []byte("13")},
		{[]byte{0x11, 0x22, 0x33}, -1, nil},
		{[]byte{0x11, 0x22, 0x33}, 0, []byte{0x22, 0x33}},
		{[]byte{0x11, 0x22, 0x33}, 1, []byte{0x11, 0x33}},
		{[]byte{0x11, 0x22, 0x33}, 2, []byte{0x11, 0x22}},
		{[]byte{0x11, 0x22, 0x33}, 3, nil},
	}
	for _, test := range tests {
		actual, _ := PopInx(test.src, test.inx)
		if !reflect.DeepEqual(actual, test.res) {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		src []byte
		res []byte
	}{
		{[]byte("123"), []byte("23")},
		{[]byte{0x11, 0x22, 0x33, 0x44}, []byte{0x22, 0x33, 0x44}},
		{[]byte{0x22, 0x33, 0x44}, []byte{0x33, 0x44}},
		{[]byte{0x33, 0x44}, []byte{0x44}},
		{[]byte{0x44}, []byte{}},
		{[]byte{}, nil},
	}
	for _, test := range tests {
		actual, _ := Pop(test.src)
		if !reflect.DeepEqual(actual, test.res) {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestDel(t *testing.T) {
	tests := []struct {
		src   []byte
		start int
		end   int
		res   []byte
	}{
		{[]byte("123"), 1, 3, []byte("1")},
		{[]byte{0x11, 0x22, 0x33, 0x44, 0x55}, -1, 3, nil},
		{[]byte{0x11, 0x22, 0x33, 0x44, 0x55}, 0, 2, []byte{0x33, 0x44, 0x55}},
		{[]byte{0x11, 0x22, 0x33, 0x44, 0x55}, 2, 2, []byte{0x11, 0x22, 0x33, 0x44, 0x55}},
		{[]byte{0x11, 0x22, 0x33, 0x44, 0x55}, 3, 2, nil},
		{[]byte{0x11, 0x22, 0x33, 0x44, 0x55}, 3, 5, []byte{0x11, 0x22, 0x33}},
		{[]byte{0x11, 0x22, 0x33, 0x44, 0x55}, 3, 6, nil},
	}
	for _, test := range tests {
		actual, _ := Del(test.src, test.start, test.end)
		if !reflect.DeepEqual(actual, test.res) {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestCombine(t *testing.T) {
	tests := []struct {
		src []byte
		val []byte
		res []byte
	}{
		{[]byte("123"), []byte("321"), []byte("123321")},
		{nil, []byte{0x33, 0x22, 0x11}, []byte{0x33, 0x22, 0x11}},
		{[]byte{}, []byte{0x33, 0x22, 0x11}, []byte{0x33, 0x22, 0x11}},
		{[]byte{0x11, 0x22, 0x33}, []byte{0x33, 0x22, 0x11}, []byte{0x11, 0x22, 0x33, 0x33, 0x22, 0x11}},
		{[]byte{0x11, 0x22, 0x33}, nil, []byte{0x11, 0x22, 0x33}},
		{[]byte{0x11, 0x22, 0x33}, []byte{}, []byte{0x11, 0x22, 0x33}},
		{[]byte{0x11, 0x22, 0x33}, []byte{}, []byte{0x11, 0x22, 0x33}},
	}
	for _, test := range tests {
		actual := Combine(test.src, test.val)
		if !reflect.DeepEqual(actual, test.res) {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestDec(t *testing.T) {
	tests := []struct {
		src []byte
		res int
	}{
		{[]byte{1}, 1},
		{[]byte{1, 2}, 258},
		{[]byte{1, 2, 3}, 66051},
	}
	for _, test := range tests {
		actual := Dec(test.src)
		if actual != test.res {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestHex(t *testing.T) {
	tests := []struct {
		src []byte
		res string
	}{
		{[]byte{1}, "01"},
		{[]byte{1, 2}, "0102"},
		{[]byte{1, 2, 3}, "010203"},
		{[]byte{1, 2, 3, 255}, "010203ff"},
	}
	for _, test := range tests {
		actual := Hex(test.src)
		if actual != test.res {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		src []byte
		res int
	}{
		{[]byte{1}, 1},
		{[]byte{1, 2}, 3},
		{[]byte{1, 2, 3}, 6},
		{[]byte{1, 2, 3, 255}, 261},
	}
	for _, test := range tests {
		actual := Sum(test.src)
		if actual != test.res {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestSum8(t *testing.T) {
	tests := []struct {
		src []byte
		res byte
	}{
		{[]byte{1}, 1},
		{[]byte{1, 2, 3}, 6},
		{[]byte{1, 2, 3, 255}, 5},
	}
	for _, test := range tests {
		actual := Sum8(test.src)
		if actual != test.res {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestCheckSum8(t *testing.T) {
	tests := []struct {
		src []byte
		val byte
		res bool
	}{
		{[]byte{1}, 1, true},
		{[]byte{1, 2, 3}, 6, true},
		{[]byte{1, 2, 3, 255}, 6, false},
	}
	for _, test := range tests {
		actual := CheckSum8(test.src, test.val)
		if actual != test.res {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestCheckSum8Merge(t *testing.T) {
	tests := []struct {
		src []byte
		res []byte
	}{
		{[]byte{1}, []byte{1, 1}},
		{[]byte{1, 2, 3}, []byte{1, 2, 3, 6}},
		{[]byte{1, 2, 3, 255}, []byte{1, 2, 3, 255, 5}},
	}
	for _, test := range tests {
		actual := CheckSum8Merge(test.src)
		if !reflect.DeepEqual(actual, test.res) {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestCheckSum16Be(t *testing.T) {
	tests := []struct {
		src []byte
		val []byte
		res bool
	}{
		{[]byte{1}, []byte{0, 1}, true},
		{[]byte{1, 2, 3}, []byte{6}, false},
		{[]byte{1, 2, 3}, []byte{0, 6}, true},
	}
	for _, test := range tests {
		actual := CheckSum16Be(test.src, test.val)
		if actual != test.res {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestCheckSum16BeMerge(t *testing.T) {
	tests := []struct {
		src []byte
		res []byte
	}{
		{[]byte{1}, []byte{1, 0, 1}},
		{[]byte{1, 2, 3}, []byte{1, 2, 3, 0, 6}},
	}
	for _, test := range tests {
		actual := CheckSum16BeMerge(test.src)
		if !reflect.DeepEqual(actual, test.res) {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestCheckSum16Le(t *testing.T) {
	tests := []struct {
		src []byte
		val []byte
		res bool
	}{
		{[]byte{1}, []byte{1, 0}, true},
		{[]byte{1, 2, 3}, []byte{6}, false},
		{[]byte{1, 2, 3}, []byte{6, 0}, true},
	}
	for _, test := range tests {
		actual := CheckSum16Le(test.src, test.val)
		if actual != test.res {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestCheckSum16LeMerge(t *testing.T) {
	tests := []struct {
		src []byte
		res []byte
	}{
		{[]byte{1}, []byte{1, 1, 0}},
		{[]byte{1, 2, 3}, []byte{1, 2, 3, 6, 0}},
	}
	for _, test := range tests {
		actual := CheckSum16LeMerge(test.src)
		if !reflect.DeepEqual(actual, test.res) {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestCheckCRCModbus(t *testing.T) {
	tests := []struct {
		src []byte
		val []byte
		res bool
	}{
		{[]byte{1, 3, 0, 0, 0, 1}, []byte{0x84, 0x0a}, true},
	}
	for _, test := range tests {
		actual := CheckCRCModbus(test.src, test.val)
		if actual != test.res {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestCheckCRCModbusMerge(t *testing.T) {
	tests := []struct {
		src []byte
		res []byte
	}{
		{[]byte{1, 3, 0, 0, 0, 1}, []byte{1, 3, 0, 0, 0, 1, 0x84, 0x0a}},
	}
	for _, test := range tests {
		actual := CheckCRCModbusMerge(test.src)
		if !reflect.DeepEqual(actual, test.res) {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestCheckCRCXmodem(t *testing.T) {
	tests := []struct {
		src []byte
		val []byte
		res bool
	}{
		{[]byte{1, 3, 0, 0, 0, 1}, []byte{0xBB, 0x53}, true},
	}
	for _, test := range tests {
		actual := CheckCRCXmodem(test.src, test.val)
		if actual != test.res {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestCheckCRCXmodemMerge(t *testing.T) {
	tests := []struct {
		src []byte
		res []byte
	}{
		{[]byte{1, 3, 0, 0, 0, 1}, []byte{1, 3, 0, 0, 0, 1, 0xBB, 0x53}},
	}
	for _, test := range tests {
		actual := CheckCRCXmodemMerge(test.src)
		if !reflect.DeepEqual(actual, test.res) {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestCheckBCC(t *testing.T) {
	tests := []struct {
		src []byte
		val byte
		res bool
	}{
		{[]byte{1, 3, 0, 0, 0, 1}, 3, true},
	}
	for _, test := range tests {
		actual := CheckBCC(test.src, test.val)
		if actual != test.res {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestCheckBCCMerge(t *testing.T) {
	tests := []struct {
		src []byte
		res []byte
	}{
		{[]byte{1, 3, 0, 0, 0, 1}, []byte{1, 3, 0, 0, 0, 1, 3}},
	}
	for _, test := range tests {
		actual := CheckBCCMerge(test.src)
		if !reflect.DeepEqual(actual, test.res) {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestBytes32ToFloatBe(t *testing.T) {
	tests := []struct {
		src []byte
		res float32
	}{
		{[]byte{0x42, 0xC8, 0x33, 0x33}, 100.1},
	}
	for _, test := range tests {
		actual := Bytes32ToFloatBe(test.src)
		if actual != test.res {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestBytes32ToFloatLe(t *testing.T) {
	tests := []struct {
		src []byte
		res float32
	}{
		{[]byte{0x33, 0x33, 0xC8, 0x42}, 100.1},
	}
	for _, test := range tests {
		actual := Bytes32ToFloatLe(test.src)
		if actual != test.res {
			t.Errorf("got %v for input %v; expected %v", actual, test.src, test.res)
		}
	}
}

func TestDeBuff(t *testing.T) {
	src, _ := hex.DecodeString("1b660014323032322f30362f33301b66000231373a34323a31380d1c2630352d322d313339b9e2b5e7b8d0d1cca1fab6aacaa7a1a1a1a10d303036b2e334bac5c2a5a1a1a1a1a1a1a1a1a1a11c2e0d1b660014323032322f30362f33301b66000231373a34323a31380d1c2630352d322d313430cec2ccbdb2e2c6f7a1fab6aacaa7a1a1a1a10d303036b2e334bac5c2a5a1a1a1a1a1a1a1a1a1a11c2e0d1b660014323032322f30362f33301b66000231373a34323a31390d1c2630352d322d313431cec2ccbdb2e2c6f7a1fab6aacaa7a1a1a1a10d303036b2e334bac5c2a5a1a1a1a1a1a1a1a1a1a11c2e0d")
	slice1, _ := hex.DecodeString("1b660014323032322f30362f33301b66000231373a34323a31380d1c2630352d322d313339b9e2b5e7b8d0d1cca1fab6aacaa7a1a1a1a10d303036b2e334bac5c2a5a1a1a1a1a1a1a1a1a1a11c2e0d")
	slice2, _ := hex.DecodeString("1b660014323032322f30362f33301b66000231373a34323a31380d1c2630352d322d313430cec2ccbdb2e2c6f7a1fab6aacaa7a1a1a1a10d303036b2e334bac5c2a5a1a1a1a1a1a1a1a1a1a11c2e0d")
	slice3, _ := hex.DecodeString("1b660014323032322f30362f33301b66000231373a34323a31390d1c2630352d322d313431cec2ccbdb2e2c6f7a1fab6aacaa7a1a1a1a10d303036b2e334bac5c2a5a1a1a1a1a1a1a1a1a1a11c2e0d")
	slice4, _ := hex.DecodeString("323032322f30362f33301b66000231373a34323a31390d1c2630352d322d313431cec2ccbdb2e2c6f7a1fab6aacaa7a1a1a1a10d303036b2e334bac5c2a5a1a1a1a1a1a1a1a1a1a11c2e0d1b66001432")
	actual := DeBuff(src, []byte{0x1b, 0x66, 0x00, 0x14}, []byte{0x1c, 0x2e, 0x0d})
	if !reflect.DeepEqual(actual, [][]byte{slice1, slice2, slice3}) {
		t.Errorf("got %v for input %v; expected %v", actual, src, [][]byte{slice1, slice2, slice3})
	}
	actual = DeBuff(slice1, []byte{}, []byte{0x1c, 0x2e, 0x0d})
	if !reflect.DeepEqual(actual, [][]byte{slice1}) {
		t.Errorf("got %v for input %v; expected %v", actual, slice1, [][]byte{slice1})
	}
	actual = DeBuff(slice1, []byte{0x1b, 0x66, 0x00, 0x14}, []byte{})
	if !reflect.DeepEqual(actual, [][]byte{slice1}) {
		t.Errorf("got %v for input %v; expected %v", actual, slice1, [][]byte{slice1})
	}
	actual = DeBuff(slice1, []byte{}, []byte{})
	if !reflect.DeepEqual(actual, [][]byte{slice1}) {
		t.Errorf("got %v for input %v; expected %v", actual, slice1, [][]byte{slice1})
	}
	test1 := []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x11, 0x22, 0x11}
	actual = DeBuff(test1, []byte{0x55}, []byte{})
	if !reflect.DeepEqual(actual, [][]byte{test1}) {
		t.Errorf("got %v for input %v; expected %v", actual, test1, [][]byte{test1})
	}
	actual = DeBuff(test1, []byte{0x11}, []byte{})
	if !reflect.DeepEqual(actual, [][]byte{{0x00}, {0x11, 0x22, 0x33, 0x44}, {0x11, 0x22}, {0x11}}) {
		t.Errorf("got %v for input %v; expected %v", actual, test1, [][]byte{{0x00}, {0x11, 0x22, 0x33, 0x44}, {0x11, 0x22}, {0x11}})
	}
	actual = DeBuff(test1, []byte{}, []byte{0x11})
	if !reflect.DeepEqual(actual, [][]byte{{0x00, 0x11}, {0x22, 0x33, 0x44, 0x11}, {0x22, 0x11}}) {
		t.Errorf("got %v for input %v; expected %v", actual, test1, [][]byte{{0x00, 0x11}, {0x22, 0x33, 0x44, 0x11}, {0x22, 0x11}})
	}
	actual = DeBuff(test1, []byte{}, []byte{0x55})
	if !reflect.DeepEqual(actual, [][]byte{test1}) {
		t.Errorf("got %v for input %v; expected %v", actual, test1, [][]byte{test1})
	}
	actual = DeBuff(slice4, []byte{0x1b, 0x66, 0x00, 0x14}, []byte{})
	if !reflect.DeepEqual(actual, [][]byte{slice4[:len(slice4)-5], slice4[len(slice4)-5:]}) {
		t.Errorf("got %v for input %v; expected %v", actual, slice4, [][]byte{slice4[:len(slice4)-5], slice4[len(slice4)-5:]})
	}
}

func TestNowTime(t *testing.T) {
	fmt.Println(Hex(NowTimeBCD()))
	fmt.Println(Hex(NowTimeYS()))
	fmt.Println(Hex(NowTimeSY()))
}
