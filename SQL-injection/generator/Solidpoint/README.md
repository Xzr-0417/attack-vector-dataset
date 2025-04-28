# SQL Payload Generator

This generator script is from the sql-injection-scanner repository in solidpoint. It uses a collaborator server to accept blind-type payloads. When the target system executes these commands, it sends requests to the collaborator server, triggering a callback. If the server receives the request, it means the command ran successfully. But in this script, we use the rand package to make a random token, part of the payload, with the seconds parameter set to 1-10s, to understand payload generation.

## Overview

### Components

#### Transformers and Encoding

- **Transformers**: Functions that modify or format payloads, such as escaping characters or applying specific command structures.
- **Encoding Handlers**: Different encoding methods for various parameter contexts (e.g., URL encoding).

#### Payload Categories

- **OOB (Out-of-Band) Payloads**: Payloads designed for out-of-band detection using techniques like pingbacks.
- **Time-based Payloads**: Payloads that introduce delays to detect vulnerabilities based on response time.

### Features

- **Customizable Payloads**: Supports different types of payloads and command structures.
- **Flexible Encoding**: Escapes and formats payloads to bypass security filters.
- **Context-aware Generation**: Generates payloads adapted to different SQL injection contexts.

## Contents

- **`main.go`**: Go script for generating the payloads.
- **Generated Payloads**: Output file containing all the generated SQL injection payloads.

## Usage

1. **Clone the repository**:

   ```bash
   git clone https://github.com/Xzr-0417/attack-vector-dataset/tree/main/SQL-injection/generator/Solidpoint
   cd sql-payload-generator
   ```

2. **Generate Payloads**:

   Run the Go script to generate the payloads:

   ```bash
   go run sql_injection_generator.go -output sql_payloads.txt
   ```

   This will create a file named `sql_payloads.txt` containing all the generated SQL injection payloads.

3. **Inspect Payloads**:

   You can view the generated payloads directly in the `sql_payloads.txt` file. These payloads can be tested against vulnerable applications to verify SQL injection vulnerabilities.

4. **Customize Parameters**:

   Modify the transformer functions, payloads, and encoding handlers in the Go script to adjust the types of payloads generated based on your testing requirements.

## Example

### Example Command

To generate payloads and save them to `my_sql_payloads.txt`:

```bash
go run main.go -output my_sql_payloads.txt
```

This will generate SQL injection payloads with the specified settings and save them to `my_sql_payloads.txt`.
