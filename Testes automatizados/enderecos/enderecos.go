package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TipoDeEnredero - tipo valido
func TipoDeEndereco(endereco string) string {

	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoLetrasMinusculas := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoLetrasMinusculas, " ")[0]
	// EX: "RUA DOS BOBOS" -> ["RUA","DOS","BOBOS"]

	var enderecoTipoValido bool = false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoTipoValido = true
		}
	}
	if enderecoTipoValido {
		return cases.Title(language.Und, cases.NoLower).String(primeiraPalavraEndereco)
	}

	return "Tipo Inválido"
}
