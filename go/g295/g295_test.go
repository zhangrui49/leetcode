package g295

import "testing"

func TestFindMedian(t *testing.T) {
	obj := Constructor()
	// obj.AddNum(-1)
	// obj.AddNum(-2)
	// obj.AddNum(-3)
	// obj.AddNum(-4)

	obj.AddNum(6)
	obj.AddNum(10)
	obj.AddNum(2)
	obj.AddNum(6)
	// obj.AddNum(5)
	// obj.AddNum(0)
	// obj.AddNum(6)
	// obj.AddNum(3)
	// obj.AddNum(1)
	// obj.AddNum(0)
	// obj.AddNum(0)

	t.Log(obj.FindMedian())
}
