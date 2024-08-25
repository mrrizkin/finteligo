# Finteligo

**Task Automation Tool**

## Requirements

Before you get started, ensure you have the following installed:

- **Node.js**
- **pnpm**
- **Go**
- **Air**
- **SQLite**
- **Templ**

## Quick Start

1. **Set Up Environment Variables**

   Copy the example environment file and modify it as needed:

   ```bash
   cp .env.example .env
   ```

2. **Install Dependencies**

   Run the following commands to install the necessary dependencies:

   ```bash
   pnpm install
   go get -u all
   ```

3. **Run the Application**

   Start the development server in one terminal:

   ```bash
   pnpm dev
   ```

   Then, in another terminal, run:

   ```bash
   air
   ```

## Building the Project

To build the project, use:

```bash
pnpm build
go build -o finteligo
```

## Running the Executable

After building, you can run the application with:

```bash
./finteligo
```
