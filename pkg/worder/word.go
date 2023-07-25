package worder

import "log"

// Create a new node for this word from the current letter node and
// the remaining word slice.
func createWordGraph(parent *LetterNode, rws []rune) (node *LetterNode) {

	node = &LetterNode{
		Letter: rws[0],
		Parent: parent,
		Child:  nil,
	}
	if parent != nil {
		parent.Child = node
	}

	// recurse for remaining letters if there is more than 1 to create
	// sub-slice from
	if len(rws) > 1 {
		createWordGraph(node, rws[1:])
	}
	return node
}

func (wg *LetterNode) Show() string {

	if wg.Child != nil {
		return string(wg.Letter) + wg.Child.Show()
	} else {
		return string(wg.Letter)
	}

}

// Create a word graph from a string
func CreateWordGraph(word string) (root *LetterNode) {
	if len(word) < 1 {
		log.Panicln("Error - provided string is too short -> ", word)
	}
	root = createWordGraph(nil, []rune(word))
	//log.Println("Create WordGraph -> ", root)
	return root
}
