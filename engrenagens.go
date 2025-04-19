package main
import(
  "os"
  "encoding/json"
  "io"
)
type Pessoa struct{
  Nome string    `json:"Nome"`
  Idade int      `json:"Idade"`
  Altura float64 `json:"Altura"`
  Id int         `json:"Id"`
}
func SalvarDados(dados []Pessoa) error {
  arquivo, err := os.Create("pessoas.json")
  if err != nil{
    return err
  }
  defer arquivo.Close()
  encoder := json.NewEncoder(arquivo)
  err = encoder.Encode(dados)
  if err != nil {
    return err
  }
  return nil
}
func CarregarDados() ([]Pessoa, error)  {
  arquivo, err := os.Open("pessoas.json")
  if err != nil{
    return nil, err
  }
  defer arquivo.Close()
  var lista []Pessoa
  decoder := json.NewDecoder(arquivo)
  err = decoder.Decode(&lista)
  if err != nil {
    if err == io.EOF{
      return []Pessoa{}, nil
    }
    return nil, err
  }
  return lista, nil
}
