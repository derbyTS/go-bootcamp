```sh
h2load -n 1000 -c 100 -m 10 -D 20s -H "Cookie: Bearer=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTUzNTcxOTMsInJvbGUiOiJhZG1pbiIsInVpZCI6MzIsInVzZXIiOiJhbGljZS5zbWl0aCJ9.wJaqBt1gEvjHUo3iHSA3z_ab3m2fXMJ-Wa6yD-leu7c" -H "Origin: https://localhost:3000" 'https://localhost:3000/students?limit=300&page=1'
```
