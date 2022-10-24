#clearGit:
#	git branch --merged | egrep -v "(^\*|master|main|dev)" | xargs git branch -d
#   https://www.git-tower.com/learn/git/faq/cleanup-remote-branches-with-git-prune