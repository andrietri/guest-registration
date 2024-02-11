# guest-registration-service


#### Structure
```
├── api/                
│   ├──  v1/    
│   │   ├──  module/
│   │   │   ├──  controller
│   │   │   ├──  dto
│   │   │   ├──  service
│   │   │   ├──  repository
├── database/           
│   ├──  database/    
│   ├──  migration/   
│   ├──  migrationSchema/    
│   ├──  model/  
├── http/           
│   ├──  httpresponse/     
│   ├──  router/     
├── utils/           
│   ├──  config/     
│   ├──  constants/     
│   ├──  helper/     
├── go.mod
└── go.sum
```

#### Architecture
- Controller --> Service --> Repository

### Collection Postman
https://api.postman.com/collections/2441723-09990d8b-72fd-462b-8778-9a2841c11eb2?access_key=PMAT-01HPC1XT64EHV9C4FNNK5EBTEV

#### Local Development Enviroment
Using Visual Studio Code IDE is recommended. Here's several steps that have to be followed : 

- Create launch.json in .vscode folder. 
- Copy-paste this configuration : 
```sh 
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${fileDirname}",
            "showLog" : true,
        }
    ]
```
- Then, run and debug. 
- Watch your Debug Console.

When debug launch for first time, you probably find the port number information in debug console. Don't worry, just use your port number that you had written in source code or enviroment file. 

### Sql-Migrate
- For complete documentation : https://github.com/rubenv/sql-migrate
- Create new file sql : sql-migrate new -config=dbconfig.yml -env="local" create_table_foo
- Check status : sql-migrate status -config=dbconfig.yml -env="local"
- Dryrun query : sql-migrate up -config=dbconfig.yml -env="local" -dryrun
- Run query : sql-migrate up -config=dbconfig.yml -env="local"

