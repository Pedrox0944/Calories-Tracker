package data

import (
	"encoding/json"
	"fmt"
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

func DeletarAlimentos(dataStr string, nomeAlimento string) error {
	caminho := filepath.Join("data",dataStr+".json")

	conteudo, err := os.ReadFile(caminho)

	if  err != nil {
		return err
	}

	var alimentos []Alimento
	if err := json.Unmarshal(conteudo, &alimentos); err != nil {
		return err
	}

	novaArray := make([]Alimento, 0)
	for _, a := range alimentos {
		if a.Alimento != nomeAlimento {
			novaArray = append(novaArray, a)
		}
	}

	if len(novaArray) == len(alimentos) {
		return fmt.Errorf("Alimento: '%s' não encontrado em %s", nomeAlimento, dataStr)
	}

	novosDados, err := json.MarshalIndent(novaArray, "","  ")
	if err != nil {
		return err
	}

	return os.WriteFile(caminho, novosDados, 0644)
}