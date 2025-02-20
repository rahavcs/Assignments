import testing
import sqlite3
import json
from app import app  

@testing.fixture
def client():
    """Creates a test client using Flask's test client."""
    app.config["TESTING"] = True
    with app.test_client() as client:
        yield client

def test_homepage(client):
    """Test if homepage loads successfully."""
    response = client.get("/")
    assert response.status_code == 200

def test_add_user(client):
    """Test adding a user via POST request."""
    user_data = {"name": "Test User", "email": "test@example.com"}
    response = client.post("/users", data=json.dumps(user_data), content_type="application/json")
    assert response.status_code == 201
    assert b"User added successfully" in response.data

def test_get_users(client):
    """Test retrieving all users."""
    response = client.get("/users")
    assert response.status_code == 200
    users = json.loads(response.data)
    assert isinstance(users, list)

def test_get_user_by_id(client):
    """Test retrieving a user by ID."""
    response = client.get("/users/1")
    assert response.status_code in [200, 404]  

def test_update_user(client):
    """Test updating an existing user."""
    user_update = {"name": "Updated User", "email": "updated@example.com"}
    response = client.put("/users/1", data=json.dumps(user_update), content_type="application/json")
    assert response.status_code in [200, 404]  

def test_delete_user(client):
    """Test deleting a user."""
    response = client.delete("/users/1")
    assert response.status_code in [200, 404] 
