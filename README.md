# idpops

## Installing the flux
 * cd products/aws/aws-account-1/us-west-2/product-java-app/foundation/flux-system
 * export KUBECONFIG=<>
 * kustomize build . | k apply -f -

## Installing the product
 

## Go Api server
Rest URL for getting all the products
      http://localhost:8080/idpops/api/v1/getProducts
To the specific product details
      http://localhost:8080/idpops/api/v1/getProductDetails?cloud=aws&account=aws-account-1&clusterName=product-java-app&region=us-west-2&product=app1

## docker image for apiserver
    docker.io/madhukirans/idpapiserver:1
