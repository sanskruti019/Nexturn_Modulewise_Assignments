from models import Customer, customers

def add_customer(name, email, phone):
    if not name or not email or not phone:
        return "Error: All fields (name, email, phone) are required."
    customers.append(Customer(name, email, phone))
    return "Customer added successfully!"

def view_customers():
    if not customers:
        return "No customers found."
    return "\n".join(customer.display() for customer in customers)
