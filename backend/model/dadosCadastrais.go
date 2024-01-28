package model

// DadosCadastrais representa a estrutura de um documento na coleção cl_dados_cadastrais
type DadosCadastrais struct {
	// Defina os campos aqui, por exemplo:
	Nome      string `bson:"nome"`
	Idade     int    `bson:"idade"`
	Profissao string `bson:"profissao"`
	// Adicione outros campos conforme a estrutura do seu documento MongoDB
}
