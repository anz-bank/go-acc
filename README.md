# Golang Coverage Action

Getting accurate code coverage information in Go [isn't exactly straight-forward](https://www.ory.sh/golang-go-code-coverage-accurate/). A tool called [go-acc](https://github.com/ory/go-acc) post-processes the coverage information produced by Go's default tooling and produces output that includes coverage from other packages.

This PR includes a new GitHub Action `coverage-action` which runs go-acc and will fail the PR if the coverage is below the `hard_target` and warn if below the `soft_target`.

Usage (.github/workflows/main.yml):
```yaml
name: Quality Analysis
on: pull_request
jobs:
  coverage:
    runs-on: ubuntu-latest
    steps:
    - name: Checkout
      uses: actions/checkout@v2
      fetch-depth: 1
    - name: Run coverage tool
      uses: anz-bank/go-acc@v0.1.0
      with:
        soft_target: 90
        hard_target: 80
```