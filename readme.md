# git-log-exec
Executes any bash command on the whole git history and produces a CSV file out of it. 

**Usecases:**
- LOC/files over time
- trend of build/compile time
- trend of technical dept (e.g. search for certain method calls)

## Examples

### Count *.go files in go/golang repo over time

```
docker run -v ~/projects/golang:/repo  --rm brainexe/git-log-exec -command="find -iname *.go | wc -l" --limit 200
time,result,commit
2020-11-13T10:36:40Z,6916,0c7f0e5448
2020-09-01T17:38:07Z,6767,ab88d97deb
2020-04-28T17:39:36Z,6764,bd01a1b756
2020-02-27T21:33:33Z,6611,62ff72d876
2019-11-07T19:14:38Z,6694,e1ddf0507c
2019-09-07T21:44:30Z,6622,a5025fdcde
2019-04-26T19:27:51Z,6818,2ae5e7f9bf
2019-02-22T17:05:17Z,6485,b35dacaac5
...
2010-05-24T22:07:47Z,1252,c95e11db56
2010-01-18T23:59:14Z,1027,16205a3534
2009-11-02T18:50:18Z,858,d00248980b
2009-07-28T21:54:49Z,623,6d3d25de21
2009-03-27T05:16:06Z,447,34050ca8de
2008-10-29T22:23:29Z,319,527669e6b1
2008-06-04T21:37:38Z,37,0cafb9ea3d
```

### Count files of homeassistant repo
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

## Docker usage

docker run --rm brainexe/git-log-exec

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