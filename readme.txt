Extract this folder


And run following commands

# please ensure the docker service is running before running the command.
docker-compose up --build -d 

docker exec -it seller_app_go /bin/sh


cd seller_app_go/


# run command to start the server
buffalo dev

# you can now run the api on
localhost:3005


Following api lists have been added in the postman collection SellerApp.postman_collection.json
Api to add an auction
Api to fetch all open auction
Api to bid for an auction



#to deactivate the expired auction, run following commands
docker exec -it seller_app_go /bin/sh

cd seller_app_go/

buffalo task close_expired_auction