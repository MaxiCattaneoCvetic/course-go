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
	Total    float64   `json:"total"`
}
