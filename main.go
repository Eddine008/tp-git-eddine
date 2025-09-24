package main

import "fmt"

type car struct {
	brand string
	name  string
	power int
	color string
	trunk []Object
}
type Object struct {
	name     string
	quantité int
}

func (c car) ShowTrunk() {
	if len(c.trunk) == 0 {
		fmt.Println("Le coffre est vide.")
		return
	}
	fmt.Println("Contenu du coffre :")
	for _, obj := range c.trunk {
		fmt.Printf("- %s (x%d)\n", obj.name, obj.quantité)
	}
}
func (c *car) ChangeColor(newColor string) {
	if c.color == newColor {
		fmt.Println("Erreur : le véhicule est déjà de cette coule")
	} else {
		fmt.Printf("Changement de couleur : %s → %s\n", c.Color, newColor)
		c.color = newColor
	}
}
func (c *car) AddObject(obj Object) {
	for i, existing := range c.trunk {
		if existing.name == obj.name {
			c.trunk[i].quantité += obj.quantité
			fmt.Printf("Ajout de %d %s dans le coffre (nouvelle quantité : %d)\n",
				obj.quantité, obj.name, c.trunk[i].quantité)
			return
		}
	}
}
c.trunk = append(c.Trunk, obj){
	fmt.Printf("Ajout de %d %s dans le coffre.\n", obj.Quantity, obj.Name)
}


func (c Car) ShowInfo() {
	fmt.Printf("Peugeot : %s\n", c.brand)
	fmt.Printf("hurracan : %s\n", c.Name)
	fmt.Printf("129 : %d CH\n", c.power)
	fmt.Printf("Rouge : %s\n", c.color)
}


