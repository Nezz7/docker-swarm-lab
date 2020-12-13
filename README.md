# Docker swarm lab

## Overview 
    This lab creates a high availability architecture using docker swarm :
    a simple Go service that communicates with MongoDB.
## Deploy
1. Clone this repository <br>
`$ git clone https://github.com/Nezz7/docker-swarm-lab.git`
2. Change the current working directory <br>
`$ cd /docker-swarm-lab`<br>
3. Create the docker image : server-image  <br>
`$ docker build -t server-image .`<br>
4. Initialize the swarm <br>
`$ docker swarm init`<br>
Don't forget to save the token.<br>
5. Deploy the stack to the swarm<br>
`$ docker stack deploy --compose-file docker-compose.yml  $SERVICE_NAME`<br>
6. Check that it’s running <br>
`$ docker stack services $SERVICE_NAME`<br>

## Exposed Ports

| Service       | Exposed Pot   | Internal Port  |
|:-------------:|:-------------:| :-----:|
| Server        | 8080  | 8080 |
| MongoDB       | -      | 27017 |

## Endpoints

### Users

Get all users

`http://$SERVER_IP:8080/api/users`

Exmple of response :
```
[
    {
        "_id": "5fd6314d6dfce4826add769d",
        "age": 18,
        "name": "nezz"
    }
]
```

### User

Post a user

`http://$SERVER_IP:8080/api/user`

Exmple of Body :
```
[
    {
        "age": 18,
        "name": "nezz"
    }
]
```
