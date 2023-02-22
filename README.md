# Bank Transaction
API Restful que permite usuário se cadastrar e realizar transações bancárias

## User Guide
<details>
    <summary>Roadmap do usuário</summary>

    - Usuário pode se cadastrar
    - Usuário pode efetuar login
    - Usuário logado pode efetuar transação
</details>

## Arquitetura & Decisões

- #### Web framework
    A escolha do framework `Fiber` se deve a sua simplicidade e features que facilitam o desenvolvimento.
    Sendo as principais: Fácil roteamento, Tratamento de erro, Validação de input. Entre outras(CORS, autenticação de endpoints, *middlewares*). É sabido que desenvolvedores `Go` não ~~costumam~~ utilizar *web frameworks*, porém o escolhido ajuda a economizar tempo e se encaixa bem aos requisitos da aplicação. Visto que o mesmo é simples, assim como o Express(NodeJS).

- #### Acompanhamento e Status
    O `Redis` está sendo utilizado para persistir *status* do processamento, que pode ser utilizado para o acompanhamento do processo.

- #### O *message broker*
    Devido a simplicidade dos requisitos para essa aplicação, o *message broker* pode se tornar uma ferramenta que acrescenta mais complexidade do que benefícios. Porém, o uso do mesmo está relacionado a manter uma única fila de transações e garantir uma por usuário, além de permitir o acompanhamento do processo em tempo real.

- #### `RabbitMQ`
    Além da familiaridade ferramenta, o `RabbitMQ` é a melhor opção para a aplicação por ser leve e robusto. A utilização do `Redis` como *message broker* poderia ser cogitada, porém a segurança e garantia de entrega do `RabbitMQ` supera a opção do banco de dados(cache).

- #### Banco de Dados
    O escopo bem definido e a simplicidade da API não necessitam de um banco de dados robusto, logo o `MySQL` atende bem as necessidades da aplicação.