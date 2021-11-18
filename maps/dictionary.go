package maps

/*
Constant exports
*/
const (
	ErrWordUnknown        = DictionaryErr("unknown word")
	ErrWordAlreadyDefined = DictionaryErr("word is already defined")
)

/*
Errors for Dictionary functions
*/
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

/*
Dictionary with words and their definitions
*/
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrWordUnknown
	}
	return def, nil
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordUnknown:
		d[word] = def
	case nil:
		return ErrWordAlreadyDefined
	default:
		return err
	}
	return nil
}
