package data

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

func SalvarAlimento(alimento Alimento)error{
	hoje := time.Now().Format("2006-04-25")
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