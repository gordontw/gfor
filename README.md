gfor - Go HOST
------------
![dataflow](gfor.png)

###Usage
```
-c string
    YAML directory (default ".")
-d    
    Debug mode
-nocache
    Cache mode [default Cached]
-check
    Health Check, will not get host [default false]
```
###DEMO
- [yaml example](src/config.yml)
- ![demo](gfor_demo.gif)

###PHP Extension
- set extension shard object
   - extension=php_gfor.so
```
<?php
host = gfor_host($group, $conf)
gfor_health($group, $conf)
?>
```
```
<?php
$ch = curl_init(); 
curl_setopt($ch, CURLOPT_URL, "http://"+gfor_host($ApiGroup, $conf)+"/"+$uri); 
curl_setopt($ch, CURLOPT_HEADER, TRUE); 
curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE); 
$head = curl_exec($ch); 
?>
```

Author: Gordon Wang <gordon.tw@gmail.com>
