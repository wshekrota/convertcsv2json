# convertcsv2json
Convert a standard headered csv to valid json with intention to use internally.

Excel allows export to csv and is a great model for this. The following text report 
will convert to json..
```
a, b, c
1, 2, 3
4, 5, 6
7, 8, 9

Will become..
[{"a":1,"b":2,"c":3},
{"a":4,"b":5 etc} .. ]
```
This is a suitable string for decode or unmarshal. Decoded being an array of hash, the column
name can be maintained as you loop through the array processing the map data elements. Using a
struct to do this is less flexible and static. By mapping the column headers to data for each 
row we make it instantly addressable using a key and line number. This is a more dynamic 
process.
To consume this interface you probably know you cannot just index it you must range over it or 
convert it to a slice of the types in question.

example output:
wshek@pop-os:~/go/src/src/github.com/wshekrota/convertcsv2json$ ./convertcsv2json 
```
<nil>    (return code from unmarshal)
[map[a:1 b:2 c:3] map[a:4 b:5 c:6] map[a:7 b:8 c:9]]  (buffer returned on unmarshal)
0 => map[a:1 b:2 c:3]  (ranging through interface returned from cvtc2j())
1 => map[a:4 b:5 c:6]
2 => map[a:7 b:8 c:9]
```
Remember cvtc2j returns a string of json.. that string is then decoded to internal.
