Version de go: ```1.23.0```

Crear el .env en la carpeta ```cmd/api-test``` con las variables para conectarse a una base de datos mysql:

``` 
DB_NAME=go-test
DB_USER=root
DB_PASSWORD=101547
```


con ```go run .``` parado en  ```cmd/api-test``` levanta el proyecto



### Seed

- Con el endpoint ```/seed-users``` se generar dos usuarios base de ejemplo solo la primera vez
- Con el endpoint ```/users``` se listan todos
- Con el endpoint ```/users/{id}``` se trae by id