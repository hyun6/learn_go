package dictionary

import "errors"

// typedef
type Dictionary map[string]string

var (
	errNotFound     = errors.New("not found")
	errNotExist     = errors.New("not exist")
	errAlreadyExist = errors.New("already exist")
)

func (d Dictionary) Search(key string) (string, error) {
	val, exist := d[key]
	if exist {
		return val, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(key string, val string) error {
	_, err := d.Search(key)
	switch err {
	case errNotFound:
		d[key] = val
	case nil:
		return errAlreadyExist
	}
	return nil
}

func (d Dictionary) Update(key string, val string) error {
	_, err := d.Search(key)
	switch err {
	case errNotFound:
		return errNotExist
	case nil:
		d[key] = val
	}
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case errNotFound:
		return errNotExist
	case nil:
		delete(d, key)
	}
	return nil
}