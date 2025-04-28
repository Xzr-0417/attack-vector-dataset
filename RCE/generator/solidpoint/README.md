# RCE Payload Generator 

This repository contains a Go-based payload generator designed for generating Remote Code Execution (RCE) payloads, useful for penetration testing, vulnerability assessments, and security research.

The payloads are crafted to test common injection points in web applications, such as URL parameters, form inputs, and other user-controlled input vectors. This tool is focused on generating various command injection patterns and techniques that can be used to simulate potential attack vectors.

## Overview

### Components

#### Transformers and Rarity

- **Transformers**: Functions that modify or format payloads, such as escaping characters or applying specific command structures.
- **Rarity**: Determines how often a transformer is applied, allowing for flexibility in payload generation complexity.

#### Payload Categories

- **Normal Payloads**: Standard command execution payloads.
- **Blind Payloads**: Specifically crafted for blind command injection scenarios.

#### Operating Systems

- **Unix/Linux**: Includes common commands and techniques applicable to Unix-like systems.
- **Windows**: Includes Windows-specific commands and cmd.exe escape patterns.

### Features

- **Customizable Payloads**: Supports different types of payloads and command structures.
- **Flexible Encoding**: Escapes and formats payloads to bypass security filters.
- **Command Chaining**: Generates payloads that chain commands or embed them in different contexts.

## Contents

- **`main.go`**: Go script for generating the payloads.
- **Generated Payloads**: Output file containing all the generated payloads.

## Usage

1. **Clone the repository**:

   ```bash
   git clone https://github.com/Xzr-0417/attack-vector-dataset/tree/main/RCE/generator/solidpoint
   cd solidpoint
   ```

2. **Generate Payloads**:

   Run the Go script to generate the payloads:

   ```bash
   go run main.go -output payloads.txt -rarity 2
   ```

   This will create a file named `payloads.txt` containing all the generated payloads.

3. **Inspect Payloads**:

   You can view the generated payloads directly in the `payloads.txt` file. These payloads can be tested against vulnerable applications to verify RCE vulnerabilities.

4. **Customize Parameters**:

   Modify the transformer functions, payloads, and rarity levels in the Go script to adjust the types of payloads generated based on your testing requirements.

## Example

### Example Command

To generate payloads with a rarity level of 1 and save them to `my_payloads.txt`:

```bash
go run main.go -output my_payloads.txt -rarity 1
```

This will generate payloads with the specified settings and save them to `my_payloads.txt`.
