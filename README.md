# BinaryBuilders Forum

This is the Backend!

An interactive fitness forum that allows users to share their fitness journeys, engage in discussions, and track progress. Users can create threads, comment on discussions, and interact with others using likes and dislikes.

---

## Features

- **Thread Management**: Create, edit, and delete threads with categories.
- **Commenting System**: Add comments to threads with real-time updates.
- **Category Filtering**: Filter threads by categories for easy navigation.
- **User Interaction**: Like or dislike threads and comments.

---

## Tech Stack

### Frontend:
- **React**: For building the user interface.
- **Material-UI**: For consistent and modern styling.
- **Axios**: For API calls.

### Backend:
- **Go (Gin)**: For RESTful API development.
- **PostgreSQL**: For relational database management.

### Deployment:
- **Render**: Hosting for the backend and database.
- **Vercel**: Hosting for the frontend.

---

## Live Demo

- **Deployed Website**: https://cvwo-winter-assignment2025-frontend.vercel.app
- **Note**: Render will temporarily spin down the backend instance when it is inactive, which can delay requests by 50 seconds or more. Please wait for a moment for it to boot up.
---

## Getting Started

### Prerequisites

- **Node.js**: [Download](https://nodejs.org/)
- **Go**: [Download](https://go.dev/)
- **PostgreSQL**: [Download](https://www.postgresql.org/)
- **PGAdmin 4**: [Download](https://www.pgadmin.org/download/) 
---

## Setup Instructions
As render only hosts postgresSQL database for 30 days, here are setup instructions if the demo does not work. It expires on February 25, 2025.

### 1. Clone the Both Back & Frontend Repository:
```bash
git clone https://github.com/WilkinsAng/CVWO-Winter-Assignment-2025-Frontend.git
git clone https://github.com/WilkinsAng/CVWO-Winter-Assignment-2025-Backend.git
```
### 2. Set Up the Database:
**Start PostgreSQL and PGAdmin**:
- Ensure PostgreSQL & PGAdmin is installed and running on your system.
- Launch pgAdmin and log in to your PostgreSQL server.

**Create the Database**:
1. In the Object Explorer on the left, right-click on Databases and select Create > Database....
2. Enter the following details:
   1. Database Name: binarybuilders 
   2. Owner: Your PostgreSQL username (default is postgres).
3. Click Save.

**Import the Schema**:
1. In the Object Explorer, expand the binarybuilders database.
2. Right-click on binarybuilders and select Query Tool.
3. Open the provided schema.sql file:
   1. Click the Open File icon in the Query Tool. 
   2. Select schema.sql from the database/ folder in your project directory.
4. Click the Run button to execute the SQL script.
  
**Seed the Categories**:
1. In the Query Tool, paste the following SQL to add initial categories:
   ```sql
   INSERT INTO categories (name) VALUES 
   ('Cardio'),
   ('Diet Plans'),
    ('Flexibility & Mobility'),
    ('General Fitness Discussions'),
    ('Healthy Recipes'),
    ('Progress Tracking'),
    ('Strength Training'),
    ('Supplements');
    ``` 
2. Run the query (click the Run button).

### 3. Set Up the Backend:
1. Navigate to the Backend folder.
2. Create a .env folder.
```plaintext
DATABASE_URL=postgres://<username>:<password>@localhost:5432/binarybuilders
PORT=8080
```
3. Install Dependencies.
```bash
go mod tidy
```
4. Start Backend Server
```bash
go run main.go
```
### 4. Set Up the Frontend:
1. Navigate to the Frontend folder.
2. Create a .env folder.
```plaintext
REACT_APP_API_URL=http://localhost:8080
```
3. Install Dependencies.
```bash
npm install
```
4. Start Backend Server
```bash
npm start
```
5. Open the app:
- Visit http://localhost:3000 in your browser.

Done By: Ang Wei Jian
