# HTTP Status Code Checker

Command line application that checks the status codes for GET requests sent to the specified target(s).

## Arguments

| Flag    | Name  | Type   | Sample Input            |
| :-----: |:-----:|:------:|:----------------------- |
| -u      | url   | String |`https://google.com`     |
| -f      | file  | String |"./path/to/urls.txt"     |
```
Note that the arguments can be combined (e.g. status-check -u https://google.com -f urls.txt).
```
## Sample Run

### File Content
```
https://yandex.com
bing.com
http://facebook.com
```
### Terminal / Command Line
```
go run main.go -u https://google.com -f urls.txt
```
### Output
```
200 | https://yandex.com
200 | http://facebook.com
200 | https://google.com
Error occurred during execution of the urls below
bing.com
```
### Error Reason
```
The protocol for communication should be specified in the url (HTTP or HTTPS).
```

