#!/bin/bash


#http://wiki.dreamhost.com/Launch_an_Instance_from_the_Command_Line

nova boot --poll --flavor lightspeed --image Ubuntu-14.04-Trusty --key-name ben --user-data ./user-data.sh my-server

# neutron floatingip-list
# neutron port-list
# neutron floatingip-associate --fixed-ip-address <Private_IPv4_Address> <Floating_IP_ID> <Port_ID> 

neutron floatingip-associate --fixed-ip-address 10.10.10.3 0f63de2b-7131-4c45-b844-d91d8fd487bc 026f93fc-87fb-4423-add6-6fcb35904776



