<?php
    error_reporting(0);

    echo gfor_host("group1","./aaa"),"\n";
    echo gfor_host("group1","./"),"\n";
    //echo gfor_host("group1", "src/aaa"),"\n";
    gfor_health("group1","./");
?>
