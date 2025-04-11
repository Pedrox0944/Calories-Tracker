package cmd

import (
	"github.com/Pedrox0944/Calories-Tracker/data"
	"fmt"
	"github.com/spf13/cobra"
)

var DataRegistro string

var listarcmd = &cobra.Command{
	Use: "Listar",
	Short: "list all the foods of the day",
	Long: "Function to list all the elements registered by the user, through this function he knows how many calories he ingested",
	Run: func(cmd *cobra.Command, args []string){
		if DataRegistro == "" {
			fmt.Println("Por favor digite o dia que você deseja olhar suas calorias no formato = --data 2025-04-25")
			return 
		}

		alimentos, err := data.CarregarAlimentos(DataRegistro)
		if err != nil {
			fmt.Println("Erro ao carregar os dados")
			return
		}

		fmt.Printf("Registro da data: %s\n", DataRegistro)

		total := 0
		for _, alimento := range alimentos{
			fmt.Printf("- %s: %d calorias\n", alimento.Alimento, alimento.Calorias)
			total += alimento.Calorias
		}
		fmt.Printf("Total: %d Calorias", total)
	},
}

func init(){
	rootCmd.AddCommand(listarcmd)
	listarcmd.Flags().StringVar(&DataRegistro,"data","","Data no formato YYYY-MM-DD")
}