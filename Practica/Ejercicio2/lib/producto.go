package lib

import "fmt"

type Producto struct {
	CodigoBarra int
	Marca       string
	Descripcion string
	Tipo        string
	Precio      float32
}
type List []Producto

var producto *Producto
var listaProductos List
var bufferList List
var buffer int

func NuevoProducto() Producto {
	producto = new(Producto)
	fmt.Print("Ingrese los datos del nuevo producto\nCodigo de barra: ")
	fmt.Scan(&producto.CodigoBarra)
	fmt.Print("Marca: ")
	fmt.Scan(&producto.Marca)
	fmt.Print("Descripcion: ")
	fmt.Scan(&producto.Descripcion)
	fmt.Print("Tipo: ")
	fmt.Scan(&producto.Tipo)
	fmt.Print("Precio: ")
	fmt.Scan(&producto.Precio)
	return *producto
}
func EliminarProducto() {
	fmt.Print("Ingrese codigo de barra del producto a eliminar: ")
	fmt.Scan(&buffer)
	for _, p := range listaProductos {
		if p.CodigoBarra != buffer {
			bufferList = append(bufferList, p)
		}
	}
	listaProductos = bufferList
}
func MostrarProducto(p Producto) {
	fmt.Println(p)
}
func BuscarProducto(cod int) bool {
	for _, p := range listaProductos {
		if p.CodigoBarra == buffer {
			return true
		}
	}
	return false
}
func TraerProducto(cod int) Producto {
	/*p := new(Producto)
	for _, *p = range listaProductos {
		if p.CodigoBarra == cod {
			fmt.Println(*p)
			return *p
		}
	}
	return *p*/
	producto := new(Producto)
	for _, producto := range listaProductos {
		if producto.CodigoBarra == cod {
			fmt.Println(producto)
			return producto
		}
	}
	return *producto
}
func ModificarProducto() {
	producto := new(Producto)
	fmt.Print("Ingrese codigo de barra del producto: ")
	fmt.Scan(&buffer)
	if BuscarProducto(buffer) == true {
		*producto = TraerProducto(buffer) ///VER SI RECIBO EL PRODUCTO QUE QUIERO
		fmt.Println(*producto)
		fmt.Print("Ingrese los datos del producto a modificar\n ")
		fmt.Print("Marca: ")
		fmt.Scan(&producto.Marca)
		fmt.Print("Descripcion: ")
		fmt.Scan(&producto.Descripcion)
		fmt.Print("Tipo: ")
		fmt.Scan(&producto.Tipo)
		fmt.Print("Precio: ")
		fmt.Scan(&producto.Precio)
	} else {
		fmt.Println("El codigo ingresado no corresponde con un producto de la lista")
	}
	fmt.Println(*producto)

}
func AgregarProductoLista(p Producto) {
	listaProductos = append(listaProductos, p)
}
func MostrarListaProducto() {
	for _, l := range listaProductos {
		fmt.Println(l)
	}
}
func MostrarJson() List {
	return listaProductos
}
