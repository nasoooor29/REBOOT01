#!/bin/bash
echo -e "\n---------------------------------------------------------------------------------------"
echo -e "Building Docker Image:"
echo "EXAMPLE: docker image build -f Dockerfile -t <name_of_the_image> ."
echo -e "---------------------------------------------------------------------------------------"
docker image build -f Dockerfile -t ascii-art-web-docker .

echo -e "\n--------------------------------------------------------------------------------------------------------------------"
echo -e "List all docker images:"
echo "EXAMPLE: docker images"
echo -e "--------------------------------------------------------------------------------------------------------------------"
docker images

echo -e "\n--------------------------------------------------------------------------------------------------------------------"
echo -e "Start the docker container:"
echo "EXAMPLE: docker container run -p <port_you_what_to_run> --detach --name <name_of_the_container> <name_of_the_image>"
echo -e "--------------------------------------------------------------------------------------------------------------------"
docker container run -p 8080:8080 -d --name ascii-art-web ascii-art-web-docker

echo -e "\n----------------------------------------------------------------------------------------------------------------------"
echo -e "View all running containers:"
echo "EXAMPLE: docker ps -a"
echo -e "----------------------------------------------------------------------------------------------------------------------"
docker ps -a

echo -e "\n-----------------------------------------------------------------------------------------------------"
echo -e "List the contents of the CWD of the container (ls -la)"
echo "EXAMPLE: docker exec -it <container_name> ls -la"
echo -e "-----------------------------------------------------------------------------------------------------"
docker exec -it ascii-art-web ls -la

echo -e "\n--------------------------------------------------------------------------------------"
echo "ASCII-ART-WEB project available at http://localhost:8080"
echo -e "--------------------------------------------------------------------------------------"