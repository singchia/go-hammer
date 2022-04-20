package tree

import (
	"sync"
	"testing"
)

func TestAddCC(t *testing.T) {
	trie := newTrieCC()
	wg := new(sync.WaitGroup)
	wg.Add(len(wordsAdd))
	for _, word := range wordsAdd {
		go func(word string) {
			defer wg.Done()
			trie.Add(word, nil)
		}(word)
	}
	wg.Wait()
	list := trie.List()
	if !stringsEqual(wordsAdd, list) {
		t.Error("wrong elems in trie")
		return
	}
}
