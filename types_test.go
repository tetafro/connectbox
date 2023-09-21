package connectbox

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalXML(t *testing.T) {
	testCases := []struct {
		name string
		data string
		in   any
		out  any
	}{
		{
			name: "GlobalSettings",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<GlobalSettings>
					<AccessLevel>1</AccessLevel>
					<SwVersion>CH7465LG-NCIP-6</SwVersion>
					<CmProvisionMode>IPv4</CmProvisionMode>
					<DsLite>0</DsLite>
					<GwProvisionMode>IPv4/IPv6</GwProvisionMode>
					<GWOperMode>IPv4/IPv6</GWOperMode>
					<ConfigVenderModel>CH7465LG</ConfigVenderModel>
					<HideRemoteAccess>True</HideRemoteAccess>
					<HideModemMode>True</HideModemMode>
					<HideCustomerDhcpLanChange>0</HideCustomerDhcpLanChange>
					<ShowDDNS>True</ShowDDNS>
					<OperatorId>ZIGGO</OperatorId>
					<AccessDenied>NONE</AccessDenied>
					<LockedOut>Disable</LockedOut>
					<CountryID>7</CountryID>
					<title>Connect Box</title>
					<Interface>1</Interface>
					<operStatus>1</operStatus>
				</GlobalSettings>`,
			in: &GlobalSettings{},
			out: &GlobalSettings{
				AccessLevel:               "1",
				SwVersion:                 "CH7465LG-NCIP-6",
				CmProvisionMode:           "IPv4",
				DsLite:                    "0",
				GwProvisionMode:           "IPv4/IPv6",
				GwOperMode:                "IPv4/IPv6",
				ConfigVenderModel:         "CH7465LG",
				HideRemoteAccess:          "True",
				HideModemMode:             "True",
				HideCustomerDHCPLANChange: "0",
				ShowDDNS:                  "True",
				OperatorID:                "ZIGGO",
				AccessDenied:              "NONE",
				LockedOut:                 "Disable",
				CountryID:                 "7",
				Title:                     "Connect Box",
				Interface:                 "1",
				OperStatus:                "1",
			},
		},
		{
			name: "CMSystemInfo",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<cm_system_info>
					<cm_docsis_mode>DOCSIS 3.0</cm_docsis_mode>
					<cm_hardware_version>5.01</cm_hardware_version>
					<cm_mac_addr>00:11:22:33:44:55</cm_mac_addr>
					<cm_serial_number>DEAP1300000A</cm_serial_number>
					<cm_system_uptime>4day(s)16h:30m:35s</cm_system_uptime>
					<cm_network_access>Allowed</cm_network_access>
				</cm_system_info>`,
			in: &CMSystemInfo{},
			out: &CMSystemInfo{
				DocsisMode:      "DOCSIS 3.0",
				HardwareVersion: "5.01",
				MacAddr:         "00:11:22:33:44:55",
				SerialNumber:    "DEAP1300000A",
				SystemUptime:    405035,
				NetworkAccess:   "Allowed",
			},
		},
		{
			name: "Multilang",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<multilang>
					<WebCapPor>0</WebCapPor>
					<Lang>en</Lang>
				</multilang>`,
			in: &Multilang{},
			out: &Multilang{
				WebCapPor: "0",
				Lang:      "en",
			},
		},
		{
			name: "Status",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<status>
					<cm_status>OPERATIONAL</cm_status>
					<Bandmode>1</Bandmode>
					<BssEnable2g>1</BssEnable2g>
					<SSID2G>home</SSID2G>
					<PreSharedKey2gLength>20</PreSharedKey2gLength>
					<BssEnable5g>1</BssEnable5g>
					<SSID5G>home</SSID5G>
					<PreSharedKey5gLength>20</PreSharedKey5gLength>
					<LanUserCount>15</LanUserCount>
				</status>`,
			in: &Status{},
			out: &Status{
				CMStatus:             "OPERATIONAL",
				Bandmode:             "1",
				BSSEnable2G:          "1",
				SSID2G:               "home",
				PreSharedKey2GLength: "20",
				BssEnable5G:          "1",
				SSID5G:               "home",
				PreSharedKey5GLength: "20",
				LANUserCount:         "15",
			},
		},
		{
			name: "Configuration",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<configuration>
					<FrequencyPlan>2</FrequencyPlan>
					<Frequency>730000000</Frequency>
				</configuration>`,
			in: &Configuration{},
			out: &Configuration{
				FrequencyPlan: "2",
				Frequency:     "730000000",
			},
		},
		{
			name: "DownstreamTable",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<downstream_table>
					<ds_num>30</ds_num>
					<downstream>
						<freq>826000000</freq>
						<pow>6</pow>
						<snr>38</snr>
						<mod>256qam</mod>
						<chid>32</chid>
						<RxMER>38.701</RxMER>
						<PreRs>13810000000</PreRs>
						<PostRs>500</PostRs>
						<IsQamLocked>1</IsQamLocked>
						<IsFECLocked>1</IsFECLocked>
						<IsMpegLocked>1</IsMpegLocked>
					</downstream>
					<downstream>
						<freq>754000000</freq>
						<pow>7</pow>
						<snr>38</snr>
						<mod>256qam</mod>
						<chid>23</chid>
						<RxMER>38.701</RxMER>
						<PreRs>13810000000</PreRs>
						<PostRs>392</PostRs>
						<IsQamLocked>1</IsQamLocked>
						<IsFECLocked>1</IsFECLocked>
						<IsMpegLocked>1</IsMpegLocked>
					</downstream>
				</downstream_table>`,
			in: &DownstreamTable{},
			out: &DownstreamTable{
				DsNum: "30",
				Downstreams: []DownstreamTableDownstream{
					{
						Freq:         "826000000",
						Pow:          "6",
						Snr:          "38",
						Mod:          "256qam",
						Chid:         "32",
						RxMER:        "38.701",
						PreRs:        "13810000000",
						PostRs:       "500",
						IsQamLocked:  "1",
						IsFECLocked:  "1",
						IsMpegLocked: "1",
					},
					{
						Freq:         "754000000",
						Pow:          "7",
						Snr:          "38",
						Mod:          "256qam",
						Chid:         "23",
						RxMER:        "38.701",
						PreRs:        "13810000000",
						PostRs:       "392",
						IsQamLocked:  "1",
						IsFECLocked:  "1",
						IsMpegLocked: "1",
					},
				},
			},
		},
		{
			name: "UpstreamTable",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<upstream_table>
					<us_num>5</us_num>
					<upstream>
						<usid>9</usid>
						<freq>13800000</freq>
						<power>41</power>
						<srate>5.120</srate>
						<mod>64qam</mod>
						<ustype>3</ustype>
						<t1Timeouts>0</t1Timeouts>
						<t2Timeouts>0</t2Timeouts>
						<t3Timeouts>8</t3Timeouts>
						<t4Timeouts>0</t4Timeouts>
						<channeltype>ATDMA</channeltype>
						<messageType>31</messageType>
					</upstream>
					<upstream>
						<usid>7</usid>
						<freq>58800000</freq>
						<power>42</power>
						<srate>5.120</srate>
						<mod>64qam</mod>
						<ustype>3</ustype>
						<t1Timeouts>0</t1Timeouts>
						<t2Timeouts>0</t2Timeouts>
						<t3Timeouts>9</t3Timeouts>
						<t4Timeouts>0</t4Timeouts>
						<channeltype>ATDMA</channeltype>
						<messageType>31</messageType>
					</upstream>
				</upstream_table>`,
			in: &UpstreamTable{},
			out: &UpstreamTable{
				UsNum: "5",
				Upstreams: []UpstreamTableUpstream{
					{
						Usid:        "9",
						Freq:        "13800000",
						Power:       "41",
						Srate:       "5.120",
						Mod:         "64qam",
						Ustype:      "3",
						T1Timeouts:  "0",
						T2Timeouts:  "0",
						T3Timeouts:  "8",
						T4Timeouts:  "0",
						Channeltype: "ATDMA",
						MessageType: "31",
					},
					{
						Usid:        "7",
						Freq:        "58800000",
						Power:       "42",
						Srate:       "5.120",
						Mod:         "64qam",
						Ustype:      "3",
						T1Timeouts:  "0",
						T2Timeouts:  "0",
						T3Timeouts:  "9",
						T4Timeouts:  "0",
						Channeltype: "ATDMA",
						MessageType: "31",
					},
				},
			},
		},
		{
			name: "SignalTable",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<signal_table>
					<sig_num>24</sig_num>
					<signal>
						<dsid>16</dsid>
						<unerrored>13810200000</unerrored>
						<correctable>325</correctable>
						<uncorrectable>0</uncorrectable>
					</signal>
					<signal>
						<dsid>12</dsid>
						<unerrored>13810000000</unerrored>
						<correctable>58</correctable>
						<uncorrectable>0</uncorrectable>
					</signal>
				</signal_table>`,
			in: &SignalTable{},
			out: &SignalTable{
				SigNum: "24",
				Signals: []SignalTableSignal{
					{
						Dsid:          "16",
						Unerrored:     "13810200000",
						Correctable:   "325",
						Uncorrectable: "0",
					},
					{
						Dsid:          "12",
						Unerrored:     "13810000000",
						Correctable:   "58",
						Uncorrectable: "0",
					},
				},
			},
		},
		{
			name: "EventLogTable",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<eventlog_table>
					<eventlog>
						<prior>notice</prior>
						<text>GUI Login Status - Login Success from LAN interface</text>
						<time>20-09-2023 14:40:41</time>
						<t>1695813641</t>
					</eventlog>
					<eventlog>
						<prior>notice</prior>
						<text>Illegal - Dropped INPUT packet</text>
						<time>20-09-2023 14:40:50</time>
						<t>1692213650</t>
					</eventlog>
				</eventlog_table>`,
			in: &EventLogTable{},
			out: &EventLogTable{
				EventLogs: []EventLogTableEventLog{
					{
						Prior: "notice",
						Text:  "GUI Login Status - Login Success from LAN interface",
						Time:  "20-09-2023 14:40:41",
						T:     "1695813641",
					},
					{
						Prior: "notice",
						Text:  "Illegal - Dropped INPUT packet",
						Time:  "20-09-2023 14:40:50",
						T:     "1692213650",
					},
				},
			},
		},
		{
			name: "FirewallLogTable",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<firewalllog_table>
					<firewalllog>
						<prior>notice</prior>
						<text>GUI Login Status - Login Success from LAN interface</text>
						<time>20-09-2023 14:40:41</time>
					</firewalllog>
					<firewalllog>
						<prior>notice</prior>
						<text>Illegal - Dropped INPUT packet</text>
						<time>20-09-2023 14:40:50</time>
					</firewalllog>
				</firewalllog_table>`,
			in: &FirewallLogTable{},
			out: &FirewallLogTable{
				FirewallLogs: []FirewallLogTableFirewallLog{
					{
						Prior: "notice",
						Text:  "GUI Login Status - Login Success from LAN interface",
						Time:  "20-09-2023 14:40:41",
					},
					{
						Prior: "notice",
						Text:  "Illegal - Dropped INPUT packet",
						Time:  "20-09-2023 14:40:50",
					},
				},
			},
		},
		{
			name: "Langsetlist",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<langsetlist>
					<langSet_support>en</langSet_support>
					<langSet_support>cz</langSet_support>
					<langSet_support>pl</langSet_support>
					<langSet_support>sk</langSet_support>
					<langSet_support>fr</langSet_support>
					<langSet_support>it</langSet_support>
				</langsetlist>`,
			in: &Langsetlist{},
			out: &Langsetlist{
				LangSetSupport: []string{"en", "cz", "pl", "sk", "fr", "it"},
			},
		},
		{
			name: "Fail",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<Fail>
					<FailCount>0</FailCount>
				</Fail>`,
			in: &Fail{},
			out: &Fail{
				FailCount: "0",
			},
		},
		{
			name: "LoginTimer",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<login_timer>
					<Flag>0</Flag>
					<AccessLevel>1</AccessLevel>
				</login_timer>`,
			in: &LoginTimer{},
			out: &LoginTimer{
				Flag:        "0",
				AccessLevel: "1",
			},
		},
		{
			name: "LANSetting",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<LANSetting>
					<UPnP>1</UPnP>
					<LanMAC>00:11:22:33:44:55</LanMAC>
					<LanIP>10.0.0.1</LanIP>
					<DMZaddr>10.0.0.1</DMZaddr>
					<DMZ>0</DMZ>
					<LanIPv6>2222:aaaa:29be:1000:6aaa:ffff:ffff:0001/64</LanIPv6>
					<LanIPv6Prefix>2222:aaaa:29be:1000::/64</LanIPv6Prefix>
					<subnetmask>10.0.0.1</subnetmask>
					<DHCP_startaddress>10.0.0.1</DHCP_startaddress>
					<DHCP_endaddress>10.0.0.1</DHCP_endaddress>
				</LANSetting>`,
			in: &LANSetting{},
			out: &LANSetting{
				UPnP:             "1",
				LANMAC:           "00:11:22:33:44:55",
				LANIP:            "10.0.0.1",
				DMZAddr:          "10.0.0.1",
				DMZ:              "0",
				LanIPv6:          "2222:aaaa:29be:1000:6aaa:ffff:ffff:0001/64",
				LanIPv6Prefix:    "2222:aaaa:29be:1000::/64",
				SubnetMask:       "10.0.0.1",
				DHCPStartAddress: "10.0.0.1",
				DHCPEndAddress:   "10.0.0.1",
			},
		},
		{
			name: "DHCPv6Info",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<DHCPv6Info>
					<AllowDHCPv6Setting>1</AllowDHCPv6Setting>
					<ipv6RAManagedflag>0</ipv6RAManagedflag>
					<ipv6_saddr>2222:aaaa:29be:1000::/64</ipv6_saddr>
					<ipv6_prefix>2222:aaaa:29be:1000::</ipv6_prefix>
					<NumberOfAddr>245</NumberOfAddr>
					<ipv6PrefixPreferredLifeTime>602400</ipv6PrefixPreferredLifeTime>
					<ipv6PrefixValidLifeTime>1502491</ipv6PrefixValidLifeTime>
					<dhcpV6AddrLifeTime>0</dhcpV6AddrLifeTime>
					<ipv6RALifetime>1800</ipv6RALifetime>
					<ipv6RAIntervaltime>180</ipv6RAIntervaltime>
				</DHCPv6Info>`,
			in: &DHCPv6Info{},
			out: &DHCPv6Info{
				AllowDHCPv6Setting:          "1",
				IPv6RAManagedFlag:           "0",
				IPv6Saddr:                   "2222:aaaa:29be:1000::/64",
				IPv6Prefix:                  "2222:aaaa:29be:1000::",
				NumberOfAddr:                "245",
				IPv6PrefixPreferredLifeTime: "602400",
				IPv6PrefixValidLifeTime:     "1502491",
				DHCPV6AddrLifeTime:          "0",
				IPv6RALifetime:              "1800",
				IPv6RAIntervaltime:          "180",
			},
		},
		{
			name: "BasicDHCP",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<BasicDHCP>
					<enableDHCPv4>1</enableDHCPv4>
					<Addr_start>10.0.0.1</Addr_start>
					<NumberOfCpes>45</NumberOfCpes>
					<LeaseTime>86400</LeaseTime>
					<LanIP>10.0.0.1</LanIP>
					<subnetmask>10.0.0.1</subnetmask>
					<ReserveIpadrr>
						<MacAddress>00:11:22:33:44:55</MacAddress>
						<LeasedIP>10.0.0.1</LeasedIP>
					</ReserveIpadrr>
					<ReserveIpadrr>
						<MacAddress>00:11:22:33:44:55</MacAddress>
						<LeasedIP>10.0.0.1</LeasedIP>
					</ReserveIpadrr>
					<BlockSubnetIP>10.0.0.1</BlockSubnetIP>
					<BlockSubnetMask>10.0.0.1</BlockSubnetMask>
					<BlockSubnetIP>10.0.0.1</BlockSubnetIP>
					<BlockSubnetMask>10.0.0.1</BlockSubnetMask>
					<HideCustomerDhcpLanChange>0</HideCustomerDhcpLanChange>
				</BasicDHCP>`,
			in: &BasicDHCP{},
			out: &BasicDHCP{
				EnableDHCPv4: "1",
				AddrStart:    "10.0.0.1",
				NumberOfCpes: "45",
				LeaseTime:    "86400",
				LanIP:        "10.0.0.1",
				SubnetMask:   "10.0.0.1",
				ReserveIPAddrs: []BasicDHCPReserveIPAddrs{
					{
						MacAddress: "00:11:22:33:44:55",
						LeasedIP:   "10.0.0.1",
					},
					{
						MacAddress: "00:11:22:33:44:55",
						LeasedIP:   "10.0.0.1",
					},
				},
				BlockSubnetIP:             []string{"10.0.0.1", "10.0.0.1"},
				BlockSubnetMask:           []string{"10.0.0.1", "10.0.0.1"},
				HideCustomerDHCPLANChange: "0",
			},
		},
		{
			name: "WANSetting",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<WANSetting>
					<NAPT_mode>1</NAPT_mode>
					<WanMAC>00:11:22:33:44:55</WanMAC>
					<wan_ipv6_addr>
						<wan_ipv6_addr_entry>bbbb:aaaa:0:5555:4444:3333:2222:0000/128</wan_ipv6_addr_entry>
						<wan_ipv6_addr_entry>bbbb:aaaa:0:5555:4444:3333/64</wan_ipv6_addr_entry>
					</wan_ipv6_addr>
					<WanDhcpv6Srv>aaaa::bbbb:cccc:eeee:dddd</WanDhcpv6Srv>
					<ipv6_LeaseTime>D:7 H:0 M:0 S:0</ipv6_LeaseTime>
					<ipv6_LeaseExpire>Mon Sep 25 03:50:03 2023</ipv6_LeaseExpire>
					<wan_ipv6_dnsaddr>
						<wan_ipv6_dnsaddr_entry>2001:9999:9999:1000::53</wan_ipv6_dnsaddr_entry>
						<wan_ipv6_dnsaddr_entry>2001:9999:9999::53</wan_ipv6_dnsaddr_entry>
					</wan_ipv6_dnsaddr>
					<WanIP>10.0.0.1</WanIP>
					<gateway_address>10.0.0.1</gateway_address>
					<LeaseTime>D:0 H:2 M:0 S:0</LeaseTime>
					<LeaseExpire>Wed Sep 20 16:30:58 2023</LeaseExpire>
					<wan_ipv4_dnsaddr>
						<wan_ipv4_dnsaddr_entry>10.0.0.1</wan_ipv4_dnsaddr_entry>
						<wan_ipv4_dnsaddr_entry>10.0.0.1</wan_ipv4_dnsaddr_entry>
					</wan_ipv4_dnsaddr>
					<dslite_enable>0</dslite_enable>
					<dslite_fqdn>aftr01.upc.nl</dslite_fqdn>
					<dslite_addr>2222:3333:4444:5555:8888:aaaa:bbbb:1111</dslite_addr>
				</WANSetting>`,
			in: &WANSetting{},
			out: &WANSetting{
				NAPTMode: "1",
				WANMAC:   "00:11:22:33:44:55",
				WANIPv6Addrs: []string{
					"bbbb:aaaa:0:5555:4444:3333:2222:0000/128",
					"bbbb:aaaa:0:5555:4444:3333/64",
				},
				WANDHCPv6Srv:    "aaaa::bbbb:cccc:eeee:dddd",
				IPv6LeaseTime:   "D:7 H:0 M:0 S:0",
				IPv6LeaseExpire: "Mon Sep 25 03:50:03 2023",
				WANIPv6DNSAddr: []string{
					"2001:9999:9999:1000::53",
					"2001:9999:9999::53",
				},
				WANIP:          "10.0.0.1",
				GatewayAddress: "10.0.0.1",
				LeaseTime:      "D:0 H:2 M:0 S:0",
				LeaseExpire:    "Wed Sep 20 16:30:58 2023",
				WANIPv4DNSAddr: []string{"10.0.0.1", "10.0.0.1"},
				DsliteEnable:   "0",
				DsliteFqdn:     "aftr01.upc.nl",
				DsliteAddr:     "2222:3333:4444:5555:8888:aaaa:bbbb:1111",
			},
		},
		{
			name: "IPFiltering",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<IPfiltering>
					<LanIP>10.0.0.1</LanIP>
					<subnetmask>255.0.0.0</subnetmask>
					<time_mode>0</time_mode>
					<GeneralTime />
					<DailyTime />
				</IPfiltering>`,
			in: &IPFiltering{},
			out: &IPFiltering{
				LanIP:       "10.0.0.1",
				SubnetMask:  "255.0.0.0",
				TimeMode:    "0",
				GeneralTime: "",
				DailyTime:   "",
			},
		},
		{
			name: "IPv6Filtering",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<IPv6filtering>
					<ipv6_prefix>2222:aaaa:1111:5555::</ipv6_prefix>
					<dir>0</dir>
					<time_mode>1</time_mode>
					<GeneralTime />
					<DailyTime />
				</IPv6filtering>`,
			in: &IPv6Filtering{},
			out: &IPv6Filtering{
				IPv6Prefix:  "2222:aaaa:1111:5555::",
				Dir:         "0",
				TimeMode:    "1",
				GeneralTime: "",
				DailyTime:   "",
			},
		},
		{
			name: "PortTrigger",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<PortTrigger />`,
			in:  &PortTrigger{},
			out: &PortTrigger{},
		},
		{
			name: "WebFilter",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<WebFilter>
					<firewallProtection>1</firewallProtection>
					<blockIpFragments>2</blockIpFragments>
					<portScanDetection>3</portScanDetection>
					<synFloodDetection>4</synFloodDetection>
					<IcmpFloodDetection>5</IcmpFloodDetection>
					<IcmpFloodDetectRate>6</IcmpFloodDetectRate>
				</WebFilter>`,
			in: &WebFilter{},
			out: &WebFilter{
				FirewallProtection:  "1",
				BlockIPFragments:    "2",
				PortScanDetection:   "3",
				SynFloodDetection:   "4",
				ICMPFloodDetection:  "5",
				ICMPFloodDetectRate: "6",
			},
		},
		{
			name: "IPv6WebFilter",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<IPv6WebFilter>
					<IPv6firewallProtection>1</IPv6firewallProtection>
					<IPv6blockIpFragments>2</IPv6blockIpFragments>
					<IPv6portScanDetection>3</IPv6portScanDetection>
					<IPv6synFloodDetection>4</IPv6synFloodDetection>
					<IPv6IcmpFloodDetection>5</IPv6IcmpFloodDetection>
					<IPv6IcmpFloodDetectRate>6</IPv6IcmpFloodDetectRate>
				</IPv6WebFilter>`,
			in: &IPv6WebFilter{},
			out: &IPv6WebFilter{
				IPv6FirewallProtection:  "1",
				IPv6BlockIPFragments:    "2",
				IPv6PortScanDetection:   "3",
				IPv6SynFloodDetection:   "4",
				IPv6ICMPFloodDetection:  "5",
				IPv6ICMPFloodDetectRate: "6",
			},
		},
		{
			name: "MACFiltering",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<MACFiltering>
					<maxInstance>32</maxInstance>
					<time_mode>0</time_mode>
					<GeneralTime />
					<DailyTime />
				</MACFiltering>`,
			in: &MACFiltering{},
			out: &MACFiltering{
				MaxInstance: "32",
				TimeMode:    "0",
				GeneralTime: "",
				DailyTime:   "",
			},
		},
		{
			name: "Forwarding",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<Forwarding>
					<LanIP>10.0.0.1</LanIP>
					<subnetmask>255.0.0.0</subnetmask>
					<UPnP>
						<LanIPAddr>10.0.0.1</LanIPAddr>
						<LanPort>9090</LanPort>
						<WanPort>9099</WanPort>
						<Protocol>3</Protocol>
						<Description>10.0.0.1:9090 to 9090 (UDP)</Description>
					</UPnP>
					<UPnP>
						<LanIPAddr>10.0.0.1</LanIPAddr>
						<LanPort>32564</LanPort>
						<WanPort>31677</WanPort>
						<Protocol>3</Protocol>
						<Description>Transmission at 51413</Description>
					</UPnP>
				</Forwarding>`,
			in: &Forwarding{},
			out: &Forwarding{
				LANIP:      "10.0.0.1",
				SubnetMask: "255.0.0.0",
				UPnPs: []ForwardingUPnP{
					{
						LANIPAddr:   "10.0.0.1",
						LANPort:     "9090",
						WANPort:     "9099",
						Protocol:    "3",
						Description: "10.0.0.1:9090 to 9090 (UDP)",
					},
					{
						LANIPAddr:   "10.0.0.1",
						LANPort:     "32564",
						WANPort:     "31677",
						Protocol:    "3",
						Description: "Transmission at 51413",
					},
				},
			},
		},
		{
			name: "LANUserTable",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<LanUserTable>
					<Ethernet>
						<clientinfo>
							<interface>Ethernet 4</interface>
							<IPv4Addr>10.0.0.1/24</IPv4Addr>
							<xmlhostname></xmlhostname>
							<xmlicon></xmlicon>
							<index>3</index>
							<interfaceid>2</interfaceid>
							<hostname>Unknown</hostname>
							<MACAddr>00:11:22:33:44:44</MACAddr>
							<method>2</method>
							<leaseTime>00:00:00:00</leaseTime>
							<speed>1000</speed>
						</clientinfo>
						<clientinfo>
							<interface>Ethernet 3</interface>
							<IPv4Addr>10.0.0.2/24</IPv4Addr>
							<xmlhostname></xmlhostname>
							<xmlicon></xmlicon>
							<index>5</index>
							<interfaceid>2</interfaceid>
							<hostname>Unknown</hostname>
							<MACAddr>00:11:22:33:44:66</MACAddr>
							<method>2</method>
							<leaseTime>00:00:45:45</leaseTime>
							<speed>1000</speed>
						</clientinfo>
					</Ethernet>
					<WIFI>
						<clientinfo>
							<interface>home</interface>
							<IPv4Addr>10.0.0.1/24</IPv4Addr>
							<xmlhostname></xmlhostname>
							<xmlicon></xmlicon>
							<index>0</index>
							<interfaceid>3</interfaceid>
							<hostname>Unknown</hostname>
							<MACAddr>00:11:22:33:44:77</MACAddr>
							<method>2</method>
							<leaseTime>00:00:47:47</leaseTime>
							<speed>54</speed>
						</clientinfo>
						<clientinfo>
							<interface>home</interface>
							<IPv4Addr>10.0.0.1/24</IPv4Addr>
							<xmlhostname></xmlhostname>
							<xmlicon></xmlicon>
							<index>1</index>
							<interfaceid>19</interfaceid>
							<hostname>Unknown</hostname>
							<MACAddr>00:11:22:33:44:88</MACAddr>
							<method>2</method>
							<leaseTime>00:00:48:48</leaseTime>
							<speed>866</speed>
						</clientinfo>
					</WIFI>
					<totalClient>9</totalClient>
					<Customer>upc</Customer>
				</LanUserTable>`,
			in: &LANUserTable{},
			out: &LANUserTable{
				Ethernet: []LANUserTableEthernet{
					{
						Interface:   "Ethernet 4",
						IPv4Addr:    "10.0.0.1/24",
						XMLHostname: "",
						XMLIcon:     "",
						Index:       "3",
						InterfaceID: "2",
						Hostname:    "Unknown",
						MACAddr:     "00:11:22:33:44:44",
						Method:      "2",
						LeaseTime:   "00:00:00:00",
						Speed:       "1000",
					},
					{
						Interface:   "Ethernet 3",
						IPv4Addr:    "10.0.0.2/24",
						XMLHostname: "",
						XMLIcon:     "",
						Index:       "5",
						InterfaceID: "2",
						Hostname:    "Unknown",
						MACAddr:     "00:11:22:33:44:66",
						Method:      "2",
						LeaseTime:   "00:00:45:45",
						Speed:       "1000",
					},
				},
				WIFI: []LANUserTableWIFI{
					{
						Interface:   "home",
						IPv4Addr:    "10.0.0.1/24",
						XMLHostname: "",
						XMLIcon:     "",
						Index:       "0",
						InterfaceID: "3",
						Hostname:    "Unknown",
						MACAddr:     "00:11:22:33:44:77",
						Method:      "2",
						LeaseTime:   "00:00:47:47",
						Speed:       "54",
					},
					{
						Interface:   "home",
						IPv4Addr:    "10.0.0.1/24",
						XMLHostname: "",
						XMLIcon:     "",
						Index:       "1",
						InterfaceID: "19",
						Hostname:    "Unknown",
						MACAddr:     "00:11:22:33:44:88",
						Method:      "2",
						LeaseTime:   "00:00:48:48",
						Speed:       "866",
					},
				},
				TotalClient: "9",
				Customer:    "upc",
			},
		},
		{
			name: "DDNS",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<DDNS>
					<Enable>0</Enable>
					<DDNSProvider>0</DDNSProvider>
					<Username></Username>
					<Password></Password>
					<Hostname></Hostname>
					<WanIP>10.0.0.1</WanIP>
				</DDNS>`,
			in: &DDNS{},
			out: &DDNS{
				Enable:       "0",
				DDNSProvider: "0",
				Username:     "",
				Password:     "",
				Hostname:     "",
				WanIP:        "10.0.0.1",
			},
		},
		{
			name: "RemoteAccess",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<RemoteAccess />`,
			in:  &RemoteAccess{},
			out: &RemoteAccess{},
		},
		{
			name: "MTUSize",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<MTUSize>
					<size>1500</size>
				</MTUSize>`,
			in: &MTUSize{},
			out: &MTUSize{
				Size: "1500",
			},
		},
		{
			name: "CMState",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<cmstate>
					<TunnerTemperature>80</TunnerTemperature>
					<Temperature>59</Temperature>
					<OperState>OPERATIONAL</OperState>
					<wan_ipv4_addr>10.0.0.1</wan_ipv4_addr>
					<wan_ipv6_addr>
						<wan_ipv6_addr_entry>bbbb:aaaa:0:5555:4444:3333:2222:0000/128</wan_ipv6_addr_entry>
						<wan_ipv6_addr_entry>bbbb::6a02:5555:feee:3333/64</wan_ipv6_addr_entry>
					</wan_ipv6_addr>
				</cmstate>`,
			in: &CMState{},
			out: &CMState{
				TunnerTemperature: 26,
				Temperature:       15,
				OperState:         "OPERATIONAL",
				WANIPv4Addr:       "10.0.0.1",
				WANIPv6Addrs: []string{
					"bbbb:aaaa:0:5555:4444:3333:2222:0000/128",
					"bbbb::6a02:5555:feee:3333/64",
				},
			},
		},
		{
			name: "WiredState1",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<wiredstate>
					<port />
					<port />
					<port>
						<Eth>3</Eth>
						<Speed>1000</Speed>
					</port>
					<port>
						<Eth>4</Eth>
						<Speed>1000</Speed>
					</port>
					<Device>2</Device>
					<ethflaplistFile>Fail</ethflaplistFile>
				</wiredstate>`,
			in: &WiredState1{},
			out: &WiredState1{
				Ports: []WiredState1Port{
					{
						Eth:   "",
						Speed: "",
					},
					{
						Eth:   "",
						Speed: "",
					},
					{
						Eth:   "3",
						Speed: "1000",
					},
					{
						Eth:   "4",
						Speed: "1000",
					},
				},
				Device:          "2",
				EthFlaplistFile: "Fail",
			},
		},
		{
			name: "WiredState2",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<wiredstate>
					<port />
					<port />
					<port>
						<Eth>3</Eth>
						<Speed>1000</Speed>
					</port>
					<port>
						<Eth>4</Eth>
						<Speed>1000</Speed>
					</port>
					<Device>2</Device>
				</wiredstate>`,
			in: &WiredState2{},
			out: &WiredState2{
				Ports: []WiredState2Port{
					{
						Eth:   "",
						Speed: "",
					},
					{
						Eth:   "",
						Speed: "",
					},
					{
						Eth:   "3",
						Speed: "1000",
					},
					{
						Eth:   "4",
						Speed: "1000",
					},
				},
				Device: "2",
			},
		},
		{
			name: "CMStatus",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<cmstatus>
					<provisioning_st>Online</provisioning_st>
					<provisioning_st_num>12</provisioning_st_num>
					<cm_comment>Operational</cm_comment>
					<ds_num>32</ds_num>
					<downstream>
						<freq>682000000</freq>
						<mod>256qam</mod>
						<chid>14</chid>
						<state>4</state>
						<status>0</status>
						<primarySettings>0</primarySettings>
					</downstream>
					<downstream>
						<freq>730000000</freq>
						<mod>256qam</mod>
						<chid>20</chid>
						<state>4</state>
						<status>0</status>
						<primarySettings>1</primarySettings>
					</downstream>
					<us_num>4</us_num>
					<upstream>
						<usid>8</usid>
						<freq>52000000</freq>
						<power>101</power>
						<srate>5.120</srate>
						<state>4</state>
					</upstream>
					<upstream>
						<usid>10</usid>
						<freq>38400000</freq>
						<power>101</power>
						<srate>5.120</srate>
						<state>4</state>
					</upstream>
					<cm_docsis_mode>DOCSIS 3.0</cm_docsis_mode>
					<cm_network_access>Allowed</cm_network_access>
					<NumberOfCpes>45</NumberOfCpes>
					<dMaxCpes>2</dMaxCpes>
					<bpiEnable>1</bpiEnable>
					<FileName>bac1020001066800000001c8</FileName>
					<serviceflow>
						<Sfid>200000001</Sfid>
						<direction>2</direction>
						<pMaxTrafficRate>32100000</pMaxTrafficRate>
						<pMaxTrafficBurst>42600</pMaxTrafficBurst>
						<pMinReservedRate>0</pMinReservedRate>
						<pMaxConcatBurst>42600</pMaxConcatBurst>
						<pSchedulingType>2</pSchedulingType>
					</serviceflow>
					<serviceflow>
						<Sfid>300000001</Sfid>
						<direction>2</direction>
						<pMaxTrafficRate>32100000</pMaxTrafficRate>
						<pMaxTrafficBurst>42600</pMaxTrafficBurst>
						<pMinReservedRate>0</pMinReservedRate>
						<pMaxConcatBurst>42600</pMaxConcatBurst>
						<pSchedulingType>2</pSchedulingType>
					</serviceflow>
				</cmstatus>`,
			in: &CMStatus{},
			out: &CMStatus{
				ProvisioningSt:    "Online",
				ProvisioningStNum: "12",
				CMComment:         "Operational",
				DsNum:             "32",
				Downstreams: []CMStatusDownstream{
					{
						Freq:            "682000000",
						Mod:             "256qam",
						Chid:            "14",
						State:           "4",
						Status:          "0",
						PrimarySettings: "0",
					},
					{
						Freq:            "730000000",
						Mod:             "256qam",
						Chid:            "20",
						State:           "4",
						Status:          "0",
						PrimarySettings: "1",
					},
				},
				UsNum: "4",
				Upstreams: []CMStatusUpstream{
					{
						Usid:  "8",
						Freq:  "52000000",
						Power: "101",
						Srate: "5.120",
						State: "4",
					},
					{
						Usid:  "10",
						Freq:  "38400000",
						Power: "101",
						Srate: "5.120",
						State: "4",
					},
				},
				CMDocsisMode:    "DOCSIS 3.0",
				CMNetworkAccess: "Allowed",
				NumberOfCpes:    "45",
				DMaxCpes:        "2",
				BpiEnable:       "1",
				FileName:        "bac1020001066800000001c8",
				ServiceFlows: []CMStatusServiceFlow{
					{
						Sfid:             "200000001",
						Direction:        "2",
						PMaxTrafficRate:  "32100000",
						PMaxTrafficBurst: "42600",
						PMinReservedRate: "0",
						PMaxConcatBurst:  "42600",
						PSchedulingType:  "2",
					},
					{
						Sfid:             "300000001",
						Direction:        "2",
						PMaxTrafficRate:  "32100000",
						PMaxTrafficBurst: "42600",
						PMinReservedRate: "0",
						PMaxConcatBurst:  "42600",
						PSchedulingType:  "2",
					},
				},
			},
		},
		{
			name: "EthFlaplist",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<ethflaplist>
					<ethflaplistFile>NULL</ethflaplistFile>
				</ethflaplist>`,
			in: &EthFlaplist{},
			out: &EthFlaplist{
				EthFlaplistFile: "NULL",
			},
		},
		{
			name: "WirelessBasic1",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<WirelessBasic>
					<NvCountry>1</NvCountry>
					<Bandmode>3</Bandmode>
					<ChannelRange>2</ChannelRange>
					<BssEnable2g>1</BssEnable2g>
					<SSID2G>home</SSID2G>
					<HideNetwork2G>2</HideNetwork2G>
					<BandWidth2G>1</BandWidth2G>
					<BssCoexistence>1</BssCoexistence>
					<TransmissionRate2g>0</TransmissionRate2g>
					<TransmissionMode2g>6</TransmissionMode2g>
					<SecurityMode2g>4</SecurityMode2g>
					<MulticastRate2G>1</MulticastRate2G>
					<ChannelSetting2G>6</ChannelSetting2G>
					<CurrentChannel2G>11</CurrentChannel2G>
					<PreSharedKey2g>qwerty</PreSharedKey2g>
					<GroupRekeyInterval2g>0</GroupRekeyInterval2g>
					<WpaAlgorithm2G>2</WpaAlgorithm2G>
					<SONAdminStatus>1</SONAdminStatus>
					<SONOperationalStatus>1</SONOperationalStatus>
					<BssEnable5g>1</BssEnable5g>
					<SSID5G>home</SSID5G>
					<HideNetwork5G>2</HideNetwork5G>
					<BandWidth5G>3</BandWidth5G>
					<TransmissionRate5g>0</TransmissionRate5g>
					<TransmissionMode5g>14</TransmissionMode5g>
					<SecurityMode5g>4</SecurityMode5g>
					<MulticastRate5G>1</MulticastRate5G>
					<ChannelSetting5G>48</ChannelSetting5G>
					<CurrentChannel5G>44</CurrentChannel5G>
					<PreSharedKey5g>qwerty</PreSharedKey5g>
					<GroupRekeyInterval5g>0</GroupRekeyInterval5g>
					<WpaAlgorithm5G>2</WpaAlgorithm5G>
				</WirelessBasic>`,
			in: &WirelessBasic1{},
			out: &WirelessBasic1{
				NvCountry:            "1",
				Bandmode:             "3",
				ChannelRange:         "2",
				BSSEnable2G:          "1",
				SSID2G:               "home",
				HideNetwork2G:        "2",
				BandWidth2G:          "1",
				BSSCoexistence:       "1",
				TransmissionRate2G:   "0",
				TransmissionMode2G:   "6",
				SecurityMode2G:       "4",
				MulticastRate2G:      "1",
				ChannelSetting2G:     "6",
				CurrentChannel2G:     "11",
				PreSharedKey2G:       "qwerty",
				GroupRekeyInterval2G: "0",
				WpaAlgorithm2G:       "2",
				SONAdminStatus:       "1",
				SONOperationalStatus: "1",
				BssEnable5G:          "1",
				SSID5G:               "home",
				HideNetwork5G:        "2",
				BandWidth5G:          "3",
				TransmissionRate5G:   "0",
				TransmissionMode5G:   "14",
				SecurityMode5G:       "4",
				MulticastRate5G:      "1",
				ChannelSetting5G:     "48",
				CurrentChannel5G:     "44",
				PreSharedKey5G:       "qwerty",
				GroupRekeyInterval5G: "0",
				WpaAlgorithm5G:       "2",
			},
		},
		{
			name: "WirelessWmm",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<WirelessWmm>
					<WMM2G>1</WMM2G>
					<Apsd2G>2</Apsd2G>
					<TransmissionMode2g>6</TransmissionMode2g>
					<WMM5G>1</WMM5G>
					<Apsd5G>2</Apsd5G>
					<TransmissionMode5g>14</TransmissionMode5g>
				</WirelessWmm>`,
			in: &WirelessWmm{},
			out: &WirelessWmm{
				WMM2G:              "1",
				Apsd2G:             "2",
				TransmissionMode2G: "6",
				WMM5G:              "1",
				Apsd5G:             "2",
				TransmissionMode5G: "14",
			},
		},
		{
			name: "WirelessSiteSurvey",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<WirelessSiteSurvey>
					<count2G>0</count2G>
					<count5G>0</count5G>
					<BandMode_2_4G />
					<BandMode_5G />
				</WirelessSiteSurvey>`,
			in: &WirelessSiteSurvey{},
			out: &WirelessSiteSurvey{
				Count2G:     "0",
				Count5G:     "0",
				BandMode24G: "",
				BandMode5G:  "",
			},
		},
		{
			name: "WirelessGuestNetwork1",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<WirelessGuestNetwork>
					<MainEnable2G>1</MainEnable2G>
					<MainEnable5G>1</MainEnable5G>
					<Interface>
						<Enable2G>1</Enable2G>
						<BSSID2G>Ziggo</BSSID2G>
						<GuestMac2G>00:11:22:33:44:55</GuestMac2G>
						<HideNetwork2G>2</HideNetwork2G>
						<SecurityMode2g>6</SecurityMode2g>
						<PreSharedKey2g></PreSharedKey2g>
						<GroupRekeyInterval2g>0</GroupRekeyInterval2g>
						<WpaAlgorithm2G>3</WpaAlgorithm2G>
					</Interface>
					<Interface>
						<Enable2G>1</Enable2G>
						<BSSID2G>we.connect.hello</BSSID2G>
						<GuestMac2G>00:11:22:33:44:55</GuestMac2G>
						<HideNetwork2G>1</HideNetwork2G>
						<SecurityMode2g>4</SecurityMode2g>
						<PreSharedKey2g>xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx</PreSharedKey2g>
						<GroupRekeyInterval2g>0</GroupRekeyInterval2g>
						<WpaAlgorithm2G>2</WpaAlgorithm2G>
					</Interface>
					<Interface5G>
						<Enable5G>2</Enable5G>
						<BSSID5G>Ziggo</BSSID5G>
						<GuestMac5G>00:11:22:33:44:55</GuestMac5G>
						<HideNetwork5G>2</HideNetwork5G>
						<SecurityMode5g>6</SecurityMode5g>
						<PreSharedKey5g></PreSharedKey5g>
						<GroupRekeyInterval5g>0</GroupRekeyInterval5g>
						<WpaAlgorithm5G>3</WpaAlgorithm5G>
					</Interface5G>
					<Interface5G>
						<Enable5G>2</Enable5G>
						<BSSID5G></BSSID5G>
						<GuestMac5G>00:11:22:33:44:55</GuestMac5G>
						<HideNetwork5G>2</HideNetwork5G>
						<SecurityMode5g>0</SecurityMode5g>
						<PreSharedKey5g></PreSharedKey5g>
						<GroupRekeyInterval5g>0</GroupRekeyInterval5g>
						<WpaAlgorithm5G>3</WpaAlgorithm5G>
					</Interface5G>
				</WirelessGuestNetwork>`,
			in: &WirelessGuestNetwork1{},
			out: &WirelessGuestNetwork1{
				MainEnable2G: "1",
				MainEnable5G: "1",
				Interfaces: []WirelessGuestNetwork1Interface{
					{
						Enable2G:             "1",
						BSSID2G:              "Ziggo",
						GuestMac2G:           "00:11:22:33:44:55",
						HideNetwork2G:        "2",
						SecurityMode2G:       "6",
						PreSharedKey2G:       "",
						GroupRekeyInterval2G: "0",
						WPAAlgorithm2G:       "3",
					},
					{
						Enable2G:             "1",
						BSSID2G:              "we.connect.hello",
						GuestMac2G:           "00:11:22:33:44:55",
						HideNetwork2G:        "1",
						SecurityMode2G:       "4",
						PreSharedKey2G:       "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
						GroupRekeyInterval2G: "0",
						WPAAlgorithm2G:       "2",
					},
				},
				Interfaces5G: []WirelessGuestNetwork1Interface5G{
					{
						Enable5G:             "2",
						BSSID5G:              "Ziggo",
						GuestMac5G:           "00:11:22:33:44:55",
						HideNetwork5G:        "2",
						SecurityMode5G:       "6",
						PreSharedKey5G:       "",
						GroupRekeyInterval5G: "0",
						WPAAlgorithm5G:       "3",
					},
					{
						Enable5G:             "2",
						BSSID5G:              "",
						GuestMac5G:           "00:11:22:33:44:55",
						HideNetwork5G:        "2",
						SecurityMode5G:       "0",
						PreSharedKey5G:       "",
						GroupRekeyInterval5G: "0",
						WPAAlgorithm5G:       "3",
					},
				},
			},
		},
		{
			name: "CMWirelessWPS1",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<cm_wirelessWPS>
					<MainEnable2g>1</MainEnable2g>
					<MainEnable5g>1</MainEnable5g>
					<WpsEnable24G>1</WpsEnable24G>
					<WpsEnable5G>1</WpsEnable5G>
					<WpsMethod24G>1</WpsMethod24G>
					<WpsMethod5G>1</WpsMethod5G>
					<WpsAPPIN24G>00000000</WpsAPPIN24G>
					<WpsAPPIN5G>00000000</WpsAPPIN5G>
					<WpsPINNUM24G></WpsPINNUM24G>
					<WpsPINNUM5G></WpsPINNUM5G>
					<WpsEnablePBC>1</WpsEnablePBC>
					<WpsEnablePIN>2</WpsEnablePIN>
					<WpsEnablePBC5G>1</WpsEnablePBC5G>
					<WpsEnablePIN5G>2</WpsEnablePIN5G>
				</cm_wirelessWPS>`,
			in: &CMWirelessWPS1{},
			out: &CMWirelessWPS1{
				MainEnable2G:   "1",
				MainEnable5G:   "1",
				WPSEnable24G:   "1",
				WPSEnable5G:    "1",
				WPSMethod24G:   "1",
				WPSMethod5G:    "1",
				WPSAPPin24G:    "00000000",
				WPSAPPin5G:     "00000000",
				WPSPinNum24G:   "",
				WPSPinNum5G:    "",
				WPSEnablePBC:   "1",
				WPSEnablePIN:   "2",
				WPSEnablePBC5G: "1",
				WPSEnablePIN5G: "2",
			},
		},
		{
			name: "CMWirelessAccessControl",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<cm_wirelessAccessControl>
					<BandMode>3</BandMode>
					<BssEnable2g>1</BssEnable2g>
					<BssEnable5g>1</BssEnable5g>
					<SSID2G>home</SSID2G>
					<SSID5G>home</SSID5G>
					<HideNetwork2G>2</HideNetwork2G>
					<HideNetwork5G>2</HideNetwork5G>
					<SecurityMode2g>4</SecurityMode2g>
					<SecurityMode5g>4</SecurityMode5g>
					<PreSharedKey2g>qwerty</PreSharedKey2g>
					<PreSharedKey5g>qwerty</PreSharedKey5g>
					<WpaAlgorithm2G>2</WpaAlgorithm2G>
					<WpaAlgorithm5G>2</WpaAlgorithm5G>
					<AccessMode24G>3</AccessMode24G>
					<AccessMode5G>3</AccessMode5G>
					<BssAccessEntry>
						<AccessStation>00:11:22:33:44:55</AccessStation>
						<AccessDeviceName></AccessDeviceName>
					</BssAccessEntry>
					<BssAccessEntry>
						<AccessStation>00:11:22:33:44:66</AccessStation>
						<AccessDeviceName></AccessDeviceName>
					</BssAccessEntry>
					<BssAccessEntry5G>
						<AccessStation5G>00:11:22:33:44:77</AccessStation5G>
						<AccessDeviceName5G></AccessDeviceName5G>
					</BssAccessEntry5G>
					<BssAccessEntry5G>
						<AccessStation5G>00:11:22:33:44:88</AccessStation5G>
						<AccessDeviceName5G></AccessDeviceName5G>
					</BssAccessEntry5G>
				</cm_wirelessAccessControl>`,
			in: &CMWirelessAccessControl{},
			out: &CMWirelessAccessControl{
				BandMode:       "3",
				BSSEnable2G:    "1",
				BSSEnable5G:    "1",
				SSID2G:         "home",
				SSID5G:         "home",
				HideNetwork2G:  "2",
				HideNetwork5G:  "2",
				SecurityMode2G: "4",
				SecurityMode5G: "4",
				PreSharedKey2G: "qwerty",
				PreSharedKey5G: "qwerty",
				WpaAlgorithm2G: "2",
				WpaAlgorithm5G: "2",
				AccessMode24G:  "3",
				AccessMode5G:   "3",
				BSSAccessEntries: []CMWirelessAccessControlBSSAccessEntry{
					{
						AccessStation:    "00:11:22:33:44:55",
						AccessDeviceName: "",
					},
					{
						AccessStation:    "00:11:22:33:44:66",
						AccessDeviceName: "",
					},
				},
				BSSAccessEntries5G: []CMWirelessAccessControlBSSAccessEntry5G{
					{
						AccessStation5G:    "00:11:22:33:44:77",
						AccessDeviceName5G: "",
					},
					{
						AccessStation5G:    "00:11:22:33:44:88",
						AccessDeviceName5G: "",
					},
				},
			},
		},
		{
			name: "ChannelMap",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<ChannelMap>
					<count2G>0</count2G>
					<MyCurrentChannel2G>11</MyCurrentChannel2G>
					<count5G>0</count5G>
					<MyCurrentChannel5G>44</MyCurrentChannel5G>
					<BandMode_2_4G>
						<W2GCH1>0</W2GCH1>
						<W2GCH2>0</W2GCH2>
						<W2GCH3>0</W2GCH3>
						<W2GCH4>0</W2GCH4>
						<W2GCH5>0</W2GCH5>
						<W2GCH6>0</W2GCH6>
						<W2GCH7>0</W2GCH7>
						<W2GCH8>0</W2GCH8>
						<W2GCH9>0</W2GCH9>
						<W2GCH10>0</W2GCH10>
						<W2GCH11>0</W2GCH11>
						<W2GCH12>0</W2GCH12>
						<W2GCH13>0</W2GCH13>
						<maxaxis2G>14</maxaxis2G>
						<total2g>0</total2g>
					</BandMode_2_4G>
					<BandMode_5G>
						<W5GCH1>0</W5GCH1>
						<W5GCH2>0</W5GCH2>
						<W5GCH3>0</W5GCH3>
						<W5GCH4>0</W5GCH4>
						<W5GCH5>0</W5GCH5>
						<W5GCH6>0</W5GCH6>
						<W5GCH7>0</W5GCH7>
						<W5GCH8>0</W5GCH8>
						<W5GCH9>0</W5GCH9>
						<W5GCH10>0</W5GCH10>
						<W5GCH11>0</W5GCH11>
						<W5GCH12>0</W5GCH12>
						<W5GCH13>0</W5GCH13>
						<W5GCH14>0</W5GCH14>
						<W5GCH15>0</W5GCH15>
						<W5GCH16>0</W5GCH16>
						<W5GCH17>0</W5GCH17>
						<W5GCH18>0</W5GCH18>
						<W5GCH19>0</W5GCH19>
						<maxaxis5G>14</maxaxis5G>
						<total5g>0</total5g>
					</BandMode_5G>
				</ChannelMap>`,
			in: &ChannelMap{},
			out: &ChannelMap{
				Count2G:            "0",
				MyCurrentChannel2G: "11",
				Count5G:            "0",
				MyCurrentChannel5G: "44",
				BandMode24G: ChannelMapBandMode24G{
					W2GCH1:    "0",
					W2GCH2:    "0",
					W2GCH3:    "0",
					W2GCH4:    "0",
					W2GCH5:    "0",
					W2GCH6:    "0",
					W2GCH7:    "0",
					W2GCH8:    "0",
					W2GCH9:    "0",
					W2GCH10:   "0",
					W2GCH11:   "0",
					W2GCH12:   "0",
					W2GCH13:   "0",
					Maxaxis2G: "14",
					Total2G:   "0",
				},
				BandMode5G: ChannelMapBandMode5G{
					W5GCH1:    "0",
					W5GCH2:    "0",
					W5GCH3:    "0",
					W5GCH4:    "0",
					W5GCH5:    "0",
					W5GCH6:    "0",
					W5GCH7:    "0",
					W5GCH8:    "0",
					W5GCH9:    "0",
					W5GCH10:   "0",
					W5GCH11:   "0",
					W5GCH12:   "0",
					W5GCH13:   "0",
					W5GCH14:   "0",
					W5GCH15:   "0",
					W5GCH16:   "0",
					W5GCH17:   "0",
					W5GCH18:   "0",
					W5GCH19:   "0",
					Maxaxis5G: "14",
					Total5G:   "0",
				},
			},
		},
		{
			name: "WirelessBasic2",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<WirelessBasic>
					<Bandmode>3</Bandmode>
					<BssEnable2g>1</BssEnable2g>
					<BssEnable5g>1</BssEnable5g>
					<WiFi_chip_status>2</WiFi_chip_status>
					<cm_status>Online</cm_status>
				</WirelessBasic>`,
			in: &WirelessBasic2{},
			out: &WirelessBasic2{
				Bandmode:       "3",
				BSSEnable2G:    "1",
				BSSEnable5G:    "1",
				WiFiChipStatus: "2",
				CMStatus:       "Online",
			},
		},
		{
			name: "WirelessGuestNetwork2",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<WirelessGuestNetwork>
					<year>0</year>
					<mouth>0</mouth>
					<day>0</day>
					<hour>0</hour>
					<minute>0</minute>
					<Interface>
						<MainEnable2G>1</MainEnable2G>
						<Enable2G>2</Enable2G>
						<BSSID2G>Ziggo-XX</BSSID2G>
						<GuestMac2G>00:11:22:33:44:55</GuestMac2G>
						<HideNetwork2G>2</HideNetwork2G>
						<SecurityMode2g>4</SecurityMode2g>
						<PreSharedKey2g>xxxxxxxxxxxx</PreSharedKey2g>
						<GroupRekeyInterval2g>0</GroupRekeyInterval2g>
						<WpaAlgorithm2G>2</WpaAlgorithm2G>
					</Interface>
					<Interface5G>
						<MainEnable5G>1</MainEnable5G>
						<Enable5G>2</Enable5G>
						<BSSID5G>Ziggo-XX</BSSID5G>
						<GuestMac5G>00:11:22:33:44:55</GuestMac5G>
						<HideNetwork5G>2</HideNetwork5G>
						<SecurityMode5g>4</SecurityMode5g>
						<PreSharedKey5g>xxxxxxxxxxxx</PreSharedKey5g>
						<GroupRekeyInterval5g>0</GroupRekeyInterval5g>
						<WpaAlgorithm5G>2</WpaAlgorithm5G>
					</Interface5G>
				</WirelessGuestNetwork>`,
			in: &WirelessGuestNetwork2{},
			out: &WirelessGuestNetwork2{
				Year:   "0",
				Mouth:  "0",
				Day:    "0",
				Hour:   "0",
				Minute: "0",
				Interface: WirelessGuestNetwork2Interface{
					MainEnable2G:         "1",
					Enable2G:             "2",
					BSSID2G:              "Ziggo-XX",
					GuestMac2G:           "00:11:22:33:44:55",
					HideNetwork2G:        "2",
					SecurityMode2G:       "4",
					PreSharedKey2G:       "xxxxxxxxxxxx",
					GroupRekeyInterval2G: "0",
					WPAAlgorithm2G:       "2",
				},
				Interface5G: WirelessGuestNetwork2Interface5G{
					MainEnable5G:         "1",
					Enable5G:             "2",
					BSSID5G:              "Ziggo-XX",
					GuestMac5G:           "00:11:22:33:44:55",
					HideNetwork5G:        "2",
					SecurityMode5G:       "4",
					PreSharedKey5G:       "xxxxxxxxxxxx",
					GroupRekeyInterval5G: "0",
					WPAAlgorithm5G:       "2",
				},
			},
		},
		{
			name: "WirelessClient",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<WirelessClient>
					<Client2G>
						<clientinfo>
							<SSID>home</SSID>
							<MAC>00:11:22:33:44:55</MAC>
							<phy_rate_tx>130000000</phy_rate_tx>
							<phy_rate_rx>130000000</phy_rate_rx>
							<phy_mode>3</phy_mode>
							<Auth_mode>3</Auth_mode>
							<RSSI>11</RSSI>
							<EncryptMethod>1</EncryptMethod>
						</clientinfo>
					</Client2G>
					<Client5G>
						<clientinfo>
							<SSID>home</SSID>
							<MAC>00:11:22:33:44:55</MAC>
							<phy_rate_tx>65000000</phy_rate_tx>
							<phy_rate_rx>65000000</phy_rate_rx>
							<phy_mode>3</phy_mode>
							<Auth_mode>3</Auth_mode>
							<RSSI>10</RSSI>
							<EncryptMethod>1</EncryptMethod>
						</clientinfo>
					</Client5G>
				</WirelessClient>`,
			in: &WirelessClient{},
			out: &WirelessClient{
				Client2G: []WirelessClientClient2G{{
					ClientInfo: []WirelessClientClient2GClientInfo{{
						SSID:          "home",
						MAC:           "00:11:22:33:44:55",
						PhyRateTx:     "130000000",
						PhyRateRx:     "130000000",
						PhyMode:       "3",
						AuthMode:      "3",
						RSSI:          "11",
						EncryptMethod: "1",
					}},
				}},
				Client5G: []WirelessClientClient5G{{
					ClientInfo: []WirelessClientClient5GClientInfo{{
						SSID:          "home",
						MAC:           "00:11:22:33:44:55",
						PhyRateTx:     "65000000",
						PhyRateRx:     "65000000",
						PhyMode:       "3",
						AuthMode:      "3",
						RSSI:          "10",
						EncryptMethod: "1",
					}},
				}},
			},
		},
		{
			name: "CMWirelessWPS2",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<cm_wirelessWPS>
					<WPS_stat>down</WPS_stat>
					<WPS_result></WPS_result>
				</cm_wirelessWPS>`,
			in: &CMWirelessWPS2{},
			out: &CMWirelessWPS2{
				WPSStat:   "down",
				WPSResult: "",
			},
		},
		{
			name: "DefaultValue",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<DefaultValue>
					<loginPwd>00000000</loginPwd>
					<WiFiSSID>Ziggo0000000</WiFiSSID>
					<WiFikey>xxxxxxxxxxxx</WiFikey>
				</DefaultValue>`,
			in: &DefaultValue{},
			out: &DefaultValue{
				LoginPwd: "00000000",
				WIFISSID: "Ziggo0000000",
				WIFIkey:  "xxxxxxxxxxxx",
			},
		},
		{
			name: "GstRandomPassword",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<GstRandomPassword>
					<PreSharedKey>aaaaaaaaaaaaaa</PreSharedKey>
				</GstRandomPassword>`,
			in: &GstRandomPassword{},
			out: &GstRandomPassword{
				PreSharedKey: "aaaaaaaaaaaaaa",
			},
		},
		{
			name: "WIFIState",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<wifistate>
					<primary24g>1</primary24g>
					<primary5g>1</primary5g>
				</wifistate>`,
			in: &WIFIState{},
			out: &WIFIState{
				Primary24G: "1",
				Primary5G:  "1",
			},
		},
		{
			name: "WirelessResetting",
			data: `<?xml version="1.0" encoding="utf-8"?>
				<WirelessResetting>
					<isWirelessResetting>0</isWirelessResetting>
				</WirelessResetting>`,
			in: &WirelessResetting{},
			out: &WirelessResetting{
				IsWirelessResetting: "0",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := xml.Unmarshal([]byte(tc.data), tc.in)
			require.NoError(t, err)
			require.Equal(t, tc.out, tc.in)
		})
	}
}

func TestParseDuration(t *testing.T) {
	testCases := []struct {
		name string
		str  string
		dur  time.Duration
		err  string
	}{
		{
			name: "valid duration",
			str:  "10day(s)20h:15m:30s",
			dur:  936930 * time.Second,
		},
		{
			name: "only seconds",
			str:  "0day(s)0h:0m:30s",
			dur:  30 * time.Second,
		},
		{
			name: "invalid duration",
			str:  "hello, world",
			err:  "invalid duration string",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			dur, err := parseDuration(tc.str)
			if tc.err == "" {
				require.NoError(t, err)
				require.Equal(t, tc.dur, dur)
			} else {
				require.ErrorContains(t, err, tc.err)
			}
		})
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	testCases := []struct {
		name       string
		fahrenheit int
		celsius    int
	}{
		{
			name:       "above zero",
			fahrenheit: 50,
			celsius:    10,
		},
		{
			name:       "below zero",
			fahrenheit: -50,
			celsius:    -45,
		},
		{
			name:       "zero fahrenheit",
			fahrenheit: 0,
			celsius:    -17,
		},
		{
			name:       "zero celsius",
			fahrenheit: 32,
			celsius:    0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			c := fahrenheitToCelsius(tc.fahrenheit)
			if c != tc.celsius {
				t.Fatalf("Wrong column\n  expected: %d\n       got: %d", tc.celsius, c)
			}
		})
	}
}
