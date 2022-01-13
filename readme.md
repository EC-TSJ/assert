# Cli

[![Home](https://godoc.org/github.com/gookit/event?status.svg)](file:///D:/EC-TSJ/Documents/CODE/SOURCE/Go/pkg/lib/cli)
[![Build Status](https://travis-ci.org/gookit/event.svg?branch=master)](https://travis-ci.org/)
[![Coverage Status](https://coveralls.io/repos/github/gookit/event/badge.svg?branch=master)](https://coveralls.io/github/)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/event)](https://goreportcard.com/report/github.com/)

> **[EN README](README.md)**

Assert para realizar las aserciones de programación.

## GoDoc

- [godoc for github](https://godoc.org/github.com/)

## Menú Principal
--- 



Se configura un menú con ese ejemplo como diseño básico. Con las reglas siguientes:


> Funciones:
  - *`Assert(bool, ...string)`*
  - *`AssertDouble(bool, bool, ...string)`*
  - *`AssertType(fnTypeCallback, ...string)`*
  - *`NotAssert(bool, ...string)`*
  - *`NotAssertDouble(bool, bool, ...string)`*
  - *`NotAssertType(fnTypeCallback, ...string)`*
  - *`CvtExpr(T) bool`*

> Tipos:
  - *fnTypeCallback* *`func(bool, ...string)`*
  - *fnAssertCallback* *`func() bool`*

## Ejemplos
```go


 	flag := CvtExpr("true")
	flag = CvtExpr("false")
	flag = CvtExpr("truef")
	flag = CvtExpr(21)
	flag = CvtExpr(0)
	flag = CvtExpr(5.45)
	flag = CvtExpr(0.00)

	Assert("Hola" == "Hola")
	NotAssert("es cadena " == cli.NullString)
	Assert("es cadena " == cli.NullString) // se sale
	NotAssert("Hola" == "Hola") // se sale

  
## Notas

Se pueden eliminar las opciones de Autor, Versión ó Descripción del menú de opciones. Para eso se pondrá en la definición de setting y en los valores susodichos en la forma siguiente:

...   

"author": "Torres Sacristán, Jesús 2020\"<0",

...

es decir, se interpondrá los caracteres \"<0 entre medias y al final de la definición




<!-- - [gookit/ini](https://github.com/gookit/ini) INI配置读取管理，支持多文件加载，数据覆盖合并, 解析ENV变量, 解析变量引用
-->
## LICENSE

**[MIT](LICENSE)**
