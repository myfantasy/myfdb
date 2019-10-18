
#### GetCurrentStruct
GET  
```
localhost:9170/c/struct/
```

### Int table

#### Create int table
POST  
```
localhost:9170/i/create_table/
```
Json  
```
{
	"name":"tst1",
	"type":"table_map_int"
}
```

#### Create int index on int table 
POST  
```
localhost:9170/i/create_index/
```
Json  
```
{
	"tbl_name":"tst1",
	"col_name":"cint1",
	"col_type":"int"
}
```

#### Create string index on int table 
POST  
```
localhost:9170/i/create_index/
```
Json  
```
{
	"tbl_name":"tst1",
	"col_name":"cstring1",
	"col_type":"string"
}
```

#### Set value
POST  
```
localhost:9170/i/set/
```
Json  
```
{
	"tbl_name":"tst1",
	"itm":{
		"key": 45,
		"fi":{
			"a":345,
			"b":34534,
			"cint1":8
		},
		"fs":{
			"m":"SSSS345",
			"n":"SSSS34534",
			"cstring1":"SSSS8"
		}
	}
}
```
Json  
```
{
	"tbl_name":"tst1",
    "full":true,
	"itms":[{
		"key": 145,
		"fia":{
			"a":[345, 324],
			"cint1":[8, 9]
		},
		"fsa":{
			"n":["SSSS34534"],
			"cstring1":["SSSS8", "SSSS9"]
		}
	},
	{
		"key": 146,
		"fia":{
			"a":[345, 24],
			"cint1":[8, 9]
		},
		"fsa":{
			"n":["SSS4"],
			"cstring1":["SSSS8", "SSSS9"]
		},
		"d":"SGVsbG8gd29ybGQh"
	}]
}
```

#### Get value
POST  
```
localhost:9170/i/get/
```
Json  
```
{
	"tbl_name":"tst1",
	"key":146
}
```
Json  
```
{
	"tbl_name":"tst1",
	"key":146,
    "short":true
}
```
Json  
```
{
	"tbl_name":"tst1",
	"keys":[146, 345]
}
```
Json  
```
{
	"tbl_name":"tst1",
	"keys":[146, 145, 45],
    "limit":2
}
```
Json  
```
{
	"tbl_name":"tst1",
	"all": true
}
```
Json  
```
{
	"tbl_name":"tst1",
	"all": true,
    "include_deleted": true,
    "limit":2
}
```

#### Get value by int index

POST  
```
localhost:9170/i/iiget/
```
Json (keys only) 
```
{
	"tbl_name":"tst1",
	"col_name":"cint1",
    "val":8
}
```
json
```
{
	"tbl_name":"tst1",
	"col_name":"cint1",
    "vals":[8],
    "limit":2
}
```
json
```
{
	"tbl_name":"tst1",
	"col_name":"cint1",
    "vals":[8],
    "short":true
}
```

#### Get value by string index

POST  
```
localhost:9170/i/siget/
```
Json (keys only) 
```
{
	"tbl_name":"tst1",
	"col_name":"cstring1",
    "val":"SSSS8"
}
```
json
```
{
	"tbl_name":"tst1",
	"col_name":"cstring1",
    "vals":["SSSS8"],
    "limit":2
}
```
json
```
{
	"tbl_name":"tst1",
	"col_name":"cstring1",
    "vals":["SSSS8"],
    "short":true
}
```





### String table

#### Create string table
POST  
```
localhost:9170/s/create_table/
```
Json  
```
{
	"name":"tst2",
	"type":"table_map_string"
}
```

#### Create int index on string table 
POST  
```
localhost:9170/s/create_index/
```
Json  
```
{
	"tbl_name":"tst2",
	"col_name":"cint1",
	"col_type":"int"
}
```

#### Create string index on string table 
POST  
```
localhost:9170/s/create_index/
```
Json  
```
{
	"tbl_name":"tst2",
	"col_name":"cstring1",
	"col_type":"string"
}
```
