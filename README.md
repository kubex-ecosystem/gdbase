#

![GDBASE Banner](docs/assets/top_banner.png)

[![Go](https://img.shields.io/badge/Go-1.19+-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](https://github.com/rafa-mori/gdbase/blob/main/LICENSE)
[![Automation](https://img.shields.io/badge/automation-zero%20config-blue)](#features)
[![Releases](https://img.shields.io/github/v/release/rafa-mori/goforge?include_prereleases)](https://github.com/rafa-mori/goforge/releases)
[![Build](https://github.com/rafa-mori/gdbase/actions/workflows/kubex_go_release.yml/badge.svg)](https://github.com/rafa-mori/gdbase/actions/workflows/kubex_go_release.yml)

---

**Modular, scalable, and automatic database management for modern systems.**

---

## **Table of Contents**

1. [About the Project](#about-the-project)
2. [Features](#features)
3. [Installation](#installation)
4. [Usage](#usage)
    - [CLI](#cli)
    - [Project Structure](#project-structure)
    - [Configuration](#configuration)
5. [Roadmap](#roadmap)
6. [Contributing](#contributing)
7. [Contact](#contact)

---

## **About the Project**

**GDBASE** is a database management solution developed in Go, designed to be **modular, scalable, and automatic**. It allows zero-configuration by default, but supports advanced customizations via configuration files. It manages local databases, Docker, and multiple databases simultaneously, making it ideal for distributed systems.

---

## **Features**

✨ **Dynamic and automatic configuration**

- Randomly generated passwords stored in the keyring.
- Automatically adjusts for occupied ports.

🗄️ **Multi-DB support**

- Redis, RabbitMQ, MongoDB, PostgreSQL, and SQLite ready to use.

🏗️ **Modular architecture**

- Models follow the `Model → Repo → Service` pattern.
- Ensures modularity and organization.

🔐 **SSH tunnel for external databases**

- `gdbase ssh tunnel` securely connects to remote databases via SSH.

⚙️ **Docker orchestration**

- Automatic container generation for portability and easy deployment.

📡 **Monitoring and events**

- Event bus for internal action tracking.

---

## **Installation**

Requirements:

- Go 1.19+
- Docker (for containerized databases)

Clone the repository and build:

```sh
# Clone the repository
git clone https://github.com/rafa-mori/gdbase.git
cd gdbase
go build -o gdbase .
```

---

## **Usage**

### CLI

Start the main server:

```sh
./gdbase start
```

See all available commands:

```sh
./gdbase --help
```

**Main commands:**

| Command      | Function                                             |
|--------------|-----------------------------------------------------|
| `start`      | Initializes `gdbase` and sets up all services       |
| `status`     | Shows status of active databases                    |
| `config`     | Creates a configuration file for customization      |
| `ssh tunnel` | Creates a secure tunnel for external DBs via SSH    |
| `docker`     | Manages Docker containers for databases             |

### Project Structure

The core implementation follows a clear and modular architecture:

```plaintext
./
├── cmd
│   ├── cli
│   ├── gen_models.go
│   ├── models.go
│   ├── main.go
│   ├── usage.go
│   └── wrpr.go
├── docs
│   └── assets
├── go.mod
├── go.sum
├── internal
│   ├── events
│   ├── models
│   └── services
├── tests
└── version
```

---

### Configuration

GDBASE can run without any initial configuration, but supports customization via YAML/JSON files. By default, everything is generated automatically on first use.

Example configuration:

```yaml
postgres:
  host: localhost
  port: 5432
  user: gdbase
  password: secure
redis:
  host: localhost
  port: 6379
```

---

## **Roadmap**

- [x] Dynamic and automatic configuration
- [x] Multi-DB support (Redis, RabbitMQ, MongoDB, PostgreSQL, SQLite)
- [x] Integrated SSH tunnel
- [x] Docker orchestration
- [ ] Plugins for new databases
- [ ] Web dashboard for monitoring

---

## **Contributing**

Contributions are welcome! Feel free to open issues or submit pull requests. See the [Contribution Guide](docs/CONTRIBUTING.md) for more details.

---

## **Contact**

💌 **Developer**:  
[Rafael Mori](mailto:faelmori@gmail.com)  
💼 [Follow me on GitHub](https://github.com/rafa-mori)  
I'm open to collaborations and new ideas. If you found the project interesting, get in touch!

---

**Made with care by the Mori family!** ❤️
