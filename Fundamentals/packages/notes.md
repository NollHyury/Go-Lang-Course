O que são pacotes em Go e como eles interagem entre eles?

Quando estamos lidando com pacotes em GO temos que criar um módulo, sendo assim quando estamos lidando com mais de um
pacote no mesmo projeto usamos um módulo;

o que é um módulo?
  Módulo é um conjunto de pacotes que compõem o projeto, tanto os pacotes que foram implementados pelo desenvolvedor, bem como
os pacotes externos;

Como criar um módulo?
    cli command: "go mod init <module-name>"
    o arquivo criado irá fazer um papel semelhante ao do packege.json do nodejs, ou seja ele ira centralizar as
dependencias da aplicação

Como fazer um build?
    cli command: "go build"
    este comando irá gerar um arquivo binário, .exe, este arquivo é o compilado de todo o projeto.

Importancia da letra maiuscula, para definir uma função publica o inicio da letra deve ser Maiúscula, funções com letra minúscula só estão visiveis dentro do pacote que ela está;

Dentro do próprio pacote é possivel referenciar sem importar!

O build não é feito automáticamente, você precisa executa-lo novamente sempre que quiser rodar o executável

Comando "go install" ira salvar o executavel dentro da pasta raiz do go lang

Para instalar pacotes externos?
    cli command: go get <package-link>
    ps: no link não pode conter https://  pois o char ":" não é aceito pela cli neste caso

Como remover os pacotes não utilizados
    cli command: go mod tidy, este command irá remover todas as os pacotes não utilizados
de dentro do arquivo go.mod