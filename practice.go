package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	ID               int
	Title            string
	ShortDescription string
	Price            float64
}

func NewProduct(title, shortDesc string, price float64) *Product {
	product := Product{
		ID:               rand.Int(),
		Title:            title,
		ShortDescription: shortDesc,
		Price:            price,
	}
	return &product
}

func (p *Product) outputProduct() {
	// print := fmt.Sprintf("The Product is called %v and the description is %v (price: %v)", p.Title, p.ShortDescription, p.Price)
	// fmt.Println(print)
	fmt.Printf("ID: %v\n", p.ID)
	fmt.Printf("Title: %v\n", p.Title)
	fmt.Printf("Description: %v\n", p.ShortDescription)
	fmt.Printf("Price: USD %.2f\n", p.Price)
}

func (p *Product) store() {
	stringId := strconv.Itoa(p.ID)
	file, _ := os.Create(stringId + ".txt")

	content := fmt.Sprintf(`
  ID: %v
  Title: %v
  Description: %v
  Price: %.2f`, p.ID, p.Title, p.ShortDescription, p.Price)

	file.WriteString(content)
	file.Close()
}

func main() {
	// personOne := NewProduct("product one", "this is product one", 10.99)
	// personTwo := NewProduct("product two", "this is product two", 12.00)
	//
	// personOne.outputProduct()
	// personTwo.outputProduct()
	createdProduct := getProduct()
	createdProduct.outputProduct()
	createdProduct.store()
}

func getProduct() *Product {
	fmt.Println("Please enter the product data.")
	fmt.Println("-----------------------------")

	reader := bufio.NewReader(os.Stdin)

	// idInput := readUserInput(reader, "Product ID: ")
	titleInput := readUserInput(reader, "Product Title: ")
	descriptionInput := readUserInput(reader, "Product Description: ")
	priceInput := readUserInput(reader, "Product Price: ")

	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	product := NewProduct(titleInput, descriptionInput, priceValue)

	return product
}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}
