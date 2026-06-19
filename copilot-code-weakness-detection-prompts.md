# Goals

Assess GitHub Copilot's efficiency and accuracy to detect code weaknesses in Go and Bash.

## Strategy

Submit a number of pull requests to the `ihcsim/harvester` repository containing deliberate code weaknesses and vulnerabilities.

Do not provide Copilot with any hints that these exploits are intentionally added to the files.

## Steps

### Go Test Cases

#### Test Case 1

Submit 3 pull requests containing one or more gosec violations found in <https://raw.githubusercontent.com/TheHackerDev/damn-vulnerable-golang/refs/heads/main/main.go>. These changes can be done to new or existing `.go` files.

Before submitting the pull requests, perform the validation steps described in the ["Validation"](#validation) section below.

#### Test Case 2

Submit 1 pull request with changes to the `go.mod` file to include the deprecated `github.com/golang/protobuf@v1.5.4` package. Add some `import` changes to a Go file to ensure the changes are not reverted by `go mod tidy`.

Before submitting the pull requests, perform the validation steps described in the ["Validation"](#validation) section below.

#### Test Case 3

Submit 1 pull request containing the [GO-2021-0113](https://pkg.go.dev/vuln/GO-2021-0113) by downgrading the `golang.org/x/text` to `v0.3.5` in the `go.mod` file. Add some `import` changes to a Go file to make sure the changes are not reverted by `go mod tidy`

Before submitting the pull requests, perform the validation steps described in the ["Validation"](#validation) section below.

### Bash Test Cases

#### Test Case 1

Submit 1 pull request containing one of more of the following code weaknesses:

| Issue | What to look for |
|-------|-----------------|
| **Unquoted variables** | Variables used without quotes (e.g., `$var` instead of `"$var"`) — can cause word splitting and glob expansion |
| **Command injection** | User input or variables interpolated into commands without sanitization — especially with `eval`, backticks, or `$(...)` |
| **Path traversal** | File paths constructed from input without validation — check for `..` sequences and absolute path handling |
| **Unsafe `eval`** | Any use of `eval` with user-controlled input or variables |

#### Test Case 2

Submit 1 pull request containing one of more of the following code weaknesses:

| Issue | What to look for |
| ------|------------------|
| **Unchecked commands** | Commands that can fail silently — ensure `set -e` or explicit error checking with `\|\| exit 1` |
| **Insecure temp files** | Predictable temp file names instead of `mktemp` — can lead to race conditions |
| **World-writable files** | Files/directories created with overly permissive permissions (666, 777) |

#### Test Case 3

Submit 1 pull request containing one of more of the following code weaknesses:

| Issue | What to look for |
| ------|------------------|
| **Missing input validation** | Script arguments (`$1`, `$2`, etc.) used without validation or bounds checking |
| **Unsafe downloads**     | `curl` or `wget` without integrity checks (checksum validation) or used with `\| sh` |
| **Secret exposure** | Secrets/credentials in script output, logs, or error messages |

These changes can be done to new or existing `.sh` files.

### Code Design Pattern

#### Test Case 1

Submit a pull request that adds a new target that invokes a `git` command to the
Makefile without the `gen-version-env` dependency rule. Without the dependency
rule, the new target will cause build or runtime failure. We want to see if
Copilot can detect this weakness and suggest adding the dependency rule.

#### Test Case 2

Submit a pull request that adds a new controller that doesn't follow the
'Controller Pattern' described in the [AGENTS.md](AGENTS.md) documentation. For
example, the new controller might not use the `config.Management` clients or
might not register its `OnChange` handler properly. We want to see if Copilot can
detect this weakness and suggest following the established pattern.

#### Test Case 3

Submit a pull request that adds a new KubeBuilder-based controlelr that doesn't
follow the 'Controller Pattern'. You can just copy and paste the sample code from

- <https://raw.githubusercontent.com/kubernetes-sigs/kubebuilder/refs/heads/master/docs/book/src/cronjob-tutorial/testdata/emptycontroller.go>
- <https://raw.githubusercontent.com/kubernetes-sigs/kubebuilder/refs/heads/master/docs/book/src/cronjob-tutorial/testdata/emptyapi.go>

Add these files to a new folder named `pkg/controller/master/empty/`.

## Validation

- Run `make test` to ensure existing tests still pass
- Run `go mod tidy && go mod vendor` to ensure the `vendor` folder is up to date with the `go.mod` changes.
- No need to run `make validate` as the goal is to test Copilot's ability to detect weaknesses, not to ensure code quality.

Your job is done once the pull requests are successfully created.

## Cleanup

Do not perform the cleanup tasks without my confirmation. When I confirm, please:

- Make sure you are working with the `ihcsim/harvester` remote repository.

- Close the pull requests you created.
- Delete the remote branches you created for the pull requests.
- Delete the local branches you created for the pull requests.
