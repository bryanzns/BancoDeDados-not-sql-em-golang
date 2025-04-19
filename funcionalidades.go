package main
import (
  "fmt"
  "log"
)
func AlterarDados() {
  pessoas, err := CarregarDados()
  if err != nil {
    log.Fatalln("erro ao carregar", err)
  }
  var escolha int
  fmt.Println(" QUAL O ID DA PESSOA QUE VOCE QUER ALTERAR?")
  fmt.Scan(&escolha)
  tem := false
  for i, p := range pessoas {
    if p.Id == escolha{
      tem = true
  fmt.Println(" O QUE VOCE DESEJA ALTERAR ")
      fmt.Println("(1) nome")
      fmt.Println("(2) idade")
      fmt.Println("(3) altura")
      fmt.Println("escolha: ")
      fmt.Scan(&escolha)
      switch escolha {
      case 1:
        var alterar string
        fmt.Println(" qual o novo nome:")
        fmt.Scan(&alterar)
        pessoas[i].Nome = alterar
      case 2:
        var alterar int
        fmt.Println(" qual a nova idade:")
        fmt.Scan(&alterar)
        pessoas[i].Idade = alterar
      case 3:
        var alterar float64
        fmt.Println(" qual a nova altura:")
        fmt.Scan(&alterar)
        pessoas[i].Altura = alterar
      }
    }
  }
  if !tem {
    log.Fatalln(" ID Esta errado")
    fmt.Println("tem: ", tem)
  }
  err = SalvarDados(pessoas)
  if err != nil {
    log.Fatalln("erro ao salvar", err)
  }
  fmt.Println("pessoa alterada com sucesso")
}
func IncluirDados() {
  pessoas, err := CarregarDados()
  if err != nil{
    log.Fatalln("erro ao carregar em incluir", err)
  }
  fmt.Println("  VAMOS COLOCAR UMA PESSOA NOVAA") 
  var (
    Nome string
    Idade int
    altura float64
  )
  fmt.Println("nome:")
  fmt.Scan(&Nome)
  fmt.Println("idade: ")
  fmt.Scan(&Idade)
  fmt.Println("altura: ")
  fmt.Scan(&altura)
  nova := Pessoa{
    Nome: Nome,
    Idade: Idade,
    Altura: altura,
    Id: len(pessoas),
  }
  pessoal := append(pessoas, nova)
  err = SalvarDados(pessoal)
  if err != nil {
    log.Fatalln("erro ao salvar dados ", err)
  }
fmt.Println("dados salvos!")
}
func ExcluirDados()  {
  var escolha int
  var novalista []Pessoa
  removido := false
  pessoa, err := CarregarDados()
  if err != nil {
    log.Fatalln("erro ao carregar dados em excluir", err)
  }
  fmt.Println(" EXCLUIR ALGUEM DA LISTA")
  fmt.Println(" qual o id da pessoa que voce quer excluir?")
  fmt.Scan(&escolha)
  for _,p := range pessoa{
    if p.Id == escolha {
      removido = true
      continue
    }
  novalista = append(novalista, p)
  }
  if !removido{
    log.Fatalln("erro no id em excluir")
  }
err = SalvarDados(novalista)
if err != nil{
  log.Fatalln("erro ao salvar dados na nova lista com exclusao", err)
}
fmt.Println("pessoa excluida com sucesso!")
}
func ListarDados()  {
  var escolha int
  var qual int
  pessoa, err := CarregarDados()
  if err != nil {
    log.Fatalln("erro ao carregar dados em listar", err)
  }
  fmt.Println(" LISTAGEM DE DADOS")
  fmt.Println("(1) quero listar todos")
  fmt.Println("(2) quero listar so 1 pessoa")
  fmt.Println("escolha: ")
  fmt.Scan(&escolha)
  if escolha == 1 {
    for _, v := range pessoa {
      fmt.Println("nome:", v.Nome)
      fmt.Println("Idade:", v.Idade)
      fmt.Println("altura:", v.Altura)
      fmt.Println("Id:", v.Id)
      fmt.Println("--------------------")
    }
  }
  if escolha == 2{
    fmt.Println(" qual o id da  pessoa? ")
    fmt.Scan(&qual)
    for _, v := range pessoa {
      if v.Id == qual{
         fmt.Println("nome:", v.Nome)
         fmt.Println("Idade:", v.Idade)
         fmt.Println("Altura:", v.Altura)
         fmt.Println("Id:", v.Id)
         fmt.Println("-----------------")
      }
    }
  }
}








