package processamento

import (
	"regexp"
	"testing"
)

// TestGeraCodigoAcesso verifica se o código de acesso gerado está no formato correto.
func TestGeraCodigoAcesso(t *testing.T) {
	codigo := GeraCodigoAcesso()

	// Regex para verificar o formato do código (por exemplo, IK-ABC1234)
	match, _ := regexp.MatchString(`^IK-[A-Z]{3}[0-9]{4}$`, codigo)

	if !match {
		t.Errorf("Código de acesso gerado está no formato incorreto: %s", codigo)
	}

	// Teste adicional para verificar a unicidade dos códigos gerados
	outroCodigo := GeraCodigoAcesso()
	if codigo == outroCodigo {
		t.Errorf("Código de acesso gerado não é único")
	}
}
