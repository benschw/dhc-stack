# Nova CLI Demo


## Allocate Floating IP

	$ nova floating-ip-pool-list
	+------------+
	| name       |
	+------------+
	| public-110 |
	+------------+

	$ neutron floatingip-create public-110
	Created a new floatingip:
	+---------------------+--------------------------------------+
	| Field               | Value                                |
	+---------------------+--------------------------------------+
	| fixed_ip_address    |                                      |
	| floating_ip_address | 173.236.248.165                      |
	| floating_network_id | b576a0f4-a0fc-4a1a-bea3-9e18bd663b64 |
	| id                  | 1ed3a606-1941-4cd4-b44d-475f1f5015b3 |
	| port_id             |                                      |
	| router_id           |                                      |
	| tenant_id           | 406b6dce0d6949f69a2e7d309cb3a3b5     |
	+---------------------+--------------------------------------+

## Launch Instance

	$ nova boot --poll --flavor lightspeed --image Ubuntu-14.04-Trusty --key-name ben --user-data ./user-data.sh my-server
	+--------------------------------------+------------------------------------------------------------+
	| Property                             | Value                                                      |
	+--------------------------------------+------------------------------------------------------------+
	| OS-DCF:diskConfig                    | MANUAL                                                     |
	| OS-EXT-AZ:availability_zone          | nova                                                       |
	| OS-EXT-STS:power_state               | 0                                                          |
	| OS-EXT-STS:task_state                | scheduling                                                 |
	| OS-EXT-STS:vm_state                  | building                                                   |
	| OS-SRV-USG:launched_at               | -                                                          |
	| OS-SRV-USG:terminated_at             | -                                                          |
	| accessIPv4                           |                                                            |
	| accessIPv6                           |                                                            |
	| adminPass                            | EFoizUaA4D95                                               |
	| config_drive                         |                                                            |
	| created                              | 2014-06-16T14:23:54Z                                       |
	| flavor                               | lightspeed (300)                                           |
	| hostId                               |                                                            |
	| id                                   | ae5a9834-15ca-42d2-b9d3-87faaeb7c933                       |
	| image                                | Ubuntu-14.04-Trusty (8363ff61-55a8-4d4f-9867-fb913e4e5e49) |
	| key_name                             | ben                                                        |
	| metadata                             | {}                                                         |
	| name                                 | my-server                                                  |
	| os-extended-volumes:volumes_attached | []                                                         |
	| progress                             | 0                                                          |
	| security_groups                      | default                                                    |
	| status                               | BUILD                                                      |
	| tenant_id                            | 406b6dce0d6949f69a2e7d309cb3a3b5                           |
	| updated                              | 2014-06-16T14:23:55Z                                       |
	| user_id                              | 4b5ae78857604bf79b3e10ca6798b761                           |
	+--------------------------------------+------------------------------------------------------------+
	Server building... 100% complete
	Finished


## Associate Floating IP

	$ nova list
	+--------------------------------------+-----------+--------+------------+-------------+--------------------------------------------------------------------------------------+
	| ID                                   | Name      | Status | Task State | Power State | Networks                                                                             |
	+--------------------------------------+-----------+--------+------------+-------------+--------------------------------------------------------------------------------------+
	| eb7043e4-4960-4753-9dfc-258ef9449738 | foo       | ACTIVE | -          | Running     | private-network=2607:f298:6050:c090:f816:3eff:fe81:ec7b, 10.10.10.2, 173.236.248.115 |
	| 5e8a6e4f-868b-4ebd-80c9-e6e2f0e051a5 | my-server | ACTIVE | -          | Running     | private-network=2607:f298:6050:c090:f816:3eff:fe77:4d5f, 10.10.10.4, 173.236.248.118 |
	| ae5a9834-15ca-42d2-b9d3-87faaeb7c933 | my-server | ACTIVE | -          | Running     | private-network=2607:f298:6050:c090:f816:3eff:fe9e:1c19, 10.10.10.5                  |
	+--------------------------------------+-----------+--------+------------+-------------+--------------------------------------------------------------------------------------+

	$ neutron floatingip-list
	+--------------------------------------+------------------+---------------------+--------------------------------------+
	| id                                   | fixed_ip_address | floating_ip_address | port_id                              |
	+--------------------------------------+------------------+---------------------+--------------------------------------+
	| 0f63de2b-7131-4c45-b844-d91d8fd487bc | 10.10.10.4       | 173.236.248.118     | aff30e1b-64b9-4011-abce-ac3ead09e5e8 |
	| 1ed3a606-1941-4cd4-b44d-475f1f5015b3 |                  | 173.236.248.165     |                                      |
	| be9461b6-30dd-4ef7-859e-d2f40b557efe | 10.10.10.2       | 173.236.248.115     | e71bec84-e35e-494f-a85a-7e4c4b19c360 |
	+--------------------------------------+------------------+---------------------+--------------------------------------+

	$ neutron port-list
	+--------------------------------------+------+-------------------+----------------------------------------------------------------------------------------------------------------+
	| id                                   | name | mac_address       | fixed_ips                                                                                                      |
	+--------------------------------------+------+-------------------+----------------------------------------------------------------------------------------------------------------+
	| 27b4bc18-ee46-4336-af9a-7b818e84d8e9 |      | fa:16:3e:fb:22:1b | {"subnet_id": "3ffd7206-ab29-4433-a618-784e391fc557", "ip_address": "10.10.10.1"}                              |
	|                                      |      |                   | {"subnet_id": "12bc5832-9455-40eb-8e0d-116cad5b59d3", "ip_address": "2607:f298:6050:c090::1"}                  |
	| aff30e1b-64b9-4011-abce-ac3ead09e5e8 |      | fa:16:3e:77:4d:5f | {"subnet_id": "3ffd7206-ab29-4433-a618-784e391fc557", "ip_address": "10.10.10.4"}                              |
	|                                      |      |                   | {"subnet_id": "12bc5832-9455-40eb-8e0d-116cad5b59d3", "ip_address": "2607:f298:6050:c090:f816:3eff:fe77:4d5f"} |
	| c6d900ef-b0a3-46d9-a558-4711ce745fbb |      | fa:16:3e:9e:1c:19 | {"subnet_id": "3ffd7206-ab29-4433-a618-784e391fc557", "ip_address": "10.10.10.5"}                              |
	|                                      |      |                   | {"subnet_id": "12bc5832-9455-40eb-8e0d-116cad5b59d3", "ip_address": "2607:f298:6050:c090:f816:3eff:fe9e:1c19"} |
	| e71bec84-e35e-494f-a85a-7e4c4b19c360 |      | fa:16:3e:81:ec:7b | {"subnet_id": "3ffd7206-ab29-4433-a618-784e391fc557", "ip_address": "10.10.10.2"}                              |
	|                                      |      |                   | {"subnet_id": "12bc5832-9455-40eb-8e0d-116cad5b59d3", "ip_address": "2607:f298:6050:c090:f816:3eff:fe81:ec7b"} |
	+--------------------------------------+------+-------------------+----------------------------------------------------------------------------------------------------------------+


	# neutron floatingip-associate --fixed-ip-address <Private_IPv4_Address> <Floating_IP_ID> <Port_ID> 

	$ neutron floatingip-associate --fixed-ip-address 10.10.10.5 1ed3a606-1941-4cd4-b44d-475f1f5015b3 c6d900ef-b0a3-46d9-a558-4711ce745fbb
	Associated floatingip 1ed3a606-1941-4cd4-b44d-475f1f5015b3
