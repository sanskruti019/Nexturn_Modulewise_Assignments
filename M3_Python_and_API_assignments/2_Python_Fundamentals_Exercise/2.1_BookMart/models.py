class Book:
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def display(self):
        return f"Title: {self.title}, Author: {self.author}, Price: {self.price}, Quantity: {self.quantity}"


class Customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone

    def display(self):
        return f"Name: {self.name}, Email: {self.email}, Phone: {self.phone}"


class Transaction:
    def __init__(self, customer_name, book_title, quantity_sold):
        self.customer_name = customer_name
        self.book_title = book_title
        self.quantity_sold = quantity_sold

    def display(self):
        return f"Customer: {self.customer_name}, Book: {self.book_title}, Quantity Sold: {self.quantity_sold}"


# Shared data stores
books = []
customers = []
sales = []
