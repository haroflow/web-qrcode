# web-qrcode

Simple web application to generate QR codes. See the [running example](https://hqrcode.herokuapp.com/)

Powered by:
- Go 1.16
- [fiber](https://github.com/gofiber/fiber)
- [skip2/go-qrcode](https://github.com/skip2/go-qrcode)

# Run locally

By default it runs on port 8080:
```
git clone https://github.com/haroflow/web-qrcode.git
cd web-qrcode
go run .
```
Open browser on localhost:8080.

If you need to change the port, heroku uses a PORT environment variable:
```
PORT=8888 go run .
```
