package gildedrose_test

import (
	"gildedrose"
	"testing"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{{"foo", 0, 0}}

	gildedrose.UpdateQuality(items)

	if items[0].Name != ("foo") {
		t.Errorf("Name: Expected %s but got %s ", "foo", items[0].Name)
	}
}
