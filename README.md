# Progetto SDCC: Algoritmi di Elezione Distribuita
## Introduzione:
Nei contesti degli algoritmi dei sistemi distribuiti, si manifesta frequentemente l'esigenza di un processo atto a fungere da coordinatore o leader. Tale necessità richiede lo sviluppo di algoritmi che, coinvolgendo tutti i processi attivi del sistema, possano coordinarli efficacemente e procedere all'elezione di un unico coordinatore. L'utilità di tali algoritmi diventa evidente se si pensa che un processo di coordinazione potrebbe andare in crash in qualsiasi istante e, in assenza di tali algoritmi, il sistema si troverebbe senza un leader che lo coordini.
### Assunzioni per il funzionamento degli algoritmi:
Nel progetto, consideriamo un sistema distribuito così fatto:
- n processi che comunicano tra loro
- comunicazione affidabile: i messaggi inviati tra i processi non vengono persi, duplicati o consegnati in ordine errato
- processi soggetti a fallimenti
- Ogni processo ha un Id unico e il processo non guasto con l'Id più alto è il processo che dovrà essere eletto come leader.
Inoltre, supponiamo che ogni processo abbia una conoscenza limitata dell'ambiente, in particolare riguardo allo stato degli altri processi. I processi possono quindi rilevare il fallimento di altri processi solo attraverso il mancato ricevimento di messaggi o attraverso segnalazioni esplicite di fallimento.
### Algoritmi di Elezione distribuita:
Nel seguente progetto sono stati implementati due algoritmi di elezione distribuita:
- Agoritmo di Bully;
- Algoritmo di Dolev, Klawe and Rodeh.
## Prequisiti e istruzioni per eseguire il codice
### Prequisiti:
- go 1.21.6, [releases of Go](https://go.dev/doc/devel/release)
- Protocol buffer compiler. Per l'installazione, vedere [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/).
- Go plugins per il protocol compiler
  1. Per l'installazione dei plugins, seguire il seguente codice:
    ```bash
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```
  2. Aggiornare il <code> PATH </code> il modo che <code>protoc</code> possa trovare i plugins:
    ```bash
    $ export PATH="$PATH:$(go env GOPATH)/bin"
    ```
- Docker 25.03 [get Docker](https://docs.docker.com/get-docker/)
- Docker Compose v2.24.6-desktop.1 [docker compose releases](https://github.com/docker/compose/releases/)
## Scelta dell'algoritmo:
Per la scelta dell'algoritmo è necessario andare a settare i valori true e false nel file config.json:
```
"algorithm": {
    "Bully": true,
    "DKR": false
}
```
## Esecuzione del codice senza istanza AWS:
Se si vuole eseguire l'applicazione senza un'istanza AWS:
* Clono la repository da git:
    ```
    git clone https://github.com/Lollogiga/ProgettoSDCC
    ```
* Avvio i container:
    ```
    cd .\codice\
   docker-compose -f compose.yaml up
    ```
## Esecuzione del codice su istanza AWS:
Nel caso in cui si voglia eseguire il codice su un'instanza EC2, bisogna:
### Connettersi all'istanza EC2:
Mi connetto all'istanza EC2 tramite ssh sulla PowerShell:
```
ssh -i <path_to_PEM> ec2-user@<ip-EC2-instance>
```
### Sull'istanza EC2:
* Installo docker:
    ```
    sudo yum update -y
    sudo yum install -y docker
    ```
* Installo docker-compose:
    ```
    sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose
    ```
* Eseguo il docker deamon:
    ```
    sudo service docker start
    ```
* Installo git e clono la repository:
    ```
    sudo yum install git -y
    git clone https://github.com/Lollogiga/ProgettoSDCC
    ```
* Eseguo docker compose:
    ```
    cd ProgettoSDCC/codice/
    sudo docker-compose -f compose.yaml up
    ```
## Verifica funzionamento programma:
Al fine di verificare il corretto funzionamento sull'istanza EC2 è possibile:
* Visualizzare l'insieme di container:
    ```
    sudo docker ps
    ```
* Stoppare un container:
    ```
  sudo docker kill <Container ID>
    ```
* Restart di un container:
    ```
  sudo docker restart <Container ID>
    ```
* Rimozione di tutti i container nello stato "Exited":
  ```
  docker rm $(docker ps --filter status=exited -q)
  ```

