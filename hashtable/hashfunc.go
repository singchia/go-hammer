package hashtable

import (
	"bytes"
	"encoding/binary"
)

func SamplingHash(key interface{}) uint64 {

	// sampling short elems from long
	sequence := func(long, short int) []int {
		if long < short {
			seq := make([]int, long)
			for i := 0; i < long; i++ {
				seq[i] = i
			}
			return seq
		} else {
			seq := make([]int, short)
			quto := long / short
			remd := long % short
			// ceil devision
			makeup := ((remd + short - 1) / short) * short
			for i, j := 0, 0; i < long && j < short; i, j = i+quto, j+1 {
				if makeup != 0 && (j+1)%makeup == 0 {
					i++
				}
				seq[j] = i
			}
			return seq
		}
	}

	switch real := key.(type) {
	case HasHash:
		return real.Hash()
	default:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, key)
		if err != nil {
			return 0
		}
		bytes := buf.Bytes()
		seq := sequence(len(bytes), 8)
		count := uint64(0)
		for i, index := range seq {
			count += uint64(int(bytes[index]) << (i * 8))
		}
		return count
	}
}
