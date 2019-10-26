# magneto


Dependencias:

    go get github.com/gorilla/mux

    go get github.com/go-sql-driver/mysql

Accesso a la api: 


Para agregar secuencias de adn:

  http://35.198.62.59:12345/api/mutant

      Ejemplo:
                curl -X POST \
                http://35.198.62.59:12345/api/mutant \
                -H 'Accept: */*' \
                -H 'Accept-Encoding: gzip, deflate' \
                -H 'Cache-Control: no-cache' \
                -H 'Connection: keep-alive' \
                -H 'Content-Length: 144' \
                -H 'Content-Type: application/json' \
                -H 'Host: 35.198.62.59:12345' \
                -H 'Postman-Token: 73c4819d-b13f-4e43-8b3a-22d7abcddc43,288f1137-91f9-4472-a530-ae9758b3a1b6' \
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

  http://35.198.62.59:12345/api/stats