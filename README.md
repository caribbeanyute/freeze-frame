# Captures screenshot from m3u8 stream

## Sample request
```bash
curl --location --request POST 'http://localhost:3001/api/thumbnail' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "http://192.168.1.101:31199/live/flow/intelcore.m3u8"
}'

```