## Our Git workflow

### Setup
You should only need to complete this setup once. 

If you'd like now is the time to set up SSH and GPG keys.

1. `git clone` the URL given on the Github repo into a directory of your choice.

### Branches
We will be using the [GitFlow](https://nvie.com/posts/a-successful-git-branching-model/) branching model.

We have **two** evergreen branches (persistent branches). 

- `master` always points to our `latest` release. Merge commits on `master` are always tagged with the version.
- `dev` is the branch you `checkout -b` and pull request merge into. `dev` should always be a `green` build (all checks are green)

We can make **three** types of temporary branches:

- `feat/branch-name`: A feature branch is used for most additions. A feature branch always branches off `dev` or another feature branch.
- `fix/branch-name`: A fix branch is identical to a `feat` branch except it indicates the branch is used only for fixing a bug.
- `hotfix/branch-name`: A hotfix branch is used for **critical** fixes such as high severity security flaws. Hotfixes are always merged directly into `master` and `dev`.

N.B1 Note all these branches are made using `checkout -b` from the appropriate base branch. This naming scheme isn't enforced by Git or Github and is only for clarity.

N.B2 We do not use release branches (we go directly from `dev` to `master` to reduce complexity)

## Typical Workflow
1. `git checkout dev` (moves to the `dev` branch)
2. `git pull` (ensures your local branch is up to date with the remote branch)
3. `git checkout -b your-feature/fix-branch-name` (checkouts to a new branch to you to do your work on) and perform a `git push --set-upstream origin /your/branch` to link your local branch to Github.
4. Make one **atomic** change to your files (typically you should contain your changes to one purpose at a time)
5. `git add -A` (stages all the changes in your working directory for commit). You can also do `git add file1 file2` to stage individual files.
6. `git commit -m "Your commit message describe your changes"` (commits your changes locally). Please use the Jira task/ticket tag that deals with your task to be able to track your progress using Jira.
7. You can share your changes with remote at any time by doing a `git push`
8. When you are done do a `git rebase -i origin/dev` (rebases your branch on the latest `dev` branch). The `-i` means you can see the changes.
9. `git push --force-with-lease` (this is a special push command which allows you to update the remote even if you have diverging histories) It *should* only push changes if it won't overwrite another person's commits BUT only do this is you are certain you won't overwrite anything. There is no easy way to undo this.
10. Open a pull request for your branch into the `dev` branch. Typically the pull request opener also merges the PR after all checks are passed.
11. Rinse and repeat.
