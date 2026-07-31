# Go (Fiber) + React Fullstack Project Template

Welcome! This is a professional-grade fullstack project template designed for learning and building web applications using **Go (Fiber)** for the backend and **React (Vite + TypeScript)** for the frontend.

## Project Structure

```text
go_workspace/
├── backend/                  # Golang API Backend
│   ├── cmd/
│   │   └── api/
│   │       └── main.go       # Entrypoint
│   ├── config/
│   │   └── config.go         # App configuration & Environment variables
│   ├── internal/             # Private application code
│   │   ├── handler/          # HTTP Request Handlers (Fiber controllers)
│   │   ├── middleware/       # Custom middlewares (CORS, Logger, etc.)
│   │   ├── model/            # Database Models & Request/Response Structs
│   │   ├── repository/       # Data Access Layer (DB queries)
│   │   └── service/          # Business Logic Layer
│   ├── pkg/                  # Shared public packages
│   │   └── utils/            # Helper utilities
│   ├── .env.example          # Environment variables template
│   └── go.mod                # Go module definition
└── frontend/                 # React Frontend (Vite + TS)
    ├── src/
    │   ├── assets/           # Static assets (images, SVGs)
    │   ├── components/       # Reusable UI components
    │   ├── hooks/            # Custom React hooks
    │   ├── pages/            # Page/Screen components
    │   ├── services/         # API integration layer
    │   ├── context/          # React contexts for global state (e.g. Auth)
    │   ├── App.css           # Global layout & page styles
    │   ├── App.tsx           # Main application setup
    │   ├── index.css         # CSS reset & design tokens variables
    │   └── main.tsx          # Frontend entry point
    ├── index.html
    └── vite.config.ts
```

## Getting Started

### Prerequisites

Make sure you have the following installed on your machine:
- **Go** (version 1.20 or later)
- **Node.js** (version 18 or later) & **npm**

---

### Running the Backend

1. Navigate to the `backend` directory:
   ```bash
   cd backend
   ```
2. Copy the environment template and configure your variables:
   ```bash
   cp .env.example .env
   ```
3. Download dependencies:
   ```bash
   go mod tidy
   ```
4. Run the development server:
   ```bash
   go run cmd/api/main.go
   ```
   The API will be available at `http://localhost:8080`.

---

### Running the Frontend

1. Navigate to the `frontend` directory:
   ```bash
   cd frontend
   ```
2. Install the dependencies:
   ```bash
   npm install
   ```
3. Run the Vite development server:
   ```bash
   npm run dev
   ```
   The frontend will be available at `http://localhost:5173`.
