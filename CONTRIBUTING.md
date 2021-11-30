# Contributing

First off, thank you for considering contributing to go-hubspot.

## Reporting bug & Feature request

If you've noticed a bug or have a feature request, please make an [issue](https://github.com/belong-inc/go-hubspot/issues).  
Some issue contents are being or have been fixed, so it is recommended to check for similar issues.  

## Code change flow

If the bug or feature request is something you think you can fix by yourself, please follow the steps below to change the code.

### Fork & create a branch

Fork repository and create branch with a descriptive name.  

### Implement fix or feature & run lint and test

When making changes to the code, keep in mind that this is an externally used library.  
Give variables and functions etc. easy-to-understand names, and avoid unnecessary exports. In addition, additional comments on why you did it would be helpful to reviewers and users.

You will also need to lint and test your changes. Make sure to run `make lint` and `make test` when the fix is complete.

**Be careful not to put your environment's api key or oauth information in git!**

### Make a pull request

When you have finished modifying and testing the code, please create a pull request with a description of the content. The lint and test will be executed by Github actions, but it's the same confirmation as local execution, so if it's already executed, no problem is expected.

When creating a pull request, we ask for the following items to be included. The items are described in the template.

- What to do  
  - Please describe the purpose and what changes of this pull request.
  - If you have any reference links, please include them.

- Background  
  - Please describe why you need to make this modification.

- Acceptance criteria  
  - Please describe what can be considered complete.

### Merging a pull request (maintainer only)

A PR can only be merged into main by a maintainer if:
- It is passing CI.
- It has been approved by maintainers.
- It has no requested changes.
- It is up to date with current main.
