# API-PERSONAL-FINANCE
### Api criada para fins de estudo e prática

### Descrição
api voltada para controle de movimentações financeiras cotidianas, como entradas e saídas(gastos e ganhos)

### Contexto de negócio
```
Docker
MySql
Golang
Fiber
JWT token
```


## Executar localmente
na primeira vez, use:
```
# docker-compose up -d --build
```


se faltar alguma dependência use o compose install dentro do container
```
# docker exec setup-php composer install
```

## Autenticação
```
Bearer Token: 
```


##  Rotas não autenticadas
```
POST: http://localhost:8080/api/v1/login
POST: http://localhost:8080/api/v1/user
```


#### endereços para acesso:
```
PhpMyAdmin: http://localhost:8888/
Endereço da api: http://localhost:8080/api/v1/
```

Developed by Neilton Rodrigues