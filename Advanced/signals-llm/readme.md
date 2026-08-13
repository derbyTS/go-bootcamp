### Step 1: Start the server in Terminal 1

```bash
go run main.go
```

### Step 2: Trigger the slow request in Terminal 2

```bash
curl http://localhost:8080/work
```

### Step 3: Immediately press Ctrl + C in Terminal 1

#### Press Ctrl + C while curl in Terminal 2 is still waiting!
