
#!/bin/bash

VERSION=$(cat version)
echo "building test-terraform-provider_${VERSION}"
go build -o test-terraform-provider_${VERSION}
