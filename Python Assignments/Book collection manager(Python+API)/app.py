from flask import Flask, request, jsonify
from db import init_db, get_db_connection

app = Flask(__name__)

#initalize database
init_db()

# checking if book exists with relation to the ID
def book_exists(book_id):
    conn = get_db_connection()
    book = conn.execute('SELECT * FROM books WHERE id = ?', (book_id,)).fetchone()
    conn.close()
    return book

# For adding notebook
@app.route('/books', methods=['POST'])
def add_book():
    data = request.get_json()
    title = data.get('title')
    author = data.get('author')
    published_year = data.get('published_year')
    genre = data.get('genre')

    if not title or not author or not published_year or not genre:
        return jsonify({"error": "Invalid data", "message": "All fields are required"}), 400

    
    valid_genres = ["Fiction", "Non-Fiction", "Mystery", "Sci-Fi"]
    if genre not in valid_genres:
        return jsonify({"error": "Invalid genre", "message": "Genre must be one of: Fiction, Non-Fiction, Mystery, Sci-Fi"}), 400

    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute('INSERT INTO books (title, author, published_year, genre) VALUES (?, ?, ?, ?)',
                   (title, author, published_year, genre))
    conn.commit()
    book_id = cursor.lastrowid
    conn.close()

    return jsonify({"message": "Book added successfully", "book_id": book_id}), 201

# Getting all the books
@app.route('/books', methods=['GET'])
def get_books():
    conn = get_db_connection()
    books = conn.execute('SELECT * FROM books').fetchall()
    conn.close()
    books_list = [dict(book) for book in books]
    return jsonify(books_list)

# Get all the books by ID
@app.route('/books/<int:book_id>', methods=['GET'])
def get_book(book_id):
    book = book_exists(book_id)
    if book is None:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404
    return jsonify(dict(book))

# Updating a book that is currently present
@app.route('/books/<int:book_id>', methods=['PUT'])
def update_book(book_id):
    book = book_exists(book_id)
    if book is None:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    data = request.get_json()
    title = data.get('title', book['title'])
    author = data.get('author', book['author'])
    published_year = data.get('published_year', book['published_year'])
    genre = data.get('genre', book['genre'])

    if not title or not author or not published_year or not genre:
        return jsonify({"error": "Invalid data", "message": "All fields are required"}), 400

    valid_genres = ["Fiction", "Non-Fiction", "Mystery", "Sci-Fi"]
    if genre not in valid_genres:
        return jsonify({"error": "Invalid genre", "message": "Genre must be one of: Fiction, Non-Fiction, Mystery, Sci-Fi"}), 400

    conn = get_db_connection()
    conn.execute('''
        UPDATE books
        SET title = ?, author = ?, published_year = ?, genre = ?
        WHERE id = ?
    ''', (title, author, published_year, genre, book_id))
    conn.commit()
    conn.close()

    return jsonify({"message": "Book updated successfully"})

# Deleting the book with respect to ID
@app.route('/books/<int:book_id>', methods=['DELETE'])
def delete_book(book_id):
    book = book_exists(book_id)
    if book is None:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    conn = get_db_connection()
    conn.execute('DELETE FROM books WHERE id = ?', (book_id,))
    conn.commit()
    conn.close()

    return jsonify({"message": "Book deleted successfully"})

if __name__ == '__main__':
    app.run(debug=True,port=5001)
