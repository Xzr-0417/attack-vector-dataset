# Server-Side Template Injection

## Concept

SSTI (Server-Side Template Injection) vulnerabilities occur when user input is embedded into server-side templates without proper validation or sanitization. This allows attackers to inject malicious code into templates, potentially leading to remote code execution or data theft. These vulnerabilities often arise in web applications that use template engines like Jinja2, Twig, or Freemarker to generate dynamic content. For example, if a web page dynamically displays user-submitted text using a template engine, it might inadvertently execute unintended expressions.

## Attack Principle

When an application directly incorporates user input into a server-side template engine, attackers can inject malicious code by manipulating the input. They exploit the template engine's ability to interpret placeholders and expressions, which might include logic like conditionals, loops, and function calls. If the template engine isn't properly configured or restricted, it can act as a gateway for attackers to execute arbitrary code.

## Attack Example

Consider a web application that uses Python’s Jinja2 template engine to display user-supplied content. Suppose the application has a page that renders a greeting with the username provided by the user. The code might look like this:

```
from flask import Flask, render_template_string, request
app = Flask(__name__)
@app.route('/greet')
def greet():
    user_input = request.args.get('name')
    template = f"Hello, {{ {user_input} }}!"
    return render_template_string(template)
```

An attacker could input `{{ 7 * 7 }}` as the username. If the application returns `49`, it indicates that the input is being evaluated by a template engine. The attacker can then escalate the attack by injecting more complex payloads like `{{ __import__('os').system('ls') }}` to execute commands on the server and list files.
