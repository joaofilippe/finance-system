package main

import (
	"fmt"
	"strconv"
)

type Money struct {
	Value int64
	Reais int64
	Cents int64
}

func (m *Money) ToMoney() error {
	valueStr := strconv.FormatInt(m.Value, 10)

	centsStr := fmt.Sprintf(valueStr[len(valueStr)-2:])

	reaisStr := fmt.Sprintf(valueStr[:len(valueStr)-2])

	reais, err := strconv.ParseInt(reaisStr, 10, 64)
	if err != nil {
		return err
	}

	cents, err := strconv.ParseInt(centsStr, 10, 64)
	if err != nil {
		return err
	}

	m.Reais = reais
	m.Cents = cents

	return nil
}