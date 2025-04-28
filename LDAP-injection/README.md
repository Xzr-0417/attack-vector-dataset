# LDAP-injection

## Concept

LDAP (Lightweight Directory Access Protocol) injection vulnerabilities allow attackers to compromise the authentication process of certain websites. This vulnerability occurs in websites that use data provided by end users to construct LDAP statements. LDAP directories are used to store access credentials in the form of objects. The information may be used by a wide range of entities, including users, roles, printers, and servers. When LDAP directories are used for website authentication purposes, threat actors can inject malicious code into user input fields. The actor can then gain unauthorized access to the LDAP directory, where the actor can view or modify usernames and passwords.

## Attack Principle

In LDAP injection attacks, a threat actor plays with the query input to get unauthorized access into the storage directory. As this directory may contain your organization’s or end-users’ e-mails, usernames, and passwords, such intrusion could have fatal results. Attackers look for security loopholes in a digital solution using LDAP services. Taking advantage of the inefficient filtering mechanism, he sends unsanitized user input data (e.g. LDAP queries to crack into the system) and tries to enter the application. Upon succeeding, he may change, delete, misuse, or add access data, depending upon what is his intention.

## Attack Example

A basic LDAP injection example is an attacker bypassing authentication. Consider an LDAP search filter that accepts two fields via a web form—USER and PASSWORD. If the LDAP filter accepts the USER parameter as is, with no sanitization of control characters, an attacker can input a username followed by control characters that break authentication. For example, if the attacker provides this as the USER value: `admin)(|(memberOf=*)`, the LDAP query is constructed as follows: `(&(USER=admin)(PASSWORD=Pwd))` becomes `(&(USER=admin)(|(memberOf=*))(PASSWORD=Pwd))`. This causes the LDAP server to only process the first part of the query, indicated in bold. The password is not evaluated at all, meaning the attacker can provide any string for password, and gain access.
