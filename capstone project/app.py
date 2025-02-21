from flask import Flask, request, jsonify, render_template_string
import sqlite3

app = Flask(__name__)

def init_db():
    conn = sqlite3.connect("users.db")
    cursor = conn.cursor()
    cursor.execute('''CREATE TABLE IF NOT EXISTS users (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        name TEXT NOT NULL,
                        email TEXT NOT NULL UNIQUE)''')
    conn.commit()
    conn.close()

init_db()

@app.route('/')
def index():
    return render_template_string('''
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>User Management</title>
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap/5.3.0/css/bootstrap.min.css">
        <script>
            async function fetchUsers() {
                let response = await fetch("/users");
                let users = await response.json();
                let userTable = document.getElementById("userTable");
                userTable.innerHTML = "";
                users.forEach((user, index) => {
                    userTable.innerHTML += `
                        <tr>
                            <td>${index + 1}</td>
                            <td>${user[1]}</td>
                            <td>${user[2]}</td>
                            <td>
                                <button class="btn btn-warning btn-sm" onclick="editUser(${user[0]}, '${user[1]}', '${user[2]}')">Edit</button>
                                <button class="btn btn-danger btn-sm" onclick="deleteUser(${user[0]})">Delete</button>
                            </td>
                        </tr>`;
                });
            }

            async function addUser() {
                let name = document.getElementById("name").value;
                let email = document.getElementById("email").value;
                if (!name || !email) {
                    alert("Please enter both name and email!");
                    return;
                }
                await fetch("/users", {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ name, email })
                });
                document.getElementById("name").value = "";
                document.getElementById("email").value = "";
                fetchUsers();
            }

            async function deleteUser(id) {
                if (confirm("Are you sure you want to delete this user?")) {
                    await fetch(`/users/${id}`, { method: "DELETE" });
                    fetchUsers();
                }
            }

            function editUser(id, name, email) {
                document.getElementById("editId").value = id;
                document.getElementById("editName").value = name;
                document.getElementById("editEmail").value = email;
                document.getElementById("editForm").style.display = "block";
            }

            async function updateUser() {
                let id = document.getElementById("editId").value;
                let name = document.getElementById("editName").value;
                let email = document.getElementById("editEmail").value;
                await fetch(`/users/${id}`, {
                    method: "PUT",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ name, email })
                });
                document.getElementById("editForm").style.display = "none";
                fetchUsers();
            }

            window.onload = fetchUsers;
        </script>
        <style>
            body {
                background: #f4f4f9;
                font-family: Arial, sans-serif;
            }
            .container {
                margin-top: 30px;
            }
            h2 {
                color: #007bff;
            }
            .card {
                padding: 15px;
                border-radius: 10px;
            }
            .table th {
                background: #007bff;
                color: white;
            }
            .btn-sm {
                margin: 2px;
            }
        </style>
    </head>
    <body>
        <div class="container">
            <h2 class="text-center">User Management</h2>
            <div class="row">
                <div class="col-md-6">
                    <div class="card bg-light">
                        <h4>Add User</h4>
                        <input type="text" id="name" class="form-control mb-2" placeholder="Name">
                        <input type="text" id="email" class="form-control mb-2" placeholder="Email">
                        <button class="btn btn-primary" onclick="addUser()">Add</button>
                    </div>
                </div>

                <div class="col-md-6" id="editForm" style="display:none;">
                    <div class="card bg-warning">
                        <h4>Edit User</h4>
                        <input type="hidden" id="editId">
                        <input type="text" id="editName" class="form-control mb-2" placeholder="Name">
                        <input type="text" id="editEmail" class="form-control mb-2" placeholder="Email">
                        <button class="btn btn-dark" onclick="updateUser()">Update</button>
                    </div>
                </div>
            </div>

            <h3 class="mt-4">Users List</h3>
            <table class="table table-bordered">
                <thead>
                    <tr>
                        <th>Sl. No</th>
                        <th>Name</th>
                        <th>Email</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody id="userTable"></tbody>
            </table>
        </div>
    </body>
    </html>
    ''')

@app.route('/users', methods=['POST'])
def add_user():
    data = request.get_json()
    conn = sqlite3.connect("users.db")
    cursor = conn.cursor()
    cursor.execute("INSERT INTO users (name, email) VALUES (?, ?)", (data['name'], data['email']))
    conn.commit()
    conn.close()
    return jsonify({"message": "User added successfully"}), 201

@app.route('/users', methods=['GET'])
def get_users():
    conn = sqlite3.connect("users.db")
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM users")
    users = cursor.fetchall()
    conn.close()
    return jsonify(users)

@app.route('/users/<int:id>', methods=['GET'])
def get_user(id):
    conn = sqlite3.connect("users.db")
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM users WHERE id=?", (id,))
    user = cursor.fetchone()
    conn.close()
    if user:
        return jsonify(user)
    return jsonify({"error": "User not found"}), 404

@app.route('/users/<int:id>', methods=['PUT'])
def update_user(id):
    data = request.get_json()
    conn = sqlite3.connect("users.db")
    cursor = conn.cursor()
    cursor.execute("UPDATE users SET name=?, email=? WHERE id=?", (data['name'], data['email'], id))
    conn.commit()
    conn.close()
    return jsonify({"message": "User updated successfully"})

@app.route('/users/<int:id>', methods=['DELETE'])
def delete_user(id):
    conn = sqlite3.connect("users.db")
    cursor = conn.cursor()
    cursor.execute("DELETE FROM users WHERE id=?", (id,))
    conn.commit()
    conn.close()
    return jsonify({"message": "User deleted successfully"})

if __name__ == '__main__':
    app.run(debug=True)
