# Scribesphere

Scribesphere is a personal blogging platform that allows users to create, read, update, and delete blog posts. It features a Go backend for the API and a React frontend for user interaction.

## Features

- Create, Read, Update, and Delete blog posts.
- CORS-enabled for frontend interaction.
- RESTful API structure.

## Technologies Used

- **Backend**: Go (Golang)
- **Frontend**: React
- **Database**: Firebase (Firestore)

## Getting Started

### Prerequisites

- Go (version 1.17 or later)
- Node.js (for the React frontend)

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/scribesphere.git
   cd scribesphere/backend
   ```

2. **Set up the Go backend**:

   ```bash
   go mod tidy
   go run main.go
   ```

3. **Set up the React frontend**:

   ```bash
   cd scribesphere  # Navigate to your frontend folder
   npm install
   npm start
   ```

## API Endpoints

- **GET /posts**: Retrieve all blog posts.
- **POST /posts**: Create a new blog post.
- **GET /post/{id}**: Retrieve a blog post by ID.
- **PUT /post/{id}**: Update a blog post by ID.
- **DELETE /post/{id}**: Delete a blog post by ID.
