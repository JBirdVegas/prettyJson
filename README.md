### Pretty Print json from the command line

Simply pass json in via a pipe, file or command arg

Example usage:

```shell
% cat test.json
{"Testing": true, "Hello": "World"}

% ./prettyJson "$(cat test.json)"
{
    "Hello": "World",
    "Testing": true
}

% cat test.json | ./prettyJson   
{
    "Testing": true,
    "Hello": "World"
}

% ./prettyJson test.json 
{
    "Testing": true,
    "Hello": "World"
}

% ./prettyJson -file test.json
{
    "Testing": true,
    "Hello": "World"
}
```
 
This app can also be used to collapse json

```shell
% cat test.json 
{
  "Testing": true,
  "Hello": "World"
}

% ./prettyJson -collapse "$(cat test.json)"
{"Hello":"World","Testing":true}

% cat test.json | ./prettyJson -collapse   
{"Testing":true,"Hello":"World"}

% ./prettyJson -collapse -file test.json
{"Testing":true,"Hello":"World"}

% ./prettyJson -collapse test.json 
{"Testing":true,"Hello":"World"}   
```
