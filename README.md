# Desafio Processo de C.I 

Realizar processo de C.I de um aplicativo em Go que executa a soma de dois valores, o processo deve ser iniciado por um pull request efetuado no repositório.

1. Criar uma aplicação em Go que executa a soma de dois valores;
2. Criar um teste unitário que valida a execução e retorno da aplicação;
3. Criar processo de C.I que executa:
3.1. Sobe a aplicação em um container Go e realiza o teste unitário
3.2. Criar uma imagem usando um container docker;
3.3. Realizar o pull da imagem criada no Container Registry do GCP.

# Instruções

##Testes no GCP

Na conta do GCP abrir o terminal e executar:
> docker run --name somaGo gcr.io/curso-codeeducation/soma-go

Após o download da imagem o terminal deverá retornar "10" que é o valor da soma.


##Testes local

Clonar o repositório
> git clone git@github.com:rogeriopontomoura/somaGo-CI.git

Entrar na pasta e executar os comandos abaixo:

-**Criar a imagem**
-> docker build -t soma-go .

-**Executar a imagem**
-> docker run --name app-soma-go soma-go

Após a criação da imagem o terminal deverá retornar "10" que é o valor da soma.