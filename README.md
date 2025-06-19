## Installation

```sh
go install github.com/<yourusername>/pace@latest
```

Then run:

```sh
pace
```
```

---

## **Step 5: (Optional) Set Up goreleaser for Prebuilt Binaries**

1. **Install goreleaser (if you don’t have it):**
   ```sh
   brew install goreleaser
   ```
   or
   ```sh
   go install github.com/goreleaser/goreleaser@latest
   ```

2. **Create a `.goreleaser.yml` in your project root:**
   ```yaml
   project_name: pace
   builds:
     - main: ./main.go
       ldflags:
         - -X main.version={{.Version}}
   archives:
     - format: tar.gz
   ```

3. **Test a local build:**
   ```sh
   goreleaser release --snapshot --clean
   ```
   This will create binaries in `dist/`.

4. **For real releases:**  
   - Push a git tag (e.g., `git tag v0.1.0 && git push --tags`)
   - Run `goreleaser release --clean` (or set up GitHub Actions for automated releases).

---

## **Let’s Start!**

**Please reply with your GitHub username** so I can tailor the instructions and code snippets for your repo.  
Once you confirm the directory rename and `go.mod` update, I’ll walk you through each next step!