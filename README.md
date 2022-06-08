# idpops

## Installing the flux
 * cd products/aws/aws-account-1/us-west-2/product-java-app/foundation/flux-system
 * export KUBECONFIG=<>
 * kustomize build . | k apply -f -

## Installing the product
 

## Go Api server
Rest URL for getting all the products
      http://localhost:8080/getProducts

## docker image for apiserver
    docker.io/madhukirans/idpapiserver:1