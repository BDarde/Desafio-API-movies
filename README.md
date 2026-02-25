<h2>desafio-api-movies</h2>

---

<h3>Tecnologias Utilizadas</h3>
<ul>
  <li>Go</li>
  <li>Gin</li>
  <li>CSV</li>
  <li>JSON</li>
  <li>YML</li>
  <li>Swagger</li>
</ul>

<h3>Objetivo</h3>
<p>A API tem como objetivo listar filmes e calcular suas métricas, alimentando-se de uma planilha fixa CSV que é serializada e armazenada em disco para otimização de performance.</p>

<h3>Estrutura de Pastas</h3>
<pre><code>
.
├── .vscode/
│   └── launch.json        # Configurações de debug do VS Code
├── Docs/
│   └── swagger.yaml       # Especificação OpenAPI/Swagger
├── Dockerfile             # Configuração para containerização
├── dockerignore           # Arquivos ignorados pelo Docker
├── go.mod                 # Definição do módulo Go
├── go.sum                 # Checksums de dependências
├── handlers.go            # Lógica dos endpoints da API
├── main.go                # Ponto de entrada (Server & Router)
├── main_test.go           # Testes principais
├── models.go              # Estruturas de dados (Structs)
├── movies.csv             # Base de dados bruta (Input)
├── README.md              # Documentação do projeto
├── utils.go               # Funções de cálculo e utilitários
└── utils_test.go          # Testes das funções utilitárias
</code></pre>

<h3>JSON Schema dos filmes em disco</h3>
<p>Estrutura do objeto após o processamento e armazenamento:</p>

```json
{
  "title": "Movie",
  "type": "object",
  "additionalProperties": false,
  "required": [
    "budget", "homepage", "id", "original_language", "overview",
    "popularity", "release_date", "revenue", "runtime", "title",
    "vote_average", "vote_count", "genre", "production_company",
    "production_country"
  ],
  "properties": {
    "budget": { "type": "integer", "minimum": 0 },
    "homepage": { "type": "string", "format": "uri" },
    "id": { "type": "integer", "minimum": 1 },
    "original_language": { "type": "string", "minLength": 2, "maxLength": 5 },
    "overview": { "type": "string", "minLength": 1 },
    "popularity": { "type": "number", "minimum": 0 },
    "release_date": { "type": "string", "format": "date" },
    "revenue": { "type": "integer", "minimum": 0 },
    "runtime": { "type": "integer", "minimum": 1 },
    "title": { "type": "string", "minLength": 1 },
    "vote_average": { "type": "number", "minimum": 0, "maximum": 10 },
    "vote_count": { "type": "integer", "minimum": 0 },
    "genre": { "type": "string", "minLength": 1 },
    "production_company": { "type": "string", "minLength": 1 },
    "production_country": { "type": "string", "minLength": 1 }
  }
}


<h3>Fluxo e Decisões Técnicas</h3>

<p><strong>Inicialização:</strong> O caminho (path) do CSV de dados é passado via argumento de linha de comando na inicialização:</p>

Powershell cmd
##criação do container
docker build -t desafio-api-movies .

##inicialização com volume do arquivo
```text
  docker run -p 8080:8080 `
  -v "${PWD}\movies.csv:/app/movies.csv" `
  desafio-api-movies `
  ./app /app/movies.csv

<p>Aplicação disponível em: <code>http://localhost:8080</code></p>


go run . [caminho_do_arquivo]
<p><strong>Estratégia de Cache:</strong> O código verifica a existência do JSON em disco. Caso não exista, ele processa o input CSV original e realiza a persistência, evitando redundância de processamento e prevenindo erros de arquivo inexistente.</p>

<p><strong>Arquitetura:</strong> O arquivo <code>main.go</code> foi criado para conter apenas a lógica de inicialização e rotas, mantendo o código limpo e com dividido conforme suas responsabilidades.</p>

<p><strong>Desacoplamento:</strong> Optou-se por criar um objeto que representa o estado dos filmes e utilizar métodos <i>handlers</i> injetados para diminuir o acoplamento de objetos globais e facilitar a execução de testes automatizados.</p>

<h3>Endpoints Disponíveis</h3>
<ul>
<li><code>GET /movies</code>: Listagem completa dos filmes.</li>
<li><code>GET /movies/:id</code>: Detalhes individuais com métricas de lucro e ROI.</li>
<li><code>GET /analytics/dashboard</code>: Análise consolidada da base de dados.</li>
<li><code>GET /analytics/top-studios</code>: Cálculos agrupados por estúdio de produção.</li>
<li><code>GET /analytics/genre-stats</code>: Cálculos agrupados por gênero cinematográfico.</li>
</ul>

<h3>Inicialização via Docker</h3>
<p>Para build e execução do ambiente containerizado:</p>



<h3>Dificuldades e Aprendizados</h3>
<ul>
  <li><strong>Testes Unitários:</strong> Desafios na implementação de cenários criativos e dificuldades técnicas na integração com o runner nativo do VS Code.</li>
  <li><strong>Gerenciamento de Tipos:</strong> Complexidade no tratamento de conversões de tipos (<code>int</code> para <code>float64</code>), exigindo atenção rigorosa para garantir a precisão dos cálculos financeiros.</li>
  <li><strong>Testes de Integração:</strong> Dificuldade em simular o fluxo completo das funções, demandando a criação de <code>main</code> individuais para validar a orquestração do sistema.</li>
  <li><strong>Cálculos de Ponto Flutuante:</strong> O raciocínio lógico para operações matemáticas com decimais gerou conflitos entre atributos, exigindo refatoração na estrutura de utilitários.</li>
 <li><strong>Documentação Swagger:</strong> A Criação da documentação swagger demandou aprendizado pois não tinha utilizada até então</li>
 <li><strong>Docker<strong>A criação do docker para subir o projeto além do volume no container<li>
</ul>