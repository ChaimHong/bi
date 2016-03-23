package bi

import (
	"log"
	"testing"
)

func Test_BitToInt(t *testing.T) {
	val := Bit2Int(101)

	log.Println(val)
}

func Test_IntToBit(t *testing.T) {
	log.Println(Int2Bit(1000))
}
