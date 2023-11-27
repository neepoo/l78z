package bloom

import "testing"

func Test_bf(t *testing.T) {
	bf, err := NewBloomFilterWithHasher(1024, 0.01, NewMurMur3Hasher())
	if err != nil {
		t.Fatal(err)
	}
	foo := []byte("foo")
	bar := []byte("bar")
	baz := []byte("baz")
	nonExist := []byte("nonExist")
	bf.Add(foo)
	bf.Add(bar)
	bf.Add(baz)
	if !bf.Test(foo) {
		t.Fatalf("%s should be in the set", foo)
	}
	if !bf.Test(bar) {
		t.Fatalf("%s  should be in the set", bar)
	}
	if !bf.Test(baz) {
		t.Fatalf("%s should be in the set", baz)
	}
	if bf.Test(nonExist) {
		t.Fatalf("%s should not be in the set", nonExist)
	}
}
