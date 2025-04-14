```
docker build -t cdn .
```

```
docker run -d \
    --restart unless-stopped \
    -p 2095:2095 \
    cdn
```