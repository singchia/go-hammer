package tree

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func stringsEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, aElem := range a {
		found := false
		for _, bElem := range b {
			if aElem == bElem {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

var (
	wordsAdd = []string{
		"a",
		"b",
		"ab",
		"abc",
		"abcd",
		"abcde",
		"abcdef",
		"abcdefg",
		"abcdefgh",
		"abcdefghi",
		"m",
		"mn",
		"mno",
		"mnop",
		"mnopq",
	}
	wordsNotAdd = []string{
		"z",
		"d",
		"zb",
		"zbc",
		"abcz",
		"abcze",
		"abczef",
		"abczefg",
		"abczefgh",
		"abczefghi",
	}
)

func prefix(word string) []string {
	words := []string{}
	for i := len(word); i > 0; i-- {
		words = append(words, word[:i])
	}
	return words
}

func TestAdd(t *testing.T) {
	trie := newTrie()
	for index, word := range wordsAdd {
		trie.Add(word, nil)
		list := trie.List()
		if !stringsEqual(wordsAdd[0:index+1], list) {
			t.Error("wrong elems in trie")
			return
		}
		t.Log(index, list)
	}
}

func TestContains(t *testing.T) {
	trie := newTrie()
	for index, word := range wordsAdd {
		trie.Add(word, nil)
		ok := trie.Contains(wordsAdd[rand.Intn(index+1)])
		if !ok {
			t.Error("elem must be in trie")
			return
		}
		t.Log(index, trie.List())
	}
	for index, word := range wordsNotAdd {
		ok := trie.Contains(word)
		if ok {
			t.Error("elem must not be in trie")
			return
		}
		t.Log(index, trie.List())
	}
}

func TestContainsPrefix(t *testing.T) {
	trie := newTrie()
	for _, word := range wordsAdd {
		trie.Add(word, nil)
	}
	for _, word := range wordsAdd {
		prefixes := prefix(word)
		for _, prefix := range prefixes {
			t.Log(prefix)
			ok := trie.ContainsPrefix(prefix)
			if !ok {
				t.Error("elem must be in trie")
				return
			}
		}
	}
}

func TestLPM(t *testing.T) {
	trie := newTrie()
	longest := wordsAdd[0]
	for _, word := range wordsAdd {
		trie.Add(word, nil)
		if len(word) > len(longest) {
			longest = word
		}
	}
	match := longest + "foo"
	node, ok := trie.LPM(match)
	if ok {
		if node.Word() != longest {
			t.Error("string not matched")
			return
		}
		t.Log(node.Word(), node.Value())
	} else {
		t.Error("elem must be matched")
		return
	}
}

func TestDelete(t *testing.T) {
	trie := newTrie()
	for _, word := range wordsAdd {
		trie.Add(word, nil)
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(wordsAdd))
	ok := trie.Delete(wordsAdd[index])
	if !ok {
		t.Error("elem should be deleted")
		return
	}
	list := trie.List()
	for _, elem := range list {
		if elem == wordsAdd[index] {
			t.Error("elem not deleted")
			return
		}
	}
	t.Log(wordsAdd[index], trie.List())
}

func BenchmarkAdd(b *testing.B) {
	trie := newTrie()
	for i := 0; i < b.N; i++ {
		trie.Add(strconv.Itoa(i), nil)
	}
}
