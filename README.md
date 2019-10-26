# magneto

Para ejecutar ir a src/magneto y ejecutar el comando
    
    go run api.go isMutant.go sequence.go stat.go
    
 
Dependencias:

    go get github.com/gorilla/mux

    go get github.com/go-sql-driver/mysql

Accesso a la api: 


Para agregar secuencias de adn:

     http://35.198.62.59:23450/api/mutant
  
      Basic Auth:
        Username: meli
        Password: 1234

      Ejemplo:
                  curl -X POST \
                  http://localhost:12345/api/mutant \
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
  
        
Para consultar los registros actuales:

    http://35.198.62.59:23450/api/stats
