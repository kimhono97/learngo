package mydict

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

var (
	errNotFound = errors.New("the key was not found")
	errFound    = errors.New("the key already exists")
)

func (d Dictionary) Search(key string) (string, error) {
	val, isExist := d[key]
	if !isExist {
		return "", errNotFound
	}
	return val, nil
}

func (d Dictionary) Add(key, val string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		return errFound
	case errNotFound:
		d[key] = val
	}
	return nil
}

func (d Dictionary) Update(key, val string) error {
	_, err := d.Search(key)
	switch err {
	case errNotFound:
		return err
	case nil:
		d[key] = val
	}
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case errNotFound:
		return err
	case nil:
		delete(d, key)
	}
	return nil
}

func (d Dictionary) String() string {
	out := "{\n"
	for key, val := range d {
		out += fmt.Sprint("\t", key, " : ", val, "\n")
	}
	out += "}"
	return out
}
