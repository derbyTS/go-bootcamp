# MariaDB Setup, Troubleshooting & Operations Guide for macOS

This guide covers installing, troubleshooting, managing, and connecting to MariaDB on macOS via Homebrew, Terminal, and DBeaver.

---

## 1. Quick Command Reference

### Homebrew Service Management

- **Start MariaDB Service:**
  ```bash
  brew services start mariadb
  ```
- **Stop MariaDB Service:**
  ```bash
  brew services stop mariadb
  ```
- **Restart MariaDB Service:**
  ```bash
  brew services restart mariadb
  ```
- **Check Service Status:**
  ```bash
  brew services list
  pgrep mariadbd || pgrep mysqld
  ```

### Terminal Login Commands

- **Connect as Root (with password prompt):**
  ```bash
  mariadb -u root -p
  ```
- **Connect as Root (without password):**
  ```bash
  mariadb -u root
  ```
- **Connect via System Admin (Socket Authentication):**
  ```bash
  sudo mariadb -u root
  ```
- **Connect directly to a specific database:**
  ```bash
  mariadb -u root -p database_name
  ```
- **Exit MariaDB Shell:**
  ```sql
  EXIT;
  ```

---

## 2. Essential SQL Commands

```sql
-- View existing databases
SHOW DATABASES;

-- Create a new database
CREATE DATABASE my_database;

-- Select/Switch to a database
USE my_database;

-- View tables in current database
SHOW TABLES;

-- Create a new dedicated user with a password
CREATE USER 'devuser'@'localhost' IDENTIFIED BY 'secure_password';

-- Grant privileges to the user
GRANT ALL PRIVILEGES ON my_database.* TO 'devuser'@'localhost';

-- Apply privilege changes
FLUSH PRIVILEGES;
```

---

## 3. Fixing Root Access & Authentication (`ERROR 1698`)

MariaDB on macOS defaults to `unix_socket` authentication for `root`, requiring `sudo` to log in. To enable password authentication for local development and DBeaver connections:

1. **Log in using `sudo`:**

   ```bash
   sudo mariadb -u root
   ```

2. **Switch root authentication to standard password:**

   ```sql
   ALTER USER 'root'@'localhost' IDENTIFIED VIA mysql_native_password USING PASSWORD('your_password');
   FLUSH PRIVILEGES;
   EXIT;
   ```

3. **Verify standard login:**
   ```bash
   mariadb -u root -p
   ```

---

## 4. Fixing Homebrew Service `Error 1` & Port Locks

If `brew services list` shows `error 1`, stale background processes or corrupted data files may be locking port `3306`.

1. **Force stop and kill all lingering process instances:**

   ```bash
   brew services stop mariadb
   sudo pkill -9 mariadbd
   sudo pkill -9 mysqld
   ```

2. **Fix ownership of Homebrew database directories:**

   ```bash
   sudo chown -R $(whoami):admin $(brew --prefix)/var/mysql
   ```

3. **Re-initialize database files (If data directory is corrupted):**

   ```bash
   rm -rf $(brew --prefix)/var/mysql/*
   mariadb-install-db
   ```

4. **Restart MariaDB:**
   ```bash
   brew services start mariadb
   ```

---

## 5. DBeaver Setup Guide

### Installation via Homebrew Cask

```bash
brew install --cask dbeaver-community
```

### Connection Settings

- **Driver:** MariaDB
- **Server Host:** `127.0.0.1` _(Use `127.0.0.1` instead of `localhost` to force TCP/IP connection over port 3306)_
- **Port:** `3306`
- **Database:** _(Leave blank or enter specific database)_
- **Username:** `root`
- **Password:** Enter the root password set in Section 3 (or leave blank if none).

---

## 6. Recovery Mode Shortcut (Optional)

If you need safe-mode recovery (`--skip-grant-tables`), add a shortcut alias to your `~/.zshrc`:

```bash
echo "alias mariadb-safe='mariadbd --skip-grant-tables --skip-networking &'" >> ~/.zshrc
source ~/.zshrc
```
