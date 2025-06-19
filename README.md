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

1. **Install goreleaser (if you don't have it):**
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

## **Let's Start!**

**Please reply with your GitHub username** so I can tailor the instructions and code snippets for your repo.  
Once you confirm the directory rename and `go.mod` update, I'll walk you through each next step!

## **Step 1: Rename Your Local Directory and Update go.mod**

1. **Rename your local directory (if needed):**
   ```sh
   mv pace-cli-app pace-announcement-cli
   cd pace-announcement-cli
   ```

2. **Update your `go.mod` file:**
   - Open `go.mod` and change the first line to:
     ```
     module github.com/fakepixels/pace-announcement-cli
     ```
   - Save the file.

## **Step 2: Push to Your GitHub Repo**

If you haven't already:

1. **Initialize git (if not already):**
   ```sh
   git init
   git remote add origin https://github.com/fakepixels/pace-announcement-cli.git
   ```

2. **Add, commit, and push your code:**
   ```sh
   git add .
   git commit -m "Initial commit"
   git branch -M main
   git push -u origin main
   ```

## **Step 3: Add a Version Flag to Your CLI**

1. **Edit your `main.go`** and add at the top (after imports):

   ```go
   import (
       "flag"
       // ... other imports
   )

   var version = "dev" // will be set by goreleaser

   func main() {
       flagVersion := false
       flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
       flagSet.BoolVar(&flagVersion, "version", false, "print version and exit")
       flagSet.BoolVar(&flagVersion, "v", false, "print version and exit (shorthand)")
       flagSet.Parse(os.Args[1:])
       if flagVersion {
           fmt.Println(version)
           return
       }
       // ... rest of your main()
   }
   ```