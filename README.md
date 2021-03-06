# API with model in Twitter application using Golang and MongoDB :boom: :whale:

## build up

```sh
docker build -t ulicode/api:v1 .
docker run --name tweeto -p 8080:8080 ulicode/api:v1

kubectl apply -f deploys.yaml
kubectl create -f deploy.yaml
kubectl expose deployment server-go --type=NodePort
```


## Database 

## Endpoints exposed

|     Endpoint      | Method |        Description         |
| :---------------: | :----: | :------------------------: |
|     /register     |  POST  |      Register a user       |
|      /login       |  POST  |       Logging to app       |
|     /profile      |  GET   |    View profile details    |
|     /profile      |  PUT   |       Update profile       |
|      /tweet       |  POST  |       Create a tweet       |
|      /tweets      |  GET   |     Get list of tweets     |
|      /tweet       | DELETE |       Delete a tweet       |
|      /avatar      |  POST  |        Up a avatar         |
|      /avatar      |  GET   |        Get a avatar        |
|      /banner      |  POST  |        Up a banner         |
|      /banner      |  GET   |        Get a banner        |
|      /follow      |  POST  |      Follow a person       |
|     /unfollow     | DELETE |     Unfollow a person      |
|     /relation     |  GET   |  Check if follow someone   |
|      /users       |  GET   |       List of users        |
| /tweets/followers |  GET   | List of tweets that follow |
