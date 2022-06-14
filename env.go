package main

import "os"

var CosBucketUrl = os.Getenv("COS_BUCKET_URL")
var CosServiceUrl = os.Getenv("COS_SERVICE_URL")
var CosSecretId = os.Getenv("COS_SECRET_KEY")
var CosSecretKey = os.Getenv("COS_SECRET_KEY")
