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
    

Accesso a la api (En GCP): 


Para agregar/checkear secuencias de adn:  http://35.198.62.59:23450/api/mutant

  
      Basic Auth:
        Username: meli
        Password: 1234

      Ejemplo:
                  curl -X POST \
                  http://35.198.62.59:23450/api/mutant \
                  -H 'Accept: */*' \
                  -H 'Accept-Encoding: gzip, deflate' \
                  -H 'Authorization: Basic bWVsaToxMjM0' \
                  -H 'Cache-Control: no-cache' \
                  -H 'Connection: keep-alive' \
                  -H 'Content-Length: 144' \
                  -H 'Content-Type: application/json' \
                  -H 'Host: localhost:12345' \
                  -H 'Postman-Token: 84b65496-7230-4974-9876-e82f1ab529e8,d30076d7-6cf1-4965-9eaa-e5ea36b08bde' \
                  -H 'User-Agent: PostmanRuntime/7.19.0' \
                  -H 'cache-control: no-cache' \
                  -d '{
                    "dna":["TCGAGCGTT",
                        "GTCATAGGG",
                        "CTTATATTT",
                        "GAAACTCAT",
                        "CCTACACAG",
                        "ACGCTAGGG",
                        "AGTCTCATC",
                        "TATTTAGAC",
                        "GGTTAAATG"]
                }'
  
        
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
