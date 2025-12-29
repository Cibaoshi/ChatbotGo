# [English](./README.md)   | [Russian](./README_RU.md)
# 🤖Gemini Chatbot (Go/REST API)

  A simple and robust console chatbot developed in Go (Golang) to interact with the Google Gemini 2.5 Flash language model via a direct REST API call. The project demonstrates fundamental skills in integrating with external AI services and correctly handling network errors.

## Technology stack and tools


-   Programming language: Go (Golang)

-  API: Google Gemini API (`gemini-2.5-flash model`)

-   Communication method: Direct HTTP requests (REST)

-   Data format: JSON

-   Go packages used: `net/http`, `encoding/json`, `os`, `log`

-   Environment: Console (CLI) application (UPD. added web interface)


#  Key features and characteristics

1. Stability and error handling
- Reliable API error handling: HTTP status checks (`200 OK`, `404`, `503`, etc.) are included, ensuring correct programme termination in case of server or API key issues.
- Panic protection: Array length checks (`Candidates`, `Parts`) are implemented before accessing index `[0]`, which completely eliminates the `panic error: runtime error: index out of range`.
- Diagnostics: Parsing and outputting information from the `PromptFeedback` field to diagnose reasons why a response may be blocked (e.g., for security reasons).

2. Architecture and interaction
- Direct REST integration: Use of the standard Go `net/http` package to independently form and send JSON requests, bypassing official SDKs.
- Structured input/output: User queries are accepted via command line arguments (`os.Args`). The model's response is extracted from a complex JSON structure and displayed in plain text.
- Key security: The API key is loaded from an environment variable (`GEMINI_API_KEY`), preventing its accidental publication in the source code. 


#  interface

When hosted, it generates a modern and attractive HTML interface in a browser window.
