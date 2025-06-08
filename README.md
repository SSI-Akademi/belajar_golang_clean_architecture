# Belajar Golang RESTFul dan Clean Architecture

## Cara Menjalankan Projek
1. Lakukan execute command
    ```bash
   go mod vendor
   go mod download
   ```

2. Ketika import atau download library baru lakukan
    ```bash
    go mod tidy
    go mod vendor
    ```

3. Silahkan setting koneksi ke database pada .env**

4. Buatlah database dengan nama articles** (sesuai key DATABASE_NAME** pada .env**) Lakukan eksekusi query pada sql editor
   ```bash
    CREATE TABLE articles(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    content VARCHAR(255) NOT NULL,
    create_at DATETIME NOT NULL,
    update_at DATETIME NOT NULL
    ) ENGINE = INNODB;
    ```

5. Run aplikasi dengan eksekusi command pada terminal : 
   ```bash
   go run cmd/main.go
   ```

6. Untuk test API gunakan postman dan import collection pada file testapi.json**