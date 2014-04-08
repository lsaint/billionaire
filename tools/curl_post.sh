#curl -X POST -H "Content-Type: application/json" -d '{"op":"gift", 
#"data":[{"uid": 50001906, "name": "幻", "tsid": 1640285, "num": 100, "time": "2014-04-10T10:07:02.0+08:00"}]}' http://localhost:40404/save
curl -X POST -H "Content-Type: application/json" -d '{"op":"gift", "param": 2}' http://localhost:40404/get
