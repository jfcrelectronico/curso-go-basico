package main

import (
	"fmt"
	"sync"
	"time"
	//"log"
	//pk "primerProyecto/src/misPaquetes" //pk es el alias para la ruta del paquete misPaquetes
	//"strconv"
	//"strings"
)

//structs son plantillas que permiten la creacion de un objeto similar a como trabaja un clase

//type-> tipo de dato
//carro-> nombre para la estructura
//struct-> indica a go que se va a desarrollar un estructura
type carro struct{
	marca string
	modelo int
}

func funcionUno(a,b int,mensaje string){//los parametros a y b son del tipo int mensaje es del tipo int
	fmt.Printf("el mensaje recibido fue: %s\n%d\n%d\n",mensaje,a,b)
}

func retonarValor(a,b int) int{//la funcion recibe 2 parametros y retona un resultado
	return a+b
}
func dobleRetorno(c,d int)(h,g int){//se reciben 2 enteros y se regresan 2 enteros
	return c*3,d*5
}

func ejemploContinue(){
	for x:=0;x<5;x++{
		if x==3{
			fmt.Println("Valor x=3")
			continue//al encontrar la palabra reservada continue se finaliza la iteracion y las demas lineas de codigo
					//que esten luego de esta NO seran ejecutadas
		}
		fmt.Printf("iteracion %d:\n",x)
		fmt.Printf("iteracion %d:\n",x+1)
		fmt.Printf("iteracion %d:\n",x+2)
	}
}
//    nombre_de_objeto_creado tipo_estructura_objeto_creado  nombre_deseado_funcion_objeto_creado
func (carro2 carro) miCarroFuncion(){
	fmt.Println(carro2.marca," mensaje funcion")
}
// nombre_de_objeto_creado apuntador_valores_objeto_creado_por_tipo  nombre_deseado_funcion_apuntador
func (carro2 *carro) apuntadorFuncion(){
	carro2.marca+="2"
}

//redefine el metodo str que retorna el objeto carro2, permitiendo hacer mas legible la informacion
func (carro2 carro) String() string {
	return fmt.Sprintf("Marca del carro: %s,y modelo: %d",carro2.marca,carro2.modelo)

}


type cuadrado struct{
	lado float64
}
type rectangulo struct{
	base float64
	altura float64
}
type calcular interface{
	area() float64
}
func (c cuadrado) area() float64{//funcion para calcular el area del cuadrado
	return c.lado*c.lado
}
func (r rectangulo) area() float64{//funcion para calcular el area del rectangulo
	return r.base*r.altura
}
// los dos metodos anteriores se denominan area pero cada uno esta aplicado a un figura diferente pero permite, usar
//una unica interface para su llamado, dicha interface se denomino calcular y tiene como unico parametro area(), que 
//es una funcion comun entre las dos figuras
func calcularArea (f calcular){//la funcion se define con una interface denominada calcular 
	fmt.Println("Area: ",f.area())
}





