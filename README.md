# Delivery Much Challenge
==============================================

Desafio de processo seletivo da Delivery Much

## Pré-requisitos
* [Git][1]
* [Docker][2]
* [Docker-Compose][3]

## Instalação

1. Dê um git clone no repositório. Exemplo: https://github.com/diniz-victor/deliverymuch-challenge.git
2. Vá até o diretório do projeto clonado e execute:
3. `make run`
* Caso ocorra o erro no módulo mux dizendo que ele não está explícito no vendor/modules.txt, adicione `## explicit`no vendor/modules.txt:
```
# github.com/gorilla/mux v1.8.0
## explicit
github.com/gorilla/mux
```
4. Verifique se a aplicação está funcionando para a rota GET localhost:8080/recipes?i=potato,tomato


Keep It Simple, Enjoy! - By Victor Diniz

[1]: https://git-scm.com/downloads
[2]: https://docs.docker.com/get-docker/
[3]: https://docs.docker.com/compose/install/