package tree

type Trie interface {
	Add(word string, value interface{}) bool
	List() []string
	Clear()
	Contains(word string) bool
	ContainsPrefix(prefix string) bool
	LPM(longerWord string) (TrieNode, bool)
	Delete(word string) bool
}

type TrieNode interface {
	Word() string
	Value() interface{}
}

func NewTrie() Trie {
	return newTrie()
}

func NewTrieCC() Trie {
	return newTrieCC()
}
