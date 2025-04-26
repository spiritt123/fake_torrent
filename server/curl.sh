
curl --location 'http://127.0.0.1:8000/add' \
--header 'Content-Type: application/json' \
--data '{
    "username":"alex",
    "file":{
        "filename":"Film12",
        "size":4096,
        "blocksize":512
    }
}'
curl --location 'http://127.0.0.1:8000/add' \
--header 'Content-Type: application/json' \
--data '{
    "username":"alex",
    "file":{
        "filename":"Film1",
        "size":2048,
        "blocksize":512
    }
}'

curl --location 'http://127.0.0.1:8000'

curl --location 'http://127.0.0.1:8000/enable' \
--header 'Content-Type: application/json' \
--data '{
    "ip":"123",
    "file":{
        "id":1,
        "filename":"Film12",
        "size":4096,
        "blocksize":512
    },
    "active_blocks":[true,true,true,true,true,true,true,true]
}'

curl --location 'http://127.0.0.1:8000/getBlock' \
--header 'Content-Type: application/json' \
--data '{
    "id":1,
    "ip":"124"
}'


