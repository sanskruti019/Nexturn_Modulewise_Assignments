from flask import Blueprint, jsonify, request
from . import db
from .models import Book

book_routes = Blueprint('book_routes', __name__)

@book_routes.route('/books', methods=['POST'])
def add_book():
    data = request.get_json()
    if not all(k in data for k in ('title', 'author', 'published_year', 'genre')):
        return jsonify({'error': 'Invalid book data', 'message': 'All fields are required, you may missed entering any of fileds: title, author, published_year, genre'}), 400

    try:
        new_book = Book(
            title=data['title'],
            author=data['author'],
            published_year=int(data['published_year']),
            genre=data['genre']
        )
        db.session.add(new_book)
        db.session.commit()
        return jsonify({'message': 'Book added successfully', 'book_id': new_book.id}), 201
    except ValueError:
        return jsonify({'error': 'Invalid data type', 'message': 'Published year must be an integer'}), 400
    except Exception as e:
        return jsonify({'error': 'Database error', 'message': str(e)}), 500

@book_routes.route('/books', methods=['GET'])
def get_books():
    books = Book.query.all()
    return jsonify([{
        'id': book.id,
        'title': book.title,
        'author': book.author,
        'published_year': book.published_year,
        'genre': book.genre
    } for book in books])

@book_routes.route('/books/<int:id>', methods=['GET'])
def get_book_by_id(id):
    book = Book.query.get(id)
    if book is None:
        return jsonify({'error': 'Book not found', 'message': 'No book exists with the provided ID'}), 404
    return jsonify({
        'id': book.id,
        'title': book.title,
        'author': book.author,
        'published_year': book.published_year,
        'genre': book.genre
    })

@book_routes.route('/books/<int:id>', methods=['PUT'])
def update_book(id):
    book = Book.query.get(id)
    if not book:
        return jsonify({'error': 'Book not found', 'message': 'No book exists with the provided ID'}), 404

    data = request.get_json()
    book.title = data.get('title', book.title)
    book.author = data.get('author', book.author)
    book.published_year = data.get('published_year', book.published_year)
    book.genre = data.get('genre', book.genre)

    db.session.commit()
    return jsonify({'message': 'Book updated successfully'})

@book_routes.route('/books/<int:id>', methods=['DELETE'])
def delete_book(id):
    book = Book.query.get(id)
    if not book:
        return jsonify({'error': 'Book not found', 'message': 'No book exists with the provided ID'}), 404

    db.session.delete(book)
    db.session.commit()
    return jsonify({'message': 'Book deleted successfully'})
