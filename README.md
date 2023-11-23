# LykoiQ
A queue server

### Quick Start
- `-p` Port e.g `8080`
- `-t` Number of threads e.d `1`
- `--debug` Run in debug mode
- `--destroy` can destroy the queue
```
-p 6060 -t 2 --debug --destroy
```

### Provider
Request path for enqueueing a message to the queue
```
http://localhost:8080/enqueue
```

Body:
```
{
    "message": "Hello Joe!"
}
```

Response:
```
{
    "id": 5,
    "created_on": "2023-11-23T21:07:07.697247Z"
}
```


### Consumer
Request path to dequeue a message from the head of queue
```
http://localhost:8080/dequeue
```


 