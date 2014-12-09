Rule Driven Record Engine [RDRE] Transformer
=====================================

RDRE: A Framework handling the workflow of transforming data records through defined rules

### Core goals:
-	simple
-	performant
-	standalone
-	extensible
-	extensive

Current status: usable (see examples), experimental

[![GoDoc](https://godoc.org/github.com/rdre/core?status.svg)](https://godoc.org/github.com/rdre/core)
[![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/rdre/core)

### Contributors
All contributiors are welcome. Support for:
- new rules (separate repository)
- custom cmd app with pre-packed rules 
will be valuable. See road map below

### Roadmap
- [ ]	cmd app (separate repository)
- [ ]	pre-packed rules for cmd app

### Examples [rdre/examples](https://github.com/rdre/examples)

Current rule syntax format [YAML based]
```javascript
rule_key_id
	[YAML definition]
```
Example:
```javascript
rename_col:
	previous: ID
	next: PERSON ID
```
Valid YAML definition or empty def defines a Rule. A rule marshalls the child data structure of its key.
For an empty definition, consider a functional rule (see examples [rdre/examples](https://github.com/rdre/examples))

### License
[MIT Public License](https://github.com/rdre/core/blob/master/LICENSE)
