from flask import Flask, request, redirect, url_for, render_template_string
import sqlite3
import datetime

app = Flask(__name__)

def get_db_connection():
    conn = sqlite3.connect('bookbuddy.db')
    conn.row_factory = sqlite3.Row  
    return conn

@app.route('/')
def home():
    current_year = datetime.datetime.now().year
    min_year = 1500

    year_options = "\n".join([f'<option value="{year}">{year}</option>' for year in range(current_year, min_year - 1, -1)])
    
    return f'''
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>BookBuddy - Manage Your Books</title>
        <style>
            body {{
                font-family: 'Arial', sans-serif;
                margin: 0;
                padding: 0;
                background: linear-gradient(120deg, #f6d365, #fda085);
                color: #333;
            }}
            header {{
                background-color: #ff6f61;
                color: white;
                padding: 20px 0;
                text-align: center;
                margin-bottom: 30px;
                box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            }}
            header h1 {{
                margin: 0;
                font-size: 3rem;
                font-weight: bold;
            }}
            header p {{
                margin: 5px 0 0;
                font-size: 1.2rem;
            }}
            main {{
                max-width: 800px;
                margin: 0 auto;
                padding: 20px;
                background: white;
                box-shadow: 0 8px 10px rgba(0, 0, 0, 0.2);
                border-radius: 12px;
            }}
            a {{
                display: inline-block;
                text-decoration: none;
                color: white;
                background-color: #ff6f61;
                padding: 10px 20px;
                border-radius: 8px;
                margin-top: 10px;
                font-size: 1.1rem;
                font-weight: bold;
                box-shadow: 0 4px 6px rgba(0, 0, 0, 0.2);
                transition: all 0.3s ease;
            }}
            a:hover {{
                background-color: #e85c50;
                box-shadow: 0 6px 8px rgba(0, 0, 0, 0.3);
            }}
            form {{
                margin-top: 20px;
            }}
            label {{
                font-size: 1.2rem;
                font-weight: bold;
                margin-bottom: 8px;
                display: block;
            }}
            input, select {{
                width: 100%;
                padding: 12px;
                margin: 10px 0;
                border: 1px solid #ddd;
                border-radius: 6px;
                font-size: 1rem;
            }}
            .input-top {{
                margin-bottom: 20px;
            }}
            .input-bottom {{
                margin-top: 20px;
            }}
            button {{
                display: inline-block;
                background-color: #ff6f61;
                color: white;
                padding: 12px 20px;
                border: none;
                border-radius: 8px;
                cursor: pointer;
                font-size: 1.1rem;
                font-weight: bold;
                box-shadow: 0 4px 6px rgba(0, 0, 0, 0.2);
                transition: all 0.3s ease;
            }}
            button:hover {{
                background-color: #e85c50;
                box-shadow: 0 6px 8px rgba(0, 0, 0, 0.3);
            }}
            footer {{
                text-align: center;
                margin-top: 20px;
                font-size: 0.9rem;
                color: #666;
            }}
        </style>
    </head>
    <body>
        <header>
            <h1>BookBuddy</h1>
            <p>Your Personal Book Management System</p>
        </header>
        <main>
            <a href="/books">View All Books</a>
            <h2>Add a New Book</h2>
            <form action="/books" method="POST">
                <!-- Title and Author Fields at the top -->
                <div class="input-top">
                    <label for="title">Title:</label>
                    <input type="text" id="title" name="title" placeholder="Enter book title" required>
                    
                    <label for="author">Author:</label>
                    <input type="text" id="author" name="author" placeholder="Enter author's name" required>
                </div>
                
                <!-- Published Year and Genre Fields at the bottom -->
                <div class="input-bottom">
                    <label for="published_year">Published Year:</label>
                    <select id="published_year" name="published_year" required>
                        {year_options}
                    </select>
                    
                    <label for="genre">Genre:</label>
                    <select id="genre" name="genre" required>
                        <option value="Fiction">Fiction</option>
                        <option value="Non-Fiction">Non-Fiction</option>
                        <option value="Mystery">Mystery</option>
                        <option value="Sci-Fi">Sci-Fi</option>
                    </select>
                </div>
                
                <button type="submit">Add Book</button>
            </form>
        </main>
        <footer>
            &copy; {datetime.datetime.now().year} BookBuddy. All rights reserved.
        </footer>
    </body>
    </html>
    '''

