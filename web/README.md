# Vite + React Project Scripts

This document provides details about the available npm scripts in the project.

---

## **Scripts**

### 1. **Development Server**

```bash
npm run dev
```

Starts the development server using Vite.

- Watches for file changes and updates the browser in real-time (Hot Module
  Replacement).
- Default server URL: `http://localhost:5173` (or as configured).

---

### 2. **Build**

```bash
npm run build
```

Builds the project for production.

- Compiles TypeScript files using `tsc -b`.
- Bundles and optimizes the app using Vite for deployment.

The output is stored in the `dist` folder.

---

### 3. **Lint**

```bash
npm run lint
```

Runs ESLint to check for code quality and adherence to linting rules.

- Outputs warnings and errors for resolution.

---

### 4. **Preview**

```bash
npm run preview
```

Serves the production build locally.

- Useful for verifying the production output before deployment.
