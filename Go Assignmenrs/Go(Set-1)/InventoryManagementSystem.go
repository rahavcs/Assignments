package main

import "fmt"

// Contents for storing in inventory
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

// Store the inventory of products
var inventory []Product

// Function to add a new product to the inventory
func addProduct(id int, name string, price float64, stock int) string {
	product := Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}
	inventory = append(inventory, product)
	return "Product added successfully!"
}

// Function to update the stock of an existing product
func updateStock(id int, newStock int) string {
	for i, product := range inventory {
		if product.ID == id {
			if newStock < 0 {
				return "Stock can't be negative."
			}
			inventory[i].Stock = newStock
			return fmt.Sprintf("Stock updated for product %s. New stock: %d", product.Name, newStock)
		}
	}
	return "Product not found."
}

// Function to search for a product by ID or Name
func searchProduct(query string) string {
	for _, product := range inventory {
		if fmt.Sprintf("%d", product.ID) == query || product.Name == query {
			return fmt.Sprintf("Found: %s (ID: %d) | Price: %.2f | Stock: %d", product.Name, product.ID, product.Price, product.Stock)
		}
	}
	return "Product not found."
}

// Function to display all products in the inventory
func displayInventory() {
	if len(inventory) == 0 {
		fmt.Println("No products in the inventory yet.")
		return
	}
	fmt.Printf("%-10s %-20s %-10s %-10s\n", "ID", "Product Name", "Price", "Stock")
	for _, product := range inventory {
		fmt.Printf("%-10d %-20s %-10.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

// Function to sort products by price (low to high)
func sortByPrice() {
	for i := 0; i < len(inventory)-1; i++ {
		for j := 0; j < len(inventory)-i-1; j++ {
			if inventory[j].Price > inventory[j+1].Price {
				inventory[j], inventory[j+1] = inventory[j+1], inventory[j]
			}
		}
	}
	fmt.Println("Inventory sorted by price (low to high).")
}

// Function to sort products by stock (low to high)
func sortByStock() {
	for i := 0; i < len(inventory)-1; i++ {
		for j := 0; j < len(inventory)-i-1; j++ {
			if inventory[j].Stock > inventory[j+1].Stock {
				inventory[j], inventory[j+1] = inventory[j+1], inventory[j]
			}
		}
	}
	fmt.Println("Inventory sorted by stock (low to high).")
}

func main() {
	fmt.Println(addProduct(1, "Apple", 1.99, 50))  
	fmt.Println(addProduct(2, "Banana", 0.99, 100)) 
	fmt.Println(addProduct(3, "Orange", 1.49, 75))  

	fmt.Println("\nCurrent Inventory:")
	displayInventory()

	fmt.Println("\nUpdating stock for product ID 1 (Apple) to 60...")
	fmt.Println(updateStock(1, 60)) 

	fmt.Println("\nSearching for product with ID 2...")
	fmt.Println(searchProduct("2")) 

	fmt.Println("\nSearching for product 'Banana'...")
	fmt.Println(searchProduct("Banana")) 

	fmt.Println("\nSorting inventory by price...")
	sortByPrice()
	displayInventory() 

	fmt.Println("\nSorting inventory by stock...")
	sortByStock()
	displayInventory() 
}
