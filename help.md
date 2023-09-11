# return all entries of the db as table to db_entries.txt

```sh
sqlite3 movies.db "select * from movies;" --table > db_entries.txt
```

```sh
go init mod
go mod tidy
```
