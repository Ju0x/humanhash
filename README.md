# HumanHash

Human-readable digests in Go

Based on: https://github.com/zacharyvoase/humanhash

⚠️ **Warning:** These should only be used to representate another digest or UUID in a more readable version. They're not safe to use as a real hashing algorithm because they don't provide any safety features from modern hashes.

## Example

Generate a Humanhash based on a SHA256 digest:
```go
// Generate SHA256 hash
s := "Example"

h := sha256.New()
h.Write([]byte(s))
digest := h.Sum(nil)

// Generate Humanhash
hash, _ := humanhash.Humanize(digest, 4, "-")

// Output:
// jig-table-kansas-lion
fmt.Println(hash)
```

Generate a UUID with a Humanhash:
```go
hash, uuid, _ := humanhash.UUID()

// Output:
// Humanhash: sixteen-maryland-island-jig UUID: 361fbc5c-8154-4b1b-983f-d313934b3d8f
fmt.Printf("Humanhash: %s UUID: %s\n", hash, uuid)
```

Set a custom Wordlist (must have 256 words)
```go
humanhash.Wordlist("custom_wordlist.txt")
```
