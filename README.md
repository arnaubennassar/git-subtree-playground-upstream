# git-subtree-playground-upstream

Playing with git subtree. This project consist of two repos, this one and [the fork repo](https://github.com/arnaubennassar/git-subtree-playground-fork)

The purpose of this repo is to:

- Simulate the implementation of niteresting packages that the fork repo will include in it's subtree
- The fork repo should be able to receive updates from this repo
- The fork repo should be able to make local modifications
- The fork repo should be able to contribute **some** changes back to this repo

## Expected flow

- [x] Creat the [init commit](https://github.com/arnaubennassar/git-subtree-playground-upstream/commit/83c36b03e7fc9b3d59b7746f64b7d0f779dfadd2)
- [x] Fork repo [init commit](https://github.com/arnaubennassar/git-subtree-playground-fork/commit/be8969dc542aa84989b918e41d1ea777a73d818d)
- [ ] Fork repo includes only the `hello` package:

```bash
# COMMANDS TO RUN IN THE FORK REPO!
# add upstream repo remote, create new tracking branch, 
git remote add -f upstream-repo git@github.com:arnaubennassar/git-subtree-playground-upstream.git
git checkout -b feature/get-hello-package-from-upstram upstream-repo/main

# split off subdir of tracking branch into separate branch
git subtree split -q --squash --prefix=hello --annotate="[upstream repo] " --rejoin -b merging/hello-package

# add separate branch as subdirectory on master.
git checkout main
git subtree add --prefix=hello --squash merging/hello-package

# run go get to import the internal dependency from the upstream package that hello consumes
# modify main.go to call the hello package
# run main.go to ensure it works
```

- [ ] Fork repo modifies the `hello` package locally
- [ ] Add changes to the `hello` package in this repo
- [ ] Fork repo get upstreram changes
- [ ] Fork repo modifies the `hello` package and creates a PR to this repo to contribute