
from book_management import add_book, view_books, search_book
from customer_management import add_customer, view_customers
from sales_management import sell_book, view_sales

def main_menu():
    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        
        try:
            choice = int(input("Enter your choice: "))
            if choice == 1:
                book_management_menu()
            elif choice == 2:
                customer_management_menu()
            elif choice == 3:
                sales_management_menu()
            elif choice == 4:
                print("Exiting... Goodbye!")
                break
            else:
                print("Invalid choice! Please try again.")
        except ValueError:
            print("Invalid input! Please enter a number.")

def book_management_menu():
    while True:
        print("\nBook Management")
        print("1. Add Book")
        print("2. View Books")
        print("3. Search Book")
        print("4. Back to Main Menu")
        
        try:
            choice = int(input("Enter your choice: "))
            if choice == 1:
                title = input("Title: ")
                author = input("Author: ")
                price = float(input("Price: "))
                quantity = int(input("Quantity: "))
                print(add_book(title, author, price, quantity))
            elif choice == 2:
                print(view_books())
            elif choice == 3:
                keyword = input("Enter title or author to search: ")
                print(search_book(keyword))
            elif choice == 4:
                break
            else:
                print("Invalid choice! Please try again.")
        except ValueError:
            print("Invalid input! Please try again.")

def customer_management_menu():
    while True:
        print("\nCustomer Management")
        print("1. Add Customer")
        print("2. View Customers")
        print("3. Back to Main Menu")
        
        try:
            choice = int(input("Enter your choice: "))
            if choice == 1:
                name = input("Name: ")
                email = input("Email: ")
                phone = input("Phone: ")
                print(add_customer(name, email, phone))
            elif choice == 2:
                print(view_customers())
            elif choice == 3:
                break
            else:
                print("Invalid choice! Please try again.")
        except ValueError:
            print("Invalid input! Please try again.")

def sales_management_menu():
    while True:
        print("\nSales Management")
        print("1. Sell Book")
        print("2. View Sales")
        print("3. Back to Main Menu")
        
        try:
            choice = int(input("Enter your choice: "))
            if choice == 1:
                customer_name = input("Customer Name: ")
                customer_email = input("Customer Email: ")
                customer_phone = input("Customer Phone: ")
                book_title = input("Book Title: ")
                quantity = int(input("Quantity: "))
                print(sell_book(customer_name, customer_email, customer_phone, book_title, quantity))
            elif choice == 2:
                print(view_sales())
            elif choice == 3:
                break
            else:
                print("Invalid choice! Please try again.")
        except ValueError:
            print("Invalid input! Please try again.")

if __name__ == "__main__":
    main_menu()
