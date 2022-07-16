package models

//DatosEntrada son los datos que se reciben via JSON
type DatosEntrada struct {
	NombreSecret     string `json:"NombreSecret"`
	Kill_level1      string `json:"Kill_level1"`
	Kill_level2      string `json:"Kill_level2"`
	NombreSecretAPI  string `json:"NombreSecretAPI"`
	NombreSecretROOT string `json:"NombreSecretROOT"`
}
