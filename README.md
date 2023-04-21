# git-subtree-playground-upstream

Playing with git subtree. This project consist of two repos, this one and git-subtree-playground-fork

The purpose of this repo is to:

- Simulate the implementation of niteresting packages that the fork repo will include in it's subtree
- The fork repo should be able to receive updates from this repo
- The fork repo should be able to make local modifications
- The fork repo should be able to contribute **some** changes back to this repo

## Expected flow

1. Creat the init commit
2. Fork repo includes only the `hello` package
3. Fork repo modifies the `hello` package locally
4. Add changes to the `hello` package in this repo
5. Fork repo get upstreram changes
6. Fork repo modifies the `hello` package and creates a PR to this repo to contribute