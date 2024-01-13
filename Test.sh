go run . --output=test1.txt "hello" standard
cat -e test1.txt

go run . --output=test2.txt "Hello There!" shadow
cat -e test2.txt