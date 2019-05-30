# Simple Tax-Calculator

### How to Run
To run this project you may use below command:
```php
docker-compose up --build
```

Above command will build Dockerfile inside `config/`:
- MySQL docker
- Golang docker

Wait until golang docker finish: 
- Download all dependencies 
- Run database schema migration 
- Data seeder

And shown below output on terminal:

```bash 
golang_service    | [GIN-debug] POST   /v1/taxes                 --> _/my_app/controllers.(*V1TaxesController).CalculateTax-fm (4 handlers)
golang_service    | [GIN-debug] POST   /v1/taxes/bulk            --> _/my_app/controllers.(*V1TaxesController).CalculateTaxBulk-fm (4 handlers)
golang_service    | 0.0.0.0:3000
golang_service    | [GIN-debug] Listening and serving HTTP on 0.0.0.0:3000
```

Project URL should be: localhost:3000

### Database Structure

Inside `src/models` is represent database structure and auto migration by gorm

### Endpoint Documentation

Endpoint documentation can be accessed via this url:
##### https://documenter.getpostman.com/view/883805/RztitAQK

