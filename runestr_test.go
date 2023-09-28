package runestr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(10, Length("abcde12345"))
	assert.Equal(0, Length(""))
	assert.Equal(1, Length("¨"))
	assert.Equal(2, Length("世界"))
}

func TestFill(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("test      ", PadRight("test", " ", 10))
	assert.Equal("test ", PadRight("test", " ", 5))
	assert.Equal("test", PadRight("test", " ", 4))
	assert.Equal("test", PadRight("test", " ", 3))
	assert.Equal("test......", PadRight("test", "..", 10))
	assert.Equal("test.....", PadRight("test", "..", 9))
	assert.Equal("test_-¨_-¨_-", PadRight("test", "_-¨", 12))
	assert.Equal("test_-¨_-¨", PadRight("test", "_-¨", 10))
	assert.Equal("test_", PadRight("test", "_-¨", 5))
	assert.Equal("test", PadRight("test", "_-¨", 4))
	assert.Equal("test世界世", PadRight("test", "世界", 7))
	assert.Equal("test世界世界世界", PadRight("test", "世界世界世界世界世界世界", 10))
	assert.Equal("test", PadRight("test", "", 10))
}

func TestLeft(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("test Left", Left("test Left", 9))
	assert.Equal("test", Left("test", 9))
	assert.Equal("test ", Left("test Left", 5))
	assert.Equal("test", Left("test Left", 4))
	assert.Equal("t", Left("test Left", 1))
	assert.Equal("", Left("test Left", 0))
	assert.Equal("", Left("", 7))
	assert.Equal("1¨2", Left("1¨2", 5))
	assert.Equal("1¨2", Left("1¨2", 3))
	assert.Equal("1¨", Left("1¨2", 2))
	assert.Equal("1", Left("1¨2", 1))
}

func TestRight(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("est Right", Right("test Right", 9))
	assert.Equal("Right", Right("test Right", 5))
	assert.Equal("t", Right("test Right", 1))
	assert.Equal("", Right("test Right", 0))
	assert.Equal("", Right("", 9))
	assert.Equal("1¨2", Right("1¨2", 5))
	assert.Equal("1¨2", Right("1¨2", 3))
	assert.Equal("¨2", Right("1¨2", 2))
	assert.Equal("2", Right("1¨2", 1))
}

func TestSplitOnNearestSpace(t *testing.T) {
	assert := assert.New(t)
	s := "Tout peut sortir d'un mot qu'en passant vous perdîtes."
	first, second := SplitOnNearestSpace(s, 9)
	assert.Equal("Tout peut", first)
	assert.Equal("sortir d'un mot qu'en passant vous perdîtes.", second)
	first, second = SplitOnNearestSpace(s, 12)
	assert.Equal("Tout peut", first)
	assert.Equal("sortir d'un mot qu'en passant vous perdîtes.", second)
	first, second = SplitOnNearestSpace(s, 16)
	assert.Equal("Tout peut sortir", first)
	assert.Equal("d'un mot qu'en passant vous perdîtes.", second)
	first, second = SplitOnNearestSpace(s, 17)
	assert.Equal("Tout peut sortir", first)
	assert.Equal("d'un mot qu'en passant vous perdîtes.", second)
	first, second = SplitOnNearestSpace(s, 18)
	assert.Equal("Tout peut sortir", first)
	assert.Equal("d'un mot qu'en passant vous perdîtes.", second)
	first, second = SplitOnNearestSpace(s, 53)
	assert.Equal("Tout peut sortir d'un mot qu'en passant vous", first)
	assert.Equal("perdîtes.", second)
	first, second = SplitOnNearestSpace(s, 54)
	assert.Equal(s, first)
	assert.Equal("", second)
	first, second = SplitOnNearestSpace(s, 55)
	assert.Equal(s, first)
	assert.Equal("", second)
	first, second = SplitOnNearestSpace(s, 100)
	assert.Equal(s, first)
	assert.Equal("", second)
	first, second = SplitOnNearestSpace("", 10)
	assert.Equal("", first)
	assert.Equal("", second)
	first, second = SplitOnNearestSpace("Toutpeutsortird'unmotqu'enpassantvousperdîtes.", 10)
	assert.Equal("Toutpeutso", first)
	assert.Equal("rtird'unmotqu'enpassantvousperdîtes.", second)
	first, second = SplitOnNearestSpace("troisième ligne", 10)
	assert.Equal("troisième", first)
	assert.Equal("ligne", second)
}

func TestRuneAtPosition(t *testing.T) {
	assert := assert.New(t)
	assert.Equal('t', RuneAtPosition("test 世界 perdîtes", 0, 'x'))
	assert.NotEqual('世', RuneAtPosition("test 世界 perdîtes", 6, 'x'))
	assert.Equal('世', RuneAtPosition("test 世界 perdîtes", 5, 'x'))
	assert.Equal('界', RuneAtPosition("test 世界 perdîtes", 6, 'x'))
	assert.Equal('î', RuneAtPosition("test 世界 perdîtes", 12, 'x'))
	assert.Equal('x', RuneAtPosition("test 世界 perdîtes", 20, 'x'))
}
