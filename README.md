# Task Manager Application

## Overview
This project is a Task Manager application with a **Frontend** built using **Vue 3**, **Vuetify**, and **Vite**, and a **Backend** implemented in **Go** with a **PostgreSQL** database.

### Features
- Create, update, delete, and view tasks.
- Visualize tasks in a calendar with highlighted task dates.
- Real-time task management and status updates.
- Responsive UI with Vuetify components.

---

## Getting Started

### Prerequisites
Ensure you have the following installed:
- [Node.js](https://nodejs.org/) (v16 or later)
- [Go](https://golang.org/) (v1.19 or later)
- [PostgreSQL](https://www.postgresql.org/)

---

## Frontend Setup

1. Navigate to the `frontend` directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Run the development server:
   ```bash
   npm run dev
   ```

4. Open your browser and navigate to:
   ```
   http://localhost:5173
   ```

---

## Backend Setup

1. Navigate to the `backend` directory:
   ```bash
   cd backend
   ```

2. Configure your PostgreSQL database:
   - Create a new database (e.g., `task_manager`).
   - Update the database credentials in the `config` file (e.g., `config.yaml` or `main.go`).

3. Run the backend server:
   ```bash
   go run main.go
   ```

4. The backend server will start on:
   ```
   http://localhost:8080
   ```

---

## API Endpoints

| Endpoint             | Method | Description           |
|----------------------|--------|-----------------------|
| `/tasks`             | GET    | Fetch all tasks       |
| `/tasks`             | POST   | Create a new task     |
| `/tasks/{id}`        | PUT    | Update a task by ID   |
| `/tasks/{id}`        | DELETE | Delete a task by ID   |
| `/entries`             | GET    | Fetch all entries       |
| `/entries`             | POST   | Create a new entrу     |
| `/entries/{id}`        | PUT    | Update a entrу by ID   |
| `/entries/{id}`        | DELETE | Delete a entrу by ID   |

---

## Project Structure

### Frontend
```
frontend/
├── src/
│   ├── components/       # Reusable Vue components
│   ├── plugins/          # Vuetify settings
│   ├── assets/           # Static assets
│   ├── main.js           # App entry point
├── public/               # Static public files
├── vite.config.js        # Vite configuration
└── package.json          # Project metadata and dependencies
```

### Backend
```
backend/
├── handlers/             # API endpoint handlers
├── models/               # Database models
├── routes/               # API route definitions
├── database/             # Database settings
├── main.go               # Entry point of the application
└── go.mod                # Go module dependencies
```

---

## Usage

1. Open the application in your browser via the frontend server.
2. Manage tasks using the provided interface.
3. View the calendar to see highlighted task dates.

---

## Contributing
Contributions are welcome! To contribute:
1. Fork this repository.
2. Create a new branch for your feature/bugfix.
3. Submit a pull request with a detailed description of your changes.

---

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## Acknowledgements
- [Vue.js](https://vuejs.org/)
- [Vuetify](https://vuetifyjs.com/)
- [Vite](https://vitejs.dev/)
- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)

