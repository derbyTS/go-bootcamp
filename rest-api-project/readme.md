### Find who runs on that port and how to kill the process

```sh
lsof -i :3000
kill -9 <pid>
```

### Create a local SSL/TLS

```sh
openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365
openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365 -subj "/C=PH/ST=Cavite/L=Bacoor/O=Dev/CN=localhost"
openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365 -config openssl.cnf

```
