# urlparser

🔑 URL modification utilities for tasks that Go's native `net/url` package cannot handle.

## Supported functions

| Function | Why? What does it do? |
| ----------- | ----------- |
| `RemoveURLQueries()` | Go's native URL library encodes special characters when reconstructing a URL after removing and parsing its query string. This function provides a one-stop solution for removing query parameters from any URL, including unusual ones. Example: *scheme://subdomain.domain.top.level.domain:123/path/more-path?query_key_1=query_value_1&query_key_2=query_value_2#fragment* ➡️ *scheme://subdomain.domain.top.level.domain:123/path/more-path#fragment* |
