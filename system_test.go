package mathsyst

import (
	"testing"
)

func TestSystem_BothWays(t *testing.T) {
	a := "01"
	s := NewSystem(a)
	for i:=0; i< 100; i++ {
		 cn:= s.FromDec(i)
		 dn, err := s.ToDec(cn)
		 if err != nil {
		 	t.Error(err)
			 return
		 }
		 if dn != i {
		 	t.Error("NOt equal: ", dn, i)
		 }
	}
}

func TestSystem_FromDec(t *testing.T) {
	dd := []struct{
		name string
		alphabet string
		toTest []int
		toMatch []string
	} {
		{
			name: "Hex",
			alphabet: Hex,
			toTest: []int{0,1,2,3,15,16,31,32},
			toMatch: []string{"0","1","2","3","f", "10","1f","20"},
		},
		{
			name: "Binary",
			alphabet: BI,
			toTest: []int{0,1,2,3,4,5,6,7},
			toMatch: []string{"0", "1","10","11", "100", "101", "110","111"},
		},
	}

	for _,d:= range dd {
		t.Run(d.name, func(t *testing.T) {
			s := NewSystem(d.alphabet)
			for i, n := range d.toTest {
				m := s.FromDec(n)
				if m != d.toMatch[i] {
					t.Errorf("the strings do not match: %s, %s", d.toMatch[i], m)
				}
			}
		})
	}

}


func TestSystem_Add(t *testing.T) {
	dd := []struct{
		name string
		alphabet string
		toTest []struct {
			n,m, f string
		}
	} {
		{
			name: "Hex",
			alphabet: Hex,
			toTest: []struct {
				n,m,f string
			}{
				{n:"0",m:"0", f: "0"},
				{"0", "1", "1"},
				{"9", "1","a"},
				{"8", "2", "a"},
				{"f", "1", "10"},
				{"f", "f", "1e"},
			},
		},
		{
			name: "Bin",
			alphabet: BI,
			toTest: []struct {
				n,m,f string
			}{
				{n:"0",m:"0", f: "0"},
				{"0", "1", "1"},
				{"10", "1","11"},
				{"10", "10", "100"},
				{"100", "10", "110"},
				{"100100", "10101", "111001"},
			},
		},
	}

	for _,d:= range dd {
		t.Run(d.name, func(t *testing.T) {
			s := NewSystem(d.alphabet)
			for _, td := range d.toTest {
				m, err := s.Add(td.n, td.m)
				if err != nil {
					t.Errorf(err.Error())
				}
				if m != td.f {
					t.Errorf("the strings do not match: %s, %s", td.f, m)
				}
			}
		})
	}
}

func TestSystem_Diff(t *testing.T) {
	dd := []struct{
		name string
		alphabet string
		toTest []struct {
			n,m, f string
		}
	} {
		{
			name: "Hex",
			alphabet: Hex,
			toTest: []struct {
				n,m,f string
			}{
				{n:"0",m:"0", f: "0"},
				{"0", "1", "-1"},
				{"9", "1","8"},
				{"-8", "2", "-a"},
				{"f", "1", "e"},
				{"f", "f", "0"},
			},
		},
		{
			name: "Bin",
			alphabet: BI,
			toTest: []struct {
				n,m,f string
			}{
				{n:"0",m:"0", f: "0"},
				{"0", "1", "-1"},
				{"10", "1","1"},
				{"10", "10", "0"},
				{"100", "10", "10"},
				{"100100", "10101", "1111"},
			},
		},
	}

	for _,d:= range dd {
		t.Run(d.name, func(t *testing.T) {
			s := NewSystem(d.alphabet)
			for _, td := range d.toTest {
				m, err := s.Diff(td.n, td.m)
				if err != nil {
					t.Errorf(err.Error())
				}
				if m != td.f {
					t.Errorf("the strings do not match: %s, %s", td.f, m)
				}
			}
		})
	}
}

func TestSystem_Mult(t *testing.T) {
	dd := []struct{
		name string
		alphabet string
		toTest []struct {
			n,m, f string
		}
	} {
		{
			name: "Hex",
			alphabet: Hex,
			toTest: []struct {
				n,m,f string
			}{
				{n:"0",m:"0", f: "0"},
				{"0", "1", "0"},
				{"1", "1","1"},
				{"-1", "1", "-1"},
				{"-1", "-1", "1"},
				{"-f", "2", "-1e"},
			},
		},
		{
			name: "Bin",
			alphabet: BI,
			toTest: []struct {
				n,m,f string
			}{
				{n:"0",m:"0", f: "0"},
				{"0", "1", "0"},
				{"10", "1","10"},
				{"10", "10", "100"},
				{"101", "10", "1010"},
				{"-10", "-10", "100"},
				{"-11", "10", "-110"},
			},
		},
	}

	for _,d:= range dd {
		t.Run(d.name, func(t *testing.T) {
			s := NewSystem(d.alphabet)
			for _, td := range d.toTest {
				m, err := s.Mult(td.n, td.m)
				if err != nil {
					t.Errorf(err.Error())
				}
				if m != td.f {
					t.Errorf("the strings do not match: %s, %s", td.f, m)
				}
			}
		})
	}
}

func TestSystem_Div(t *testing.T) {
	dd := []struct{
		name string
		alphabet string
		toTest []struct {
			n,m, f string
		}
	} {
		{
			name: "Hex",
			alphabet: Hex,
			toTest: []struct {
				n,m,f string
			}{
				{n:"14",m:"2", f: "a"},
				{"-f", "5", "-3"},
			},
		},
		{
			name: "Bin",
			alphabet: BI,
			toTest: []struct {
				n,m,f string
			}{
				{n:"10",m:"10", f: "1"},
				{"111", "-10", "-11"},
			},
		},
	}

	for _,d:= range dd {
		t.Run(d.name, func(t *testing.T) {
			s := NewSystem(d.alphabet)
			for _, td := range d.toTest {
				m, err := s.Div(td.n, td.m)
				if err != nil {
					t.Errorf(err.Error())
				}
				if m != td.f {
					t.Errorf("the strings do not match: %s, %s", td.f, m)
				}
			}
		})
	}
}