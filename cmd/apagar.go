package cmd

import(
	"github.com/Pedrox0944/Calories-Tracker/data"
	"fmt"
	"github.com/spf13/cobra"
)

var dataDel string
var alimentoDel string

var DeletarCmd = &cobra.Command{
	Use: "Apagar",
	Short: "Remove date or food",
	Long: "Function to remove date or element that the user considers to be wrong or useless",

	Run: func(cmd *cobra.Command, args []string) {
		if dataDel == "" || alimentoDel == "" {
			fmt.Println("Por favor digite a data ou o nome do alimento que você deseja apagar usando --data ou --alimento")
			return 
		}

		err := data.DeletarAlimentos(dataDel, alimentoDel)

		if err != nil {
			fmt.Printf("Erro ao deletar o alimento")
			return 
		}

		fmt.Printf("Alimento removido com sucesso!\n")
	},
}

func init(){
	rootCmd.AddCommand(DeletarCmd)

	DeletarCmd.Flags().StringVar(&dataDel,"data","","Data no formato YYYY-MM-DD")
	DeletarCmd.Flags().StringVar(&alimentoDel, "alimento","","Nome do alimento")

}