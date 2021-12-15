mysqlStart:
	cd ./docker/mysql; docker-compose -f dev-docker-compose.yml up -d
#clearGit:
#	git branch --merged | egrep -v "(^\*|master|main|dev)" | xargs git branch -d
#   https://www.git-tower.com/learn/git/faq/cleanup-remote-branches-with-git-prune