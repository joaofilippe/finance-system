package main

import (
	"testing"
)

func Test_toMoney(t *testing.T) {
	tests := []struct {
		name  string
		money Money
		reais int64
		cents int64
	}{
		{
			name: "testando a conversãode R$ 1.255,98",
			money: Money{
				Value: 125598,
			},
			reais: 1255,
			cents: 98,
		},
		{
			name: "testando a conversãode R$ 255,99",
			money: Money{
				Value: 25599,
			},
			reais: 255,
			cents: 99,
		},
		{
			name: "testando a conversãode R$ 65.255,00",
			money: Money{
				Value: 6525500,
			},
			reais: 65255,
			cents: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.money.ToMoney()

			if tt.money.Cents != tt.cents {
				t.Errorf("a conversão de centavos não foi bem sucedida\nresultado: %d, esperado: %d", tt.money.Cents, tt.cents)
			}

			if tt.money.Reais != tt.reais {
				t.Errorf("a conversão de reais não foi bem sucedida\nresultado: %d, esperado: %d", tt.money.Reais, tt.reais)
			}
		})
	}
}
