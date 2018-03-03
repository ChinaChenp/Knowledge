<?php
$a = $argv[1];
if ($a >1) {
    echo "yes\n";

} elseif ($a == 1) {
    echo "equal\n";
}
 else {
    echo "no\n";
}

$total = 10;

for ($i = 0; $i < $total; $i++) {
    if ($i == 5) {
        break;
    }
    echo "for going\n";
}

// array
$people = array(
    array("name" => "chenping", "age" => 15),
    array("name" => "xunbo", "age" => 25)
);

$array_size = count($people);
for ($i = 0; $i < $array_size; $i++) {
    echo var_dump($people[$i]);
}

$arr = array(1, 2, 3, 4, 5, 6);
foreach($arr as &$value) {
    $value *= 2;
}
unset($value);

foreach($arr as $v) {
    echo $v;
    echo "\n";
}

function add_str($str) {
    return array($str .= "you\n", "kao\n");
}

$re = add_str("fuck");
var_dump($re);

$file_path = "shell.php";
if (file_exists($file_path)) {
    $fp = fopen($file_path, "r");
    $txt= fread($fp,filesize($file_path));
    echo $txt;
} else {
    echo "shit\n";
}

if ($a == 5){
    echo "a equals 5";
    echo "...";
} elseif ($a == 6){
    echo "a equals 6";
    echo "!!!";
}else{
    echo "a is neither 5 nor 6";
}
?>
