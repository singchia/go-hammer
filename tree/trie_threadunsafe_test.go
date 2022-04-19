package tree

import "testing"

func TestAdd(t *testing.T) {
	trie := newTrie()
	trie.Add("a", nil)
	trie.Add("ab", nil)
	trie.Add("abc", nil)
	words := trie.List()
	t.Log(words)
}
