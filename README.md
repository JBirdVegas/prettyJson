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