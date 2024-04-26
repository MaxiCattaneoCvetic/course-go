// Generamos unas estructuras que representaran un carrito de compras

package commerce

// generamos una estructura de tipo Type que represente un producto
// se representa nombre de atributo - tipo de dato - como se llamara en el json

//producto
type Product struct {
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Type  Type     `json:"type"`
	Count uint16   `json:"count"` // cantidad de unidades que se agregaron al carro
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

// tipo
type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

//carrito
type Car struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

// metodos

func (product Product) TotalPrice() float64 {
	return float64(product.Count) * product.Price
}

func (c *Car) AddProduct(product ...Product) { // con el ... definimos que vamos a recibir varios productos
	c.Products = append(c.Products, product...) // con los .. pero para el otro lado desestructuramos los productos
}

// en el primero caso los productos vienen en forma de array {1, 2, 3} en el segundo caso lo desestructuramos {1, 2, 3} -> {1}, {2}, {3}
func (c Car) Total() float64 {
	var total float64
	for _, product := range c.Products {
		total += product.TotalPrice()
	}
	return total
}

func NewCar(UserID uint) Car {
	return Car{UserID: UserID}
}
