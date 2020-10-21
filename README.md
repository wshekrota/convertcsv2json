# convertcsv2json
Convert a standard headered csv to valid json with intention to use internally.

Excel allows export to csv and is a great model for this. The following text report 
will convert to json..
a, b, c
1, 2, 3
4, 5, 6
7, 8, 9

Will become..
[{"a":1,"b":2,"c":3},
{"a":4,"b":5 etc} .. ]

This is a suitable string for decode or unmarshal. Decoded being an array of hash the column
name can be maintained as you loop through the array processing the map data elements. Using a
struct to do this is less flexible and static. By mapping the column headers to data for each 
row we make it instantly addressable using a key and line number. This is a more dynamic 
process
