![image](https://github.com/verdantjuly/minevillages-monitoring/assets/131671804/38ef69b8-9202-4527-94d9-8e249d9ffde1)
# Mine Villages Monitoring Service
This is Monitoring Service of Mine Villages Project made by Golang.

## Versions
- GO v1.22.0
- Redis v7.2.4
- ElasticSearch v7.17.4

## Features
### Health-Check
- Check Server's Health with HTTP request
  
### Logging
- Index Document with Elastic Search when server is unhealthy

### Alert
- Send Email per 1h when unhealthy occurs and consists.
- Use Redis TTL 1h to lock sending email every time.
