# urlparser

🔑 Whatever URL modification stuff that Go native url library can't do.

## Supported functions

| Function | Why? What does it do? |
| ----------- | ----------- |
| `RemoveURLQueries()` | Go native url library has one fault which is encoding special characters when reconstructing the URL after query removal process and parsing return. This function is intended to save your headache and be the one-stop solution for any kind of URL query removal from any kind of URLs, even the weirdest ones. Example: *scheme://subdomain.domain.top.level.domain:123/path/more-path?query_key_1=query_value_1&query_key_2=query_value_2#fragment* ➡️ *scheme://subdomain.domain.top.level.domain:123/path/more-path#fragment* |