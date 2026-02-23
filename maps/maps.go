package main

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound        = errors.New("Key not found")
	ErrKeyExists       = errors.New("Key already exists")
	ErrKeyDoesNotExist = errors.New("Key does not exist")
)

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		d[key] = value
	case ErrNotFound:
		return ErrKeyDoesNotExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		delete(d, key)
	case ErrNotFound:
		return ErrKeyDoesNotExist
	default:
		return err
	}
	return nil
}
