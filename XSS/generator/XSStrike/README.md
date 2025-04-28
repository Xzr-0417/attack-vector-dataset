# XSS Payload Generator

This generator script originates from the XSStrike repository. It integrates utils.py, jsContexter.py, and config.py into generator.py. In the main section, it simulates occurrences (the contexts where XSS tags appear, such as HTML and attributes) and response (HTTP response content). The goal is to grasp the mechanics of payload generation.

### Overview

#### Core Configuration
The script includes a core configuration section that defines:
- **XSS Check String**: A unique string (`v3dm0s`) used to identify successful payload execution.
- **Blacklisted Tags**: Tags such as `iframe`, `title`, and `textarea` that are typically filtered by security mechanisms.
- **Allowed Tags**: Tags like `html`, `a`, and `details` that can be leveraged for XSS.
- **Fillings**: Character sequences used to obfuscate payloads and bypass simple filters.
- **Event Handlers**: Events like `ontoggle`, `onmouseover`, and `onpointerenter` that can trigger JavaScript execution.
- **JavaScript Functions**: Functions like `confirm()` and `prompt()` that can be used to test for XSS vulnerabilities.

#### JavaScript Context Closure
The `jsContexter` function is designed to close JavaScript contexts properly. It analyzes the context before the XSS check string and generates the necessary closing characters to ensure the payload executes correctly.

#### Utility Functions
- **Random Case Modification**: The `randomUpper` function randomizes the case of characters in a string to bypass case-sensitive filters.
- **Script Extraction**: The `extractScripts` function identifies where the XSS check string appears in script tags within a response.

#### Payload Generation Logic
The `generator` function is the core of the payload generation process. It:
- Analyzes the context in which the XSS check string appears (HTML, attribute, script).
- Generates payloads specifically designed for each context.
- Prioritizes payloads based on their likelihood of success.

### Features

- **Context-Aware Payloads**: Generates payloads tailored to different contexts:
  - **HTML Context**: Payloads that break out of HTML and inject script tags.
  - **Attribute Context**: Payloads that close attributes and inject event handlers.
  - **Script Context**: Payloads that execute within JavaScript code.

- **Obfuscation Techniques**: Uses various methods to obfuscate payloads:
  - Random case modification.
  - URL encoding and other character transformations.
  - Insertion of filler characters and comments.

- **Event Handler Utilization**: Leverages different event handlers to trigger JavaScript execution in various scenarios.

- **Prioritized Payloads**: Generates payloads with different priority levels based on their likelihood of success.

### Usage

1. **Clone the Repository**:

```bash
git clone https://github.com/Xzr-0417/attack-vector-dataset/tree/main/XSS/generator/XSStrike
cd xss-generator
```

2. **Generate Payloads**:

Run the Python script to generate XSS payloads:

```bash
python xss_generator.py
```

3. **Inspect Payloads**:

The script will output generated payloads to the console. You can integrate these payloads into your testing workflow to identify XSS vulnerabilities in web applications.

4. **Customize Parameters**:

Modify the parameters in the core configuration section of the Python script to adjust the types of payloads generated based on your testing requirements.

### Example Output

When run with the provided test cases, the script generates payloads like:

```
Generated XSS Payloads:

▶ Priority 10 (共 24 条):
  01. </TEXTAREA><html%09ontoggle=%09coFirM()%0D>v3dm0s
  02. </TEXTAREA><html%09ontoggle=%09(confirm)()%0D>v3dm0s
  03. </TEXTAREA><html%09ontoggle=%09a=prompt,a()%0D>v3dm0s
  ...
```

These payloads can be tested against web applications to verify XSS vulnerabilities.

### Note

This tool is designed for authorized security testing and research purposes only. Always ensure you have proper authorization before testing any system.
