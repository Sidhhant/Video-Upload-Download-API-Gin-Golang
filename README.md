# Video-Upload-Download-API-Gin-Golang
API for video download and upload using Gin Framework 

## Steps to setup project
### Run
`docker-compose up`

### Stop
`docker-compose down --volume`

### API
```bash
GET    /v1/health                
GET    /v1/files                 
POST   /v1/files                  
DELETE /v1/files/:fileid         
GET    /v1/files/:fileid          
```

Use Curl or Postman to test the API. 
Please feel free to contribute or raise issue.