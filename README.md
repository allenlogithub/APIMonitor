# APIMonitor
#### Usage
- To know the parameters:  
`./main_linux_amd64 -h`

- Parameters:  
`-config-path=/path/to/the/file -output-path=/directory/of/the/output/folder`  
(The default output folder will be same as the input file if not passing the -output-path)

#### Run the Source Code
- Build Environment:  
	 &nbsp; `bash docker/build.sh`  
	 &nbsp; `bash docker/run.sh`  
	 &nbsp;	`bash docker/execute.sh`  

- Run the Source Code on Linux Machine:  
   &nbsp;	`bash run_go.sh` (need to modify the path inside)  

- Compile the Source Code on Linux Machine:  
   &nbsp;	`bash build_linux.sh`  

#### Input File  
- Follow the docs/examples.json  
General App Settings:  
		&nbsp;	`async`  
    &nbsp;&nbsp;&nbsp;&nbsp;  (true: async on each API, false: async on each API set)  
		&nbsp;	`domain`  
    &nbsp;&nbsp;&nbsp;&nbsp;  (Target server)  
		&nbsp;  `rounds`  
    &nbsp;&nbsp;&nbsp;&nbsp;  (Rounds for the execution)  
		&nbsp;  `workers`  
    &nbsp;&nbsp;&nbsp;&nbsp;  (worker number)  
API Settings: (cases)  
		&nbsp;  `route`  
    &nbsp;&nbsp;&nbsp;&nbsp;  Route  
		&nbsp;  `request_type`  
    &nbsp;&nbsp;&nbsp;&nbsp;  GET, POST, ...  
		&nbsp;  `est_elapse`  
    &nbsp;&nbsp;&nbsp;&nbsp;  millisecond, will show TimeOut if the runtime was longer than it  
		&nbsp;  `url_params`  
    &nbsp;&nbsp;&nbsp;&nbsp;  "String_key": "String_value"  
		&nbsp;  `headers`  
    &nbsp;&nbsp;&nbsp;&nbsp;  "String_key": "String_value"  
		&nbsp;  `form_params`  
    &nbsp;&nbsp;&nbsp;&nbsp;  "String_key": "value"  
