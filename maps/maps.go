package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("unknown word")
	ErrWordAlreadyExist = DictionaryErr("word already exist")
	ErrWordNotExist     = DictionaryErr("word does not exist")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

func (d Dictionary) Search(key string) (string, error) {
	v, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}

func (d Dictionary) Add(key, value string) error {
	if _, ok := d[key]; !ok {
		d[key] = value
		return nil
	} else {
		return ErrWordAlreadyExist
	}

}

func (d Dictionary) Update(key, value string) error {
	_, ok := d[key]

	switch ok {
	case true:
		d[key] = value
		return nil
	case false:
		return ErrWordNotExist
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
