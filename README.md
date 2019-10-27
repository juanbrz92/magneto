# magneto

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

Para ejecutar ir a src/magneto y ejecutar cualquiera de los siguientes comandos
    
    go run api.go isMutant.go sequence.go stat.go
    
    go build y luego ./mangeto
    
Para realizar test ir a src/magneto

    go test -v -cover
    

Accesso a la api (En GCP): 


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
                  --header 'Content-Length: 100' \
                  --header 'Content-Type: application/json' \
                  --header 'Host: 35.198.62.59:23450' \
                  --header 'Postman-Token: 8a43c7d2-dbb7-45ef-b438-0453df138d79,59ea29e3-1c48-428d-9071-f2dc7bdefaf6' \
                  --header 'User-Agent: PostmanRuntime/7.19.0' \
                  --header 'cache-control: no-cache' \
                  --data '{\n	"dna":["ACATACCA","GTGTTACA","GGCAAgAG","ATGGGCTC","GTGCCGTA","AAGAGGAG","ATGATGGG","TCCTTCCT"]\n}'
  
        
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
