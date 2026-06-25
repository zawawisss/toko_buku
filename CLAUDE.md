# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

The project consists of a frontend (React + Vite) and a backend (Go/Gin with PostgreSQL).

### Frontend (`toko_buku/`)
- `npm run dev` – Start the Vite development server with hot module replacement (http://localhost:5173).
- `npm run build` – Type-check (`tsc -b`) then bundle for production via Vite.
- `npm run lint` – Run ESLint across the source files.
- `npm run preview` – Preview the production build locally.

### Backend (`backend/`)
- `go run main.go` – Run the Go server directly (requires a running PostgreSQL instance matching `.env`).
- Docker: `docker compose up` – Starts PostgreSQL and the Go service (see `docker-compose.yml`).
- The API listens on `:8080` by default; CORS is configured to allow the Vite dev server (`http://localhost:5173`).

### Testing
No test scripts are configured by default. For the frontend, consider adding Vitest (`npm i -D vitest @vitest/ui happy-dom`) and adding a `"test": "vitest"` script. For the backend, you can add Go tests in the `backend/` directory.

## High-Level Architecture

- **Frontend**: React 19 with TypeScript, bundled by Vite. The React Compiler is enabled via `@rolldown/plugin-babel` with `reactCompilerPreset` for automatic memoization and optimized renders.
- **State Management**: No global state library is added yet; component state can be lifted or managed with React Context/Zustand as needed.
- **Routing**: Not set up; add `react-router-dom` if navigation between pages is required.
- **Styling**: Global CSS in `src/index.css`; component‑scoped styles can be added via CSS Modules, styled‑components, or Tailwind (install as needed).
- **Backend**: Go 1.22+ Gin web framework, communicating with a PostgreSQL database. Provides a RESTful `/books` resource (CRUD) used by the frontend.
- **Communication**: The frontend accesses the backend via relative or absolute URLs (e.g., `fetch('/books')`). During development, the Vite proxy is not configured; ensure the backend is reachable at `http://localhost:8080` and CORS allows the origin.
- **Tooling**: TypeScript project references (`tsconfig.json` → `tsconfig.app.json` & `tsconfig.node.json`) enable separate type‑checking for app and Vite/node code. ESLint is configured with React Hooks and Refresh plugins.

## Backend Development

### Project Structure
```
backend/
├─ main.go          # Gin HTTP server, DB setup, CRUD handlers
├─ go.mod           # Go module definition
├─ go.sum
├─ .env             # Environment variables (see below)
├─ docker-compose.yml # PostgreSQL + service definition
└─ books.json       # Sample data (unused by the API; kept for reference)
```

### Environment Variables
Create `.env` in the `backend/` folder (the file is already present but may need adjustment):
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=awi
DB_PASSWORD=awi
DB_NAME=toko_buku
```
When using `docker compose up`, the service name `postgres` can be used as the host.

### API Endpoints
- `GET /books` – Returns JSON array of all books.
- `POST /books` – Creates a new book; expects `{judul, penulis, harga, stok}`; returns the created object with `id`.
- `PUT /books/:id` – Updates a book; same payload as POST.
- `DELETE /books/:id` – Deletes a book; returns a status message.

### Running the Backend
1. **With Docker (recommended for local development)**:
   ```bash
   cd backend
   docker compose up
   ```
   This starts a PostgreSQL container and compiles/runs the Go API.

2. **Directly with Go** (requires a local PostgreSQL):
   ```bash
   cd backend
   go run main.go
   ```
   Ensure the `.env` variables point to a reachable database.

The backend automatically runs `initDB()` on startup, creating the `buku` table if it does not exist and seeding it with sample data when the table is empty.

## Frontend Development Guidelines

- **Adding a Component**: Place new `.tsx` files in `src/components/` and import them where needed (e.g., `App.tsx`).
- **State & Data Fetching**: Use `fetch` or a library like `axios` to call the backend endpoints. Consider wrapping requests in a custom hook.
- **Styling**: Edit `src/index.css` for global styles. For component‑scoped styling, add a CSS module (`ComponentName.module.css`) or styled‑components.
- **Environment Variables**: Vite exposes variables prefixed with `VITE_` to client code. Define them in a `.env` file at the project root (e.g., `VITE_API_URL=http://localhost:8080`) if you need to configure the API base.
- **TypeScript**: The project uses strict mode; ensure new files follow the existing `tsconfig` conventions.

## Common Tasks

- **Adding a Page/Route**: Install `react-router-dom` (`npm i react-router-dom`) and set up `<BrowserRouter>`, `<Routes>`, and `<Route>` components in `main.tsx` or `App.tsx`.
- **Styling with Tailwind**: Install `tailwindcss`, `postcss`, `autoprefixer`, configure via `tailwind.config.cjs`, and add the directives to `index.css`.
- **Linting Fixes**: Run `npm run lint`; address any reported ESLint errors. For type‑aware rules, consider extending the ESLint config with `typescript-eslint` recommendations.
- **Building for Production**: `npm run build` outputs to `dist/`; preview with `npm run preview`.

## Notes

- Keep secrets out of the repository: never commit `.env` files containing real credentials; they are already listed in `.gitignore`.
- When adding new dependencies, run `npm install <package>` (frontend) or `go get <module>` (backend) and commit the updated lockfiles (`package-lock.json`, `go.sum`).
- After changing TypeScript configuration (`tsconfig*.json`), run `npm run build` to verify type safety before committing.
- The React Compiler may affect debugging; consult the Vite and React Compiler documentation if you encounter unexpected behavior in development.