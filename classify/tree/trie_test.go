package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie(t *testing.T) {
	t1 := Constructor()
	t1.Insert("app")
	t1.Insert("apple")
	t1.Insert("beer")
	t1.Insert("add")
	t1.Insert("jam")
	t1.Insert("rental")
	assert.False(t, t1.Search("apps"))
	assert.True(t, t1.Search("app"))
	assert.False(t, t1.Search("ad"))
	assert.False(t, t1.Search("applepie"))
	assert.False(t, t1.Search("rest"))
	assert.False(t, t1.Search("jan"))
	assert.False(t, t1.Search("rent"))
	assert.True(t, t1.Search("beer"))
	assert.True(t, t1.Search("jam"))

	assert.False(t, t1.StartsWith("apps"))
	assert.True(t, t1.StartsWith("app"))
	assert.True(t, t1.StartsWith("ad"))
	assert.False(t, t1.StartsWith("applepie"))
	assert.False(t, t1.StartsWith("rest"))
	assert.False(t, t1.StartsWith("jan"))
	assert.True(t, t1.StartsWith("rent"))
	assert.True(t, t1.StartsWith("beer"))
	assert.True(t, t1.StartsWith("jam"))
}

// ["MapSum","insert","sum","insert","insert","sum"]
//[[],["apple",3],["ap"],["app",2],["apple",2],["ap"]]
func TestMapSum(t *testing.T) {
	t1 := ConstructorMapSum()
	t1.Insert("apple", 3)
	assert.Equal(t, t1.Sum("ap"), 3)
	t1.Insert("app", 2)
	t1.Insert("apple", 2)
	assert.Equal(t, t1.Sum("ap"), 4)
}
