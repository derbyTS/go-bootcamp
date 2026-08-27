### How to Enable TLS in Postman

1. Go to **Settings** > **Certificates**.
2. Add your **PEM File** and toggle **Enable CA Certificates** to **ON**.
3. Click on **Add Certificate**.
4. Enter the following details:
   - **Host:** `localhost:3000`
   - **CRT file:** `cert.pem`
   - **KEY file:** `key.pem`

### Run in Postman

1. HTTP Verb: Get
2. URL: https://localhost:3000/orders
3. Enable SSL certificates verification: ON
4. Automatically follow redirects: ON
