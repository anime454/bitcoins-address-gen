#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os
import subprocess 

def read_file (filename) :
    with open(filename) as f:
        lines = f.read().splitlines()
    for l in lines :
        res = subprocess.check_output("ruby address-gen.rb {}".format(l), shell=True);
        res = res.rstrip().decode("utf-8")
        os.system("printf '{} : ' >> result.txt".format(l))
        os.system("printf '{} \n' >> result.txt".format(res))
read_file("words.txt")



