package tree

type Trie interface {
	Add(words ...string) bool
	Clear()
	Contains(word string) bool
	ContainsPrefix(prefix string) bool
	LPM(word string) *TrieNode
	Delete(word string) bool
}

type TrieNode interface {
	Word() string
	Custom() interface{}
}
