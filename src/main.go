package main

import (
	"fmt"
	"log"
	"strconv"
)

func funcionUno(a,b int,mensaje string){//los parametros a y b son del tipo int mensaje es del tipo int
	fmt.Printf("el mensaje recibido fue: %s\n%d\n%d\n",mensaje,a,b)
}

func retonarValor(a,b int) int{//la funcion recibe 2 parametros y retona un resultado
	return a+b
}
func dobleRetorno(c,d int)(h,g int){//se reciben 2 enteros y se regresan 2 enteros
	return c*3,d*5
}
func main (){
	//declaracion de constantes
	const nombre float64 = 3.14	
	const nombre2 =3.1416
	//zero values : valores por defecto que tendran las variables al ser declaradas
	var a int //valor por defecto 0
	var b float64 // valor por defecto 0.000
	var c string // valor por defecto " "
	var d bool // valor por defecto False
	base :=12 //los dos puntos indican que la variable no fue creada anteriormente
	var altura int=14 //declarar variable y define tipo de dato inicializandola
	var area int //se declara la variable pero no se incializa
	fmt.Println(nombre)
	fmt.Println(nombre2)  
	fmt.Println(base)
	fmt.Println(altura)
	fmt.Println(area)
	fmt.Printf("variable a: %d\nvariable b:%f\nvariable c:%s\nvariable d: %v\n",a,b,c,d)
	var mensaje string = fmt.Sprintf("mensaje a guardar")
	fmt.Println(mensaje)	
	fmt.Printf("El tipo de variable de mensaje es: %T\n",mensaje)//imprimir tipo de variable
	funcionUno(5,7,"prueba funcion")
	valorRetorno:=retonarValor(7,8)
	fmt.Println(valorRetorno)
	r1,r2:=dobleRetorno(3,4)
	fmt.Printf("valor r1: %d\nvalor r2: %d\n",r1,r2)
	_,r3:=dobleRetorno(3,4)//la funcion retorna 2 valores pero solo se tiene en cuenta uno
	fmt.Printf("valor de r3 sin tomar en cuenta el valor de la primera variable: %d",r3)
	for w:=0;w<10;w++{//bluce for 
		fmt.Println(w)
	}
	//for while
	conteo:=0
	for conteo<10{
		fmt.Println(conteo)
		conteo++
	}
	//conteo infinito si no se tiene un condicional en el for continuara la ejecucion del codigo de forma infinita
	/*infinito:=0
	for{
		fmt.Printf("mensaje %d",infinito)
		infinito++
	}*/

	if conteo==10{
		fmt.Println("valor igual a 10")
	}else{
		fmt.Println("valor diferente de 10")
	}
	//condicion AND
	if r1==10 && r2==11{
		fmt.Println("los valores son iguales")
	}else{
		fmt.Println("los valores son diferentes")
	}
	//condicion OR
	if r1==9 || r2==11{
		fmt.Println("Algun valor es igual")
	}else{
		fmt.Println("los valores son diferentes")
	}

	valor, err :=strconv.Atoi("3")//convierte un string en entero

	if err !=nil{//nil -> no initialized or not error
		log.Fatal(err)//para la ejecucion del programa y muestra un mensaje en la consola
	}
	fmt.Println(valor)

	//modulo:=5%2
	//switch modulo{
	//la estructura superior e inferior son iguales	
	//switch modulo:=5%2; modulo{
	modulo:=5%2
	switch {
	//las estructuras superiores son iguales	
	case modulo==0:
		fmt.Println("numero es par")
	break
	default:
		fmt.Println("numero es impar")
	break
	}
}