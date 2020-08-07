# odin

## directory 

```
├── app                         
│   ├── common            
│   │   ├── Const.go     
│   │   ├── util.go     
│   │   └── var.go     
│   ├── entity    
│   │   └── user.go 
│   ├── repository
│   │   ├── memoryDao
│   │   │   └── user.go  
│   │   ├── pikaDao     
│   │   │   └── user.go 
│   │   ├── redisDao   
│   │   │   └── user.go 
│   │   └── repository.go  
│   ├── service     
│   │   ├── serviceBridge.go
│   │   └── service.go  
│   ├── serviceImpl 
│   │   ├── HelloService.go 
│   │   └── UserService.go 
│   ├── serviceInit.go    
│   └── serviceInterface 
│       └── interface.go  
├── bin                 
│   └── odin           
├── cmd               
│   └── odin
│       └── main.go  
├── conf         
│   ├── conf_dev.ini   
│   ├── conf.ini      
│   ├── conf_release.ini  
│   └── conf_online.ini        
├── examples 
│   └── main.go
├── go.mod                     
├── Makefile
├── README.md
└── version            
    └── version.go   
```

