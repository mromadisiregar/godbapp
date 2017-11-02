## Godbapp

This is sample of executing go code with mongo on docker.
---

**How to run this code?**
```
~ docker-compose build
~ docker-compose up -d
```

**Check the counter**
```
~ curl 172.19.0.3:8080       
Halaman telah dibuka 0 kali
<br>
~ curl 172.19.0.3:8080
Halaman telah dibuka 1 kali
<br>
~ curl 172.19.0.3:8080
Halaman telah dibuka 2 kali
<br>
~ curl 172.19.0.3:8080
Halaman telah dibuka 3 kali
<br>
```