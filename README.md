# Patent_Page

Patent_Page is an open-source, self-hosted toolkit for experimenting with AI-assisted patent drafting, prior-art retrieval, patent comparison, and structured report generation.

The repository combines a Vue-based web interface, a Go-based administration backend, Python/FastAPI services, local LLM integration through Ollama, patent retrieval helpers, and automated Word document generation.

> Project status: early-stage / research-oriented. Some routes and integrations are still prototype implementations and may require configuration before production use.

## Features

- Patent and prior-art retrieval workflows.
- AI-assisted patent disclosure drafting.
- Prior-art comparison and novelty/innovation analysis.
- Local LLM integration through Ollama.
- Structured `.docx` report generation with `python-docx`.
- Web/API components for integrating patent workflows into applications.
- Vue-based user and administration interfaces.
- Go/Gin-based administration backend with database and Swagger support.

## Repository Structure

```text
Patent_Page/
├── PatentAdminPlat-UI-master/       # Vue user-facing frontend
├── PatentAdminPlatAdmin-UI-main/   # Vue administration frontend
├── patent-admin-plat-master/       # Go/Gin administration backend
├── ollama/                         # Python/FastAPI + Ollama patent AI modules
│   ├── main.py
│   ├── patent_disclosure_generator.py
│   └── New_innovation_analysis_report.py
├── chonghu.py                      # Patent retrieval-related helper code
├── Image.py                        # Image-related helper code
├── 搭建教程.pdf                    # Existing setup guide
├── LICENSE
└── README.md
```

## Core AI Workflow

The Python modules under `ollama/` demonstrate a workflow that can:

1. Receive patent-related technical content.
2. Retrieve related patent information from an external patent service.
3. Build prompts using the submitted invention content and retrieved prior art.
4. Call locally deployed language models through Ollama.
5. Generate patent disclosure or novelty-analysis content.
6. Export structured Word documents.

The current scripts reference models such as `qwen2:0.5b` and `llama2`. You may replace them with other Ollama-compatible models after reviewing prompt and output behavior.

## Technology Stack

### Frontend

- Vue 2
- Vue Router
- Vuex
- Element UI
- Axios

### Backend and services

- Go 1.19
- Gin
- GORM
- FastAPI
- Python

### AI and document processing

- Ollama
- `python-docx`
- Requests
- BeautifulSoup / lxml

## Quick Start

### 1. Clone the repository

```bash
git clone https://github.com/1308100560/Patent_Page.git
cd Patent_Page
```

### 2. Start the Vue frontend

```bash
cd PatentAdminPlat-UI-master
npm install
npm run dev
```

The frontend is based on Vue 2 / Vue CLI. See the README files inside the frontend directories for additional details.

### 3. Build and start the Go backend

```bash
cd patent-admin-plat-master
bash build_linux.sh
```

Initialize the database:

```bash
bin/PatentAdminPlat-linux-amd64 migrate -c config/settings_dev.yml
```

Start the server:

```bash
bin/PatentAdminPlat-linux-amd64 server -c config/settings_dev.yml
```

Database connection and other runtime settings should be reviewed in the backend configuration before deployment.

### 4. Run the Python/FastAPI prototype

Install the Python dependencies required by the current modules, for example:

```bash
pip install fastapi uvicorn python-multipart jinja2 python-docx requests beautifulsoup4 lxml ollama
```

Install and start Ollama, then pull the model used by the patent disclosure generator:

```bash
ollama pull qwen2:0.5b
```

The novelty-analysis script also references `llama2`:

```bash
ollama pull llama2
```

Run the FastAPI development server:

```bash
cd ollama
uvicorn main:app --reload
```

The current `ollama/main.py` contains prototype routes. The patent document generation call in the form submission route is currently commented out, so review and enable the intended workflow before using it as a complete service.

## Security Considerations

Patent_Page processes user-supplied content, performs external HTTP requests, interacts with local LLM services, and generates files. Before exposing any component to untrusted users or the public internet, review at least the following areas:

- Input validation and output encoding.
- Authentication and authorization.
- External HTTP request handling and timeouts.
- File paths and generated document handling.
- Dependency versions and known vulnerabilities.
- Model prompt/input boundaries and sensitive-data handling.
- Database credentials and deployment configuration.

Do not treat generated patent content or novelty analysis as legal advice. Human review is required before relying on generated material for intellectual-property decisions or patent filings.

## Development Status and Roadmap

Planned maintenance areas include:

- Improve installation and deployment documentation.
- Add automated tests for the AI and document-generation workflows.
- Improve error handling and input validation.
- Review and update dependencies.
- Harden API and file-generation security.
- Simplify repository structure and remove development artifacts that are not required for distribution.
- Establish tagged releases and release notes.

## Contributing

Issues and pull requests are welcome. When contributing, please keep changes focused, explain the motivation for the change, and include relevant testing or verification information where possible.

For security-sensitive issues, avoid publishing secrets, credentials, personal information, or exploit details that could put deployments at risk.

## License

The original code in this repository is distributed under the MIT License. See the root [`LICENSE`](./LICENSE) file.

This repository also contains third-party open-source components. Their original copyright notices and license files remain applicable and must be preserved. In particular:

- `PatentAdminPlat-UI-master/` contains code based on `vue-element-admin`, Copyright (c) 2017-present PanJiaChen, MIT License.
- `patent-admin-plat-master/` contains code from the go-admin ecosystem, Copyright (c) 2020 go-admin-team, MIT License.
- Other bundled third-party components may contain their own license notices; those notices take precedence for the corresponding third-party code.

The root MIT License does not replace or remove third-party copyright and license obligations.
