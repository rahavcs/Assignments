books = []

def add_book():
    try:
        title = input("Enter book title: ")
        author = input("Enter book author: ")
        price = float(input("Enter book price: "))
        quantity = int(input("Enter book quantity: "))

        if price <= 0 or quantity <= 0:
            print("Price and quantity must be positive numbers.")
            return

        books.append({"title": title, "author": author, "price": price, "quantity": quantity})
        print("Book added successfully!")
    except ValueError:
        print("Invalid input. Please enter numbers for price and quantity.")

def view_books():
    if not books:
        print("No books available.")
    else:
        print("Available Books:")
        for book in books:
            print(f"Title: {book['title']}, Author: {book['author']}, Price: {book['price']}, Quantity: {book['quantity']}")
