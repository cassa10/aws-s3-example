# AWS s3 golang

## Requisitos:

- Go 1.20
- Credenciales AWS
- Haber creado un bucket de AWS S3


## Ejecución

1) Setear las siguientes envs

```
AWS_BUCKET=<insert-bucket>
AWS_REGION=<insert-region>
AWS_SECRET_ID=<insert-secret_id>
AWS_SECRET_KEY=<insert-secret_key>
```

2) Utilizar o modificar las variables o argumentos de la funcion main:

```go
func main() {
	logScript(func() {
		awsRepo := setUpAwsRepo()
		
		// add any repo call with vars example:
		ListBucketObjets(awsRepo)
	})
}
```