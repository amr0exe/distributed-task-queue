

**CREATE TASK**
```
curl -X POST http://localhost:8080/task \
-H "Content-Type: application/json" \
-d '{"title":"learn distributed systems"}'
```

**GET ALL TASK**
```
curl http://localhost:8080/all
```

**DELETE TASK**
```
curl -X DELETE http://localhost:8080/task/id_here
```

**Update Task**
```
curl -X PUT http://localhost:8080/task/id_here \
-H "Content-Type: application/json" \
-d '{"title":"halo systems", "is_completed": true}'
```