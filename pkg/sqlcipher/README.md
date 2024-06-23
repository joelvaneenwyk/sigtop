# About sqlite3.c and sqlite3.h

The `sqlite3.c` and `sqlite3.h` files were generated from the SQLCipher source
using the following procedure.

Clone the SQLCipher repository and check out the `v4.5.7` tag:

```bash
git clone -b v4.5.7 https://github.com/sqlcipher/sqlcipher.git
cd sqlcipher
```

Apply `sqlcipher.diff`:

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
