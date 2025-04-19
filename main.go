package main
import (
  "fmt"
  "os"
)
func main()  {
  for {
  fmt.Println("-----------------")
  fmt.Println("     BANCO DE DADOS")
  fmt.Println("(1)- incluir")
  fmt.Println("(2)- excluir")
  fmt.Println("(3)- alterar")
  fmt.Println("(4)- listar")
  fmt.Println("(5)- sair")
  fmt.Println("faça sua escolha:")
  var escolha int
  fmt.Scan(&escolha)
  switch escolha {
  case 1:
  IncluirDados()
  case 2:
  ExcluirDados()
  case 3:
  AlterarDados()
  case 4:
  ListarDados()
  case 5:
  os.Exit(0)
    break
    }
  }
}
