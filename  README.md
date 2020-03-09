#trivag case study for Backend Software Engineer - User Flow position

## General info
1. The task is implemented using [Go](https://golang.org/) as programming language.
2. It is implemented as CLI tool where you will pass the required formats in the terminal.
3. You can use the IDE you want but I recommend [GoLand](https://www.jetbrains.com/go/)

## Installation and setup
You can run the project in two ways,either locally in your machine or by using [Docker](https://www.docker.com/)

##### Local run
1. Unzip the compressed file to $GOROOT/src folder.
2. Run the following command, you can pass the type of format you want i.e only 3 formats are supported 
    ```
    go run *.go xml yaml json
    ```
3. You can also run through the following commands as build, run the first to generate the binary then run the binary using the second command
    ```
    go build .
    ```
    ```
    ./trivago-BackendSoftwareEngineer-UserFlow xml yaml json
    ```
##### Docker run
You can use docker for running the project as follows
    ```
    docker-compose up
    ```
## Notes
1. The project does not have any unit tests but this is due to lack of my time.
2. The output files will be generated in the output folder found in the project folder structure.
3. I added a docke compose file to make the running easy from your side.
   

## Author
[Mohamed Elkashif](mailto:mohamedelkashif922@gmail.com)

## License
[MIT](https://choosealicense.com/licenses/mit/)      