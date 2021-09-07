package g165

import "testing"

func TestCompareVersion(t *testing.T) {
	t.Log(compareVersion("1.01", "1.001"))
}
