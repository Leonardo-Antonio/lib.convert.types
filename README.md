# Libreria de conversion de tipos de datos

manera de convertir tipos de dato mas sencilla que el reflect

- Install
```shell
go get github.com/Leonardo-Antonio/lib.convert.types
```

- Ejemplo: 


Con Reflect
```go
var value interface{} = "hola"
valueConvertido := reflect.ValueOf(value).String()
```

Con la lib.convert.types
```go
var value interface{} = "hola"
valueConvertido, _ := convert.ToString(value)
```
