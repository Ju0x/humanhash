# HumanHash

Human-readable digests in Go
Based on: https://github.com/zacharyvoase/humanhash

**Warning:** These should only be used to representate another digest UUID in a more readable version. They're not safe to use as a real hashing algorithm because they don't provide any safety features from modern hashes.

## Example

Generate a Humanhash based on a SHA256 digest:
```
// Generate SHA256 hash
s := "Example"

h := sha256.New()
h.Write([]byte(s))
digest := h.Sum(nil)

// Generate Humanhash
// Result: jig-table-kansas-lion
hash, err := humanhash.Humanize(digest, 4, "-")
fmt.Println(hash)
```

Set a custom Wordlist (must have 256 words)
```
humanhash.Wordlist("custom_wordlist.txt")
```