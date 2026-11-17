## install

```sh
go install github.com/mailhog/MailHog@latest
~/go/bin/MailHog # This is the path where it i install
export PATH=$PATH:$(go env GOPATH)/bin # Add to ~/.zshrc
```

## 1. Quick Command Reference

- **Start Mailhog:**
  ```bash
  Mailhog
  ```
- **Stop Mailhog:**
  ```bash
  pkill MailHog
  ```
- **Check if running:**

  ```bash
  pgrep -l MailHog
  ```

- **Browser Link:**

```markdown
[http://localhost:8025/](http://localhost:8025/)
```
