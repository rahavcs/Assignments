customers = []

def add_customer():
    name = input("Enter customer name: ")
    email = input("Enter customer email: ")
    phone = input("Enter customer phone: ")

    if not name or not email or not phone:
        print("All fields are required!")
        return

    customers.append({"name": name, "email": email, "phone": phone})
    print("Customer added successfully!")

def view_customers():
    if not customers:
        print("No customers found.")
    else:
        print("Customer List:")
        for customer in customers:
            print(f"Name: {customer['name']}, Email: {customer['email']}, Phone: {customer['phone']}")
