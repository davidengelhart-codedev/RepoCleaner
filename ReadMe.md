## Installation and Usage Guide for RepoCleaner

### Description

The idea behind RepoCleaner is to help organizations get an idea of GitHub private repositories that have not been used (had a commit) in a set interval set by the user (i.e which GitHub repositories have not had a commit in the last 12 months). 

The push towards EaC (Everything as Code) or IaC (Infrastructure as Code) has placed emphasis on ensuring an organizations actively commited repos (which should contain most everything in code) are only listed within GitHub for view (especially for on-boarding of new employees). 

Currently the CLI program lists the private repos matching the organization and the last commit for each longer than the month interval specified; and saves to a CSV file for someone later to archive repos manually within GitHub--the manual process is not ideal but gives the company information to decide if the repos are being used or should be archived and read-only. 

Future state work will be to auto-archive these repos in the GitHub org upon user confirmation 'y' or 'n' from the CLI; however it is important to verify these repositories are actually not in use before archiving and making read-only--primary reason this part is not built out as of yet. 

### Enviromental Variables to Set

**GITHUB_API** --> This is your API key created in GitHub. this can be created per user or as an org level account similar to those used in CI/CD pipelines. 

For detailed instructions on how to obtain the personal access api token follow this [link](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)

**GITHUB_ORG** --> This is the name of your GitHub organization. Currently this CLI only supports private repos based upon the organization structure. 

**REPO_MONTHS** --> This is the number of months back from today you want to search for repos which have no commits/pushes. This should be a one or two digit number. 

Below is an example setting the enviroment variables in the terminal or your bash_profile etc. 

*obvisously all the value below are fake and use the values of your org and api key*

```shell
export GITHUB_API=ghp_fgnbsfdf87209bn2854280f43

export GITHUB_ORG=HUDSUCKER

export REPO_MONTHS=12
```

### Usage of the Go CLI RepoCleaner

To run any Go program you merely navigate to the project root and type in the terminal

```
go run nameOfProgram.go
```

or in this case for this project we can type

```
go run repo-clean.go
```

Go program can also be complied into an executable binary with all dependencies included using the `build` command as follows:

```
go build repo-clean.go
```

This creates a executable binary in the project root called `repo-clean` and to run the binary you type:

```
./repo-clean
```

*The current project already has this binary created; make sure enviroment variables are set and execute the command above to run the binary.* 



#### CSV file

The CSV file of the repopositories and the last push date after the time of the internal (number of months) specified in the enviromental variable will be in the root directory called `expired_repos.csv`

### Troubleshooting

If you get an error with failing to create csv file, make sure your project area is writable

The most common error you will experience may be: 

```
GET https://api.github.com/orgs//repos?type=private: 404 Not Found []
```

This error basically means you did not set your enviromental variables and they cannot be found. Ensure they are set and sourced to your current terminal session. 



### Deveopment Notes

This is my first Go program and recently learned the language. My past development background has been in Ruby, Python, TypeScript/JavaScript, PHP etc. I know I could have create a more modular project with separate functions outside the main() as well as separate files/packages etc. 



### Credits

Much of this program utilizes the API methods from this [repo](https://github.com/google/go-github)

