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
- [x] Fork repo includes only the `hello` package
  - [merge commit for hello package](https://github.com/arnaubennassar/git-subtree-playground-fork/commit/43291d3dc6200be55ef8a029cf8f77092bf0b449)
  - [main.go mods](https://github.com/arnaubennassar/git-subtree-playground-fork/commit/7093a15e155e96fc1afad9235be125fe419870c5)

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

# run go get to import the internal dependency from the upstream package that hello consumes:
➜  git-subtree-playground-fork git:(main) ✗ go get
# modify main.go to call the hello package
# run main.go to ensure it works:
➜  git-subtree-playground-fork git:(main) ✗ go run .
hello from the fork repo
hello
Hello from the internal dependency
```

- [x] Add changes to the `hello` package in this repo, in [this commit](https://github.com/arnaubennassar/git-subtree-playground-upstream/commit/c9fb994db3f4457d50d7b76a2412cf8c5ad263b1)
- [x] Fork repo gets upstream modifications, [in this commit](https://github.com/arnaubennassar/git-subtree-playground-fork/commit/6d4493c5ec096429ea0c8b11e93ce796732136a7)

```bash
git fetch --all

# re-create tracking branch
git branch -D feature/get-hello-package-from-upstram
git checkout -b feature/get-hello-package-from-upstram upstream-repo/main

# update the separate branch with changes from upstream
git subtree split -q --squash --prefix=hello --annotate="[upstream repo] " --rejoin -b merging/hello-package

# switch back to main and use subtree merge to update the subdirectory
git checkout main
git subtree merge -q --prefix=hello --squash merging/hello-package

# run main.go to ensure it works:
➜  git-subtree-playground-fork git:(main) go run .
hello from the fork repo
BRAND NEW HELLO FUNC, but still... hello
Hello from the internal dependency
```

- [x] Upstream modifies `hello`, [in this commit](https://github.com/arnaubennassar/git-subtree-playground-upstream/commit/0e12ea034bb3d9b7e4aed0588abeb477d9d98d7d)
- [x] Fork repo gets upstream modifications, in a separated branch so we can do a [PR to main](https://github.com/arnaubennassar/git-subtree-playground-fork/pull/1)

```bash
git fetch --all

# create update branch
git checkout -b feature/update-upstream

# re-create tracking branch
git branch -D feature/get-hello-package-from-upstram
git checkout -b feature/get-hello-package-from-upstram upstream-repo/main


# update the separate branch with changes from upstream
git subtree split -q --squash --prefix=hello --annotate="[upstream repo] " --rejoin -b merging/hello-package

# switch back to feature/update-upstream and use subtree merge to update the subdirectory
git checkout feature/update-upstream
git subtree merge -q --prefix=hello --squash merging/hello-package

# push changes
git push --set-upstream origin feature/update-upstream

# Open PR, approve and merge on GH

# switch back to main and use subtree merge to update the subdirectory
git checkout main
git pull

# run main.go to ensure it works:
➜  git-subtree-playground-fork git:(main) go run .
hello from the fork repo
Simple... hello
Hello from the internal dependency
```

- [x] Fork repo modifies the `hello` package locally in [this commit](https://github.com/arnaubennassar/git-subtree-playground-fork/commit/6ed73f251133dab55fa7b0c70229d6ae4ce939d8)
- [x] This repo makes changes to `hello` that will create conflict later on, in [this commit](https://github.com/arnaubennassar/git-subtree-playground-upstream/commit/b7c1013544c60314563758983324ddd05eaf7412)
- [x] Fork repo get changes from upstream as explained before, but needs to fix conflicts. [PR to main](https://github.com/arnaubennassar/git-subtree-playground-fork/pull/2)
- [x] Fork repo creates a [PR](https://github.com/arnaubennassar/git-subtree-playground-upstream/pull/1) to the upstream with the changes made in the previous step

```bash
git fetch --all

# create branch from main of the upstream repo
git checkout -b feature/push-to-upstream upstream-repo/main

# get changes from the fork main branch of the hello package
git checkout --patch main hello

# fix deps with go get
go get

# push changes
git add .
git commit -m "Contribute to upstream"
git push upstream-repo HEAD
# in the upstream repo create a PR from the branch with the same name (feature/push-to-upstream)
```