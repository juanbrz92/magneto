# magneto

Se detalla a continuación los rquisitos para ejecutar el programa localmente. Por otro lado mas abajo se detalla el acceso a la api directamente.

Requisitos:

    Tener instalado go, ver en: https://golang.org/doc/install
    Tener un servicio de base de datos mysql en el puerto 3306
    Crear una base de datos:
    
        CREATE DATABASE `magneto`
    
    Crear la tabla donde se alamacenaran los datos:
    
        CREATE TABLE `sequence` (
        `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
        `dna` varchar(500) NOT NULL,
        `result` varchar(45) DEFAULT NULL,
        PRIMARY KEY (`id`),
        UNIQUE KEY `sequence_dna_idx` (`dna`)
        ) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=latin1
        
Dependencias:

    go get github.com/gorilla/mux

    go get github.com/go-sql-driver/mysql
    
Configuraciones extras:

Modificar el archivo /src/magneto/api.go las lineas, definir el usuario, contraseña de la base de datos y puerto desde donde la api escucha las peticiones.

    var host = "http://localhost"
    var port = "12345"
    var connectionString = "usuario:contraseña@tcp(127.0.0.1:3306)/magneto?charset=utf8&parseTime=True&loc=Local"

Para ejecutar ir a src/magneto y ejecutar cualquiera de los siguientes comandos
    
    go run api.go isMutant.go sequence.go stat.go
    
    go build y luego ./mangeto
    
Para realizar test ir a src/magneto

    go test -v -cover

Del cual se obtiene

    ...
    ...
    PASS
    coverage: 87.4% of statements
    ok      magneto 0.610s
    

Acceso a la api (En GCP): 


Para agregar/checkear secuencias de adn:  http://35.198.62.59:23450/api/mutant

  
      Basic Auth:
        Username: meli
        Password: 1234

      Ejemplo:
                  curl --request POST \
                  --url http://35.198.62.59:23450/api/mutant \
                  --header 'Accept: */*' \
                  --header 'Accept-Encoding: gzip, deflate' \
                  --header 'Authorization: Basic bWVsaToxMjM0' \
                  --header 'Cache-Control: no-cache' \
                  --header 'Connection: keep-alive' \
                  --header 'Content-Length: 97' \
                  --header 'Content-Type: application/json' \
                  --header 'Host: 35.198.62.59:23450' \
                  --header 'Postman-Token: 02fd5ea9-a9b4-4b5a-abd6-eb8b1003efa1,54e29b28-014e-489a-9071-4ea81f5de222' \
                  --header 'User-Agent: PostmanRuntime/7.19.0' \
                  --header 'cache-control: no-cache' \
                  --data '{"dna":["ACATACCA","GTGTTACA","GGCAAgAG","ATGGGCTC","GTGCCGTA","AAGAGGAG","ATGATGGG","TCCTTCCT"]}'
  
        
Para consultar los registros actuales:  http://35.198.62.59:23450/api/stats
    
    Basic Auth:
        Username: meli
        Password: 1234
    
    Ejemplo:
    
                curl -X GET \
                http://35.198.62.59:23450/api/stats \
                -H 'Authorization: Basic bWVsaToxMjM0' \
                -H 'Postman-Token: f87360b3-96a5-4b56-bcd8-738cd5d9530e' \
                -H 'cache-control: no-cache'
                
Tambien se adjuntó un archivo json, con la collection de Postman
