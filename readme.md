# git-log-exec
Executes any bash command on the whole git history and produces a CSV file out of it. 

**Usecases:**
- LOC/files over time
- trend of build/compile time
- trend of technical dept (e.g. search for certain method calls)

## Example
```
cd ~/projects/home-assistant
git-log-exec -out loc.csv -command="find homeassistant -type f | wc -l" 
```

**Result**
```
time,result,commit
2019-12-22T23:41:22+01:00,12064,48d35a455
2019-12-04T00:46:38+01:00,12052,564c468c2
2019-11-27T20:52:03+01:00,12006,d7a66e6e4
2019-11-25T04:57:40+01:00,11876,c38240673
2019-11-13T15:32:22+01:00,11876,15ce73835
2019-11-03T20:36:02+01:00,11779,5fd9b474d
2019-10-25T01:42:54+02:00,11666,643b3a98e
2019-10-21T09:55:53+02:00,11607,c1fccee83
2019-10-17T15:03:05+02:00,11551,8350e1246
2019-10-12T21:57:18+02:00,11486,17b1ba2e9
2019-10-07T21:49:54+02:00,11420,1febb32dd
``` 

## Options
```
  -command string
    	command to execute
  -directory string
    	git workspace (by default current directory)
  -limit int
    	max number of commits to check (default 500)
  -branch string
    	git branch to check (default "master")
  -output string
    	output csv file (default "out.csv")
  -after string
    	optional begin date of history search
  -before string
    	optional end date of history search
``` 