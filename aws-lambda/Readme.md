AWS Lambda https://github.com/aws/aws-lambda-go

GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main

Start Routine auswählen (default ist "hello") und "main" eingeben