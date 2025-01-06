from models import Book, books

def add_book(title, author, price, quantity):
    try:
        price = float(price)
        quantity = int(quantity)
        if price <= 0 or quantity <= 0:
            raise ValueError("Price and Quantity must be positive numbers.")
        books.append(Book(title, author, price, quantity))
        return "Book added successfully!"
    except ValueError as e:
        return str(e)

def view_books():
    if not books:
        return "No books available."
    return "\n".join(book.display() for book in books)

def search_book(search_term):
    results = [book.display() for book in books if search_term.lower() in book.title.lower() or search_term.lower() in book.author.lower()]
    return "\n".join(results) if results else "No matching books found."
