package digits

import (
	"fmt"
)

type Digits []rune

func (d Digits) Less(i, j uint64) bool {
	return d[i] < d[j]
}

func (d Digits) Swap(i, j uint64) {
	d[i], d[j] = d[j], d[i]
}

func (d Digits) Len() uint64 {
	return uint64(len(d))
}

func (dptr *Digits) Decrement() {
	d := *dptr
	for i := uint64(0); i < d.Len(); i++ {
		if d[i] != '0' {
			d[i]--
			break
		}
		d[i] = '9'
		i++
	}
	zero := d.Len()
	i := zero - 1
	for d[i] == '0' {
		zero = i
		i--
	}
	*dptr = d[:zero]
}

func (dptr *Digits) Increment() {
	d := *dptr
	for i := uint64(0); i < d.Len(); i++ {
		if d[i] != '9' {
			d[i]++
			return
		}
		d[i] = '0'
		i++
	}
	*dptr = append(d, '1')
}

func (d Digits) reverse() {
	i := d.Len() - 1
	j := uint64(0)
	for i > j {
		d[i], d[j] = d[j], d[i]
	}
}

// Scan is a support routine for fmt.Scanner
func (dptr *Digits) Scan(s fmt.ScanState, ch rune) error {
	s.SkipSpace() // skip leading space characters
	*dptr = make(Digits, 0)
	for {
		r, _, err := s.ReadRune()
		switch {
		case err != nil:
			dptr.reverse()
			return err
		case '0' <= r && r <= '9':
			*dptr = append(*dptr, r)
		default:
			s.UnreadRune()
			dptr.reverse()
			return nil
		}
	}
}
