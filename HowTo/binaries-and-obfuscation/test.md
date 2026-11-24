```sh
h2load -n 1000 -c 100 -m 10 -D 20s -H "Cookie: Bearer=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTU0NDM0OTksInJvbGUiOiJhZG1pbiIsInVpZCI6MzIsInVzZXIiOiJhbGljZS5zbWl0aCJ9.MqRPxpsMG9rnyMybgAnx-Y3SfIufGgsHrlSLCvhtm0Q" -H "Origin: https://localhost:3000" 'https://localhost:3000/students?limit=300&page=1'
```

## Binaries

./binaries/rest_api_macOs_with_middleware_http2_tls
./binaries/rest_api_macOs_no_TLS_with_middleware
./binaries/rest_api_macOs_with_TLS_No_Middleware
./binaries/rest_api_macOs_no_TLS_No_Middleware

## Obfuscation

./obfuscated/rest_api_macOs_with_middleware_http2_tls
./obfuscated/rest_api_macOs_no_TLS_with_middleware
./obfuscated/rest_api_macOs_with_TLS_No_Middleware
./obfuscated/rest_api_macOs_no_TLS_No_Middleware

---

---

---

---

---

---

---

---

## Binaries

./binaries/rest_api_macOs_with_middleware_http2_tls
h2load -n 1000 -c 100 -m 10 -D 20s -H "Cookie: Bearer=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTU1MjA4NDYsInJvbGUiOiJhZG1pbiIsInVpZCI6MzIsInVzZXIiOiJhbGljZS5zbWl0aCJ9.18q3t0fEHZAhfLcsfZKOg11yqo_xUz1ZhiJ53F2fT3c" -H "Origin: https://localhost:3000" 'https://localhost:3000/students?limit=300&page=1'
finished in 20.01s, 2183.55 req/s, 68.10MB/s
requests: 43671 total, 44671 started, 43671 done, 43671 succeeded, 0 failed, 0 errored, 0 timeout
status codes: 43959 2xx, 0 3xx, 0 4xx, 0 5xx
traffic: 1.33GB (1428183589) total, 1.34MB (1402218) headers (space savings 96.32%), 1.33GB (1425194739) da
ta
min max median p95 p99 mean sd +/- sd
request : 3.42ms 1.68s 423.92ms 843.61ms 1.06s 451.34ms 215.55ms 68.85%
connect : 6.76ms 46.25ms 18.24ms 43.71ms 46.25ms 21.91ms 11.21ms 66.00%
TTFB : 50.12ms 380.89ms 157.65ms 280.79ms 380.89ms 167.03ms 68.69ms 66.00%
req/s : 20.30 23.55 21.85 22.55 23.55 21.83 0.50 69.00%

---

./binaries/rest_api_macOs_no_TLS_with_middleware
h2load --h1 -n 1000 -c 100 -m 10 -D 20s -H "Cookie: Bearer=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTU1Mjg2NjYsInJvbGUiOiJhZG1pbiIsInVpZCI6MzIsInVzZXIiOiJhbGljZS5zbWl0aCJ9.cx-0WYR0ZHFNwAPBdE3j8EYNcFKEOPhWR0sKR1Ag2Cw" -H "Origin: https://localhost:3000" 'http://localhost:3000/students?limit=300&page=1'
finished in 20.03s, 1948.85 req/s, 62.17MB/s
requests: 38977 total, 39977 started, 38977 done, 38977 succeeded, 0 failed, 0 errored, 0 timeout
status codes: 38997 2xx, 0 3xx, 0 4xx, 0 5xx
traffic: 1.21GB (1303812136) total, 32.75MB (34342935) headers (space savings 0.00%), 1.18GB (1264321737) d
ata
min max median p95 p99 mean sd +/- sd
request : 19.72ms 747.99ms 507.93ms 610.09ms 653.01ms 506.23ms 72.08ms 77.67%
connect : 8us 292us 146us 284us 292us 155us 83us 56.00%
TTFB : 13.75ms 77.95ms 48.06ms 75.02ms 77.95ms 47.85ms 11.09ms 81.00%
req/s : 18.35 20.75 19.45 20.40 20.75 19.49 0.49 72.00%

---

./binaries/rest_api_macOs_with_TLS_No_Middleware
h2load -n 1000 -c 100 -m 10 -D 20s -H "Cookie: Bearer=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTU1MjE2NjMsInJvbGUiOiJhZG1pbiIsInVpZCI6MzIsInVzZXIiOiJhbGljZS5zbWl0aCJ9.FcVaArYvfbv5mhlPescd8QdWJmWs92_UwB-Vf1MM1ts" -H "Origin: https://localhost:3000" 'https://localhost:3000/students?limit=300&page=1'
finished in 20.01s, 10087.25 req/s, 312.27MB/s
requests: 201745 total, 202745 started, 201745 done, 201745 succeeded, 0 failed, 0 errored, 0 timeout
status codes: 201746 2xx, 0 3xx, 0 4xx, 0 5xx
traffic: 6.10GB (6548715305) total, 639.39KB (654738) headers (space savings 95.43%), 6.09GB (6540791029) d
ata
min max median p95 p99 mean sd +/- sd
request : 807us 697.42ms 83.13ms 231.92ms 323.49ms 98.66ms 68.54ms 73.61%
connect : 5.51ms 38.99ms 20.22ms 37.97ms 38.99ms 21.01ms 10.62ms 57.00%
TTFB : 40.23ms 126.10ms 74.57ms 123.59ms 126.10ms 78.51ms 19.63ms 62.00%
req/s : 96.39 104.70 100.71 103.88 104.70 100.86 1.61 72.00%

