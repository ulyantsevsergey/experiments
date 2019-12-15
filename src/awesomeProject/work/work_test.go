package job

import "testing"

func TestRandStringRunes(t *testing.T) {
	var lengthShould, lengthReal  int
	var str string
	for lengthShould=1; lengthShould<=3; lengthShould++{
		str = RandStringRunes(lengthShould)
		lengthReal = len(str)
		if lengthShould != lengthReal{
			t.Error("Length of RandStringRunes returned string is incorrect", lengthReal)
		}
	}
}

