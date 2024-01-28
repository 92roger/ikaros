package model

// TratadorDeCampoNulo é uma interface para structs que precisam tratar campos nulos.
type TratadorDeCampoNulo interface {
	TratarCamposNulos()
}
