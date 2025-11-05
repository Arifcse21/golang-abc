package shopinventory

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"github.com/fatih/color"

)

func SetShopName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your shop name: ")
	shopName, _ := reader.ReadString('\n')
	color.Green("Your shop name is: %s", shopName)
	return shopName
}

func AddItem(inventory []string) []string {
    fmt.Println("Adding item to inventory")

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter item name: ")
    itemName, _ := reader.ReadString('\n')
    itemName = strings.TrimSpace(itemName)
    fmt.Println("Item name is: ", itemName)

    var itemQuantity int64
    for {
        fmt.Print("Enter item quantity: ")
        itemQuantityStr, _ := reader.ReadString('\n')
        itemQuantityStr = strings.TrimSpace(itemQuantityStr)
        
        parsedQuantity, err := strconv.ParseInt(itemQuantityStr, 10, 64)
        if err != nil {
            fmt.Println("Invalid input for quantity. Please enter a valid number.")
            continue
        }
        itemQuantity = parsedQuantity
        fmt.Println("Item quantity is: ", itemQuantity)
        break
    }

    var itemPrice float64
    for {
        fmt.Print("Enter item price: ")
        itemPriceStr, _ := reader.ReadString('\n')
        itemPriceStr = strings.TrimSpace(itemPriceStr)
        
        parsedPrice, err := strconv.ParseFloat(itemPriceStr, 64)
        if err != nil {
            fmt.Println("Invalid input for price. Please enter a valid number.")
            continue
        }
        itemPrice = parsedPrice
        fmt.Println("Item price is: ", itemPrice)
        break
    }
    
    itemData := "Name: " + itemName + " Price: " + fmt.Sprintf("%.2f", itemPrice) + " Quantity: " + fmt.Sprintf("%d", itemQuantity)
    inventory = append(inventory, itemData)
	return inventory
}

func ViewInventory(inventory []string) {
    if len(inventory) == 0 {
        color.Red("Inventory is empty \n\n")
        return
    }
    fmt.Println("Viewing inventory:")
    for i, item := range inventory {
        fmt.Printf("%d. %s\n\n", i+1, item)
    }
}

func RunShopInventory() {
	shopName := ""
	for {
		shopName = SetShopName()
		if shopName == "" {
			fmt.Println("You must enter a shop name.")
		} else {
			break
		}	
	}

	var option string = ""
	var inventory []string = []string{}


	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Add an item")
		fmt.Println("2. View inventory")
		fmt.Println("3. Exit")
		fmt.Print("Enter your option: ")
		fmt.Scanln(&option)
		fmt.Println("")

		switch option {
		case "1":
			inventory = AddItem(inventory)
		case "2":
			ViewInventory(inventory)
		case "3":
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
	 

}