/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Pedrox0944/Calories-Tracker/data"
	"github.com/spf13/cobra"
)

var Nome string
var Calorias int

var RefeicaoCmd = &cobra.Command{
	Use:   "Refeicao",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example: calories_tracker refeicao --Nome "Alimento" --Calorias 10

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Nome == "" || Calorias <= 0 {
			fmt.Println("Por favor forneça um nome de alimento e suas calorias corretamente")
			return
		}

		alimento := data.Alimento{
			Alimento: Nome,
			Calorias: Calorias,
		}

		err := data.SalvarAlimento(alimento)
		if err != nil {
			fmt.Printf("Erro ao salvar o alimento: %v\n", err)
			return 
		}

		fmt.Printf("Refeicão registrada com sucesso! Alimento: %s - Calorias: %d\n", alimento.Alimento, alimento.Calorias)
	},
}

func init() {
	rootCmd.AddCommand(RefeicaoCmd)

	RefeicaoCmd.Flags().StringVar(&Nome, "Nome", "", "Nome do alimento")
	RefeicaoCmd.Flags().IntVar(&Calorias, "Calorias", 0, "Quantidade de calorias")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// RefeicaoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// RefeicaoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
