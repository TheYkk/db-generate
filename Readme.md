# Generate leaked passwords db
[Source of passwords](https://haveibeenpwned.com/Passwords)
```
wget https://downloads.pwnedpasswords.com/passwords/pwned-passwords-sha1-ordered-by-count-v7.7z
```

After download zip we need to extract it

```
7z x pwned-passwords-sha1-ordered-by-count-v7.7z
```

This operation will generate `pwned-passwords-sha1-ordered-by-count-v7.txt`

Take First 1 Million passwords

```
head -n 1000000 pwned-passwords-sha1-ordered-by-count-v7.txt > pass.txt
```

Generate db

```
go run . pass.txt
```