/*
func main (){

	//lista de interfaces
	//permite asignarle a una lista diferentes tipos de elementos como lo hace python
	multi_Variables:=[]interface{}{"jony",10,5.0,true,10+5i}
	fmt.Printf("%v,\n tipo 1: %T\n tipo 2: %T\n tipo 3: %T\n tipo 4: %T\n tipo 5: %T",
	multi_Variables,multi_Variables[0],multi_Variables[1],multi_Variables[2],multi_Variables[3],multi_Variables[4])


	cuadro:=cuadrado{lado:5}//instanciar un objeto del tipo cuadrado
	rect:=rectangulo{base:6,altura:7}//instanciar un objeto del tipo rectangulo

	calcularArea(cuadro)//llama a la funcion calcularArea que recibe el objeto que cuenta con un metodo area
	calcularArea(rect)

	
	/*
	defer fmt.Println("finaliza programa")//la palabra reservada defer indica que esta linea de codigo sera
										  //la ultima que se ejecute antes de cerrar finalizar la funcion
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
	*/
	//conteo infinito si no se tiene un condicional en el for continuara la ejecucion del codigo de forma infinita
	/*infinito:=0
	for{
		fmt.Printf("mensaje %d",infinito)
		infinito++
	}*/
		/*
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
	case modulo==0://establece la codicion en el case
		fmt.Println("numero es par")
	break
	default:
		fmt.Println("numero es impar")
	break
	}

	ejemploContinue()

	array:= [5] int{3,4,5,6,7}// array se indica la cantidad de elementos que va a tener
							  //los arrays son INMUTABLES luego de ser declarados no pueden ser cambiados
	//array[0]=1
	//array[1]=2
	fmt.Println(array,' ',array[1])
	fmt.Println(len(array),cap(array))//para los arreglos len y cap dan como resultado el mismo valor

	slice := []int{1,2,3,4,5}//es un arreglo de elementos donde no se define su capacidad maxima desde la declaracion
	fmt.Println(slice[1:])	
	fmt.Println(slice[:3])
	fmt.Println(slice[:])
	fmt.Println(slice[1:4])
	slice = append(slice, 8)
	fmt.Println(slice)
	lista:=[]int{1,2,3,5}
	slice = append(slice,lista...)//para agregar una serie de elementos a un slice se deben almacenar en otra variable 
								 //y luego con la estructura append hacer mension a dicha variable, go de forma automatica
								//agregara los ...
	fmt.Println(slice)

	for pos,item := range slice{//recorre item por item un arreglo de elementos, donde se captura la posicion y el valor almacenado en la misma
								//si no se desea capturar la posicion se usa el metodo de escape _
		fmt.Println(pos,item)
	}

	var texto string="jony"
	var texto2 string="FREdy"
	fmt.Println(texto[0])//retorna el valor ascii de la letra del string texto en la posicion 0
	fmt.Println(string(texto[0]))//retorna el simbolo del valor ascii en la posiicon cero del string
	fmt.Println(strings.ToLower(texto2))//convierte todos los caracteres de un string en minusculas

	//map son iguales a los diccionarios
	diccionarios:=make(map[string]string)//estructura llave valor usando la palabra reservada map

	diccionarios["jony"]="carmona"
	diccionarios["natalia"]="gomez"

	fmt.Println(diccionarios)

	for nombre,apellido:=range diccionarios{//recorre cada una de las combinaciones llave valor del diccionario
		fmt.Println(nombre,apellido)
	}

	fmt.Println(diccionarios["jony"])//imprime el valor de la llave jony

	llave,ok:=diccionarios["natalia"]//determinar si una llave existe de ser verdadero se carga en la variable ok ->true
	fmt.Println(llave,ok)

	llaves,ok:=diccionarios["fredy"]//determinar si una llave existe de ser verdadero se carga en la variable ok ->false
	fmt.Println(llaves,ok)

	//declarar o instanciar una variable del tipo struct
	micarro:=carro{marca:"mazda",modelo:2022}
	fmt.Println(micarro)	

	

	//declara una instancia desde el paquete misPaquetes para ello se debe correr en la terminar el comando go mod init
	//motosf:=misPaquetes.MotosPublic{Cilindraje:"1500",Pais:"Chile"}
	motosf:=pk.MotosPublic{Cilindraje:"1500",Pais:"Chile"}//se crea el objeto del tipo struct del paquete misPaquetes usando el alias pk
	fmt.Println(motosf)

	//punteros -> permiten acceder a una direccion de memoria especifica
	// para golang se usa el & para indicar posicion de memoria y * para indicar la lectura del valor en la posicion de memoria deseada

	a1 :=20
	b1 := &a1

	fmt.Println(a1)//valor de la varible a1
	fmt.Println(b1)//posicion en memoria de la variable a1 se almacena en la variable b1
	fmt.Println(*b1)//valor almacenado en la posicion de memoria a1 asignado a la variable b1
	*b1=80//al modificar el valor asignado a la variable b1 que esta apuntando a la direccion de memoria de la variable a1
	      //se cambiara el valor de la variable a1
	fmt.Println(a1)

	//declarar o instanciar una variable del tipo struct
	var carro2 carro
	carro2.marca="ferrari"
	fmt.Println(carro2)

	carro2.miCarroFuncion()//se crea una funcion para el objeto carro 2
	fmt.Println(carro2)
	carro2.apuntadorFuncion()
	fmt.Println(carro2)
	carro2.apuntadorFuncion()
	fmt.Println(carro2)
	carro2.apuntadorFuncion()
	fmt.Println(carro2)
	
}*/

//goroutines
//permite la creacion de sistemas con multiples hilos que se ejecutaran de manera concurrente en el codigo haciendo el mismo mas eficiente

func mensaje (texto string,wg *sync.WaitGroup){//
	defer wg.Done()// indica que la subrutina finalizo
	time.Sleep(time.Second*4)
	fmt.Println(texto)	
}
func contador2 (conteos int,wg *sync.WaitGroup){//
	defer wg.Done()// indica que la subrutina finalizo
	for i:=0;i<conteos;i++{
		fmt.Println("conteo: ",i)
		time.Sleep(time.Second*1)
	}	
}
//la funcion main esta corriendo dentro de una goroutine que una 
//vez termina la misma muere
func main(){
	//time.Sleep(time.Second*1) //retardo de 1 segundo

	var wg sync.WaitGroup //acumula un conjunto de gotoutines para ser liberados poco a poco
	fmt.Println("Mensaje dentro de la funcion main")
	wg.Add(2)//agregar una subrutina al waitgroup, que espera que se ejecute la subrutina antes que la subrutina main muera 
	go mensaje("Mensaje desde subrutina",&wg)//la palabra resevada go genera que esta linea de codigo se ejecute de forma
											 //concurrente con la funcion main en este caso	
	go contador2(3,&wg)
	go func(text string){
		fmt.Println(text)
	}("mensaje desde funcion anonima")	
	wg.Wait()//se le indica a la goroutine del main que espere hasta que las subrutinas del waitgroup acaben	
}