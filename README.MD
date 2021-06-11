# EncodedApi

## Instructions

```bash
go run main.go
```

- Description: Starts the server

```bash
go test -v ./.../
```

- Execute all the tests recursively if you do on the project root

```bash
go test
````

- Execute the tests files in the folder that you\'re doing the command

## Design Pattern

 - Clean Ach pattern
    
## /upload

  - Description: Upload an image to a temp dir file using multipart/form-data then return the path for this file
  - Curl to test:
    ```bash
    curl    --request POST  
            --url http://localhost:3030/upload  
            --header 'Content-Type: multipart/form-data; 
            boundary=---011000010111000001101001'   
            --form 'file=@test.bmp'
    ```


## /write-message-on-image
    
- Description: Receives a message and a file path via POST of a JSON then encodes the message in the given file and return the new path
- Curl to test: 
    ``` bash 
    curl    --request POST \
            --url http://localhost:3030/write-message-on-image \
            --header 'Content-Type: application/json' \
            --data '{
	                   "path": "./assets/raw/test.bmp",
	                   "phrase": "test"
                    }'
    ```

## /get-image
    
- Description: Receives a encoded image name via GET then return the readable buffer for this image (utf-8 encoded)
- Curl to test: 
    ```bash 
    curl    --request GET \
            --url http://localhost:3030/get-image/test-encoded.bmp 
            --output ./assets/encoded/test.bmp
    ```



## /decode-message-from-image

- Description: Receives a path to an encoded image, decode the message in this file and return the message
- Curl to test: 
    ```bash 
    curl    --request GET \
            --url http://localhost:3030/decode-message-from-image/test-encoded.bmp
    ```
