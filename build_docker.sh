docker buildx rm mbuilder
docker create --config ../buildkitd.toml mbuilder
docker buildx use mbuilder
# docker buildx build -f Dockerfile -t gofly:latest .
docker buildx build -f Dockerfile --platform linux/arm/v6 -t 192.168.86.166:5000/gofly:latest . --push