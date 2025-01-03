from bookmanagement import add_book, view_books
from customermanagement import add_customer, view_customers
from salesmanagement import sell_book, view_sales

def main():
    while True:
        print("Welcome to the BookArena!!!\nHow May I Assist You Today?")
        print("\n1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        choice = input("\nEnter your choice: ")

        if choice == "1":
            print("\nBook Management")
            print("1. Add Book")
            print("2. View Books")
            sub_choice = input("\nEnter your choice: ")
            if sub_choice == "1":
                add_book()
            elif sub_choice == "2":
                view_books()
            else:
                print("Invalid choice.")

        elif choice == "2":
            print("\nCustomer Management")
            print("1. Add Customer")
            print("2. View Customers")
            sub_choice = input("\nEnter your choice: ")
            if sub_choice == "1":
                add_customer()
            elif sub_choice == "2":
                view_customers()
            else:
                print("Invalid choice.")

        elif choice == "3":
            print("\nSales Management")
            print("1. Sell Book")
            print("2. View Sales")
            sub_choice = input("\nEnter your choice: ")
            if sub_choice == "1":
                sell_book()
            elif sub_choice == "2":
                view_sales()
            else:
                print("Invalid choice.")

        elif choice == "4":
            print("Exiting... Goodbye!")
            break

        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main()
