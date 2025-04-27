# RCE Payload Generator

This repository contains a Python-based payload generator designed for generating over **207,000 Remote Code Execution (RCE)** payloads, useful for penetration testing, vulnerability assessments, and security research.

The payloads are crafted to test common injection points in web applications, such as URL parameters, form inputs, and other user-controlled input vectors. This tool is focused on shell command injection, Java, Python, and PowerShell invocations, among others, to simulate potential attack vectors that an attacker might use in real-world exploitation scenarios.

## Components

### Prefixes

Prefixes are used to construct commands and include:

* **Generic shell operators and command separators**: Such as `;`, `|`, `&&`, etc., for chaining commands.
* **Space bypasses and special encoding characters**: Like `${IFS}`, `%0A` (URL-encoded newline), etc.
* **Language-specific invocation prefixes**: E.g., `Runtime.getRuntime().exec("`, `powershell -Command "`, etc., for executing commands within specific programming languages.

### Payloads

Payloads are the core commands or scripts intended for execution. They are categorized as follows:

* **Linux-specific commands**: Including basic system enumeration commands like `id`, `whoami`, file operations, network commands, etc.
* **Windows-specific commands**: Such as `ipconfig /all`, `tasklist`, `net user`, etc.
* **Cross-platform commands**: Commands that can be executed across different platforms or within specific programming languages.

### Suffixes

Suffixes are appended to the end of commands and serve purposes such as:

* **Command comments and termination**: Using symbols like `#`, `//`, `/*` to add comments or terminate command execution.
* **Output redirection and control**: E.g., `2>&1` for redirecting standard error to standard output, `> /tmp/out.txt` for redirecting output to a file.
* **Command separation and execution**: Including `%0A` for newline separation, `;` for separating commands, etc.
* **Command execution with time delays**: Such as `; sleep 5` to introduce a delay after command execution.
* **Character escaping and special character handling**: Using `\` to escape special characters.

---

## Encoding Functions

* **`encode_all`**: Utilizes `urllib.parse.quote` to perform URL encoding on the entire string, ensuring only safe characters are retained.
* **`encode_special`**: Applies URL encoding only to special characters within the string, leaving other characters unchanged.
* **`escape_chars`**: Escapes special characters by adding a backslash (`\`) before each one.

---

## Mode Dictionary

The mode dictionary defines different encoding and escaping modes:

* **`normal`**: No encoding or escaping is applied to the input string; it is returned as-is.
* **`all_encoded`**: Corresponds to the `encode_all` function for full URL encoding.
* **`special_encoded`**: Corresponds to the `encode_special` function for encoding only special characters.
* **`escaped`**: Corresponds to the `escape_chars` function for escaping special characters.

---

## Generation Function

The generation function combines `parameters`, `modes`, and `payloads` to produce all possible permutations of parameters, encoding modes, and payloads.

* It uses nested loops and `itertools.product` to generate combinations of parameters, encoding modes, payloads, prefixes, and suffixes.
* For each combination, it extracts `param`, `mode_name`, and `payload`, retrieves the corresponding transformation function from the mode dictionary, and generates payloads in four different combinations:
  * `param + payload`
  * `param + prefix + payload`
  * `param + payload + suffix`
  * `param + prefix + payload + suffix`

---

## Main Function

The main function executes the payload generation process:

* It iterates over each generated payload from the `generate_all()` function.
* Each payload is printed to the console and simultaneously written to a file named `rce.txt`, with each payload on a new line.


### Contents

- **`generate_payloads.py`**: Python script for generating the payloads.
- **`rce.txt`**: Generated payloads saved in a text file (over 207,000 unique payloads).

### Features

- **Customizable Payloads**: The script supports different types of payloads including:
  - Basic enumeration commands (e.g., `whoami`, `id`, `ls -la`).
  - Advanced discovery commands (e.g., `find / -name '*.conf'`).
  - OS command execution via various injection points (e.g., `bash -c`, `powershell -Command`).
  - HTTP-based exploits (e.g., `curl`, `wget`, `powershell` commands).
  - Language-specific injection payloads (Java, Python, etc.).
  
- **Flexible Encoding**: Payloads are encoded in various formats to bypass security filters and firewalls:
  - URL encoding (`%` encoding for special characters).
  - Special character encoding to avoid detection by simple filters.
  - Escaped characters for payload obfuscation.

- **Combination of Prefixes and Suffixes**: The script generates payloads by applying various shell prefixes and suffixes to simulate different attack scenarios, such as chaining commands, adding comments, or redirecting output.

### Usage

1. **Clone the repository**:

   ```bash
   git clone https://github.com/inpentest/rce
   cd rce
   ```

2. **Generate Payloads**:

Run the Python script to generate the payloads:

```bash
python generate_payloads.py
This will create a file named rce.txt containing all the generated payloads.
```

3. **Inspect Payloads**:

You can view the generated payloads directly in the rce.txt file. These payloads can be tested against vulnerable applications to verify RCE vulnerabilities.

4. **Customize Parameters**:

Modify the parameters, prefixes, payloads, and suffixes lists in the Python script to adjust the types of payloads generated based on your testing requirements.

### Important Note
This tool is intended for ethical penetration testing and security research only. Ensure you have explicit permission to test the target systems before using the generated payloads. Unauthorized use of these payloads may be illegal and unethical.

### License
This repository is licensed under the MIT License. See LICENSE for more information.
