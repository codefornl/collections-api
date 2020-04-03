# Contributing Guide

Welcome! We love receiving contributions from our community, so thanks for stopping by! There are many ways to contribute, including submitting bug reports, improving documentation, submitting feature requests, reviewing new submissions, or contributing code that can be incorporated into the project.
This document describes our development process. Following these guidelines shows that you respect the time and effort of the developers managing this project. In return, you will be shown respect in addressing your issue, reviewing your changes, and incorporating your contributions.

## Questions, Feature requests and bugs
If you have any questions regarging this project, please use the issue tracker in this repository. Please create a new GitHub issue for any major changes and enhancements that you wish to make. Please provide the feature you would like to see, why you need it, and how it will work. Discuss your ideas transparently and get community feedback before proceeding.
**If you find a security vulnerability, do NOT open an issue. Email milo@codefor.nl instead.**
Before you submit your issue, please [search the issue archive](https://github.com/codefornl/collections-api/issues) - maybe your question or issue has already been identified or addressed.
If you find your issue is still legitimate, you can help us by submitting an issue to our GitHub issue tracker. Even better, you can submit a Pull Request with a fix.

## Improving Documentation
If you want to help improve the docs, it's a good idea to let others know what you're working on to minimize duplication of effort. Create a new issue (or comment on a related existing one) to let others know what you're working on. If you're making a small change (typo, phrasing) don't worry about filing an issue first. For large fixes, please build and test the documentation before submitting the PR to be sure you haven't accidentally introduced any layout or formatting issues.

## Contributing Code
Unsure where to begin contributing? You can start by looking through beginner and help-wanted issues: Beginner issues - issues which should only require a few lines of code, and a test or two. Help wanted issues - issues which should be a bit more involved than beginner issues.

Working on your first open source project or pull request? Here are some helpful tutorials:
* [How to Contribute to an Open Source Project on GitHub][2]
* [Make a Pull Request][3]
* [First Timers Only][4]

### Getting Started

Read the README.md for instructions on how to set up a development environment. When you are ready you can checkout the master branch - you want your new branch to come from master:
```
git checkout master
```
Next, create a new branch named newfeature (give your branch its own simple informative name)
```
git branch newfeature
```
Make sure to switch to your new branch
```
git checkout newfeature
```
For more information on the GitHub fork and pull-request processes, [please see this helpful guide][5].
You should be using the master branch for the most stable release. We do releases every week or two and send out notes. If you want to keep up with the latest changes, we work in the `dev` branch.  If you are using dev, keep an eagle-eye on commits and/or join our daily stand-up. 

> New to git? [please read this blog post][7].

### Finding an Issue

The list of outstanding feature requests and bugs can be found on our on our [GitHub issue tracker][6]. Pick an unassigned issue that you think you can accomplish and add a comment that you are attempting to do it.

Provide notes on different kinds of issues or labels
> `starter` labeled issues are deemed to be good low-hanging fruit for newcomers to the project
> `help-wanted` labeled issues may be more difficult than `starter` and may include new feature development
> `doc` labeled issues must only touch content in the `docs` folder.

### Review Process

The core team looks at Pull Requests on a regular basis. Except for critical, urgent or very small fixes, we try to leave pull requests open for most of the day or overnight if something comes in late in the day, so that multiple people have the chance to review/comment.  Anyone who reviews a pull request should leave a note to let others know that someone has looked at it.  For larger commits, we like to have a +1 from someone else on the core team and/or from other contributor(s).  Please note if you reviewed the code or tested locally -- a +1 by itself will typically be interpreted as your thinking its a good idea, but not having reviewed in detail.

### Addressing Feedback

Once a PR has been submitted, your changes will be reviewed and constructive feedback may be provided. Feedback isn't meant as an attack, but to help make sure the highest-quality code makes it into our project. Changes will be approved once required feedback has been addressed.

If a maintainer asks you to "rebase" your PR, they're saying that a lot of code has changed, and that you need to update your fork so it's easier to merge.

To update your forked repository, follow these steps:

```
# Fetch upstream master and merge with your repo's master branch
git fetch upstream
git checkout master
git merge upstream/master

# If there were any new commits, rebase your development branch
git checkout newfeature
git rebase master
```

If too much code has changed for git to automatically apply your branches changes to the new master, you will need to manually resolve the merge conflicts yourself.

Once your new branch has no conflicts and works correctly, you can override your old branch using this command:

```
git push -f
```

Note that this will overwrite the old branch on the server, so make sure you are happy with your changes first!

## Community

All our primary communication takes place in this repository. Incidental communication can be done in the slack channels that can be reached through praatmee.codefor.nl

[0]: CODE_OF_CONDUCT.md
[1]: style_guidelines.md
[2]: https://egghead.io/series/how-to-contribute-to-an-open-source-project-on-github
[3]: http://makeapullrequest.com/
[4]: http://www.firsttimersonly.com
[5]: https://gist.github.com/Chaser324/ce0505fbed06b947d962
[6]: link/to/your/project/issue/tracker
[7]: http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html
