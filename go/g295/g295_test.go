package g295

import "testing"

func TestFindMedian(t *testing.T) {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	t.Log(obj.FindMedian())
}
