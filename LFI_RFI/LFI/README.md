# Local File Inclusion

## Concept

Attackers exploit vulnerabilities in web applications by manipulating user - controllable inputs to include local files, thereby accessing sensitive server - side information or executing malicious code.

## Attack Principle

When applications use user inputs to generate file paths or URLs without strict validation and sanitization, attackers can input malicious paths to make the application include local files, such as system configuration files and password files.

## Attack Instance

Take a website with the URL http://example.com/index.php?page=home  as an example. The`page` parameter specifies the page file to be included. If an attacker changes the URL to http://example.com/index.php?page=../../etc/passwd  and the application has an LFI vulnerability, the server's`/etc/passwd` file content might be displayed.
