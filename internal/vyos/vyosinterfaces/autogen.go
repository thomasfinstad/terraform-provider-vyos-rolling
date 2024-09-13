// Autogenerated by Makefile. DO NOT EDIT.
package vyosinterface

import "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"

// GetInterfaces returns all autogenerated interface definitions
func GetInterfaces() []schemadefinition.InterfaceDefinition {
	return []schemadefinition.InterfaceDefinition{
		container(),
		firewall(),
		highavailability(),
		interfaces_bonding(),
		interfaces_bridge(),
		interfaces_dummy(),
		interfaces_ethernet(),
		interfaces_geneve(),
		interfaces_input(),
		interfaces_loopback(),
		interfaces_macsec(),
		interfaces_openvpn(),
		interfaces_pppoe(),
		interfaces_pseudoethernet(),
		interfaces_sstpc(),
		interfaces_tunnel(),
		interfaces_virtualethernet(),
		interfaces_vti(),
		interfaces_vxlan(),
		interfaces_wireguard(),
		interfaces_wireless(),
		interfaces_wwan(),
		loadbalancing_reverseproxy(),
		loadbalancing_wan(),
		nat(),
		nat_cgnat(),
		netns(),
		pki(),
		policy(),
		policy_localroute(),
		policy_route(),
		protocols_babel(),
		protocols_bfd(),
		protocols_bgp(),
		protocols_eigrp(),
		protocols_failover(),
		protocols_igmpproxy(),
		protocols_isis(),
		protocols_mpls(),
		protocols_nhrp(),
		protocols_openfabric(),
		protocols_ospf(),
		protocols_pim(),
		protocols_rip(),
		protocols_ripng(),
		protocols_rpki(),
		protocols_segmentrouting(),
		protocols_static(),
		protocols_static_arp(),
		protocols_static_multicast(),
		protocols_static_neighborproxy(),
		qos(),
		service_aws_glb(),
		service_broadcastrelay(),
		service_configsync(),
		service_conntracksync(),
		service_consoleserver(),
		service_dhcprelay(),
		service_dhcpserver(),
		service_dns_dynamic(),
		service_dns_forwarding(),
		service_eventhandler(),
		service_https(),
		service_ids_ddosprotection(),
		service_ipoeserver(),
		service_lldp(),
		service_mdns_repeater(),
		service_monitoring_telegraf(),
		service_monitoring_zabbixagent(),
		service_ndpproxy(),
		service_ntp(),
		service_pppoeserver(),
		service_routeradvert(),
		service_saltminion(),
		service_sla(),
		service_snmp(),
		service_ssh(),
		service_stunnel(),
		service_suricata(),
		service_tftpserver(),
		service_webproxy(),
		system_acceleration(),
		system_configmanagement(),
		system_conntrack(),
		system_console(),
		system_domainname(),
		system_domainsearch(),
		system_flowaccounting(),
		system_frr(),
		system_hostname(),
		system_ip(),
		system_lcd(),
		system_login(),
		system_login_banner(),
		system_logs(),
		system_nameserver(),
		system_option(),
		system_proxy(),
		system_sflow(),
		system_statichostmapping(),
		system_sysctl(),
		system_syslog(),
		system_taskscheduler(),
		system_timezone(),
		system_updatecheck(),
		system_wireless(),
		vpn_ipsec(),
		vpn_openconnect(),
		vpn_pptp(),
		vpn_sstp(),
		vrf(),
	}
}
