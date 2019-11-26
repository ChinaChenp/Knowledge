#!/bin/bash

for((i=1; i<5; i++))
do
    {
        echo "ok $i";
        sleep $i;
    }&
done

wait

{
    echo "ok 2";
    sleep 2; 
}&

{
    echo "ok 3";
    sleep 3; 
}&

{
    echo "ok 1";
    sleep 1; 
}&

wait