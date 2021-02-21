package maps

import "errors"

// Dictionary type
type Dictionary map[string]string

// Search method
func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", errors.New("Word missing")
	}

	return definition, nil
}

// Add method
func (d Dictionary) Add(key string, value string) {
	d[key] = value
}
