package ecommerce

import (
	"fmt"
)

// errors
type MissingItem struct {
	operation, item string
}

func (mi MissingItem) Error() string {
	return fmt.Sprintf("%s: %s: does not exist in database", mi.operation, mi.item)
}

type ItemAlreadyExists struct {
	operation, item string
}

func (iae ItemAlreadyExists) Error() string {
	return fmt.Sprintf("%s: %s: already exists in database", iae.operation, iae.item)
}

// types
type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func NewDatabase() database {
	return database{}
}

func (db database) Get(item string) (database, error) {
	if item == "" {
		return db, nil
	}

	price, ok := db[item]
	if !ok {
		return nil, MissingItem{"get", item}
	}

	return database{item: price}, nil
}

func (db database) GetAll() database {
	results, _ := db.Get("")
	return results
}

func (db database) Insert(item string, price float32) error {
	_, ok := db[item]
	if ok {
		return ItemAlreadyExists{"insert", item}
	}

	db[item] = dollars(price)
	return nil
}

func (db database) Update(item string, price float32) error {
	_, ok := db[item]
	if !ok {
		return MissingItem{"update", item}
	}

	db[item] = dollars(price)
	return nil
}

func (db database) Delete(item string) error {
	_, ok := db[item]
	if !ok {
		return MissingItem{"delete", item}
	}

	delete(db, item)
	return nil
}
