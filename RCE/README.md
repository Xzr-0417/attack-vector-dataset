# Remote Code Execution

## Concept

RCE vulnerabilities allow attackers to remotely inject and execute arbitrary code on a target system. This can happen when an application uses untrusted user input without proper validation and sanitization.

## Attack Principle

When an application processes user input directly into code execution functions like Python’s `eval()` or PHP’s `system()`, it creates a pathway for attackers to inject malicious code. For example, if an application allows users to input a calculation formula, an attacker might input a command that executes system-level operations.

## Attack Example

Consider a web application that provides a calculator feature. The application uses the `eval()` function to evaluate user input as Python code. If the input is not properly sanitized, an attacker can input something like `__import__('os').system('ls')` to list the files in the system directory.
