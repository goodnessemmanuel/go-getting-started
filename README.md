# Go getting-started

## This is only a learn and practice repo

Learning source - Pluralsight

**Go workspace during set up should look similar to this description after installation**

```
<go-workspace> (i.e. parent directory name)
    - bin
    - pkg
    - src
        - bitbucket.com
        - github.com
            - <github-username>
                - <go-project-folder1>
                - <go-project-folder2>
                ...
```

***NB: workspace directory can be anywhere on your system but remember to define GO_PATH env, also add `<go-workspace>/bin` to your system path***

### Go commands

- **go build**  
 The go build command compiles the packages e.g. `<go-project-folder1>`, along with their dependencies,
  but it doesn't install the results (.i.e. the build file is not added to the bin directory of your go workspace).

- **go install**  
  The go install command compiles and installs the packages. I prefer this command because I can run my app without 'go run' by using just the build file name

***NB: you must run the commands from your project directory e.g `<go-project-folder1>`***
