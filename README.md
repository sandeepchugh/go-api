# go-api
Demo api using golang

## Database

Pull the db image from docker hub
```shell
docker pull mariadb
```

Running the db server
```shell
docker run --detach --name customers --env MARIADB_USER=appuser --env MARIADB_PASSWORD=mysecretpassword! --env MARIADB_ROOT_PASSWORD=mysecretrootpassword!  mariadb:latest
```
