# SQL-Injection 

## Concept

SQL injection vulnerabilities allow attackers to manipulate an application's SQL queries by injecting malicious SQL code into user inputs. This can lead to unauthorized access to sensitive data, data tampering, or even complete database compromise.

## Attack Principle

When an application uses user inputs directly in SQL statements without proper validation or sanitization, it creates an entry point for attackers. Attackers insert malicious SQL code into input fields or parameters, which the application then executes as part of its SQL queries, allowing them to alter, retrieve, or delete unauthorized data.

## Attack Example

Consider a login page where the application constructs an SQL query using the username and password provided by the user. If the application does not properly sanitize the user input, an attacker can input a username like `' OR '1'='1` and a password like `' OR '1'='1`. The resulting SQL query becomes `SELECT * FROM users WHERE username = '' OR '1'='1' AND password = '' OR '1'='1'`, which would bypass authentication and grant the attacker access to the system.