---

./binaries/rest_api_macOs_no_TLS_No_Middleware
h2load --h1 -n 1000 -c 100 -m 10 -D 20s -H "Cookie: Bearer=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTU1MjcyODAsInJvbGUiOiJhZG1pbiIsInVpZCI6MzIsInVzZXIiOiJhbGljZS5zbWl0aCJ9.cIyRhsSP-fAbTeqNG2ky-n7jOeQX_NnNi5HKnzNAztU" -H "Origin: https://localhost:3000" 'http://localhost:3000/students?limit=300&page=1'
finished in 20.00s, 11297.60 req/s, 350.70MB/s
requests: 225952 total, 226952 started, 225952 done, 225952 succeeded, 0 failed, 0 errored, 0 timeout
status codes: 225952 2xx, 0 3xx, 0 4xx, 0 5xx
traffic: 6.85GB (7354737600) total, 18.32MB (19205920) headers (space savings 0.00%), 6.82GB (7325589792) d
ata
min max median p95 p99 mean sd +/- sd
request : 6.12ms 221.26ms 86.76ms 118.76ms 137.51ms 88.29ms 17.42ms 70.69%
connect : 6us 3.40ms 2.68ms 3.39ms 3.40ms 2.19ms 1.19ms 70.00%
TTFB : 9.08ms 28.34ms 15.29ms 26.38ms 28.34ms 16.35ms 5.03ms 63.00%
req/s : 110.35 117.05 112.87 115.40 117.05 112.97 1.37 67.00%

---

## Obfuscation

./obfuscated/rest_api_macOs_with_middleware_http2_tls
h2load -n 1000 -c 100 -m 10 -D 20s -H "Cookie: Bearer=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTU1MjE0NjUsInJvbGUiOiJhZG1pbiIsInVpZCI6MzIsInVzZXIiOiJhbGljZS5zbWl0aCJ9.o124Sqhl4GqppA5X6TnRHYMjYDPAIHFNjYslyNCf5x0" -H "Origin: https://localhost:3000" 'https://localhost:3000/students?limit=300&page=1'
finished in 20.01s, 2192.15 req/s, 68.39MB/s
requests: 43843 total, 44843 started, 43843 done, 43843 succeeded, 0 failed, 0 errored, 0 timeout
status codes: 44143 2xx, 0 3xx, 0 4xx, 0 5xx
traffic: 1.34GB (1434150894) total, 1.33MB (1397543) headers (space savings 96.34%), 1.33GB (1431160203) da
ta
min max median p95 p99 mean sd +/- sd
request : 6.45ms 1.73s 421.35ms 843.46ms 1.05s 448.94ms 214.25ms 69.00%
connect : 9.88ms 46.64ms 27.16ms 42.34ms 46.64ms 25.65ms 10.50ms 56.00%
TTFB : 62.35ms 367.67ms 160.47ms 277.06ms 367.67ms 173.75ms 63.02ms 64.00%
req/s : 20.70 22.80 22.00 22.70 22.80 21.92 0.49 65.00%

---

./obfuscated/rest_api_macOs_no_TLS_with_middleware
h2load --h1 -n 1000 -c 100 -m 10 -D 20s -H "Cookie: Bearer=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTU1Mjg2NjYsInJvbGUiOiJhZG1pbiIsInVpZCI6MzIsInVzZXIiOiJhbGljZS5zbWl0aCJ9.cx-0WYR0ZHFNwAPBdE3j8EYNcFKEOPhWR0sKR1Ag2Cw" -H "Origin: https://localhost:3000" 'http://localhost:3000/students?limit=300&page=1'
finished in 20.03s, 1993.30 req/s, 63.59MB/s
requests: 39866 total, 40866 started, 39866 done, 39866 succeeded, 0 failed, 0 errored, 0 timeout
status codes: 39885 2xx, 0 3xx, 0 4xx, 0 5xx
traffic: 1.24GB (1333498551) total, 33.50MB (35123159) headers (space savings 0.00%), 1.20GB (1293144006) d
ata
min max median p95 p99 mean sd +/- sd
request : 30.73ms 795.03ms 496.35ms 599.16ms 645.32ms 495.16ms 71.64ms 77.16%
connect : 9us 268us 145us 260us 268us 145us 74us 59.00%
TTFB : 30.81ms 94.64ms 45.23ms 60.31ms 94.64ms 44.98ms 9.75ms 74.00%
req/s : 19.10 20.90 19.95 20.60 20.90 19.93 0.42 63.00%

---

