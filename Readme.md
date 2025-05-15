

Banco de Dados NoSQL em Go

Um mini banco de dados não relacional construído do zero com a linguagem Go. Persistência de dados em arquivos JSON, ideal para aprendizado e pequenos projetos que dispensam servidores externos.





---

✨ Funcionalidades

CRUD completo: crie, leia, atualize e exclua dados.

Armazenamento em JSON: persistência local e simples.

Design modular: código dividido por responsabilidade.

Interface no terminal: experiência interativa via CLI.



---

⚙️ Como usar

1. Clone o projeto



git clone https://github.com/bryanzns/BancoDeDados-not-sql-em-golang.git
cd BancoDeDados-not-sql-em-golang

2. Execute a aplicação



go run main.go


---

🧠 Estrutura

.
├── main.go                # Interface principal e menu
├── engrenagens.go         # Mecanismos internos (engine)
├── funcionalidades.go     # Lógica das operações (CRUD)
├── pessoas.json           # Base de dados persistente
└── go.mod                 # Dependências e versão do Go


---

🛠 Tecnologias

Go — linguagem principal

JSON — formato leve para persistência de dados



---

📌 Objetivo

Este projeto nasceu para aprender e demonstrar como é possível construir um sistema de banco de dados simples, funcional e sem dependências externas, usando apenas a linguagem Go e arquivos JSON.


---

🤝 Contribuindo

Quer colaborar? Fique à vontade para abrir uma issue ou um pull request com melhorias, correções ou novas ideias.


---

📄 Licença

Distribuído sob a licença MIT. Veja o arquivo LICENSE para mais informações.

