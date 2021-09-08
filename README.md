# web-api-template-go
concern:
- development difficulty
  - DRY-ness
  - easy to understand structures
  - simple folder structure
  - no over-to-the-top layer/abstraction
- no interface hell
- error handling
  - api error
  - logic error
  - generic error e.g. database-query error etc
- code convention for common library
  - function name
  - functionality
  - should not hardcode any value, only default/failover value that can be replaced from actual service
  - logging
- service mode (develop, staging or production)
