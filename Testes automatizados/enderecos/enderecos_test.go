// Teste de unidade
package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEnredeco(t *testing.T) {

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Praça ABC", "Tipo Inválido"},
		{"Rua ABC", "Rua"},
		{"Rua ABC", "Avenida"},
		{"Avenida ABC", "Avenida"},
		{"Praça ABC", "Avenida"},
		{"Rodovia ABC", "Estrada"},
		{"Praça ABC", "Rodovia"},
		{"Rodovia ABC", "Rodovia"},
		{"Rua ABC", "Estrada"},
		{"Estrada ABC", "Estrada"},
	}

	for _, cenario := range cenarioDeTeste {
		TipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if TipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				TipoDeEnderecoRecebido,
				cenario.retornoEsperado)
		}
	}
}

// func TestTipoDeEnredeco(t *testing.T) {
// 	enderecoParaTeste := "Avenida Paulista"

// 	tipoDeEnderecoEsperado := "Avenida"

// 	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

// 	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
// 		t.Errorf("O tipo recebido é diferente do esperado!\n Valor esperado: %s\nRecebido: %s",
// 			tipoDeEnderecoEsperado,
// 			tipoDeEnderecoRecebido)
// 	}
// }
