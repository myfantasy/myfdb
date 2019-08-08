# myfdb
Simple kv db  
Based on github.com/myfantasy/myfdbstorage  

### Abount
Application is simple key value storage  
Contains 2 type storage:  
1. Simple Int64 key value  
1. Simple string key value  

## Methods
### Create table  
```
POST /crtbli - for int64 key
POST /crtbls - for string key
Body:
{
	"name":"tst",
	"type":"simple",
	"stored":true
}
```
name: table name  
type: now only simple  
stored: store on disk (true) or only in memory (false)   

if OK returns code: 200 body: ok  

### Set value
```
POST /seti/tst - for int64 key
POST /sets/tst - for string key
Body:
{
	"key":98,
	"value":"abrs",
    "is_base64":false
}


```
second segment id table name 'tst'  

key: item key (id)  
value: item value  
is_base64: value is base 64 encoded (true) or text (false)  

if OK returns code: 200 body: ok  

### Del value
```
POST /deli/tst - for int64 key
POST /dels/tst - for string key
Body:
{
	"key":98
}


```
second segment id table name 'tst'  

key: item key (id)   

if OK returns code: 200 body: true|false  

### Get value
```
POST /geti/tst - for int64 key
POST /gets/tst - for string key
Body:
{
	"key":98
}


```
second segment id table name 'tst'  

key: item key (id)   

if OK returns code: 200 body: bytes  
if NotFound returns code: 404 body: error text  


## Settings info
### Port (default)
port ":7171"  
`OutputInternalErrors: true`  
### Folders (default)
Default storage: `data/struct.json` storage data  
String storage folder: `data/S/`  
Int64 storage folder: `data/I/`  
### Timeout (default)
dump_timeout: 1*time.Second
