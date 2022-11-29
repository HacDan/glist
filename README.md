# go-list

```bash
list_a="one two"
list_b="one three"

# union "one two three"
list -u "$list_a" "$list_b"

# intersect "one"
list -x "$list_a" "$list_b"

# complement "two three"
list -n "$list_a" "$list_b"
```
