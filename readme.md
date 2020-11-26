# git-log-exec
Executes any bash command on the whole git history and produces a CSV file out of it. This can be visualized quickly into nice graphs/chars.

**Usecases:**
- LOC/files over time
- trend of build/compile time
- trend of technical dept (e.g. search for certain method calls)

## Examples
### 1) Count files on remote repo
Note: The docker container will first clone the repository (repo needs to be public+accessable!).
```
docker run --rm brainexe/git-log-exec https://github.com/innogames/slack-bot.git -command="find -iname *.go | wc -l" --limit 10 
Clone https://github.com/innogames/slack-bot.git into local docker container
Cloning into '/repo'...
Commits: 120, Step size: 11 (11 commits to check)
 100% |████████████████████████████████████████|  [0s:0s]
Wrote file out.csv
time,result,commit
2019-07-21T19:18:39Z,122,ddeb55b
2019-09-18T10:58:37Z,136,f77d5f8
2019-09-27T12:35:50Z,137,0fb8ad5
2019-10-25T08:21:38Z,143,82ed091
2019-12-05T20:32:41Z,143,a0ff65d
2020-01-09T18:42:03Z,148,f84f181
2020-01-22T07:13:53Z,162,e5c6f2a
2020-07-30T14:14:31Z,162,431ea0f
2020-10-02T12:18:27Z,162,41615c2
2020-10-16T20:16:26Z,162,ba85f3f
2020-11-07T18:31:12Z,169,16469c3
```

### 2) Count LOC on local repo over time
Using golang/go repo as example:
```
docker run -v $(pwd):/repo  --rm brainexe/git-log-exec -command="find . -type f -exec wc -l {} + | awk '{sum += \$1} END {print sum}'" --limit 10
time,result,commit
2010-01-07T03:36:28Z,14234142,7a5852b50e
2011-07-18T02:59:16Z,14736460,a8e0035bd3
2012-08-20T08:56:41Z,15163804,a8357f0160
2014-01-09T20:21:24Z,15342160,9847c065f4
2015-04-07T23:22:12Z,15610076,2f16ddc580
2016-03-21T04:07:09Z,16031222,39af1eb96f
2017-03-21T04:04:46Z,16797796,0dafb7d962
2018-05-03T15:23:13Z,17437888,4704149e04
2019-09-04T21:52:18Z,18419724,aae0b5b0b2
2020-11-10T04:11:42Z,18359650,1642cd78b5

```

### 3) Count files of homeassistant repo
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
