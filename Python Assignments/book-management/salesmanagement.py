from bookmanagement import books

sales = []

def sell_book():
    customer_name = input("Enter customer name: ")
    book_title = input("Enter book title: ")
    try:
        quantity = int(input("Enter quantity: "))

        for book in books:
            if book["title"].lower() == book_title.lower():
                if book["quantity"] >= quantity:
                    book["quantity"] -= quantity
                    sales.append({"customer": customer_name, "book": book_title, "quantity": quantity})
                    print("Sale successful!")
                    return
                else:
                    print(f"Not enough stock. Only {book['quantity']} available.")
                    return

        print("Book not found.")
    except ValueError:
        print("Invalid input. Please enter a valid number for quantity.")

def view_sales():
    if not sales:
        print("No sales records found.")
    else:
        print("Sales Records:")
        for sale in sales:
            print(f"Customer: {sale['customer']}, Book: {sale['book']}, Quantity: {sale['quantity']}")
