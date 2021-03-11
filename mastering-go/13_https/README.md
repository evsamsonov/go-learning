
## Генерация сертификата клиента
```bash
openssl req -x509 -nodes -newkey rsa:2048 -keyout client.key -out client.crt -days 365 -subj "/"
```

## Генерация TLS сертификата сервера
```bash
openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```
