# SQL Cipher

The `sqlite3.c` and `sqlite3.h` files were generated from the [SQLCipher](https://github.com/sqlcipher/sqlcipher.git) source
using the following procedure.

Clone the [SQLCipher repository](https://github.com/sqlcipher/sqlcipher.git) and check out the `v4.6.1` tag:

```bash
git clone -b v4.6.1 https://github.com/sqlcipher/sqlcipher.git
cd sqlcipher
```

Apply [`sqlcipher.diff`](./sqlcipher.diff) unified patch:

```bash
patch < /path/to/sigtop/sqlcipher/sqlcipher.diff
```

Generate `sqlite3.c` and `sqlite3.h`:

```bash
./configure --enable-tempstore=yes CFLAGS=-DSQLITE_HAS_CODEC
make sqlite3.c
```

Move `sqlite3.c` and `sqlite3.h` into place:

```bash
mv sqlite3.[ch] /path/to/sigtop/sqlcipher
```
