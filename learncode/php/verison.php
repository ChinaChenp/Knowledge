<?php
$foo = 25;
$bar = &$foo;      // 合法的赋值
//$bar = &(24 * 7);  // 非法; 引用没有名字的表达式

function test()
{
   return 25;
}

//$bar = &test();    // 非法
?>
