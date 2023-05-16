test:
	# source: https://gist.github.com/gregohardy/d148db9e01aa721ea42daf4abbba725e
	echo "\033[34mRunning Go test...\033[39m"
	go test ./... -count=1
	echo ""
	echo "\033[34mRunning Python test...\033[39m"
	python3 -m unittest solution/*_test.py