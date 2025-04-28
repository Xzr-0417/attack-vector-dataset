# Cross Site Scripting

## Concept

XSS vulnerabilities enable attackers to inject malicious scripts into web pages viewed by other users. These scripts can execute in users' browsers and access any cookies, session tokens, or other sensitive data the browser keeps for that site. XSS can be categorized into reflected (non-persistent), stored (persistent), and DOM - based types.

## Attack Principle

When a web application doesn't properly validate or sanitize user - input data that's reflected back to the user's browser, it creates an opening for XSS attacks. Attackers can inject malicious scripts into input fields or URLs. When a user visits a page with the injected script, the script executes in their browser, allowing the attacker to hijack sessions, deface web pages, or spread web worms.

## Attack Example

Consider a simple comment section on a website where users can leave comments. If the site doesn't sanitize user input, an attacker can post a comment containing a malicious script, such as:

```html
<script>alert(document.cookie)</script>
```

When another user views the comment, the script will execute in their browser, revealing their cookie information. Attackers can also use more sophisticated scripts to steal session tokens or redirect users to malicious sites.
