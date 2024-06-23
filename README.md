# Video Streaming 
Simple golang app to let you upload videos to S3. On AWS you can connect S3 to Cloudfront to serve the videos through a Cloudfront

## Running 
Copy contents of `.env.example` into a `.env` file and update the content as required.
Then run  
```bash
go run cmd/app/main.go -region=<region of s3 bucket> -bucket=<bucket name>
```

## Routes
```http
GET /upload 
body-> file: <file>
```
