package assert

import (
	"ec-tsj/core"
	"fmt"
	"os"
)

type (
	T                = core.T
	fnAssertCallback func(bool, ...string)
	fnTypeCallback   func() bool
)

//!+
/**
 * Realiza una aserción doble. Se realiza con dos expresiones Go, si las expresiones
 * se evalúan a True, no hace nada. Si se evalúan a False,  envía un mensaje de error
 * y aborta la ejecución del programa. Evalúa las dos expresiones, si una cualquiera
 * de las dos, da False ó True, eso es lo que será.
 * @param {bool} Es una expresión golang
 * @param {bool} Expresion golang
 * @param {...string}
 */
func AssertDouble(f1 bool, f2 bool, msg ...string) {
	if !f1 || !f2 {
		_fnExit("AssertDouble", msg)
	}
}

//!-

//!+
/**
 * Realiza una aserción. Se realiza con una expresión Go, si la expresión
 * se evalúa a True, no hace nada. Si se evalúa a False envía un mensaje de
 * error y aborta la ejecución del programa.
 * @param {bool} Es una expresión golang
 * @param {...string}
 */
func Assert(f bool, msg ...string) {
	if !f {
		_fnExit("Assert", msg)
	}
}

/**
 * Realiza una comprobación de tipo. Si la función devuelve True no hace nada, y si False
 * envía el mensaje y para la ejecución del programa.
 * <CODE>
 *    ...
 *    AssertType(func() bool {return who.Kind(element) == reflect.Kind()}, "cadena no string")
 *    ...
 * </CODE>
 * @param {fnTypeCallback} valor a comprobar
 * @param {...string} mensaje opcional
 */
func AssertType(fn fnTypeCallback, msg ...string) {
	if !fn() {
		_fnExit("AssertType", msg)
	}
}

/**
 * Realiza una aserción. Se realiza con una expresión Go, si la expresión
 * se evalúa a True, no hace nada. Si se evalúa a False envía un mensaje de
 * error y aborta la ejecución del programa.
 * @param {bool} Es una expresión golang
 * @param {...string}
 */
var AssertTrue fnAssertCallback = Assert

// helper' Assert and NotAssert fn's
var _fnExit = func(s string, msg ...T) {
	opt := core.ArgOptional(core.Literals().NullString, msg).([]string)

	fmt.Fprintf(os.Stderr, "%s%s. %s\n", "Error: Aserción fallida: ", s,
		core.IIf(len(opt) == 0, core.Literals().NullString, opt))
	os.Exit(-1)
}

/**
 * Realiza una No Aserción. Se realiza con una expresión Go, si la expresión
 * se evalúa a False, no hace nada. Si se evalúa a True envía un mensaje de
 * error y aborta la ejecución del programa.
 * @param {bool} Es una expresión golang
 * @param {...string}
 */
func NotAssert(f bool, msg ...string) {
	if f {
		_fnExit("NotAssert", msg)
	}
}

/**
 * Realiza una No Aserción. Se realiza con una expresión Go, si la expresión
 * se evalúa a False, no hace nada. Si se evalúa a True envía un mensaje de
 * error y aborta la ejecución del programa.
 * @param {bool} Es una expresión golang
 * @param {...string}
 */
var AssertFalse fnAssertCallback = NotAssert

//!-
//!+
/**
 * Realiza una NO aserción doble. Se realiza con dos expresiones Go, si las expresiones
 * se evalúan a False, no hace nada. Si se evalúan a True envía un mensaje de
 * error y aborta la ejecución del programa. Evalúa las dos expresiones,
 * Si una cualquiera de las dos, da False ó True, eso es lo que será.
 * @param {bool} Es una expresión golang
 * @param {bool} Expresion golang
 * @param {...string}
 */
func NotAssertDouble(f1 bool, f2 bool, msg ...string) {
	if f1 || f2 {
		_fnExit("NotAssertDouble", msg)
	}
}

//!-
/**
 * Realiza una comprobación de tipo. Si la función devuelve True no hace nada, y si False
 * envía el mensaje y para la ejecución del programa.
 * <CODE>
 *    ...
 *    NotAssertType(func() bool {return who.Kind(element) != reflect.Kind()}, "cadena no string")
 *    ...
 * </CODE>
 * @param {fnTypeCallback}
 * @param {...string} mensaje
 */
func NotAssertType(fn fnTypeCallback, msg ...string) {
	if fn() {
		_fnExit("NotAssertType", msg)
	}
}

//!+
/**
 * Convierte una expresión a bool
 * @param {T}
 * @return {bool}
 */
func CvtToBool(i T) bool {
	switch i.(type) {
	case nil:
		return false
	case string:
		c := i.(string)
		return c != ""
	case int, int32, int64, int8, int16, uint, uint32, uint64, uint8, uint16, uintptr:
		c := i.(int)
		return c != 0
	case float64, float32:
		c := i.(float64)
		return c != 0.00
	case complex64, complex128:
		c := i.(complex128)
		return c != 0.00+0i
	}

	return false
}

//!-
