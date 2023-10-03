# gohorloge

## Build

Pour Builder
```shell
set GOOS=linux
set GOARCH=arm
set GOARM=5
go build
```

## Installation du service

 Installation
```shell
sudo cp /usr/rep/horloge.service /etc/systemd/system/horloge.service
sudo systemctl enable horloge.service
sudo systemctl start horloge.service
```

Voir l'état
```shell
sudo systemctl status horloge.service
```

Consulter les logs
```shell
journalctl -u horloge.service
```
