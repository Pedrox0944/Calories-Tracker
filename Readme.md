O calories tracker é um projeto pessoal feito para testar meus conhecimentos prévios em Go e me ajudar com minha dieta diária

-> Compilando 
Compile o go usando o comando **go build -o nomedoexecutavel**
Ele fará a instalação do Go na versão 1.23.7(Versão usada no desenvolvimento do programa) e do Cobra-CLI

-> Como usar 
* Cadastrando um alimento
**./nomedoexecutavel Refeicao --Nome "nome do alimento" --Calorias quantidade --Proteinas quantidade**
Será criado um arquivo json com a nome sendo a data de criação e lá será gravado o alimento, suas respectivas calorias e suas proteínas
exemplo: *./Caloriestracker Refeicao --Nome "Arroz com batata" --Calorias 250 --Proteinas 10*

* Listando suas refeições
**./nomedoexecutavel Listar --Data YYYY-MM-DD**
Vai ser exibido no seu terminal todas as suas refeições com suas calorias e a soma total de todas elas
exemplo: *./Caloriestracker Listar --Data 2025-04-14*
 
* Apagando alimentos
**./nomedoexecutavel Apagar --Data YYYY-MM-DD --Alimento "nome do alimento"**
Esse comando apagará o alimento que voçê deseja, é importante escrever o nome do alimento de forma perfeita para não ficar dando erros
exemplo: *./Caloriestracker Apagar --Data 2025-04-14 --Alimento "Arroz com batata"*
 