# github-actions-validator

[![Go Reference](https://pkg.go.dev/badge/github.com/Cardinal-Cryptography/github-actions-validator.svg)](https://pkg.go.dev/github.com/Cardinal-Cryptography/github-actions-validator) [![Go Report Card](https://goreportcard.com/badge/github.com/Cardinal-Cryptography/github-actions-validator)](https://goreportcard.com/report/github.com/Cardinal-Cryptography/github-actions-validator) ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/Cardinal-Cryptography/github-actions-validator?sort=semver)

Quick tool to validate workflows and actions in .github directory

## Checks
See the checks that are performed on all the workflow and action files.  These are separate into errors
and warnings.  Each check has a code where as one starting with `E` indicates an error, `N` indicates
a warning about invalid naming convention and, finally `W` is any other warning.
Additionally, code will contain either `A` if it is an action where the issue is found, and `W` if 
issue occurs in a workflow.

### Errors

| Code | Description |
|------|-------------|
| EA809 | Called step with id '%s' does not exist |
| EA811 | Called step with id '%s' output '%s' does not exist |
| EW203 | Job '%s' has invalid value '%s' in 'needs' field |
| EW201 | Called variable '%s' is invalid |
| EW202 | Called input '%s' does not exist |
| EW203 | Job '%s' has invalid value '%s' in 'needs' field |
| EW801 | Path to external action '%s' is invalid |
| EW802 | Path to local action '%s' is invalid |
| EW803 | Call to non-existing local action '%s' |
| EW804 | Required input '%s' missing for local action '%s' |
| EW805 | Input '%s' does not exist in local action '%s' |
| EW806 | Required input '%s' missing for external action '%s' |
| EW807 | Input '%s' does not exist in external action '%s' |
| EW808 | Call to non-existing external action '%s' |
| EW809 | Called step with id '%s' does not exist |
| EW810 | Called step with id '%s' does not exist |
| EW811 | Called step with id '%s' output '%s' does not exist |
| EW254 | Called variable '%s' does not exist in provided list of available vars (when -z provided) |
| EW255 | Called secret '%s' does not exist in provided list of available secrets (when -s provided) |

### Warnings

| Code | Description |
|------|-------------|
| WW101 | Called env var '%s' not found in global, job or step 'env' block - check it |
| WW201 | Called var '%s' may not need to be in double quotes |

### Naming convention warnings

| Code | Description |
|------|-------------|
| NA101 | Action directory name should contain lowercase alphanumeric characters and hyphens only |
| NA102 | Action file name should have .yml extension |
| NA103 | Action name is empty |
| NA104 | Action description is empty |
| NA301 | Action input name should contain lowercase alphanumeric characters and hyphens only |
| NA302 | Action input must have a description |
| NA501 | Action output name should contain lowercase alphanumeric characters and hyphens only |
| NA502 | Action output must have a description |
| NW101 | Workflow file name should contain alphanumeric characters and hyphens only |
| NW102 | Workflow file name should have .yml extension |
| NW103 | Env variable name '%s' should contain uppercase alphanumeric characters and underscore only |
| NW104 | Workflow name is empty |
| NW106 | When workflow has only one job, it should be named 'main' |
| NW107 | Called variable name '%s' should contain uppercase alphanumeric characters and underscore only |
| NW301 | Workflow input name should contain lowercase alphanumeric characters and hyphens only |
| NW302 | Workflow input must have a description |
| NW501 | Workflow job name should contain lowercase alphanumeric characters and hyphens only |
| NW502 | Env variable name '%s' should contain uppercase alphanumeric characters and underscore only |
| NW701 | Env variable name '%s' should contain uppercase alphanumeric characters and underscore only |


## Building
Run `go build -o github-actions-validator` to compile the binary.

### Building docker image
To build the docker image, use the following command.

    docker build -t github-actions-validator .


## Running
Check below help message for `validate` command:

    Usage:  github-actions-validator validate [FLAGS]

    Runs the validation on files from a specified directory

    Required flags: 
      -p,	 --path  	Path to .github directory
    
    Optional flags: 
      -s,	 --secrets-file  	Check if secret names exist in this file (one per line)
      -z,	 --vars-file  		Check if variable names exist in this file (one per line)

Use `-p` argument to point to `.github` directories.  The tool will search for any actions in the `actions`
directory, where each action is in its own sub-directory and its filename is either `action.yaml` or
`action.yml`.  And, it will search for workflows' `*.yml` and `*.yaml` files in `workflows` directory.

Additionally, all the variable names (meaning `${{ var.NAME }}`) as well as secrets (`${{ secret.NAME }}`)
in the workflow can be checked against a list of possible names.  Use `-z` and `-s` arguments with paths
to files containing a list of possible variable or secret names, with names being separated by new line or
space.

### Example of checking secrets

    % cat ~/secrets-list.txt 
    MY_SECRET_1
    MY_SECRET_2
    % ./github-actions-validator validate -p /path/to/.github -s ~/secrets-list.txt | grep '^EW25'
    EW255: workflow my-workflow.yml                              Called secret 'GITHUB_TOKEN' does not exist in provided list of available secrets


### Using docker image
Note that the image has to be present, either built or pulled from the registry.
Replace path to the .github directory.

    docker run --rm --name tmp-gha-validator \
      -v /Users/me/my-repo/.github:/dot-github \
      github-actions-validator \
	  validate -p /dot-github


## Exit code
Currently, tool always exit with code 0.  To check if there are any errors, please use `grep` to filter
the output for errors.

