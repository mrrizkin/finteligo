# Finteligo

**Advanced AI-powered fintech solution**

## Requirements

Before you get started, ensure you have the following installed:

- **Node.js**
- **pnpm**
- **Go**
- **Air**
- **SQLite**

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
go build -o finteligo ./main.go
```

## Running the Executable

After building, you can run the application with:

```bash
./finteligo
```

## Systemd Service

To run the application as a systemd service, create a new service file:

```bash
nano /etc/systemd/system/finteligo.service
```

Then, add the following content:

```ini
[Unit]
Description=Finteligo Service
After=network-online.target
Wants=network-online.target systemd-networkd-wait-online.service

StartLimitIntervalSec=500
StartLimitBurst=5

[Service]
Restart=on-failure
RestartSec=5s

WorkingDirectory=/path/to/finteligo
ExecStart=/path/to/finteligo/finteligo

[Install]
WantedBy=multi-user.target
```

Finally, start and enable the service:

```bash
systemctl daemon-reload
systemctl start finteligo
systemctl enable finteligo
```
