package gildedrose

import (
	"testing"
)

func Test_Foo(t *testing.T) {
	var items = []*Item{{"foo", 0, 0}}

	UpdateQuality(items)

	if items[0].Name != ("foo") {
		t.Errorf("Name: Expected %s but got %s ", "foo", items[0].Name)
	}
}
