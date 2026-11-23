```sh
h2load -n 1000 -c 100 -m 10 -D 20s -H "Cookie: Bearer=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTUzNTk4MTEsInJvbGUiOiJhZG1pbiIsInVpZCI6MzIsInVzZXIiOiJhbGljZS5zbWl0aCJ9.55IZZ3fnX8AWKi5AvzNCiiXAF9oTp9cW2NBCX_hXAlc" -H "Origin: https://localhost:3000" 'https://localhost:3000/students?limit=300&page=1'
```
