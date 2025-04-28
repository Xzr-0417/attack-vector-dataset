# PHP-injection

## Concept

PHP unserialize vulnerabilities arise when PHP unserializes user-supplied data. An attacker can craft malicious serialized data to inject objects and trigger unintended behaviors, such as calling dangerous methods or achieving remote code execution. These vulnerabilities are common in PHP applications using the unserialize() function with untrusted input.

## Attack Principle

When an application uses PHP’s unserialize() function on user input without proper validation, attackers can inject serialized data that creates unexpected objects. These objects may have magic methods (like __wakeup() or __destruct()) that execute code during deserialization. Attackers exploit these methods to perform harmful actions on the server.

## Attack Example

Consider a PHP application that accepts serialized data from a user input parameter and deserializes it. The code might look like this:

```php
<?php
$data = unserialize($_GET['data']);
?>
```

An attacker can craft a serialized payload that creates an object with a malicious magic method. For example, if the application uses a class with a vulnerable __destruct() method that deletes a file, the attacker could send a serialized string representing an instance of that class. Upon deserialization, the __destruct() method would be triggered, deleting the specified file on the server.
