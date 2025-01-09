# 📋 Todo-App

This is a simple task management application.

## 🚀 **Authentication**

To work with the application, you first need to register and log in.

### 📌 **User Registration (Sign-Up)**  
**POST:** `/v1/auth/sign-up`  

### 📌 **User Login (Sign-In)**  
**POST:** `/v1/auth/sign-in`  

**After successfully logging in, you will receive a JWT token**, which must be used to work with the lists API.

---

## 🗑️ **CRUD Operations for Lists**

### 📌 **Get All Lists**  
**GET:** `/v1/api/lists`

### 📌 **Get List by ID**  
**GET:** `/v1/api/lists/{id}`

### 📌 **Delete List by ID**  
**DELETE:** `/v1/api/lists/{id}`

### 📌 **Update List by ID**  
**PUT:** `/v1/api/lists/{id}`

---

## 💂‍♂️ **Working with Items in Lists**

### 📌 **Create a New Item in a List**  
**POST:** `/v1/api/lists/{id}/items`  
Creates a new item in the list with the identifier `{id}`.

### 📌 **Get Items from a List**  
**GET:** `/v1/api/lists/{id}/items`  
Returns all items from the list with the identifier `{id}`.

### 📌 **Get Item by ID**  
**GET:** `/v1/api/items/{id}`

### 📌 **Delete Item by ID**  
**DELETE:** `/v1/api/items/{id}`

### 📌 **Update Item by ID**  
**PUT:** `/v1/api/items/{id}`

---

