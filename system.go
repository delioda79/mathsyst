package mathsyst

import (
	"errors"
	"math"
)

const AlphaNum = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Hex = "0123456789abcdef"
const BI = "01"
const Oct = "01234567"

type System struct {
	decTo []string
	toDec map[string]int
}

func (s System) ToDec(n string) (int, error) {
	sign := float64(1)
	if n[0:1] == "-" {
		sign = -1
		n = n[1:]
	}
	base := float64(len(s.toDec))
	var res float64
	for i:=0;i< len(n); i++ {
		d, ok := s.toDec[n[i:i+1]]
		if !ok {
			return 0, errors.New("character not in alphabet")
		}
		p := float64(len(n)-1-i)

		res += math.Pow(base,p)*float64(d)*sign
	}

	return int(res), nil
}

func (s System) FromDec(n int)  string {
	base := float64(len(s.toDec))
	var res string
	var p float64
	if n == 0 {
		return s.decTo[0]
	}
	if n <0 {
		res +="-"
		n= n*-1
	}
	for {
		if n == 0 {
			for i:=p-1; i>=0; i-- {
				res += s.decTo[0]
			}
			break
		}
		cur := p
		l := math.Log(float64(n))/math.Log(base)
		p = math.Floor(l)

		if cur-p > 1 {
			for i:=0; i<int(cur-p)-1; i++ {
				res += s.decTo[0]
			}
		}

		c := math.Floor(float64(n)/math.Pow(base,p))
		char := s.decTo[int(c)]
		res += char

		n = n-int(c*(math.Pow(base,p)))
	}

	return res
}

func (s System) Add(n,m string) (string, error) {
	dn, err := s.ToDec(n)
	if err != nil {
		return "", err
	}

	dm, err := s.ToDec(m)
	if err != nil {
		return "", err
	}

	return s.FromDec(dn+dm), nil
}

func (s System) Diff(n,m string) (string, error) {
	dn, err := s.ToDec(n)
	if err != nil {
		return "", err
	}

	dm, err := s.ToDec(m)
	if err != nil {
		return "", err
	}

	return s.FromDec(dn-dm), nil
}

func (s System) Mult(n,m string) (string, error) {
	dn, err := s.ToDec(n)
	if err != nil {
		return "", err
	}

	dm, err := s.ToDec(m)
	if err != nil {
		return "", err
	}

	return s.FromDec(dn*dm), nil
}

func (s System) Div(n,m string) (string, error) {
	dn, err := s.ToDec(n)
	if err != nil {
		return "", err
	}

	dm, err := s.ToDec(m)
	if err != nil {
		return "", err
	}

	return s.FromDec(dn/dm), nil
}

func (s System) WithLeadingZeros(n string, c int) string {
	var f string
	zs := int(math.Max(0, float64(c-len(n))))
	for i:=0; i<zs; i++ {
		f += s.decTo[0]
	}
	return f+n
}

func NewSystem(al string) System {
	s := System{
		toDec: map[string]int{},
	}
	for i:=0; i < len(al); i++ {
		ch := al[i:i+1]
		s.decTo = append(s.decTo, ch)
		s.toDec[ch]=i
	}

	return s
}

