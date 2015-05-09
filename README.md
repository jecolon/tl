# tl
tl is a go command line transliteration utility. Transliteration in this context is the replacement of non ASCII characters with similar ASCII "equivalents". For example, accented vowels used in many languages are replaced with the unaccented form: á é í ó ú ñ -> a e i o u n .

```bash
$ ./tl [-f, --file filename]
```

# Sample run:
```bash
$ echo "áéíóúñ" | ./tl 
aeioun
```