@app.route('/books', methods=['POST'])
def add_book():
    
    title = request.form['title']
    author = request.form['author']
    published_year = request.form['published_year']
    genre = request.form['genre']

    conn = get_db_connection()
    conn.execute('INSERT INTO books (title, author, published_year, genre) VALUES (?, ?, ?, ?)', 
                 (title, author, published_year, genre))
    conn.commit()
    conn.close()

    return redirect(url_for('view_books'))  

@app.route('/books', methods=['GET'])
def view_books():
    conn = get_db_connection()
    books = conn.execute('SELECT * FROM books').fetchall()
    conn.close()

    if not books:
        return '''
            <h1>No books available</h1>
            <a href="/">Go back to Home</a>
        '''

    books_html = '''
        <table border="1" style="width:100%; border-collapse: collapse; margin-top: 20px;">
            <tr>
                <th>Title</th>
                <th>Author</th>
                <th>Published Year</th>
                <th>Genre</th>
            </tr>
    '''

    for book in books:
        books_html += f'''
            <tr>
                <td>{book['title']}</td>
                <td>{book['author']}</td>
                <td>{book['published_year']}</td>
                <td>{book['genre']}</td>
            </tr>
        '''
    
    books_html += '''
        </table>
        <br>
        <a href="/">Back to Home</a>
    ''' 

    return f'''
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>BookBuddy - View All Books</title>
        <style>
            body {{
                font-family: 'Arial', sans-serif;
                margin: 0;
                padding: 0;
                background: linear-gradient(120deg, #f6d365, #fda085);
                color: #333;
            }}
            header {{
                background-color: #ff6f61;
                color: white;
                padding: 20px 0;
                text-align: center;
                margin-bottom: 30px;
                box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            }}
            header h1 {{
                margin: 0;
                font-size: 3rem;
                font-weight: bold;
            }}
            header p {{
                margin: 5px 0 0;
                font-size: 1.2rem;
            }}
            main {{
                max-width: 800px;
                margin: 0 auto;
                padding: 20px;
                background: white;
                box-shadow: 0 8px 10px rgba(0, 0, 0, 0.2);
                border-radius: 12px;
            }}
            table {{
                width: 100%;
                border-collapse: collapse;
                margin-top: 20px;
            }}
            th, td {{
                padding: 12px;
                text-align: left;
                border-bottom: 1px solid #ddd;
            }}
            th {{
                background-color: #ff6f61;
                color: white;
                font-weight: bold;
            }}
            tr:hover {{
                background-color: #f5f5f5;
            }}
            a {{
                display: inline-block;
                text-decoration: none;
                color: white;
                background-color: #ff6f61;
                padding: 10px 20px;
                border-radius: 8px;
                margin-top: 10px;
                font-size: 1.1rem;
                font-weight: bold;
                box-shadow: 0 4px 6px rgba(0, 0, 0, 0.2);
                transition: all 0.3s ease;
            }}
            a:hover {{
                background-color: #e85c50;
                box-shadow: 0 6px 8px rgba(0, 0, 0, 0.3);
            }}
            footer {{
                text-align: center;
                margin-top: 20px;
                font-size: 0.9rem;
                color: #666;
            }}
        </style>
    </head>
    <body>
        <header>
            <h1>BookBuddy</h1>
            <p>Your Personal Book Management System</p>
        </header>
        <main>
            <h1>All Books</h1>
            {books_html}
        </main>
        <footer>
            &copy; {datetime.datetime.now().year} BookBuddy. All rights reserved.
        </footer>
    </body>
    </html>
    '''

if __name__ == '__main__':
    app.run(debug=True)