./obfuscated/rest_api_macOs_with_TLS_No_Middleware
h2load -n 1000 -c 100 -m 10 -D 20s -H "Cookie: Bearer=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTU1MjcyODAsInJvbGUiOiJhZG1pbiIsInVpZCI6MzIsInVzZXIiOiJhbGljZS5zbWl0aCJ9.cIyRhsSP-fAbTeqNG2ky-n7jOeQX_NnNi5HKnzNAztU" -H "Origin: https://localhost:3000" 'https://localhost:3000/students?limit=300&page=1'
finished in 20.01s, 11059.10 req/s, 342.35MB/s
requests: 221182 total, 222182 started, 221182 done, 221182 succeeded, 0 failed, 0 errored, 0 timeout
status codes: 221182 2xx, 0 3xx, 0 4xx, 0 5xx
traffic: 6.69GB (7179623874) total, 696.29KB (713000) headers (space savings 95.46%), 6.68GB (7170941622) d
ata
min max median p95 p99 mean sd +/- sd
request : 888us 769.75ms 75.55ms 212.48ms 297.93ms 90.03ms 62.97ms 74.07%
connect : 9.28ms 31.81ms 21.80ms 30.66ms 31.81ms 20.61ms 6.52ms 57.00%
TTFB : 35.42ms 132.49ms 65.69ms 105.63ms 132.49ms 69.42ms 20.17ms 69.00%
req/s : 107.24 114.64 110.38 113.83 114.64 110.57 1.76 67.00%

---

./obfuscated/rest_api_macOs_no_TLS_No_Middleware
h2load --h1 -n 1000 -c 100 -m 10 -D 20s -H "Cookie: Bearer=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTU1MjcyODAsInJvbGUiOiJhZG1pbiIsInVpZCI6MzIsInVzZXIiOiJhbGljZS5zbWl0aCJ9.cIyRhsSP-fAbTeqNG2ky-n7jOeQX_NnNi5HKnzNAztU" -H "Origin: https://localhost:3000" 'http://localhost:3000/students?limit=300&page=1'
finished in 20.00s, 11721.10 req/s, 363.85MB/s
requests: 234422 total, 235422 started, 234422 done, 234422 succeeded, 0 failed, 0 errored, 0 timeout
status codes: 234422 2xx, 0 3xx, 0 4xx, 0 5xx
traffic: 7.11GB (7630436100) total, 19.00MB (19925955) headers (space savings 0.00%), 7.08GB (7600228083) d
ata
min max median p95 p99 mean sd +/- sd
request : 7.67ms 234.01ms 83.33ms 115.96ms 135.86ms 85.12ms 17.51ms 71.16%
connect : 5us 3.86ms 2.47ms 3.78ms 3.86ms 2.15ms 1.24ms 63.00%
TTFB : 10.82ms 35.29ms 18.01ms 29.94ms 35.29ms 18.61ms 5.46ms 70.00%
req/s : 113.80 121.05 117.42 119.34 121.05 117.21 1.49 66.00%

---

| Build Type                | Binary Variant              | Protocol     | Throughput (`req/s`) | Succeeded (`2xx`) | Failed (`4xx`) | Request Median | Request p99 | Data Transferred |
| :------------------------ | :-------------------------- | :----------- | :------------------- | :---------------- | :------------- | :------------- | :---------- | :--------------- |
| **Standard (`go build`)** | `with_middleware_http2_tls` | HTTP/2 + TLS | `2,183.55 req/s`     | `43,959`          | `0`            | `423.92ms`     | `1.06s`     | `1.33 GB`        |
| **Standard (`go build`)** | `no_TLS_with_middleware`    | HTTP/1.1     | `1,948.85 req/s`     | `38,997`          | `0`            | `507.93ms`     | `653.01ms`  | `1.18 GB`        |
| **Standard (`go build`)** | `with_TLS_No_Middleware`    | HTTP/2 + TLS | `10,087.25 req/s`    | `201,746`         | `0`            | `83.13ms`      | `323.49ms`  | `6.09 GB`        |
| **Standard (`go build`)** | `no_TLS_No_Middleware`      | HTTP/1.1     | `11,297.60 req/s`    | `225,952`         | `0`            | `86.76ms`      | `137.51ms`  | `6.82 GB`        |
| **Obfuscated (`garble`)** | `with_middleware_http2_tls` | HTTP/2 + TLS | `2,192.15 req/s`     | `44,143`          | `0`            | `421.35ms`     | `1.05s`     | `1.33 GB`        |
| **Obfuscated (`garble`)** | `no_TLS_with_middleware`    | HTTP/1.1     | `1,993.30 req/s`     | `39,885`          | `0`            | `496.35ms`     | `645.32ms`  | `1.20 GB`        |
| **Obfuscated (`garble`)** | `with_TLS_No_Middleware`    | HTTP/2 + TLS | `11,059.10 req/s`    | `221,182`         | `0`            | `75.55ms`      | `297.93ms`  | `6.68 GB`        |
| **Obfuscated (`garble`)** | `no_TLS_No_Middleware`      | HTTP/1.1     | `11,721.10 req/s`    | `234,422`         | `0`            | `83.33ms`      | `135.86ms`  | `7.08 GB`        |
