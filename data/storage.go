package data

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

func SalvarAlimento(alimento Alimento)error{
	hoje := time.Now().Format("2006-01-02")
	caminho := filepath.Join("data", hoje+".json")

	var alimentos []Alimento
	if _, err := os.Stat(caminho); err == nil{
		conteudo, _ := os.ReadFile(caminho)
		json.Unmarshal(conteudo, &alimentos)
	}

	alimentos = append(alimentos, alimento)
	novosdados, err := json.MarshalIndent(alimentos,"","  ")
	if err != nil {
		return err
	}
	return os.WriteFile(caminho, novosdados, 0644)
}

func CarregarAlimentos(data string) ([]Alimento, error){
	caminho := filepath.Join("data", data+".json")

	var alimentos[] Alimento
	conteudo, err := os.ReadFile(caminho)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(conteudo, &alimentos)
	if  err != nil {
		return nil, err
	}

	return alimentos, nil
}