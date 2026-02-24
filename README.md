# 📌 Nome do Projeto

O nome do projeto foi seleciona pois sua função era um desafio dos conhecimentos ja utilizados, por isso o nome ficou simplório, apenas Desafio.

O projeto como citado, visa testar os conhecimentos back-end na linguagem GO, com a criação de uma API que faz uma leitura e analise com calculos dos dados de filmes.
---

## 🚀 Tecnologias Utilizadas

- Go
- Gin
- CSV
- JSON


## 🎯 Objetivo

O Codigo em questão, faz a leitura de um CSV utilizando o pacote csv padrão da linguagem go, assim então, serializa os dados e faz um marshal para uma struct onde contém todas as informações dos filmes (MOVIES).

Assim temos o controle dos dados em código para armazenar em disco, escolha pessoal pela rapidez de desenvolvimento e a diminuição de tecnologias.


Com os dados salvos em disco, podem fazer todas as analises e calculos solicitados.

## Estrutura e dificuldades

A pasta contendo a base de dados deve ser informada via linha de comando no argumento de inicialização após go run . [caminho do arquivo csv], no caso da estrutura, como esta, ja pronta a pasta esta na raiz do projeto com o nome [movies.csv], após isso o codigo ira criar um arquivo json com os filmes para reutulizar, ficando salvo em disco.

JSON SCHEMA:
{
  "title": "Movie",
  "type": "object",
  "additionalProperties": false,
  "required": [
    "budget",
    "homepage",
    "id",
    "original_language",
    "overview",
    "popularity",
    "release_date",
    "revenue",
    "runtime",
    "title",
    "vote_average",
    "vote_count",
    "genre",
    "production_company",
    "production_country"
  ],
  "properties": {
    "budget": {
      "type": "integer",
      "minimum": 0
    },
    "homepage": {
      "type": "string",
      "format": "uri"
    },
    "id": {
      "type": "integer",
      "minimum": 1
    },
    "original_language": {
      "type": "string",
      "minLength": 2,
      "maxLength": 5
    },
    "overview": {
      "type": "string",
      "minLength": 1
    },
    "popularity": {
      "type": "number",
      "minimum": 0
    },
    "release_date": {
      "type": "string",
      "format": "date"
    },
    "revenue": {
      "type": "integer",
      "minimum": 0
    },
    "runtime": {
      "type": "integer",
      "minimum": 1
    },
    "title": {
      "type": "string",
      "minLength": 1
    },
    "vote_average": {
      "type": "number",
      "minimum": 0,
      "maximum": 10
    },
    "vote_count": {
      "type": "integer",
      "minimum": 0
    },
    "genre": {
      "type": "string",
      "minLength": 1
    },
    "production_company": {
      "type": "string",
      "minLength": 1
    },
    "production_country": {
      "type": "string",
      "minLength": 1
    }
  }
}

Build da imagem => docker build -t desafio-api-movies .

docker run -p 8080:8080 desafio-api-movies

A aplicação estará disponível => http://localhost:8080

A API disponibiliza via servidor GIN os seguintes end points
/movies => Listagem completa dos filmes
"/movies/:id Lista filmes individuais com metricas adicionais, como Retorno de investimento e lucro
"/analytics/dashboard => Analise completa e soma de dados da base
"/analytics/top-studios => Calculos e divisões por studios
"/analytics/genre-stats => Calculos e divisões por genero


As divisões de pastas foram feitas seguuindo a logica de divisão por estrutuas (strucs a serem usadas) => models.go
, funções handlers para cada endpoint => handler.go
, funções utilitarios calculos, paginação e filtragem e etc. = util.go

As dificuldades encontradas ficaram na criação de testes unitários, falta de criatividade e especifidade para abraanger todos os casos de testes. Ainda nos casos de testes não consegui rodar todos de uma vez como a IDE do vs-code disponibiliza via botão no arquivo de teste, outra dificuldade pelo tempo sem desenvolver foi a escolha dos tipos de dados paras os calculos pois alguns atributos eram de tipo int e na hora da conversão precisavam ser convertidos para float64, o que gerou um pouco de conflito entre atributos, junto com o raciocinio da logica de calculos de ponto flutuante.

